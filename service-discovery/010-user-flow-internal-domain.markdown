---
layout: single
title: Create an Internal Domain
permalink: /service-discovery/user-workflow-internal-domains
sidebar:
  title: "C2C Service Discovery"
  nav: sidebar-service-discovery
---

## Assumptions
- You have a CF deployed with silk release
- You have appA  talking to appB via c2c networking and policy

## What

In the previous story you created an internal route for appA to talk to appB
using the domain "apps.internal", but what if we wanted to create our own
internal domain (may I suggest meow.meow.meow)?

## How

1. Start off where you left off from the previous story "User Flow Container to
   Container Networking"). You should have appA talking to appB via an overlay
   IP using `watch  "curl -sS appB.apps.internal:8080"` inside of the appA
   container in one terminal.

1. In another terminal, create a new internal domain `cf create-shared-domain
   meow.meow.meow --internal` Check that it worked
   ```bash
   $ cf domains
   Getting domains in org o as admin...
   name                         status   type   details
   meow.meow.meow               shared          internal
   ```

1. Using `cf map-route`, create and map a route for appB that uses our new
   internal domain "meow.meow.meow". May I suggest the route,
   appB.meow.meow.meow?

1. In the terminal that is in the container for appA, use this new internal
   route to curl appB `watch "curl -sS appB.meow.meow.meow:8080"` What? "Could
   not resolve host"???? Why doesn't it work like our other internal route?
   Unlike other domains, internal domains require one more step in order for
   them to work.

1. Download the manifest for your CF. Look at the property `internal_domains`
   on the `bosh_dns_adapter` job. It probably looks like this:
   ```yaml
   internal_domains:
   - apps.internal.
   ```

So unfortunately, there is a deploy time dependency for internal domains. I
know, this makes me sad too. Let's dig in why this is.

⚠️ Warning, Architecture Description Ahead (please follow along with the diagram
below):

1. When an app makes ANY network request that requires DNS lookup (that
is, any request to a URL, not an IP), the DNS lookup first hits Bosh DNS.
2. Bosh DNS then checks to see if the domain of the URL being requested matches
any of the internal domains that it knows about from the Bosh DNS Adapter
`internal_domains` property. There is no reason why this couldn't be dynamic,
Bosh DNS Adapter *could* make an API call to CAPI to figure out what the
up-to-date internal domains are. But it doesn't, so it's not dynamic.
3. Then the Bosh DNS Adapter calls out to the Service Discovery Controller, which
keeps track of what internal route maps to what overlay IP, very similar to the
routes table in the GoRouter.

![href](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/architecture-diagram.png?raw=true)

🤔 **Make your internal domain work**
1. Add our internal domain meow.meow.meow to the bosh manifest.

2. Redeploy your environment.

3. In the terminal that is in the container for appA, use this new internal
   route to curl appB
   ```bash
   watch "curl -sS appB.meow.meow.meow:8080"
   ```

## Expected Result

appA should be able to successfully reach appB using the internal route with
our brand new internal domain.

## Resource

* [configuring internal domain
  docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains)
