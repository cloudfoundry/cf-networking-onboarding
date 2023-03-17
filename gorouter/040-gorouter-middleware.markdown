---
layout: single
title: Gorouter Middleware via Handlers
permalink: /gorouter/middleware-handlers
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What
In the previous story you learned that there is a component called proxy round tripper that handles the custom routing logic for gorouter. But what component(s) is in charge of all of the other logic? Where are all those VCAP headers added? Where are the access logs created? Where are the metrics emitted?

All of this per request logic is done in **handlers**. Handlers are middleware and in this story we are going to learn more about them.

## How
**Look at the code**
1. All of these handlers are set up in order in proxy.go. [Take a look](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191).
  ❓Can you guess what most of the handlers are for?
  ❓Are there any handlers that you don't know what they are for?

**Look at the docs**
1. These handlers are implemented via [negroni](https://github.com/urfave/negroni). Negroni is a BYOR (Bring your own Router) middleware-focused library that is designed to work directly with golang's net/http package.

## Links
* [Code where the handlers are set up](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191)
* [negroni github docs](https://github.com/urfave/negroni)
