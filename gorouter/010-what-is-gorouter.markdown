---
layout: single
title: What is Gorouter?
permalink: /gorouter/what-is-gorouter
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What

So what is Gorouter?

[Gorouter](https://github.com/cloudfoundry/gorouter) is a bosh job in [routing release](https://github.com/cloudfoundry/routin

* subscribing to route registration messages from nats and keeping an up-to-date routing table
* acting as a reverse proxy and routing to backends 👈👈👈
* serving a healthcheck endpoint

In this section we are going to focus on the second bullet: how gorouter acts as a reverse proxy and routes to backends.

## How
1. 📚Skim through [bosh configuration options for gorouter](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/

## Links
* [Gorouter](https://github.com/cloudfoundry/gorouter)
* [routing release](https://github.com/cloudfoundry/routing-release)
* [bosh configuration options for gorouter](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/gorouter/spec)
