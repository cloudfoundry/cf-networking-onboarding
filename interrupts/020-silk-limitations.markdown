---
layout: single
title: Silk Limitations
permalink: /interrupts/silk-limitations
sidebar:
  title: "Interrupts"
  nav: sidebar-interrupts
---
## Interrupt

_Hello,_

_I've started reading about the container to container networking and I've noticed that the "batteries included pack"
the Silk CNI plugin and the VXLAN Policy Agent use VXLAN._

_In the
[description](https://docs.cloudfoundry.org/concepts/understand-cf-networking.html
) of the VXLAN Policy agent it is written that the outbound traffic is tagged
with the unique identifier of the source app using the VXLAN Group-Based Policy
(GBP) header. Checking the [VXLAN GBP header
documentation](https://tools.ietf.org/html/draft-smith-vxlan-group-policy-05)
I've read that the "Group Policy ID" field is 16bit long which means 2^16= 64k
unique source apps can be tagged. The VXLAN GBP header also has the VXLAN
Network Identifier (VNI) field. I've also read in the [VXLAN
RFC](https://tools.ietf.org/html/rfc7348) in section 4. VXLAN that "Only VMs
within the same VXLAN segment can communicate with each other. Each VXLAN
segment is identified through a 24-bit segment ID, termed the 'VXLAN Network
Identifier (VNI)'."_

_Three things are not clear to me:_
1. _Does this setup use only one VXLAN? The docs talk about one overlay network
   to which all the Diego cells are connected._
2. _Does the limitation that the VXLAN GBP header brings, means that only 64k
   unique source apps can be used for container to container networking per cf
   deployment?_
3. _Can we somehow overcome the limitation brought by the "Group Policy ID"
   field's size of 16bit in the VXLAN GBP by using the current setup and adding
   another VXLAN overlay or we have to re-implement all of the swappable
   components?_

## Restate the question
Sometimes interrupts misuse networking or CF terms. It's your job to get to
the bottom of what they are asking for. In 2 sentences or less, restate their
question.

## Answer
Now that we're on the same page about what is being asked, what is the answer?

