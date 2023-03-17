---
layout: single
title: A Formal Introduction to Handlers 1 - 5
permalink: /gorouter/handlers-part-1
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What
In the past few stories you have learned about what handlers are and what their code looks like. In this story you are going to be formally introduced to the first five handlers in Gorouter.

## How
1. Read the summary for each handler.
1. Where there is a ✨ make sure you take a look at the code.

## The first five handlers
**1\. ✨[Panic Check](https://github.com/cloudfoundry/gorouter/blob/main/handlers/paniccheck.go)**
    If a request/response causes a panic this handler will catch it and log. This way the gorouter continues working for other requests.

**2\.✨ [Request Info](https://github.com/cloudfoundry/gorouter/blob/main/handlers/requestinfo.go)**
    This handler creates the [RequestInfo struct](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/requestinfo.go#L22-L32). There is one per request. The handlers use this struct to pass information to other handlers. For example, this handler sets the [reqInfo.StartedAt time to time.Now()](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/requestinfo.go#L51). Later the access log handler will use this information.

**3\. [Proxy Writer](https://github.com/cloudfoundry/gorouter/blob/main/handlers/proxywriter.go)**
  This handler creates the response writer.

**4\. [Vcap Request ID Header](https://github.com/cloudfoundry/gorouter/blob/main/handlers/request_id.go)**
  This handler generates and sets the "X-Vcap-Request-Id" on the request. This header value is unique per each request and is used for debugging and tracing requests/responses through different components.

**5\. [HTTP Start Stop](https://github.com/cloudfoundry/gorouter/blob/main/handlers/httpstartstop.go)**
    This handler emits HTTP Start Stop events. You can see these events by [installing the firehose nozzle](https://docs.cloudfoundry.org/loggregator/cli-plugin.html), sending traffic to an app, and watching for this event: `cf nozzle --filter HttpStartStop | grep gorouter `

## Questions
❓Why do you think that the panic handler is first?

## Links
* [Code where the handlers are set up](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191)
