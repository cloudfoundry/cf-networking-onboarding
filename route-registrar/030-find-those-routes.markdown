---
layout: single
title: Find Those Routes!
permalink: /route-registrar/find-those-routes
sidebar:
  title: "Route Registrar"
  nav: sidebar-route-registrar
---

## Assumptions
- You have my-http-server deployed with route registrar routes setup from the
  previous story

## What

In this story you are going to follow how route registrar registers routes (say
that 5 times fast).

In the HTTP Routes section you learned how the Route Emitter on the Diego Cell
repeatedly sends route registration messages to NATS. Then GoRouter subscribes
to those NATS messages and then populates its route table. The same thing
happens with component routes. But instead of the Route Emitter emitting routes
(teehehe), it's Route Registrar that is emitting routes repeatedly.

## How

🤔 **Look at NATS**

0. Subscribe to the the NATS messages for your component route from the
my-http-server VM.
* You can find the NATS username, password, and host on the my-http-server VM
  at `/var/vcap/jobs/route_registrar/config/registrar_settings.json`
* See the story [Route Propagation - Part 3 - Route Emitter and NATS](../http-routes/route-propagation-pt-3) if you
  help.
* ❓ How do the component route NATS messages compare to the app route NATS messages?

🤔 **Look at the routes table**
0. Bosh ssh onto the router VM.
0. Look at the GoRouter routes table and find your component route.
* See the story [Route Propagation - Part 4 -
  GoRouter](../http-routes/route-propagation-pt-4) if you need a reminder on
  how to this.
* ❓How does the component route compare to the app routes?

## ❓ Bonus Question

So if GoRouter routes off-platform users to other components, how do
off-platform users route to GoRouter?!?!

## Expected Result
There are (almost) no differences between the app routes and the component
routes. GoRouter does not know the difference between them and treats them the
same.
