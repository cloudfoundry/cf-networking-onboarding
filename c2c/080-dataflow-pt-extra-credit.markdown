---
layout: single
title: Dataflow Extra Credit - Marks
permalink: /c2c/dataflow-extra-credit-marks
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## What
You may have noticed a discrepancy in the diagram and the steps about _where_
the mark is located. In the diagram it shows the mark on the _underlay_ packet.
But the steps say that the _overlay_ packet is marked. Also the iptables rules
seem to add the mark to the _overlay_ packet. So which is it? Well, both, kind
of.

**The mark for the overlay packet is _not_ part of the packet itself.** This
mark is just a bit of metadata about the packet that the kernel keeps track of.
This mark exists only as long as it's handled by the Linux kernel. So if a
packet is marked before it is sent to a different host, the host will not
receive the mark information.

**When VTEP on the source host encapsulates the overlay packet, the mark gets
recorded as a header in the underlay packet.** When the VTEP on the destination
host decapsulates the underlay packet, it sets the mark on the kernel for the
overlay packet.
