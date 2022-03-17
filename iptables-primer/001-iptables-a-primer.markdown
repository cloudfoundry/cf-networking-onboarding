---
layout: single
title: Iptables - A Primer
permalink: /iptables/primer
sidebar:
  title: "Iptables"
  nav: sidebar-iptables
---

## What

Think of iptables as a firewall implementation. (Iptables can do more than
firewall things, but we have to start somewhere.)

ASGs and c2c network policies are implemented using iptables rules.

All network traffic on linux machines is evaluated against the iptables rules
written on the machine. To organize these rules, there are tables and chains.
Tables have chains and chains have rules. Users (that's you!) can create custom
chains and rules for their own needs. Traffic hits different tables and chains
depending on if the traffic is ingress, egress, or East/West.

Let's look at the anatomy of an iptables rule.

![iptables rule anatomy](https://deliveryimages.acm.org/10.1145/2070000/2062737/10822f2.png)

The example above can be translated as:
- (**table/chain**) For any traffic that is evaluated against the filter table
  on the INPUT chain...
- (**match**)...if that traffic is using the tcp protocol...
- (**match**)...and if that traffic is sending traffic to port 80...
- (**jump rule/target**) ... then ACCEPT that traffic.

Once a packet hits ACCEPT, then it stops evaluating in that table and chain. It
would also stop if it hit a DROP or REJECT target.

In this story you are going to skim/read a couple great resources on iptables
rules.

## How
1. Give the [iptables man page](http://ipset.netfilter.org/iptables.man.html) a
   skim. At least the description, targets, and tables sections. It leaves out
   a lot of commonly used features that are covered in [the iptables-extensions
   man page](http://ipset.netfilter.org/iptables-extensions.man.html).

1. Check out [Julia Evan's Blog on
   iptables](https://jvns.ca/blog/2017/06/07/iptables-basics/).

## Optional Videos
Here are some thorough videos on iptables rules. If you are the type of person
who likes getting lots of information up front, watch them now.  If you are the
type of person who likes to experiment first, skip these videos and watch them
after the next story. (Don't worry I'll remind you.)

1. 🎬 Watch this video ["iptables: Packet Processing" by Dr. Murphy's
   Lectures](https://www.youtube.com/watch?v=yE82upHCxfU) _length 14:22_
1. 🎬 Watch this video ["iptables: Tables and Chains" by Dr. Murphy's
   Lectures](https://www.youtube.com/watch?v=jgH976ymdoQ) _length 10:34_

## Resources
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)
* [Julia Evans iptables basics](https://jvns.ca/blog/2017/06/07/iptables-basics/)
* ["iptables: Packet Processing" by Dr. Murphy's Lectures](https://www.youtube.com/watch?v=yE82upHCxfU)
* ["iptables: Tables and Chains" by Dr. Murphy's Lectures](https://www.youtube.com/watch?v=jgH976ymdoQ)
