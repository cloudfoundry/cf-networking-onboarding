Stale Routes Reenactment

## Assumptions
- You have a CF deployed

## What
Recap from http-route track: the GoRouter redirects traffic for a particular CF HTTP route to a Diego Cell IP and port. However, when one app is deleted, another app may later use the same Diego Cell IP and port.

In the happy path case, the following steps should happen in the following order:

- appA is pushed and is at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appA to 10.0.1.2:12345
- appA is deleted
- the Route Emitter sends an unregister route message to the GoRouter
- 👍 **the GoRouter does NOT forward traffic for appA to 10.0.1.2:12345**
- appB is pushed and is now at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appB to 10.0.1.2:12345

That's a lot of components working together to make sure routes are sent to the correct place. Now imagine this situation where the Route Emitter is not sending out messages fast enough.

- appA is pushed and is at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appA to 10.0.1.2:12345
- appA is deleted
- ~~the Route Emitter sends an unregister route message to the GoRouter~~
- ~~the GoRouter does NOT forward traffic for appA to 10.0.1.2:12345~~
- appB is pushed and is now at 10.0.1.2:12345
- ~~the Route Emitter sends a register route message to the GoRouter~~
- **the GoRouter continues to forward traffic for appA to 10.0.1.2:12345, where appA used to be, where appB currently is**
- 😱 the user tried to access appA, but they were routed to appB instead

Yikes! This is an example of misrouting due to "stale" routes.

Let's set up a situation with stale routes manually and see how dangerous it is. Then in the next story we will look at how Route Integrity fixes this issue.

The Networking Program has had to fix several bugs related to stale routes. To replicate the bugs and test our fixes, we came up with the following method to simulate stale routes.
It took a surprisingly long time to come up with this idea. Enjoy our hard work. :)

## How

🤔 **cause stale routes**
1. Route Integrity is turned on by default with CF Deployment. Let's turn it off and see what happened before Route Integrity. But first, let's save the current manifest for the next story.
 ```
bosh manifest > /tmp/env-with-route-integrity.yml
 ```

1. Redeploy your CF with Route Integrity turned off, using [this opsfile](https://github.com/cloudfoundry/cf-deployment/pull/745).

1. Push one instance of [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) and call it appA.

1. Ensure that appA has an HTTP route.

1. Push one instance of [dora](https://github.com/cloudfoundry/cf-acceptance-tests/tree/master/assets/dora) and name it appB.

1. Ensure that there is an HTTP Route mapped to appB.

1. Make sure an instance of appA and appB are on the same Diego Cell. This can be achieved through scaling appB if it doesn't happen at first.

1. Curl appA. See that you get the expected result for appA.

1. In previous http-route stories, you learned that the route table on the GoRouter maps routes to particular Diego Cell IPs and ports. Then the GoRouter sends traffic to the Diego Cell IP and port. Then the traffic is rerouted to an app's overlayIP and app port using iptables rules on the nat table. Change the iptables rules so that the Diego Cell port that maps to appA now, incorrectly, redirects traffic to appB. This will imitate route staleness. Go back to the http-route DNAT stories if you need help doing this.

1. Curl the route for appA.

### Expected Result
Even though you curled the route for appA, because the routes for appA were "stale" you got back the information for appB.

## Questions

❓ Why is this dangerous?
❓ What is the implications for this on a shared production environment, like PWS?
❓ What could cause routes to become stale? What could exacerbate this problem?
❓ Before Route Integrity, what could app devs do to prevent stale routes from accidentally releasing their data?

L: route-integrity
L: deploy
L: questions
---

Route Integrity Demonstration

## Assumptions
- You have a CF deployed

## What

You saw the dangers of stale routes in the last story. This story will explore how Route Integrity solves this issue. Route Integrity is enabled by enforcing TLS (or mTLS) between the GoRouter and the apps it routes to.
By default only TLS (not mTLS) is enforced. This is because mTLS will break the cf ssh command.

(Side Note: Confusingly, this feature is called "Route Integrity" by engineers, but is called "(m)TLS to apps" or "encrypting traffic from the GoRouter to backends" in documentation. I will continue calling it Route Integrity, because it is more succinct and sounds cooler.)

Let's set up the same experiment as the previous story and see that even when you imitate route staleness, that you are no longer able to see the data from the incorrect app.

## How
🤔 **Send traffic to a stale route**

1. Deploy your CF with Route Integrity enabled with TLS. Use the manifest you saved in the last story at `/tmp/env-with-route-integrity.yml`.

1. Push an app, appA with an HTTP route.

1. Push a different app, appB, that will respond differently than appA. Ensure there is an HTTP Route mapped to appB. Make sure an instance of appA and appB are on the same Diego Cell. You can achieve this by scaling appB, but make sure there is only 1 instance of appA.

1. Curl appA. See that you get the expected result.

1. Change the iptables rules so that the Diego Cell port that maps to appA now, incorrectly, redirects traffic to appB. This will imitate route staleness.

1. In another terminal, tail the gorouter.stdout.log.

1. Curl the route for appA.

### Expected Result
When you curl appA, you should get a 503 status code.
```
$ curl appA.cf-app.com
503 Service Unavailable
```

In the GoRouter bosh logs you should see a line about pruning the bad route.

```
{"log_level":3,"timestamp":1554761639.745829,"message":"prune-failed-endpoint","source":"vcap.gorouter.registry","data":{"route-endpoint":{"ApplicationId":"bf1cdd21-ed5b-429b-93ef-f74a4aadc55f","Addr":"10.0.1.13:61010","Tags":{"component":"route-emitter"},"RouteServiceUrl":""}}}

{"log_level":3,"timestamp":1554761639.7473993,"message":"backend-endpoint-failed","source":"vcap.gorouter","data":{"route-endpoint":{"ApplicationId":"bf1cdd21-ed5b-429b-93ef-f74a4aadc55f","Addr":"10.0.1.13:61010","Tags":{"component":"route-emitter"},"RouteServiceUrl":""},"error":"x509: certificate is valid for 0c2ca10a-0c19-4288-68e8-8d1f, not e5fdb45c-1b41-4602-45ac-b49e","attempt":1,"vcap_request_id":"3673ccd9-44ee-4d81-4a45-0ea75a96121c"}}
```

Route Integrity doesn't actually prevent route staleness, but it *does* prevent leaking data from appB when you are trying to access appA.

## Look at the code
GoRouter analyses different types of errors from the backend. In the case of misrouting, we are running into this [HostnameMismatch error](https://github.com/cloudfoundry/gorouter/blob/master/proxy/fails/basic_classifiers.go#L44-L51).

Depending on the type of error, GoRouter may choose to retry or not. Look at these [failure classifiers](https://github.com/cloudfoundry/gorouter/blob/master/proxy/fails/classifier_group.go#L5-L17).

## Questions
1. ❓Are HostnameMismatch errors retried?
1. ❓If you had multiple instances of appA what do think would've happened? (Extra Credit: Try it out. Were you correct?)
1. ❓Route Integrity was introduced to fix misrouting due to route staleness. What other benefits does Route Integrity provide?
1. ❓Keep curling the HTTP route for appA several times over 1 minute. Sometimes you get a 503 and sometimes you get a 404. Why?

## Links
[CF Docs - TLS to Apps](https://docs.cloudfoundry.org/concepts/http-routing.html#tls-to-back-end)

L: route-integrity
L: deploy
L: questions
---

Checking the SAN for Route Integrity

## Assumptions
- You have a CF deployed with Route Integrity enabled
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE

## What
In the last story you saw that Route Integrity ensures that appA.cf-app.com is *actually going to* appA and not misrouting to appB due to stale routes. But how does it  work?

Route Integrity is implemented by giving the GoRouter a [CA certificate](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/gorouter/spec#L106-L107) at deploy time. At runtime, when an app container is created, Diego generates a certificate from that CA for each app instance and puts it in the app container. The certificate has a unique subject alternative name (SAN) per app instance.

Diego runs a sidecar Envoy inside of the app container which intercepts all incoming traffic from the GoRouter. It uses the generated certificate to terminate TLS before forwarding traffic to the app. When the GoRouter checks the certificate from the Envoy sidecar, it checks that the SAN matches the app instance id that it has for that IP and port. This way GoRouter is able to make sure that appA.cf-app.com is *actually going to* appA and not misrouting to appB due to stale routes.

All of this work means that the app developer does not need to do anything in order for the GoRouter to app traffic to be sent over TLS. It should "just work".

Let's see how the GoRouter compares the application certificate SAN to make sure it is routing correctly.

## How
🤔 **Find the SAN in the Routes table**
1. Bosh ssh onto the Router VM.
1. Look at the routes table (see the story "Route Propagation - Part 4 - GoRouter" if you need a reminder on how to do this).
1. Find the entry for APP_A_ROUTE.
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
 ```

📝  **Look at the SAN in the certificate**

1. Get into the app container
 ```
cf ssh appA
 ```
1. get the location of the certificate
 ```
env | grep CF_INSTANCE_CERT
 ```
1. Look at the certificate
 ```
cat $CF_INSTANCE_CERT
 ```
1. Use openssl to read the certificate
 ```
openssl x509 -text -in  $CF_INSTANCE_CERT
 ```

You should see a section that looks like this
 ```
 X509v3 Subject Alternative Name:
          DNS:16b6f1fc-347c-4565-71e0-bf7f,   <-------  This matches the server_cert_domain_san
          IP Address:10.255.96.8              <-------  The overlay ip of the app instance
 ```

### Expected Result
You should see the same SAN value stored in GoRouter routes table and in the app's certificate.

❓If you had multiple instances of appA would they have the same SAN? (Extra Credit: Try it out. Were you correct?)

## Links
[CF Docs - TLS to Apps](https://docs.cloudfoundry.org/concepts/http-routing.html#tls-to-back-end)
[Using Instance Identity Credentials](https://docs.cloudfoundry.org/devguide/deploy-apps/instance-identity.html)

L: route-integrity
L: questions
---

[RELEASE] Route Integrity ⇧
L: route-integrity
