---
layout: single
title: "How Many Data Flow Options are There?"
permalink: /http-routes/data-flow-options
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---
## What

When we talk about the data flow from client to Cloud Foundry app, we often
draw it like this:
```
+----+    +----------+         +-------+     +-----+
| LB +--->+ Gorouter +-------->+ Envoy +---->+ App |
+----+    +----------+         +-------+     +-----+
```

But that diagram is very general. Often that is okay because the details don't
always matter. But sometimes the details _do_ matter (like with HTTP/2).

In this story we will look at some more specific data flow diagrams for Cloud
Foundry.

## How

First we need to understand what L4 (TCP) LBs and L7 (HTTPS) LBs _are_.
1. 📚 Read [this medium article about "TCP vs HTTP(S) Load
   Balancing."](https://medium.com/martinomburajr/distributed-computing-tcp-vs-http-s-load-balancing-7b3e9efc6167)
1. 📚 Read [these cloudfoundry docs on TLS Termination Options for HTTP
   Routing](https://docs.pivotal.io/application-service/2-10/adminguide/securing-traffic.html).

## ❓ Questions
Look at the following diagrams and think about the following questions for
each:
* What connections (the arrows between boxes) are encrypted?
* Which are not?

**With a L4 LB in front**
```
+-------+    +----------+         +-------+     +-----+
| L4 LB +--->+ Gorouter +-------->+ Envoy +---->+ App |
+-------+    +----------+         +-------+     +-----+
```

**With a L7 LB in front**
```
+-------+    +----------+         +-------+     +-----+
| L7 LB +--->+ Gorouter +-------->+ Envoy +---->+ App |
+-------+    +----------+         +-------+     +-----+
```

**With an HAProxy in front**
```
+----------+    +----------+         +-------+     +-----+
| HA Proxy +--->+ Gorouter +-------->+ Envoy +---->+ App |
+----------+    +----------+         +-------+     +-----+
```

**With an L4 LB and an HAProxy in front**
```
+-------+    +----------+    +----------+         +-------+     +-----+
| L4 LB +--->| HA Proxy +--->+ Gorouter +-------->+ Envoy +---->+ App |
+-------+    +----------+    +----------+         +-------+     +-----+
```

**With an L7 LB and an HAProxy in front**
```
+-------+    +----------+    +----------+         +-------+     +-----+
| L7 LB +--->| HA Proxy +--->+ Gorouter +-------->+ Envoy +---->+ App |
+-------+    +----------+    +----------+         +-------+     +-----+
```
