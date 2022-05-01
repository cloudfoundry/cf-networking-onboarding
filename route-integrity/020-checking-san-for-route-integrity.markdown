---
layout: single
title: Checking the SAN
permalink: /route-integrity/checking-the-san
sidebar:
  title: "Route Integrity"
  nav: sidebar-route-integrity
---

## Assumptions
- You have a CF deployed with Route Integrity enabled
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE

## What
In the last story you saw that Route Integrity ensures that appA.cf-app.com is
*actually going to* appA and not misrouting to appB due to stale routes. But
how does it  work?

Route Integrity is implemented by giving the GoRouter a [CA
certificate](https://github.com/cloudfoundry/routing-release/blob/f25df8de6aaa3fca02bd51343df70bd800d0ab75/jobs/gorouter/spec#L125-L126)
at deploy time. At runtime, when an app container is created, Diego generates a
certificate from that CA for each app instance and puts it in the app
container. The certificate has a unique subject alternative name (SAN) per app
instance.

Diego runs a sidecar Envoy inside of the app container which intercepts all
incoming traffic from the GoRouter. It uses the generated certificate to
terminate TLS before forwarding traffic to the app. When the GoRouter checks
the certificate from the Envoy sidecar, it checks that the SAN matches the app
instance id that it has for that IP and port. This way GoRouter is able to make
sure that appA.cf-app.com is *actually going to* appA and not misrouting to
appB due to stale routes.

All of this work means that the app developer does not need to do anything in
order for the GoRouter to app traffic to be sent over TLS. It should "just
work".

Let's see how the GoRouter compares the application certificate SAN to make
sure it is routing correctly.

## How
🤔 **Find the SAN in the Routes table**
1. Bosh ssh onto the Router VM.
2. Look at the routes table (see the story ["Route Propagation Part 4 - GoRouter"](../http-routes/route-propagation-pt-4) if you need a reminder on how to do this).
3. Find the entry for APP_A_ROUTE.

   It should look something like this
   ```
   "proxy.apps.mitre.c2c.cf-app.com": [
     {
       "address": "10.0.1.18:61022",
       "tls": true,
       "ttl": 120,
       "tags": {
         "component": "route-emitter"
       },
       "private_instance_id": "16b6f1fc-347c-4565-71e0-bf7f",
       "server_cert_domain_san": "16b6f1fc-347c-4565-71e0-bf7f"       <------- This is the SAN that GoRouter is matching against.
     },
   }
   ```

📝  **Look at the SAN in the certificate**

1. Get into the app container
   {% include codeHeader.html %}
   ```bash
   cf ssh appA
   ```
1. get the location of the certificate
   {% include codeHeader.html %}
   ```bash
   env | grep CF_INSTANCE_CERT
   ```
1. Look at the certificate
   {% include codeHeader.html %}
   ```bash
   cat $CF_INSTANCE_CERT
   ```
1. Use openssl to read the certificate
   {% include codeHeader.html %}
   ```bash
   openssl x509 -text -in  $CF_INSTANCE_CERT
   ```

   You should see a section that looks like this
   ```
   X509v3 Subject Alternative Name:
     DNS:16b6f1fc-347c-4565-71e0-bf7f,   <-------  This matches the server_cert_domain_san
     IP Address:10.255.96.8              <-------  The overlay ip of the app instance
   ```

### Expected Result
You should see the same SAN value stored in GoRouter routes table and in the
app's certificate.

## ❓ Question
If you had multiple instances of appA would they have the same SAN? (Extra
Credit: Try it out. Were you correct?)

## Resources
* [CF Docs - TLS to Apps](https://docs.cloudfoundry.org/concepts/http-routing.html#tls-to-back-end)
* [Using Instance Identity Credentials](https://docs.cloudfoundry.org/devguide/deploy-apps/instance-identity.html)
