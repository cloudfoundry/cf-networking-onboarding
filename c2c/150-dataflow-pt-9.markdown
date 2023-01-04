---
layout: single
title: Dataflow Part 9 - Same Cell vs Different Cell
permalink: /c2c/dataflow-pt-9
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and named appA and appB (the fewer apps you have deployed the
  better)

## Review
This track of stories is going to go through the steps (listed below) that were
covered in the dataflow overview.  The steps and diagram will be at the top of
each story in case you need to orient yourself. Higher quality diagram
[here](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/c2c-data-plane.png).

![c2c traffic flow](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on
   Diego Cell 2). This packet is called the overlay packet (aka the c2c
   packet).
1. The packet exits the app container through the veth interface.
1. The overlay packet is marked with a ...mark... that is unique to the source
   app.
1. Because the packet is an overlay packet, it is sent to the silk-vtep
   interface on the Diego Cell. This interface is a VXLAN interface.
1. The overlay packet is encapsulated inside of an underlay packet. This
   underlay packet is addressed to the underlay IP of the Diego Cell where the
   destination app is located (appA in this case).
1. The underlay packet exits the cell.
1. The packet then travels over the physical underlay network to the correct
   Diego Cell.
1. The packet arrives to the correct Diego Cell
1. The underlay packet is decapsulated. Now it's just the overlay packet again.
1.  **Iptables rules check that appA is allowed to talk to appB based on the
    mark on the overlay packet.   <------- CURRENT STORY**
1. If traffic is allowed, the overlay network directs the traffic to the
   correct place.

## What

The diagram above shows the container to container networking traffic flow
between two apps on different Diego Cells. But what about when the apps are on
the same Diego Cell?

In this story, you are going to do some investigation to figure out what the
container to container traffic flow is for apps on the same Diego Cell.

## ❓ Questions

* What chain (INPUT, OUTPUT, or FORWARD) are the container networking policy
  iptables rules appended to?
* Does overlay traffic between apps on different Diego Cells hit this chain?
* Does overlay traffic between apps on the same Diego Cell hit this chain?
* Does overlay traffic between apps on the same Diego Cell get encapsulated?
* What steps listed in the review section apply for apps on the same Diego
  Cell?

## Expected Result

At the end of this story, you should know the traffic flow between two apps on
the same Diego Cell.

