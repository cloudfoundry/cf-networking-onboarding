---
layout: single
title: Vocab
permalink: /route-registrar/vocab
sidebar:
  title: "Route Registrar"
  nav: sidebar-route-registrar
---

## What

Cloud Foundry components (like CAPI, Diego, and Policy Server) are deployed
into a private network. This means they can only be accessed within their
network. But what if you want to make an endpoint to a CF component available
outside of Cloud Foundry?

**Route Registrar** makes Cloud Foundry components available outside of Cloud
Foundry. These routes get registered with the GoRouter, just like app routes.
Or, if you like analogies:
```
Route Registrar:CF Components::cf map-route:CF Apps
```

In this Route Registrar series of stories you are going to create your own
instance group to run a HTTP server. Then you are going to use route registrar
to map a route to the server.

## Vocab 💬

**Off-Platform** - Anything not in a Cloud Foundry deployment. This term is
usually used when talking about where traffic originates. For example, traffic
from Wendy (the end user) is off-platform traffic. Your local machine is off
platform.

**On-Platform** - Anything that is within a Cloud Foundry deployment. This term
is usually used when talking about where traffic originates. For example, when
CAPI sends information to Diego, this is on-platform traffic.

**App Routes** - These are routes that resolve to a CF app.

**Component Routes** - These are routes that resolve to a bosh component. For
example, uaa.beanie.c2c.cf-app.com is a component route that resolves to UAA.

## Note
This set of stories uses the instance group `my-http-server`, which you will
create. It is handy to use this VM with nearly nothing on it so that there is
much less traffic coming/going to it. However, all of this work _could_ be
done on any VM.
