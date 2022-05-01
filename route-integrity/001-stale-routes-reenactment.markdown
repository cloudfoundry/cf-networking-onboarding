---
layout: single
title: Stale Routes Reenactment
permalink: /route-integrity/stale-routes-reenactment
sidebar:
  title: "Route Integrity"
  nav: sidebar-route-integrity
---

## What
Recap from [HTTP Routes module](../http-routes/user-workflow): the GoRouter
redirects traffic for a particular CF HTTP route to a Diego Cell IP and port.
However, when one app is deleted, another app may later use the same Diego Cell
IP and port.

In the happy path case, the following steps should happen in the following order:

- appA is pushed and is at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appA to 10.0.1.2:12345
- appA is deleted
- the Route Emitter sends an unregister route message to the GoRouter
- 👍 **the GoRouter does NOT forward traffic for appA to 10.0.1.2:12345**
- appB is pushed and is now at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appB to 10.0.1.2:12345

That's a lot of components working together to make sure routes are sent to the
correct place. Now imagine this situation where the Route Emitter is not
sending out messages fast enough.

- appA is pushed and is at 10.0.1.2:12345
- the Route Emitter sends a register route message to the GoRouter
- the GoRouter forwards traffic for appA to 10.0.1.2:12345
- appA is deleted
- ~~the Route Emitter sends an unregister route message to the GoRouter~~
- ~~the GoRouter does NOT forward traffic for appA to 10.0.1.2:12345~~
- appB is pushed and is now at 10.0.1.2:12345
- ~~the Route Emitter sends a register route message to the GoRouter~~
- **the GoRouter continues to forward traffic for appA to 10.0.1.2:12345, where appA used to be, where appB currently is**
- 😱 the user tried to access appA, but they were routed to appB instead

Yikes! This is an example of misrouting due to "stale" routes.

Let's set up a situation with stale routes manually and see how dangerous it
is. Then in the next story we will look at how Route Integrity fixes this
issue.

The Networking team has had to fix several bugs related to stale routes. To
replicate the bugs and test our fixes, we came up with the following method to
simulate stale routes. This is only one of several possible methods to force
routes stale.

## How

🤔 **cause stale routes**
1. Route Integrity is turned on by default with CF Deployment. Let's turn it
   off and see what happened before Route Integrity. But first, let's save the
   current manifest for the next story.
   ```bash
   bosh manifest > /tmp/env-with-route-integrity.yml
   ```

1. Redeploy your CF with Route Integrity turned off, using [this opsfile](../opsfiles/disable-routing-integrity.yml).

1. Push one instance of [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) and call it appA.

1. Ensure that appA has an HTTP route.

1. Push one instance of [dora](https://github.com/cloudfoundry/cf-acceptance-tests/tree/develop/assets/dora) and name it appB.

1. Ensure that there is an HTTP Route mapped to appB.

1. Make sure an instance of appA and appB are on the same Diego Cell. This can be achieved through scaling appB if it doesn't happen at first.

1. Curl appA. See that you get the expected result for appA.

1. In previous http-route stories, you learned that the route table on the GoRouter maps routes to particular Diego Cell IPs and ports. Then the GoRouter sends traffic to the Diego Cell IP and port. Then the traffic is rerouted to an app's overlayIP and app port using iptables rules on the nat table. Change the iptables rules so that the Diego Cell port that maps to appA now, incorrectly, redirects traffic to appB. This will imitate route staleness. Go back to the http-route DNAT stories if you need help doing this.

1. Curl the route for appA.

### Expected Result
Even though you curled the route for appA, because the routes for appA were "stale" you got back the information for appB.

## ❓Questions

* Why is this dangerous?
* What is the implications for this on a multi-tenant production environment?
* What could cause routes to become stale in the real world? What could
  exacerbate this problem?
* Before Route Integrity, what could app devs do to prevent stale routes from
  accidentally releasing their data?
