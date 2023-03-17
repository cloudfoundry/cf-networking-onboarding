---
layout: single
title: A Formal Introduction to Handlers 11-17
permalink: /gorouter/handlers-part-3
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What
In the past few stories you have learned about what handlers are and what their code looks like. In this story you are going to be formally introduced to the last 7 handlers in Gorouter.

## How
1. Read the summary for each handler.
1. Where there is a ✨ make sure you take a look at the code.

## The next seven handlers

**11\. [w3c](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/w3c.go)**
    [W3c is a type of tracing header](https://www.w3.org/TR/trace-context/). If you enable the [`router.tracing.enable_w3c` bosh property](https://github.com/cloudfoundry/routing-release/blob/4dd3ff8ebded5667232bfa0a7a7a0f5e89b3a8c1/jobs/gorouter/spec#L245-L250), then this handler will attach w3c headers on the request.

**12\. ✨[Protocol Check](https://github.com/cloudfoundry/gorouter/blob/main/handlers/protocolcheck.go)**
    This handler checks to make sure that the request is a valid HTTP protocol.

**13\. ✨[Lookup](https://github.com/cloudfoundry/gorouter/blob/main/handlers/lookup.go)**
    This handler looks up the host in the route table and gets the route pool. The route pool contains all backends (IPs and Ports) for that route. The handler sets the route pool on the RequestInfo struct.

**14\. [Client Cert](https://github.com/cloudfoundry/gorouter/blob/main/handlers/clientcert.go)**
    This handler handles the "X-Forwarded-Client-Cert" and either passes it on to the app or not [based on configuration](https://docs.google.com/spreadsheets/d/1Zlws0TJibQLbjDZWXKeRYrSyM9sOSyTbNhA7DX-_fAA/edit#gid=0).

**15\. [XForwarded Proto](https://github.com/cloudfoundry/gorouter/blob/main/handlers/x_forwarded_proto.go)**
    This handler handles the "X-Forwarded-Proto" header based on configuration.

**16\. [Route Service](https://github.com/cloudfoundry/gorouter/blob/main/handlers/routeservice.go)**
    This handler handles requests for route services and request that are coming from a route service. (Route services are explained in another section).

**17\. ✨ [Proxy](https://github.com/cloudfoundry/gorouter/blob/main/proxy/proxy.go#L223-L266)**
    This handler handles web socket requests

18\. That's it! Next comes the routing decisions with the Proxy Round Tripper! That's in the next story.

## Questions
❓If someone tries to send HTTP/2 traffic and the [router.enable_http2](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/gorouter/spec#L108) property is not set, where will it fail?

## Links
* [Code where the handlers are set up](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191)
