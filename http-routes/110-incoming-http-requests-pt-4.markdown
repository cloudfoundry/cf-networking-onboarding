---
layout: single
title: Incoming HTTP Requests Part 4 - Route Integrity and Envoy
permalink: /http-routes/incoming-http-requests-pt-4
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
A **proxy** is a process that sits in-between the client and the server and
intercepts traffic before forwarding it on to the server. Proxies can add extra
functionality, like caching or SSL termination.

In this case, Envoy is a sidecar proxy (Envoy can be other types of proxies
too, but forget about that for now). The sidecar Envoy is only present when
Route Integrity is turned on (which is done by default).

Route Integrity is when the GoRouter sends all app traffic via TLS. As part of
the TLS handshake, the GoRouter validates the certificate's SAN against the ID
found in its route table to make sure it is connecting to the intended app
instance. This makes communication more secure and prevents stale routes in the
route table from causing misrouting, which is a large security concern. Read
more about how Route Integrity prevents misrouting
[here](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting).

The Envoy sidecar is the process that actually terminates the TLS traffic from
the GoRouter making Route Integrity possible. Then the Envoy proxies it
onto...can it be? finally?! YES! THE APP!!

Let's look at how the Envoy sidecar is configured to proxy traffic to the app.
## How

📝 **Look at envoy config**
1. Ssh onto AppA
1. Hit the Envoy `help` endpoint: `curl localhost:61003/help` These are all of
   the endpoints you can hit. Try `/clusters` what do you see?
1. Run `curl localhost:61003/config_dump`. This gives you all of the
   information about how the Envoy is configured.
1. Search for the `CONTAINER_ENVOY_PORT`, in the example it is 61001.
   This is where the DNAT rules forwarded the traffic to, as we saw in the last story.
   Find a listener called `listener-8080` that looks similar to the following:
   ```
   {
     "listeners": [
       {
         "address": {
           "socket_address": {
             "address": "0.0.0.0",
             "port_value": 61001                       <---- This listener is listening on port 61001.
           }                                                 That's the CONTAINER_ENVOY_PORT we know and love!
         },
         "filter_chains": [
           {
             "filters": [
               {
                 "config": {
                   "cluster": "0-service-cluster",     <---- This is the name of the cluster where Envoy will forward traffic
                 }                                           that is sent to the CONTAINER_ENVOY_PORT, let's call this CLUSTER-NAME
                   "stat_prefix": "0-stats"
                 },
                 "name": "envoy.tcp_proxy"
               }
             ],
             "tls_context": {
               "require_client_certificate": true      <---- This means Route Integrity is turned on
             }
           }
         ],
         "name": "listener-8080"                       <---- The name of the listener
       }
     ]
   }

   ```
1. In the same config_dump output, find the cluster, `CLUSTER-NAME`, that is referenced above. It should look something like this:
   ```
   {
     "clusters": [
       {
         "hosts": [
           {
             "socket_address": {           <---- This is the port that the app is listening on
               "address": "10.255.116.6",        inside of the container, should match OVERLAY_IP
               "port_value": 8080                and CONTAINER_APP_PORT
             }
           }
         ],
         "name": "0-service-cluster"       <---- This is the name of the cluster, CLUSTER-NAME
       }
     ]
   }
   ```

So the traffic gets sent to the `OVERLAY_IP:CONTAINER_ENVOY_PORT`, then the envoy forwards it on to `OVERLAY_IP:CONTAINER_APP_PORT`!

We made it! We finally made it to the end! Everything is set up and someone can use that route you made!

## Expected Result
Look at the Envoy's 8080 listener and related cluster and see how network traffic is sent to the app.

## Resources
* [Route Integrity/Misrouting Docs](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting)
* [What is Envoy?](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy)
