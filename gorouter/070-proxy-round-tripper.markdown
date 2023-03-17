---
layout: single
title: The Proxy Round Tripper
permalink: /gorouter/proxy-round-tripper
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
The Proxy Round Tripper

## Assumptions
* You have completed all of the previous stories in this track.

## What
In the past stories you were _thoroughly_ introduced to all of the handlers. Those handlers are the first layers of the gorouter onion.

![gorouter onion img](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/the-gorouter-onion.png)

In this story we are going to go deeper into the Proxy Round Tripper.

## How

**Read docs**
1. You have seen this Proxy Round Tripper before in the story "http.ReverseProxy in gorouter". In that story you looked at where the proxy created an http.ReverseProxy. It set the Proxy Round Tripper (prt) as the http.ReverseProxy.Transport. [Take a look at this code again here](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L145).

1. Take a look again at the [golang docs for the http.ReverseProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy).

1. Based on those docs you can see that our Proxy Round Tripper has to match the interface for an http.RoundTripper. Look at the [golang docs for the http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper).
  ❓Based on the [golang docs for the http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper) what function must our Proxy Round Tripper implement?
  ❓What arguments does this function take?
  ❓What values does this function return?


**Look at code**
1. Look at the [Proxy Round Tripper code](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/round_tripper/proxy_round_tripper.go).
1. Find and record the line numbers where the following events happen:
  * It selects a backend endpoint
  * It forwards the request to a backend (app or component)
  * It retries if it failed to connect to a backend
  * It forwards the request to a route service
  * It retries if it failed to connect to a route service
  * It sets the VCAP_ID for sticky sessions. [Learn about sticky sessions here](https://github.com/cloudfoundry/routing-release/blob/develop/docs/session-affinity.md).


**But this proxy round tripper code looks fairly high level? What actually makes the connection to the backend?**

Our Proxy Round Tripper does not re-implement the low-level stuff related to sending traffic. Instead the Proxy Round Tripper wraps a different round tripper that uses the default Trasport
1. Look at [where the other round tripper with the default Transport is made](https://github.com/cloudfoundry/gorouter/blob/1e285091233eec98592cb11bad7d23c8dcbc90c4/proxy/proxy.go#L109-L118).
1. Read (skim) more about http.Transport and its version of roundTrip [here](https://golang.org/src/net/http/transport.go).


## Hint
[This code in Proxy Round Tripper`reqInfo.RouteServiceURL == nil`](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/round_tripper/proxy_round_tripper.go#L132) means that this app is not bound to a  route service and that the gorouter should send the request onto a backend.

## Links
* [golang docs for the http.ReverseProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy)
* [golang docs for the http.RoundTripper](https://golang.org/pkg/net/http/#RoundTripper)
* [Proxy Round Tripper code](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/round_tripper/proxy_round_tripper.go).
* [Learn about sticky sessions here](https://github.com/cloudfoundry/routing-release/blob/develop/docs/session-affinity.md)
* [where the other round tripper with the default Transport is made](https://github.com/cloudfoundry/gorouter/blob/1e285091233eec98592cb11bad7d23c8dcbc90c4/proxy/proxy.go#L109-L118).
* [http.Transport](https://golang.org/src/net/http/transport.go)
