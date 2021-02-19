Gorouter the Code (TM)

⚠️⚠️⚠️
**WARNING: The "Gorouter the Code" stories are very new. ** You will probably find some typos. You will probably find some things that aren't as clear as they could be.
Please open a [PR or an issue](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md) when you find these problems. This is my first attempt at making stories that dive into the implementation of a component. Good luck! All feedback is a 🎁 and is very appreciated.
⚠️⚠️⚠️

**The goal of this section is**
 * introduce the golang concepts gorouter uses as a reverse proxy
 * give you an understanding of what code is evaluated for each request/response
 * give you time and prompts to encourage you to look around gorouter code 
 * get you more comfortable poking around the gorouter code

L: gorouter-the-code
---
What is Gorouter?

## Assumptions
* You have completed the stories in the "http-routes" section.

## What

So what is Gorouter? 

[Gorouter](https://github.com/cloudfoundry/gorouter) is a bosh job in [routing release](https://github.com/cloudfoundry/routing-release). It is in charge of a handful of things, including: 

* subscribing to route registration messages from nats and keeping an up-to-date routing table
* acting as a reverse proxy and routing to backends 👈👈👈
* serving a healthcheck endpoint

In this section we are going to focus on the second bullet: how gorouter acts as a reverse proxy and routes to backends.

## How
1. 📚Skim through [bosh configuration options for gorouter](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/gorouter/spec) to start getting the sense of what this component can do.

## Links
* [Gorouter](https://github.com/cloudfoundry/gorouter) 
* [routing release](https://github.com/cloudfoundry/routing-release)
* [bosh configuration options for gorouter](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/gorouter/spec)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
---
What is a reverse proxy?

## Assumptions
* You have completed all the previous stories in this track.

## What 

In the last story we learned that part of Gorouter's job is to be a reverse proxy.

But what _is_ a reverse proxy??? Let's find out!

## How

1. 🎬Watch this Hussein Nasser's video ["Proxy vs Reverse Proxy Server Explained"](https://www.youtube.com/watch?v=SqqrOspasag&ab_channel=HusseinNasser) _length 14:17_

and/or 

1. 📚Read [CloudFlare's "What Is A Reverse Proxy?"](https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/)

## Questions
❓Hussein mentioned 5 use cases for a reverse proxy: caching, load balancing, ingress, canary deployment, and microservices. Which use cases do you think Gorouter is used for?

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
http.ReverseProxy in Gorouter

## Assumptions
* You have completed all of the other stories in this track

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

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
Gorouter middleware via handlers

## Assumptions
* You have completed all of the previous stories in this track.

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

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
Gorouter is an onion

## Assumptions
* You have completed all of the previous stories in this track

## What 
In the last story you learned that there are handlers that act as middleware in gorouter. But how do these handlers work?

![middleware onion](https://storage.googleapis.com/cf-networking-onboarding-images/onion-layers-abstract.png)

**Gorouter is like an onion**
This middleware is organized like an onion. A request goes through all of the middleware layers to get to the app (the center of the onion). Then the app sends a response, which goes through all the layers again in reverse order. This means that the handlers have two chances to do things during each request/response: once when the request is on the way to the app and once when the response in on the way back to the client.

![gorouter onion img](https://storage.googleapis.com/cf-networking-onboarding-images/the-gorouter-onion.png)

**But what does this look like in the code?**

Each handler has a function with the signature: `ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)`. Inside the function each handler calls `next(rw, r)`. The stuff inside the function _before_ the `next(rw, r)` is done when request is at this handler on its way to the app. The stuff inside the function _after_ the `next(rw, r)` is done once the response has made it back to this handler on its way back to the client. Some handlers only do things when the request is on its way to the app. Some handlers only do things when the response is returning to the client. Some handlers do stuff in both cases.

**Handlers need to share information**
In order for handlers to pass information between each other there is a [Request Info](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/requestinfo.go#L22-L32) struct. There is one Request Info struct per request. The handlers can look up this struct using the request. [See here for an example in the lookup handler](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/lookup.go#L102). 

## How

🤔 **Examine the lookup handler**

1. Look at [the lookup handler](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/lookup.go#L55-L109).
  ❓ What does this handler do? 
  ❓ What does this handler set on the Request Info struct? 
  ❓ When does this handler do its logic? When the request is on its way to the app or when the response is on its way to the client? Or both? How do you know?


## Extra Credit 
🤔 Find another handler that only contains logic that occurs when the request is on its way to the app. 
🤔 Find a handler that only contains logic that occurs when the response is on its way to the client. 

## Links
* [code for all handlers](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers)
* [request info struct definition](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/requestinfo.go#L22-L32)
* [the lookup handler code](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/handlers/lookup.go#L55-L109)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
A formal introduction to handlers 1 - 5

## Assumptions
* You have completed all of the previous stories in this track.

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

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
A formal introduction to handlers 6 - 10

## Assumptions
* You have completed all of the previous stories in this track.

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

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---
A formal introduction to handlers 11 - 17

## Assumptions
* You have completed all of the other stories in this track.

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
❓If someone tries to send HTTP/2 traffic where will it fail? 

## Links
* [Code where the handlers are set up](https://github.com/cloudfoundry/gorouter/blob/68fb24bfe35a379fee6591651b96660dc9712a80/proxy/proxy.go#L162-L191)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---

The Proxy Round Tripper

## Assumptions
* You have completed all of the previous stories in this track.

## What 
In the past stories you were _thoroughly_ introduced to all of the handlers. Those handlers are the first layers of the gorouter onion. 

![gorouter onion img](https://storage.googleapis.com/cf-networking-onboarding-images/the-gorouter-onion.png)

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

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? [Open this file in GitHub](https://github.com/pivotal/cf-networking-program-onboarding/edit/master/gorouter-the-code.prolific.md). Search for the phrase you want to edit. Make the fix!_

L: gorouter-the-code
L: questions
---

[RELEASE] Gorouter the Code (TM) ⇧

Congrats! You made it! Time to start making some PRs to Gorouter!
L: gorouter-the-code
