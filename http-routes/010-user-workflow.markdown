---
layout: single
title: HTTP routes user workflow
permalink: /http-routes/user-workflow
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a CF deployed
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed, named appA and appB

## What

**Routes** are the URLs that can be used to access CF apps.

**Route Mappings** are the join table between routes (URLs) and the apps they
send traffic to. Apps can have many routes.  And routes can send traffic to
many apps. So Route Mappings is a many-to-many mapping.

## How

📝 **Create a route that maps to two apps**
0. By default `cf push` creates a route. Look at all of the routes for all of your apps.
{% include codeHeader.html %}
   ```bash
   cf apps
   ```
0. Use curl to hit appA.
   ```bash
   curl APP_A_URL
   ```
   It should respond with something like
   ```json
   {"ListenAddresses":["127.0.0.1","10.255.116.44"],"Port":8080}
   ```
   We'll get into the listen addresses later, but for now the most important thing to know is that the 10.255.X.X address is the overlay IP address. This IP is unique per app instance.
0. Create your own route (`cf map-route --help`) and map it to both appA **AND** appB.
0. Curl your new route over and over again `watch "curl -sS MY-NEW-ROUTE`".

## Expected Result
You have a route that maps to both appA and appB. See that the overlay IP
changes, showing that you are routed evenly(ish) between all the apps mapped to
the routes.
