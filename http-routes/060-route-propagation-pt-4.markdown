---
layout: single
title: Route Propagation Part 4 - GoRouter
permalink: /http-routes/route-propagation-pt-4
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
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
So the Route Emitter emits routes via the NATS message Bus. GoRouter subscribes to those messages and keeps a route table that is uses to route network traffic bound for CF apps and CF components.

Let's take a look at that route table.
## How

📝 **look at route table**
0. Bosh ssh onto the router vm and become root.
0. Get the username and password for the routing api
 ```
 head /var/vcap/jobs/gorouter/config/gorouter.yml
 ```
0. Get the routes table
 ```
 curl "http://USERNAME:PASSWORD@localhost:8080/routes" | jq .
 ```
0. Scroll through and look at the routes.
  ❓How does this differ from the route information you saw in Cloud Controller?
   For example, you should see routes for CF components, like UAA and doppler.
   This because the GoRouter is in charge of routing traffic to CF apps *AND* to CF components.
0. Find APP_A_ROUTE in the list of routes. Let's dissect the most important bits.
    ```
    "proxy.meow.cloche.c2c.cf-app.com": [   <------ The name of the route! This should match APP_A_ROUTE
        {
          "address": "10.0.1.12:61014",     <------ This is where GoRouter will send traffic for this route. This should match DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT
          "tls": true                       <------ This means Route Integrity is turned on, so the GoRouter will use send traffic to this app over TLS
        }
      ]
    ```
    See how the traffic is being sent to `10.0.1.12:61014` or DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT?
    This means all traffic is being sent to the sidecar envoy via TLS, this is because route integrity is enabled.
    ❓What port do you think would be listed here if route integrity was not enabled?

### Expected Result
Access the route table on the router vm. Inspect app routes and CF component routes.

See that the GoRouter sends traffic for this route to DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT.

The route has now been propagated all the way to the Gorouter! In the next stories we will learn what happens when somone uses that route.
## Resources

[GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR.
Amelia will be eternally grateful. How? Open [this file in
GitHub](https://github.com/cloudfoundry/cf-networking-onboarding). Search for
the phrase you want to edit. Make the fix!_
