---
layout: single
title: Route Propagation Part 0 - Creating a Route Overview
permalink: /http-routes/route-propagation-pt-0
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---
## What
In the previous story you used the CF CLI to create and map routes. But what
happens under the hood to make all of this work? (hint: iptables is involved)
There are two main data flows for routes, (1) when an app dev pushes a new app
with a route and (2) when an internet user connects to an app using the route.

Let's focus on the first case. Here is what happens under the hood:

Each step marked with a ✨ will be explained in more detail in a story in this track.
**When an app dev pushes a new app with a route**
1. ✨ The app dev pushes an app with a route using the CF CLI.
1. Cloud Controller receives this request and sends this information to Diego.
1. Diego schedules the container create on a specific Diego Cell.
1. Garden creates the container for your app.
1. ✨ Diego deploys a sidecar envoy inside of the app container, which will
   proxy traffic to your app.
1. ✨ When the container is being set up, iptables rules are created on the
   Diego Cell to send traffic that is intended for the app to the sidecar
   proxy.
1. ✨ When the app is created, Diego sends the route information to the Route
   Emitter. The Route Emitter sends the route information to GoRouter via NATS.
1. ✨ The GoRouter keeps a mapping of routes -> ip:ports in a routes table,
   which is consulted when someone curls the route.

## How
The following stories will look at how many components (Cloud Controller, Diego
BBS, Route Emitter, Nats, GoRouter, DNAT Rules, Envoy) work together to make
routes work.

1. 🤔 Step through steps above and follow along on [the HTTP Routing section of
   this
   diagram](https://miro.com/app/board/o9J_kyWPVPM=/?moveToWidget=3074457346471397934).

## Expected Result
You can talk about route propagation at a high level.

## Logistics
In the next few stories, you are going to need to remember values from one
story to another, there will be a space provided at the bottom of each story
for your to record these values so you can store them.  It can be annoying to
scroll up and down in the story as you use the values, so it could be helpful
to store these values in a doc outside of tracker.

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
