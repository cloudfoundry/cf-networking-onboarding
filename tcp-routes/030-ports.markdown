---
layout: single
title: Route Ports! Backend Ports! App Ports! Oh my!
permalink: /tcp-routes/ports
sidebar:
  title: "TCP Routes"
  nav: sidebar-tcp-routes
---

## What
Earlier in this track of TCP work you created a TCP route. This TCP route
needed a port.  This is called a route port. There are 3 different types of
ports in the Cloud Foundry ecosystem. They can be extremely confusing, so it is
nice if there is clear vocabulary so you can easily describe which port you are
talking about.

The three types of ports are:

**Route Port** - this is the port that a client makes a connection to. For HTTP
routes, this is always 80 (for http) or 443 (for https). For TCP routes, this
port is configured and unique.

**Backend Port** - this is the high number port on the Diego Cell that the
GoRouter or TCP Router proxies traffic to. This port is unique per app on each
Diego Cell.

**App Port** - this is the port where the app is listening inside of the
container. In CF this defaults to 8080.

Let's look at a diagram of these ports.

![ports for TCP traffic](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/tcp-trafficflow-ports.png)

![ports for HTTP traffic](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/http-traffic-flow-ports.png)

## ❓ Questions
* Look at the help text for mapping a route (`cf map-route --help`). What are the different flags allowed for HTTP routes vs TCP routes?
* How do these different flags align with what you learned in this story?

## Resources
* [configuring app ports in
  CF](https://docs.cloudfoundry.org/devguide/custom-ports.html)
