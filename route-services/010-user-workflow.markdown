---
layout: single
title: Using Route Services
permalink: /route-services/user-workflow
sidebar:
  title: "Route Services"
  nav: sidebar-route-services
---

## Assumptions
- You have a CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE

## What

In the Route Propagation tack we talked about the flow of network traffic when
an internet user tries to access your app. A very simplified version of this
flow looked like this:
1. The traffic hits a load balancer in front of the CF Foundation
1. The load balancer sends it to one of the GoRouters
1. The GoRouter sends the traffic to the app

What is left out in that simplified flow, is that the app developer also has
the option to introduce Route Services into this flow. From the [CF
docs](https://docs.cloudfoundry.org/services/route-services.html): "Cloud
Foundry application developers may wish to apply transformation or processing
to requests before they reach an application. Common examples of use cases
include authentication, rate limiting, and caching services."

When you add Route Services to the flow it looks like this:
1. The traffic hits a load balancer in front of the CF Foundation
1. The load balancer sends it to one of the GoRouters
1. GoRouter sees there is a Route Service to apply
1. GoRouter appends a header to the traffic saying that the traffic has hit
   the GoRouter once and is being redirected to a Route Service
1. GoRouter redirects the traffic to the Route Service (which may or may not
   be an app inside of the same CF)
1. The Route Service applies a transformation to traffic
1. The Route Service redirects the traffic back to the load balancer
1. The load balancer sends it to one of the GoRouters
1. GoRouter sees the traffic has a header saying it has already been
   redirected to a Route Service
1. The GoRouter sends the traffic to the app

![user provided route service](https://docs.cloudfoundry.org/services/images/route-services-user-provided.png)

In this story we are going to create a user-provided route service to rate
limit traffic going to appA.

## How

📝 **Take control measurements**
1. Download boom, a benchmarking tool.
{% include codeHeader.html %}
   ```bash
   go get github.com/rakyll/boom
   ```

1. Use boom to see what percentage of requests to APP_A_ROUTE return status 200.
 This command sends 100 requests, 10 concurrently with 10 QPS (queries per second)
{% include codeHeader.html %}
   ```bash
   boom -n 100 -c 10 -q 10 http://APP_A_ROUTE
   ```

   You should see that all requests returned 200 OK.
   ```
   Status code distribution:
     [200]	100 responses
   ```

📝 **Use a Route Service**
0. Set
   [this](https://github.com/cloudfoundry/routing-release/blob/2e1cc8b89df0b569102489f7eda159107094fc9f/jobs/gorouter/spec#L145-L147)
   gorouter bosh property and redeploy:
{% include codeHeader.html %}
   ```yaml
   router.ssl_skip_validation: true
   ```
0. Follow [these
   instructions](https://github.com/cloudfoundry-samples/ratelimit-service) to
   deploy a rate limiting Route Service and bind it to appA.  This example
   Route Service is old and you may get errors that it is running on an old go
   version. Fix it! Even better, submit a PR with your fix.

🤔 **Take measurements with Route Service**

1. Bind the Route Service to appA.
1. Make the Route Service seriously limit traffic.
1. Run the benchmarking tests.

## Expected Results
* See that without the Route Service there is no rate limiting to appA. All
  responses should have status code 200.
* See that with the Route Service there is rate limiting to appA. Some
  responses should have status code 200 and the rest should have code 429.
```
Status code distribution:
  [200]	2 responses
  [429]	98 responses
```

Delete the rate limiter app and unbind the route service before the next story.

## Resources
* [CF docs - Route Services](https://docs.cloudfoundry.org/services/route-services.html)
* [Rate Limiting Route Service](https://github.com/cloudfoundry-samples/ratelimit-service)
