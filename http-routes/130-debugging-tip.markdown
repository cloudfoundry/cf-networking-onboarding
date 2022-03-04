---
layout: single
title: "Debugging Tip: Skip the Load Balancer"
permalink: /http-routes/debugging-tip
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## What
In this story we are going to learn how to remove the LB (load balancer) from the data flow.

Here is simplified diagram of the data flow of an http route:
```
+----+    +----------+         +-------+     +-----+
| LB +--->+ Gorouter +-------->+ Envoy +---->+ App |
+----+    +----------+         +-------+     +-----+
```

When to do this: 
* when you are having problems connecting to an app and you want to start picking off items on by one that are _not_ the problem.
* when one particular gorouter is having problems and you want to send traffic to just that gorouter.
* when you are debugging and want to point your traffic at a particular gorouter so you can find the logs easier.

## How

📝**Send HTTP traffic using LB**
1. Curl the route for your app!
  ```
  curl APP_A_ROUTE -v
  ```
1. Save this output.

❓Do you see a host header on the request? How did that get there?

📝**Send HTTP traffic without using LB**
1. Get the IP for your router VM.
1. Send the traffic to the Gorouter IP and set the route in the host header:
  ```
  curl GOROUTER_IP -H "Host: APP_A_ROUTE" -v
  ```
1. Huh. That timed out and failed. 
1. Ssh onto any bosh VM and try again. 

❓Why did it fail the first time and succeed when you were ssh-ed onto any bosh VM?
❓How does the output for this curl differ from the one you did in the first section?

### Expected Result
You were able to send traffic to a specific gorouter bypassing the LB.

## Resources
- [Host Header Docs](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host)

--- 
🙏 _If this story needs to be updated: please, please, PLEASE submit a PR.
Amelia will be eternally grateful. How? Open [this file in
GitHub](https://github.com/cloudfoundry/cf-networking-onboarding). Search for
the phrase you want to edit. Make the fix!_
