---
layout: single
title: User Workflow
permalink: /c2c/user-workflow
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with silk release and at least 2 diego cells
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE

## What

In the HTTP route module you learned about the dataflow for HTTP traffic. But what if appA (an app within CloudFoundry) wants to talk to appB (another app within CloudFoundry)? Before Container to Container networking (c2c) there was no shortcut for intra-CF traffic. Without c2c, if appA wanted to talk to appB, then the traffic had to travel outside of the CF foundation, hit the external load balancer, and go through the normal HTTP traffic flow.

```
Life Before C2C
                                                  +------+
                                                  |      |
                                                  |      |
          +------------------------------------+  | AppA |
          |                                       |      |
          |                                       |      |
          |                                       +------+
          |
          v                                       +------+
                              +-----------+       |      |
  +------------------+        |HTTP Router|       |      |
  |HTTP Load Balancer| +----> |(GoRouter) | +---->| AppB |
  +------------------+        +-----------+       |      |
                                                  |      |
                                                  +------+
```

So many extra hops for apps within the same CF! Especially if the apps are on the same Diego Cell!

Those hops come with extra latency. It can also come with an extra security risk. If appB  *only* needs to be accessed by appA (for example if appA is the frontend microservice and appB is the backend microservice for appA), appB would still need to be accessible via an HTTP route! This is exposing appB to more attack vectors than it should be.

With container to container networking (c2c) apps within a foundation can talk directly to each other. Now appA can talk to appB without leaving the CF foundation. And appB doesn't need to be accessible via an HTTP route (just an internal one, we'll get to that later in the service discovery track).

```
Life with C2C

+------+        +------+
|      |        |      |
|      |        |      |
| AppA | +----> | AppB |
|      |        |      |
|      |        |      |
+------+        +------+
```

In order for appA to be *able* to talk to appB, it needs to have permission. You will need to create a network policy.

Let's ignore the technical implementation for now and go through the user workflow.

## How

📝**Use Container to Container Networking**

1. Push another
   [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
   app named appB. Use the `--no-route` so that no HTTP route is created for
   appB.
   {% include codeHeader.html %}
   ```bash
   cf push appB --no-route
   ```
1. Get the overlay IP for appB.  (An explanation of what "overlay" means awaits
   in future stories! For now just know that each app instance has a unique
   overlay IP that c2c uses.)
   {% include codeHeader.html %}
   ```bash
   cf ssh appB -c "env | grep CF_INSTANCE_INTERNAL_IP"
   ```

1. Get onto the container for appA and curl the appB internal IP and app port.
   {% include codeHeader.html %}
   ```bash
   cf ssh appA
   ```
   {% include codeHeader.html %}
   ```bash
   watch  "curl -Ssk <value of CF_INSTANCE_INTERNAL_IP>:8080"
   ```
You should get a `Connection refused` error because there is no network policy yet.

1.  In another terminal, add a network policy from appA to appB, with protocol tcp, on port 8080.
   {% include codeHeader.html %}
   ```bash
   cf add-network-policy appA appB --protocol tcp --port 8080
   ```

## Expected Result
After you add the policy, the curl from inside of the appA container to appB should succeed.
If it doesn't work, check that you created the policy in the correct direction, from appA --> appB, not the other way around.
