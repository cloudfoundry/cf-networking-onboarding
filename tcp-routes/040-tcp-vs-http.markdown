---
layout: single
title: TCP vs HTTP Routes
permalink: /tcp-routes/tcp-vs-http
sidebar:
  title: "TCP Routes"
  nav: sidebar-tcp-routes
---
## What

Earlier in this track, you learned that all CF HTTP Routes resolve to the IP of
the HTTP load balancer and that all TCP Routes resolve to the IP of the TCP
Load balancer.

But if _all_ routes resolve to the IP of a load balancer, then how does traffic
actually get sent to an application????

Well, it depends if the route is HTTP or TCP.

One major difference between the HTTP protocol and the TCP protocol, is that
HTTP packets contain more headers. These headers can carry all sorts of
information, including what URL the request is being made to. (The client can
even set arbitrary headers themselves!) GoRouter is able to look at this
headers, see the requested URL, and route appropriately based on the
information in the Routes Table. Because of this, all HTTP routes can have
identical route ports (80 or 443).

![gorouter routing](https://storage.googleapis.com/cf-networking-onboarding-images/gorouter-traffic-routing.png)

TCP is barebones. TCP packets have very limited headers that only include the
source and destination ports of the packets. Because of this, the TCP Router
has _no_ knowledge of the URL. So the TCP Router needs something else to
differentiate between apps. It can't be the destination IP because all TCP
routes have the same destination IP. The only thing left, is the destination
port. Because of this, all TCP routes must have unique route ports.

![tcp router routing](https://storage.googleapis.com/cf-networking-onboarding-images/tcp-traffic-routing.png)

## How
📝 **Inspect HTTP headers**
1. Curl the networking api and look at the request headers
   ```bash
   cf curl /networking/v1/external/policies -v
   ```
   You should get a response that looks like this:
   ```
   REQUEST: [2019-04-24T11:01:49-07:00]
   GET /networking/v1/external/policies <---- this is the header that contains the URL path
   HTTP/1.1
   Host: api.beanie.c2c.cf-app.com      <---- this is the header that contains the URL base
   Accept: application/json
   Authorization: [PRIVATE DATA HIDDEN]
   Content-Type: application/json
   User-Agent: go-cli 6.43.0+815ea2f3d.2019-02-20 / darwin
   ```

## ❓ Question
What would happen if two TCP routes had the same route port?

## Resources
* [tcp header format](https://www.freesoft.org/CIE/Course/Section4/8.htm)
* [http headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
