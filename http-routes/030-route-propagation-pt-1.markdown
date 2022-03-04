---
layout: single
title: Route Propagation Part 1 - Cloud Controller
permalink: /http-routes/route-propagation-pt-1
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- I recommend deleting all other apps

## What
The Cloud Controller API (CC API) maintains the database of all apps, domains, routes, and route mappings.
However Cloud Controller does not keep track of *where* those apps are deployed. Nor does CC track the IPs and ports each
route should, well, *route* to. That's the job of the router, often called GoRouter.

CC keeps track of the desired state. The user wants a route called MY_ROUTE that sends traffic to appA.
But the user doesn't (shouldn't) care about the logistics needed to make that route happen. That is the responsibility of other components.

Let's look at what information Cloud Controller *does* keep track of.

## How

0. 🤔 Map a route to appA. Let's call this route APP_A_ROUTE. I recommend _deleting_ all other routes.

0. 🤔 Look at the domains, routes, destinations (route mappings), and apps via the Cloud Controller's API.
    To look at all the domains you can curl using `cf curl /v3/domains`. Use the [API docs](https://v3-apidocs.cloudfoundry.org/) to figure out the endpoints for the other resources.

This is all of the information that CC contains about routes. Note there are no IPs anywhere. Note that all of these routes are for CF apps, none of them are for CF components.

### Expected Result
You can view data from CC about the route APP_A_ROUTE that you created.

## Recorded Values
Record the following values that you generated or discovered during this story.
```
APP_A_ROUTE=<value>
```

## Resources
[Cloud Controller API docs](https://v3-apidocs.cloudfoundry.org/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR.
Amelia will be eternally grateful. How? Open [this file in
GitHub](https://github.com/cloudfoundry/cf-networking-onboarding). Search for
the phrase you want to edit. Make the fix!_
