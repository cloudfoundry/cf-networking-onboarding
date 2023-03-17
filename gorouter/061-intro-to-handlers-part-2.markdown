---
layout: single
title: A Formal Introduction to Handlers 6 - 10
permalink: /gorouter/handlers-part-2
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What
In the past few stories you have learned about what handlers are and what their code looks like. In this story you are going to be formally introduced to the next five handlers in Gorouter.

## How
1. Read the summary for each handler.
1. Where there is a ✨ make sure you take a look at the code.

## The next five handlers
**6\. ✨[Access Log](https://github.com/cloudfoundry/gorouter/blob/main/handlers/access_log.go)**
    This handler creates and emits the access logs. There is one access log per request/response. You can find these logs at /var/vcap/sys/log/gorouter/access.log.

**7\. [Reporter](https://github.com/cloudfoundry/gorouter/blob/main/handlers/reporter.go)**
    This handler emits a metric containing the status code of the response and sets information about the response latency on the RequestInfo struct.

**8\. [HTTP Rewrite](https://github.com/cloudfoundry/gorouter/blob/main/handlers/http_rewrite.go)**
    This handler alters the headers on the request. It adds [these headers](https://github.com/cloudfoundry/routing-release/blob/4dd3ff8ebded5667232bfa0a7a7a0f5e89b3a8c1/jobs/gorouter/spec#L218-L220) and [removes these headers](https://github.com/cloudfoundry/routing-release/blob/4dd3ff8ebded5667232bfa0a7a7a0f5e89b3a8c1/jobs/gorouter/spec#L226-L228) based on bosh properties.

**9\. [Proxy Healthcheck](https://github.com/cloudfoundry/gorouter/blob/main/handlers/proxy_healthcheck.go)**
    This handler responds to healthcheck requests.

**10\. [Zipkin](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/zipkin.go)**
    Zipkin is a distributed tracing system. If you enable the [`router.tracing.enable_zipkin` bosh property](https://github.com/cloudfoundry/routing-release/blob/4dd3ff8ebded5667232bfa0a7a7a0f5e89b3a8c1/jobs/gorouter/spec#L242-L244), then this handler will attach zipkin headers on the request.


## Questions
❓If an invalid response results in a panic in golang's transport code, do you think that there will be a log line in the access logs? Why or why not? Consult the onion diagram if you need.

## Links
* [Code where the handlers are set up](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191)
