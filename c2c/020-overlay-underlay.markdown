---
layout: single
title: Overlay vs Underlay
permalink: /c2c/overlay-vs-underlay
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## What

The **underlay network** is the network that you're given to work with, on top
of which you build a virtual **overlay network**. Routing for overlay networks
is done at the software layer.

Overlay networks are used to create layers of abstraction that can be used to
run multiple separate layers on top of the physical network — the bottom
underlay network. These are general definitions that are not specific to Cloud
Foundry.

In CF, the software that creates the overlay network uses a combination of
tools including VXLAN and iptables rules. Much more about this in the stories
that follow!

** Your ** underlay network is often someone else's overlay network, that
engineer just works on a lower abstraction layer and might work on an IaaS
rather than a PaaS, for example. It's all relative! 🤯

Routing to an app using the Diego Cell IP and port is done on what we will
refer to as the **underlay network**. Container to container networking (c2c)
is done on what we will refer to as the **overlay network**.

## Resources
- [Difference between overlay and
underlay](https://ipwithease.com/difference-between-underlay-network-and-overlay-network/)

