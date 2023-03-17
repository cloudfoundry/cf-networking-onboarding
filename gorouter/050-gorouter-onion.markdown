---
layout: single
title: Gorouter is an Oonion
permalink: /gorouter/gorouters-have-layers-onions-have-layers
sidebar:
  title: "Gorouter the Code"
  nav: sidebar-gorouter
---
## What 
In the last story you learned that there are handlers that act as middleware in gorouter. But how do these handlers work?

![middleware onion](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/onion-layers-abstract.png)

**Gorouter is like an onion**
This middleware is organized like an onion. A request goes through all of the middleware layers to get to the app (the center of the onion). Then the app sends a response, which goes through all the layers again in reverse order. This means that the handlers have two chances to do things during each request/response: once when the request is on the way to the app and once when the response in on the way back to the client.

![gorouter onion img](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/the-gorouter-onion.png)

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
