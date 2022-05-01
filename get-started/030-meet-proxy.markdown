---
layout: single
title: Meet the Proxy App
permalink: /get-started/meet-proxy
sidebar:
  title: "Getting Started"
  nav: sidebar-getting-started
---

## What

In these onboarding stories you will be using the proxy app a lot. Basically
every story, so let's get familiar with it. Proxy is little golang app that is
surprisingly powerful. In this story we are going to test out some of its
functions.

Often with container to container (c2c) networking stories, you need an app to
make a request to another app. Or you want to time how long DNS resolution
(turning a URL into an IP address) is taking. You could do this with `cf ssh`
and then running `curl` or `dig` (and we will in later stories!), but the proxy
app was created so we didn't have to do that. It has different endpoints that
will ...proxy... traffic to a given destination and that will do DNS resolution
of a URL you give it, among other things.  Using proxy gives a better mirror
how users use our product.

Let's check out proxy's power.

## How

📝 **Push a proxy app**

1. Clone the [cf-networking-release repo](https://github.com/cloudfoundry/cf-networking-release)
{% include codeHeader.html %}
   ```bash
   git clone https://github.com/cloudfoundry/cf-networking-release
   ```
1. Go to the proxy app
{% include codeHeader.html %}
   ```bash
   cd ~/workspace/cf-networking-release/src/example-apps/proxy
   ```
1. Push the app and name it appA
{% include codeHeader.html %}
   ```bash
   cf push appA
   ```

When you push an app, an HTTP route is automatically created. Let's call this route PROXY_ROUTE.

🤔 **Use the proxy app**

Skim through [proxy's
README](https://github.com/cloudfoundry/cf-networking-release/blob/develop/src/example-apps/proxy/README.md)
and look at all the endpoints that it has.

NOTE: you may need to use the -k option for these curl commands, if you get a certificate error.

1. Use the `/dig/URL_TO_DIG` endpoint to do DNS resolution for google.com.
1. Use the `/digudp/URL_TO_DIG` endpoint to do a DNS resolution  for google.com over udp. Dig usually uses tcp. This is a great way to test if udp traffic is working. (What are tcp and udp? Check out the resource below!)
1. Use the `/proxy/URL` endpoint to send traffic to neopets.com.
1. Use at least two more endpoints.

### Expected Outcome

Now you know the power of proxy!

## Resources
* [tcp vs udp](https://www.vpnmentor.com/blog/tcp-vs-udp/)
* [proxy's
  README](https://github.com/cloudfoundry/cf-networking-release/blob/develop/src/example-apps/proxy/README.md)

