---
layout: single
title: Instance Specific Routes
permalink: /route-registrar/instance-specific-routes
sidebar:
  title: "Route Registrar"
  nav: sidebar-route-registrar
---

## Assumptions
- You have two instances of my-http-server deployed. (See story `Life Without
  Route Registrar` for help if needed).
- There is an HTTP server on both instances of my-http-server. (See story `Life
  Without Route Registrar` for help if needed).

## What

In the `Life With Route Registrar` story you got the component route working
and load balancing between two instances of my-http-server. But what if you
want to be able to target a specific instance?

In this story you are going to create instance specific component routes.

## How

🤔 **Make and use instance specific routes**

1. Update the [prepend_instance_index property](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/route_registrar/spec#L95-L96) in your bosh manifest to turn on instance specific routing.

1. Redeploy

1. Use the new routes!

1. Prove that you are hitting only one instance and that you can choose which instance you are hitting.

🤔 **Check the gorouter routing table**

1. Look at the gorouter routes table and find your instance component routes.

## ❓ Question
How do these routes differ from the route you saw in the `Life With Route
Registrar` story?
