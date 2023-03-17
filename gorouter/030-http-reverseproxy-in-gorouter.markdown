---
layout: single
title: http.ReverseProxy in Gorouter
permalink: /gorouter/http-reverseproxy-in-gorouter
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What
In the last story you learned what a reverse proxy is at a high level. In this story you will look at how golang implements a reverse proxy and how gorouter uses that reverse proxy struct.

## How

**Read some docs**

1. 📚Read [these golang docs on the Reverse Proxy struct ](https://golang.org/pkg/net/http/httputil/#ReverseProxy)

**Look at Gorouter code**
1. Gorouter's `main.go`, like many `main.go`s, is where everything is set up and initialized, but not much happens. Skim through [main.go](https://github.com/cloudfoundry/gorouter/blob/main/main.go).
  ❓ Do you see anything interesting in there? Did anything catch your eye?
1. In `main.go` it initializes a new proxy. [Look at the code here](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/main.go#L184-L196).

1. When this new proxy is created, it creates a new http.ReverseProxy. [Look at the code here](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L143-L149).

1. Golang's http.ReverseProxy struct has 7 configurable properties. Gorouter configures 5 of them.
  ❓ Which http.ReverseProxy properties does gorouter configure?
  ❓ Which http.ReverseProxy properties does gorouter leave alone?

1. The ~~meatiest~~ tofu-iest bit of configuration that gorouter does is that it assigns its own Proxy Round Tripper to the http.ReverseProxyTransport. [Look at the code here](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L145).

1. This Proxy Round Tripper handles gorouter's custom routing logic and is a wrapper around the default Transport roundTrip function. [Take a quick look at the code for the Proxy Round Tripper](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/round_tripper/proxy_round_tripper.go). We will looker closer at this component in a later story.

## Links
* [Golang docs on the Reverse Proxy struct ](https://golang.org/pkg/net/http/httputil/#ReverseProxy)
* [Gorouter main.go](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/main.go#L184-L196)
* [Code where http.ReverseProxy is made](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L143-L149)
* [Proxy Round Tripper Code](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/round_tripper/proxy_round_tripper.go)
