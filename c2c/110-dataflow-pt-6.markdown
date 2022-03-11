---
layout: single
title: Dataflow Part 6 - VLAN VXLAN VTEP V-What?
permalink: /c2c/dataflow-pt-6
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## What
In the explanation of the data flow for container to container (c2c)
networking, step 4 says: *Because the packet is an overlay packet, it is sent
to the silk-vtep interface on the Diego Cell. This interface is a VXLAN
interface.*

What does that even mean!??!!?  Let's define some terms.

**LAN** - Local Area Network - a small network that is usually for connecting
personal computers. Each computer in a LAN is able to access data and devices
anywhere on the LAN. This means that many users can share devices, such as
laser printers, as well as data. ([paraphrased from
here](https://www.webopedia.com/TERM/L/local_area_network_LAN.html)). Often,
used by gamers in [LAN parties](https://en.wikipedia.org/wiki/LAN_party) and by
offices.

**VLAN** - Virtual Local Area Network - a network that *appears* to be on the
same LAN, even though the machines are physically separated. For example, the
Pivotal LA and SF (and probably other west coast offices) were on the same VLAN.
This allowed Pivots in these offices to SSH onto other machines in the VLAN. But
those outside of the VLAN cannot SSH onto machines inside of the VLAN.

**VXLAN** - Virtual eXtensible Local Area Network - VXLAN was developed to
provide the same network services as VLAN does, but with greater extensibility
and flexibility. Container to container (c2c) networking uses VXLAN to create
the overlay network.

**VTEP**- VXLAN EndPoints - are VXLAN tunnels that encapsulate outgoing overlay
traffic into underlay packets and decapsulate incoming underlay packets into
overlay packets. The VTEP that c2c uses is called silk-vtep.

## Resources
* [LAN](https://www.webopedia.com/TERM/L/local_area_network_LAN.html)
* [VLAN wiki](https://en.wikipedia.org/wiki/Virtual_LAN)
* [VLAN vs VXLAN](http://www.fiber-optic-transceiver-module.com/vxlan-vs-vlan-which-is-best-fit-for-cloud.html)
