---
layout: single
title: User Flow interal routes
permalink: /service-discovery/user-workflow-internal-routes
sidebar:
  title: "C2C Service Discovery"
  nav: sidebar-service-discovery
---

## Assumptions
- You have a CF deployed with silk release
- You have appA  talking to appB via c2c networking and policy (see previous
  story "User Flow Container to Container Networking")

## What

In the c2c module you were able to use c2c networking to get appA to talk
directly to appB using the overlay IP for appB (explanation of what overlay is
will come soon, I swear).

But IPs are ugly and URLs are pretty. Also, what happens when you restage an
app? Also, how do I load balance across many instances when I can only use IPs?

In order to fix these problems, we implemented Service Discovery, which is
apart of cf-networking-release. Service Discovery is also sometimes called
app-sd. Service Discovery is a fancy way of saying we handle the URL -> IP
translation for internal routes. Now appA can "discover" where the "service"
appB is, without having to know the IP.

Service Discovery is implemented using "internal routes" these routes will
*only* work from one CF app to another. They will not be accessible from
clients outside of CF.

## How

1. You should have appA talking to appB via an overlay IP using `watch  "curl
   CF_INSTANCE_INTERNAL_IP:8080"` inside of the appA container in one terminal.

1. In another terminal, run `cf restart appB` Predictably, the curl from appA
   to appB fails when appB is stopped. But it should come back when appB starts
   running again, right? ...Right? WHY IS IT STILL FAILING?

1. Recheck the overlay IP for appB `cf ssh appB -c "env | grep
   CF_INSTANCE_INTERNAL_IP"` What?! It moved! Use this new overlay IP to curl
   appB from appA and see that it still works. Lesson learned: IPs suck. Let's
   use Service Discovery instead.

1. When you create a route, any route, you have to supply a domain. To create
   an internal route, it must use an internal domain. We'll get into why in
   another story. For now, run `cf domains` and see that you should have a
   domain (or two) that is labeled `internal`.  Note the name of an internal
   domain that DOES NOT CONTAIN THE WORD "istio". You probably have the
   internal domain "apps.internal". Let's use that. (If you don't have a non
   istio internal domain, follow the resource at the bottom of this story to
   add a custom internal domain).

1. Using `cf map-route`, create and map a route for appB that uses the domain
   "apps.internal". May I suggest the route, appB.apps.internal?

1. In the terminal that is in the container for appA, run `watch  "curl
   appB.apps.internal:8080"`.

1. Restart appB.

1. Scale appB. Can you tell which instance you are hitting?

## Expected Result

Now that you are using internal routes to communicate via c2c, it shouldn't
matter that appB is restarted. As long as appB is running, appA should be able
to access it thanks to Service Discovery. When there are multiple instances of
appB, the internal route will automatically load balance between all of the
instances.

## Resources

* [internal domain docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains)
* [servide discovery docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md)
