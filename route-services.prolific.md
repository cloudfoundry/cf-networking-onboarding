User Workflow: Route Services

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE

## What

In the Route Propagation tack we talked about the flow of network traffic when an internet user tries to access your app. A very simplified version of this flow looked like this:
1. The traffic hits a load balancer in front of the CF Foundation
1. The load balancer sends it to one of the GoRouters
1. The GoRouter sends the traffic to the app

What is left out in that simplified flow, is that the app developer also has the option to introduce Route Services into this flow. From the [CF docs](https://docs.cloudfoundry.org/services/route-services.html): "Cloud Foundry application developers may wish to apply transformation or processing to requests before they reach an application. Common examples of use cases include authentication, rate limiting, and caching services."

When you add Route Services to the flow it looks like this:
1. The traffic hits a load balancer in front of the CF Foundation
1. The load balancer sends it to one of the GoRouters
1. **GoRouter sees there is a Route Service to apply**
1. **GoRouter appends a header to the traffic saying that the traffic has hit the GoRouter once and is being redirected to a Route Service**
1. **GoRouter redirects the traffic to the Route Service (which may or may not be an app inside of the same CF)**
1. **The Route Service applies a transformation to traffic**
1. **The Route Service redirects the traffic back to the load balancer**
1. **The load balancer sends it to one of the GoRouters**
1. **GoRouter sees the traffic has a header saying it has already been redirected to a Route Service**
1. The GoRouter sends the traffic to the app

![user provided route service](https://docs.cloudfoundry.org/services/images/route-services-user-provided.png)

In this story we are going to create a user-provided route service to rate limit traffic going to appA.

## How

📝 **Take control measurements**
1. Download boom, a benchmarking tool.
 ```
 go get github.com/rakyll/boom
 ```

1. Use boom to see what percentage of requests to APP_A_ROUTE return status 200.
 This command sends 100 requests, 10 concurrently with 10 QPS (queries per second)
 ```
 boom -n 100 -c 10 -q 10 http://APP_A_ROUTE
 ```

 You should see that all requests returned 200 OK.
 ```
 Status code distribution:
   [200]	100 responses
 ```

📝 **Use a Route Service**
0. Set [this](https://github.com/cloudfoundry/routing-release/blob/2e1cc8b89df0b569102489f7eda159107094fc9f/jobs/gorouter/spec#L145-L147) gorouter bosh proprty and redeploy: 
   ```
   router.ssl_skip_validation: true
   ```
0. Follow [these instructions](https://github.com/cloudfoundry-samples/ratelimit-service) to deploy a rate limiting Route Service and bind it to appA.
This example Route Service is old and you may get errors that it is running on an old go version. Fix it! Even better, submit a PR with your fix.

🤔 **Take measurements with Route Service**

1. Bind the Route Service to appA. See instructions from previous link.
1. Make the Route Service seriously limit traffic.
1. Run the benchmarking tests.

### Expected Result
See that without the Route Service there is no rate limiting to appA. All responses should have status code 200.
See that with the Route Service there is rate limiting to appA. Some responses should have status code 200 and the rest should have code 429.
```
Status code distribution:
  [200]	2 responses
  [429]	98 responses
```

Delete the rate limiter app and unbind the route service before the next story.

## Resources
[CF docs - Route Services](https://docs.cloudfoundry.org/services/route-services.html)
[Rate Limiting Route Service](https://github.com/cloudfoundry-samples/ratelimit-service)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: route-services
L: user-workflow
---

Exploiting Route Services

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You unbound the rate limiter route service from the previous story

## What

In the previous story when you created your own Route Service you ran a command that looked something like this...

`cf create-user-provided-service MY_SERVICE -r MY_SERVICE_URL`

The route MY_SERVICE_URL could be inside or outside of your Cloud Foundry. However, one takes precedence.
Let's say that MY_SERVICE_URL exists inside *and* outside of your Cloud Foundry. (Before a fix) GoRouter would
"hairpin" and always default to the MY_SERVICE_URL that exists *inside* the Cloud Foundry.

Let's exploit this.

## How

🤔 **Pretend that you are an innocent dev**

1. You need to deploy to an older version of routing-release from before this attack vector was fixed. Deploy CF with [routing-release 0.186.0](https://bosh.io/releases/github.com/cloudfoundry-incubator/cf-routing-release?version=0.186.0).

1. Create a Route Service that sends all traffic to whitehouse.gov
 ```
 cf create-user-provided-service realwhitehouse -r https://www.whitehouse.gov/
 ```

1. Bind this new Route Service to your APP_A_ROUTE
 ```
cf bind-route-service DOMAIN  --hostname HOSTNAME realwhitehouse`
 ```

1. See that when you curl APP_A_ROUTE, you now just get www.whitehouse.gov (I don't know why you would want this. But it's an easy way to show this attack vector.)

😈 **Pretend that you are a malicious dev**

1. Push an app called fakewhitehouse using `--no-route`

1. Create the domain whitehouse.gov (`cf create-shared-domain --help`)

1. Map the route www.whitehouse.gov to the app fakewhitehouse
 ```
cf map-route fakewhitehouse whitehouse.gov --hostname www
 ```

1. In one terminal, watch the logs for the fakewhitehouseapp (`cf logs --help`)

1. In another terminal curl APP_A_ROUTE

### Expected Result
You should see in the fakewhitehouse app logs that traffic was redirected to fakewhitehouse
You should see by the response, that traffic was never sent to the real whitehouse.gov

❓What happened?
❓Why is this really, really bad?
❓How could you exploit this on a shared deployment like PWS?

**Extra Credit**
1. Find the code in GoRouter that let this happen. Maybe start your search [here](https://github.com/cloudfoundry/gorouter/blob/f6879c04bac67c1e467f14b79496b9832869df91/proxy/round_tripper/proxy_round_tripper.go#L126-L196).

## Resources
[CF docs - Route Services](https://docs.cloudfoundry.org/services/route-services.html)
[Rate Limiting Route Service](https://github.com/cloudfoundry-samples/ratelimit-service)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: route-services
L: questions
---

[RELEASE] Route Services ⇧
L: route-services
