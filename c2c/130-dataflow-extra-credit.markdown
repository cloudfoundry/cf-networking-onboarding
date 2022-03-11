---
layout: single
title: Extra Credit - How are only some packets decapsulated?
permalink: /c2c/dataflow-extra-credit-decapsulation
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA on Diego Cell 1
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appB on Diego Cell 2

## What

So far the overlay packet has been encapsulated into an underlay packet and
then it is sent to a second Diego Cell. Once the underlay packet gets to the
Diego Cell it is gets decapsulated by the VTEP. But how does "it" know to send
these specific underlay packets to the silk-vtep interface to be decapsulated
and not other packets?

Let's figure it out by inspecting the packets with tcpdump! Tcpdump is a CLI
tool that allows you to inspect all of the traffic flowing through your
container.

### How

🤔 **Send traffic via the overlay from appA to appB**
1. In terminal 1, use `watch` to continuously curl appB from appA using appB's overlay IP and app port.

📝 **Look at the underlay traffic**
1. In terminal 2, ssh onto Diego Cell 2, where appB is running.
1. The underlay packet is from Diego Cell 1 to Diego Cell 2, so use tcpdump to look at all traffic from Diego Cell 1.
```
tcpdump -n src DIEGO_CELL_1_IP -v
```

## ❓ Questions
1. What do you notice about all of the traffic? What do they have in common? Based on this information how do you think only this traffic is being decapsulated?
1. What protocol is this traffic using? Is that surprising to you?

## Expected outcome
You should see that all traffic to be decapsulated is sent to the same port.
This is how some traffic is decapsulated by the VTEP but not others.

You should also notice that all of the traffic is sent via UDP. WHAT? Read
[here](https://blog.ipspace.net/2012/01/vxlan-runs-over-udp-does-it-matter.html)
for more details on _that_.
