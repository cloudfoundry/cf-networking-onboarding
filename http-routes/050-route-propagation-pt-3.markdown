---
layout: single
title: Route Propagation Part 3 - Route Emitter and NATS
permalink: /http-routes/route-propagation-pt-3
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What

There is one Route Emitter per Diego Cell and its job is to... emit routes.
According to the ever helpful [Diego Design
Notes](https://github.com/cloudfoundry/diego-design-notes) the Route Emitter
"monitors DesiredLRP state and ActualLRP state via the BBS. When a change is
detected, the Route Emitter emits route registration and unregistration
messages to the GoRouter via the NATS message bus." Even when no change is
detected, the Route Emitter will periodically emit the entire routing table as
a kind of heartbeat.

For this story, let's look at the messages that the Route Emitter is publishing
via NATS. Subscribing to these NATs messages can be a helpful debugging
technique.

## How

📝 **subscribe to NATs messages**
0. Bosh ssh onto the Diego Cell where your app is running and become root
0. Get the NATS cli
    ```
wget https://github.com/nats-io/natscli/releases/download/v0.0.32/nats-0.0.32-linux-amd64.zip
unzip nats-0.0.32-linux-amd64.zip
chmod +x nats-0.0.32-linux-amd64
mv nats-0.0.32-linux-amd64/nats /usr/bin

    ```
0. Get NATS username, password, and server address
    ```
    jq . /var/vcap/jobs/route_emitter/config/route_emitter.json | grep nats
    ```
0. Use the nats cli to connect to nats: `nats sub "*.*" -s nats://NATS_USERNAME:NATS_PASSWORD@NATS_ADDRESS/ --tlscert <cert file from json> --tlskey <key file from json> --tlsca <ca file from json>`. The `"*.*"` means that you are subscribing to all NATs messages.
    The Route Emitter registers routes every 20 seconds (by default) so that the GoRouter (which subscribes to these messages) has the most up-to-date information about which IPs map to which apps and routes. Depending on how many routes there are, this might be a lot of information.

0. When you successfully connect to nats, plus a few seconds of waiting, you
   should see a message that contains information about the route you created.
   It will look something like this and contain APP_A_ROUTE:
 ```
   [#32] Received on [router.register] :
{
    "host": "10.0.1.12",
    "port": 61012,
    "tls_port": 61014,
    "uris": [
        "proxy.meow.cloche.c2c.cf-app.com"     <--- This should match APP_A_ROUTE
      ],
    "app": "6856799f-aebf-4e2b-81a5-28c74dfb6162",
     "private_instance_id": "a0d2b217-fa7d-4ac1-65a2-7b19",
     "private_instance_index": "0",
    "server_cert_domain_san": "a0d2b217-fa7d-4ac1-65a2-7b19",
    "tags": {
         "component": "route-emitter"
     }
}
 ```
   Nats is used by multiple services, and the `router.register` messages are intermixed with other messages.  Modify the sub filter from `"*.*"` to `"router.register"` to view only router messages.

## ❓ Questions
* Do the values in the NATS message match the values you recorded previously
  from BBS? Which ones are present? Which ones aren't there?
* How does it compare to the information in Cloud Controller?

## Expected Result
Inspect NATs messages. Look at what route information is sent to the GoRouter.

## Resources
* [NATS message bus repo](https://github.com/nats-io/gnatsd)
* [NATS ruby gem repo](https://github.com/nats-io/ruby-nats)
