---
layout: single
title: Get Started
permalink: /get-started/directional-networking-terms
sidebar:
  title: "Directional Networking Terms"
  nav: sidebar-getting-started
---

## What

Network traffic can go in many directions, and there are many jargony ways to refer to them. Let's define them.

**Ingress Traffic** is traffic that originates outside the local network that is transmitted to somewhere inside the local network. Remember ingress traffic is coming *in*.

**Egress Traffic** is traffic that originates inside the local network that is transmitted somewhere outside of the local network. Remember egress traffic is *e*xiting the local network.

```
 INGRESS TRAFFIC EXAMPLE          EGRESS TRAFFIC EXAMPLE

         +-----+                          +-----+
         |     |                          |     |
     +---+     +------+               +---+     +------+
   +-+                |             +-+                |
 +-+   The Internet   ++          +-+   The Internet   ++
 |                   +-+          |                   +-+
 +-------------------+            +-------------------+

               V                                ^
               |                                ^
               |                                |
+----------------------+         +----------------------+
| Container    |       |         | Container    |       |
|              |       |         |              |       |
| +-------+    |       |         | +-------+    |       |
| |       |    |       |         | |       |    |       |
| | MyApp | <--+       |         | | MyApp | >--+       |
| |       |            |         | |       |            |
| +-------+            |         | +-------+            |
+----------------------+         +----------------------+

```

**North/South** traffic is any communication between two different networks.  Both Ingress and Egress are examples of North/South traffic.

**East/West** traffic is any communication within one network.

```
                  EAST/WEST TRAFFIC EXAMPLE

+-------------------------------------------------------------+
| My Local Network                                            |
|                                                             |
|                                                             |
|                                                             |
| +----------------------+         +----------------------+   |
| | Container1           |         | Container2           |   |
| |                      |         |                      |   |
| | +-------+            |         | +-------+            |   |
| | |       |            |         | |       |            |   |
| | | MyApp | +----------------->  | | MyApp |            |   |
| | |   1   |            |         | |   2   |            |   |
| | +-------+            |         | +-------+            |   |
| +----------------------+         +----------------------+   |
|                                                             |
+-------------------------------------------------------------+

```

## ❓Questions

How would you use ingress, egress, north/south, and east/west to describe the following situations:
- You visit neopets in your browser.
- Your pair `ssh`-es onto your computer.
- You set up a local netcat server and send traffic to it from your terminal.

