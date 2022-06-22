---
layout: single
title: Incoming HTTP Requests Part 0 - HTTP Traffic Overview
permalink: /http-routes/incoming-http-requests-pt-0
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## What
In the previous stories you learned what happens when an app dev pushes a new app with a route. In this story you will learn an at a high level what happens when an end user connects to an app using the route.

Each step marked with a ✨ will be explained in more detail in a story in this track.

**When an internet user sends traffic to your app**
1. The user visits your route in the browser or curls it via the command line.
1. The traffic first hits a load balancer in front of the CF Foundation.
1. The load balancer sends it to one of the GoRouters.
1. ✨ The GoRouter consults the route table and sends it to the listed IP and
   port. If Route Integrity is enabled, it sends this traffic via TLS. (This
   was explored in the previous story!)
1. ✨ The traffic makes its way to the correct Diego Cell, where it hits
   iptables DNAT rules that reroutes the traffic to the sidecar envoy for the
   app.
1. ✨ The Envoy terminates the TLS from the GoRouter and then sends the traffic
   on to the app.

## How
The following stories will look at how many components (Cloud Controller, Diego
BBS, Route Emitter, Nats, GoRouter, DNAT Rules, Envoy) work together to make
routes work.

1. 🤔 Step through steps above and follow along on [the HTTP Routing section of
   this diagram](https://realtimeboard.com/app/board/o9J_kyWPVPM=/)

## Expected Result
You can talk about HTTP network traffic flow with fellow CF engineers.

## Logistics
In the next few stories, you are going to need to remember values from one
story to another, there will be a space provided at the bottom of each story
for your to record these values so you can store them.

## Resources for the entire route propagation track
**Cloud Controller**
* [Cloud Controller V2 API docs](https://apidocs.cloudfoundry.org)
* [Cloud Controller V3 API docs](http://v3-apidocs.cloudfoundry.org)

**Diego**
* [cfdot docs](https://github.com/cloudfoundry/cfdot)
* [diego design notes](https://github.com/cloudfoundry/diego-design-notes#what-are-all-these-repos-and-what-do-they-do)
* [diego bbs API docs](https://github.com/cloudfoundry/bbs/tree/master/doc)

**NATs**
* [NATS message bus repo](https://github.com/nats-io/gnatsd)
* [NATS ruby gem repo](https://github.com/nats-io/ruby-nats)

**GoRouter**
* [GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)
* [Detailed Diagram of several Route Related User Flows](https://realtimeboard.com/app/board/o9J_kyWPVPM=/)

**Iptables**
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)

**Route Integrity**
* [Route Integrity/Misrouting Docs](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting)

**Envoy**
* [What is Envoy?](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy)
