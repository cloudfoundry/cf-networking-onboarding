---
layout: single
title: Digging Routes
permalink: /tcp-routes/dig
sidebar:
  title: "TCP Routes"
  nav: sidebar-tcp-routes
---

## Assumptions
- You have a CF deployed
- You have a TCP server deployed named tcp-app
- You have a TCP route mapped to tcp-app called TCP_ROUTE
- You have a 2
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed, that are named appA and appB
- You have a HTTP route mapped to appA called APP_A_ROUTE and one mapped to
  appB called APP_B_ROUTE

## What
`dig` is one of many network utilities that can be very helpful for debugging.
Dig does a DNS lookup for a URL.

Let's play around with dig, TCP Routes, and HTTP Routes.

## How

📝 **Dig with a good URL**
1. Run `dig neopets.com` If you see `ANSWER 1` (which you should) that means
   that it was able to resolve the route. `ANSWER 0` means it was unable to
   resolve the route.  In the `ANSWER SECTION` you should see an IP. (Likely,
   23.96.35.235)
1. Put the IP that neopets.com resolved to in a browser. Is it neopets?

📝 **Dig with a bogus URL**
1. Resolve a bogus URL so you get ANSWER 0.

📝 **Dig with CF HTTP routes**
1. Use dig to resolve APP_A_ROUTE. Let's call this APP_A_IP
1. Curl that IP.
* ❓ What happens?

1. Use dig to resolve APP_B_ROUTE. Let's call this APP_B_IP
1. Curl that IP.
* ❓ What happens?
* ❓Why are APP_A_IP and APP_B_IP the same?
* ❓Why doesn't the IP resolve to either of the apps?

📝 **Dig with CF TCP routes**
1. Use dig to resolve the TCP Route. Let's call this TCP_APP_IP
1. Curl that IP. What happens?
* ❓Is TCP_APP_IP the same as APP_A_IP?
* ❓Why or why not?
* ❓How does traffic get to the apps if the IPs don't work?

🤔 **Sleuthing in your IAAS**
1. In your IAAS GUI, find what infrastructure that these IPs map to.

## Expected Results
All CF HTTP Routes resolve to the same IP. All CF TCP Routes resolve to the
same IP. You should find the load balancers that map to these IPs.

## Extra Credit
❓Why does `dig neopets.com` work, but `dig http://neopets.com` does not work?

## Resource
* [understanding the dig command](https://mediatemple.net/community/products/dv/204644130/understanding-the-dig-command)
