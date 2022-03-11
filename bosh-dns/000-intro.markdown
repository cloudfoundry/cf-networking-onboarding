---
layout: single
title: Bosh DNS Introduction
permalink: /bosh-dns/intro
sidebar:
  title: "Bosh DNS"
  nav: sidebar-bosh-dns
---

## Assumptions
- You have completed the Route Registrar stories

## What

As you learned in the [Route Registrar module](../route-registrar/intro), Route
Registrar is used to create routes that make CF components available to
off-platform users. You _could_ use these routes for talking from one CF
component to another, _but you should not_.  Using Route Registrar in place of
internal component routes (1) exposes CF components unnecessarily to the big
bad internet and (2) adds extra hops through the load balancer and GoRouter,
which cause latency.

So what do you do when one component wants to talk to another?

You use Bosh DNS!

In this Bosh DNS story series you are going to add a custom URL for an HTTP
server and trace the DNS requests for Bosh DNS routes and non-Bosh DNS routes.

## Vocab

**DNS** - stands for **D**omain **N**ame **S**ystem. DNS translates domain
names (urls) to IP addresses. You can do a DNS request in the terminal by using
the command line tool `dig`. With dig you can learn that neopets.com resolves
to 23.96.35.235.

**Bosh DNS** - provides native DNS Support for deployments (in our case, Cloud
Foundry). It gives the bosh release creator (Rahcel, that's (probably) you!) an
easy way to reference other VMs with load balancing. For example, you could set
up Bosh DNS such that cc.cf.internal resolves to the IP for the VM where the
Cloud Controller lives. This Bosh DNS server is typically deployed on every
Bosh deployed machine in a deployment (that is, every Cloud Foundry VM) using a
Bosh addon.

**Bosh DNS Routes** - routes available via Bosh DNS that route to Cloud Foundry
components.

**Alias** - a custom Bosh DNS Route. For example you can make the alias
meow.meow send traffic to the api VM. Fun!

**External URLs** - non-CF urls. Anything on the internet. Example:
neopets.com.

## Notes
- This set of stories uses the instance group `my-http-server` that was created
  in the Route Registrar set of stories. It is handy to use this VM with nearly
  nothing on it so that there is much less traffic on it.
- The networking program does not maintain Bosh DNS. However, DNS resolution is
  an important part of the networking traffic flow, so it is important to
  understand how it works.

## Resource
- [Bosh DNS Docs](https://bosh.io/docs/dns/)

