---
layout: single
title: TCP Routes Table
permalink: /tcp-routes/routes-table
sidebar:
  title: "TCP Routes"
  nav: sidebar-tcp-routes
---

## Assumptions
- You have a CF deployed
- You have a TCP server deployed named tcp-app
- You have a TCP route mapped to tcp-app called TCP_ROUTE

## What

The TCP traffic flow is nearly identical to the HTTP traffic flow. The big
difference is that instead of an HTTP load balancer there is a TCP load
balancer and instead of GoRouter there is a TCP Router.

Go back to [this story in the http routes module](../http-routes/incoming-http-requests-pt-0) to review this flow.

In [this story in the http routes module](../http-routes/route-propagation-pt-4) you learned how to look at the
route table for the GoRouter. In this story you are going to look at the
analogous route table for the TCP Router.

## How

📝 **Try to list tcp routes**
1. List tcp routes via the routing api.
   ```bash
   cf curl /routing/v1/tcp_routes
   ```
   Most likely you will get the error message:
   ```json
   {"name":"UnauthorizedError","message":"Token is expired"}
   ```

🤔 **Get correct permissions**

Based on the [routing api docs](https://github.com/cloudfoundry/routing-api/blob/master/docs/api_docs.md#list-tcp-routes),
you need to have a client with routing.routes.read permissions.

There is probably already a client deployed with the correct permissions. Find
out the name and password for this user from the bosh manifest.

1. Download your manifest
   ```bash
   bosh manifest > /tmp/my-env.yml
   ```

1. Search for `routing.routes.read`. You should find uaa client properties that
   look like this:
    ```yaml
    routing_api_client:
      authorities: routing.routes.write,routing.routes.read,routing.router_groups.read
      authorized-grant-types: client_credentials
      secret: ((uaa_clients_routing_api_client_secret))
    ```
  The name of the client is: routing_api_client. The password is in credhub under
  the key uaa_clients_routing_api_client_secret.

1. Use the credhub CLI to get the password.

📝 **Use uaac to get the oath token**

1. Run `uaac` to see if you have the uaa CLI installed.

1. If you don't have it installed, install it.
   ```bash
   gem install cf-uaac
   ```

1. Target your uaa. (To determine this url you can run `cf api` and replace api with uaa.)
   ```bash
   uaac target uaa.<YOUR-DOMAIN>
   ```

1. Get the client information for the routing_api_client. It will prompt you for a password.
   ```bash
   uaac token client get routing_api_client
   ```

1. Get the bearer token
   ```bash
   uaac context
   ```
   You will see something like this (this one is truncated):
   ```
   client_id: routing_api_client
   access_token: eyJhbGciOiJ <------- This is the BEARER_TOKEN that you will need. Yours will be longer.
   token_type: bearer
   expires_in: 43199
   scope: routing.router_groups.read routing.routes.write routing.routes.read
   ```

📝 **Get tcp routes**
1. This time when you curl, pass in the bearer token as a header.
   ```bash
   cf curl /routing/v1/tcp_routes -H "Authorization: bearer BEARER_TOKEN" | jq .
   ```

## Expected Outcome

You should see one TCP route that looks like the one below (this one is edited for brevity):
```
{
    "router_group_guid": "e47c747a-d655-4ea8-5f1a-b59f21ad7852",
    "backend_port": 61004,         <--------- This is the backend port
    "backend_ip": "10.0.1.12",     <--------- This is the Diego Cell IP
    "port": 1025,                  <--------- This is the route port
    "isolation_segment": ""
}
```
## ❓ Questions
* Go back to the story _Route Propagation - Part 4 - GoRouter_  and look at the example HTTP route table entry.
  What differences do you see between the TCP routes and the HTTP routes?
* How does this difference match with what you understand about TCP and HTTP?

## Resource
* [routing api docs](https://github.com/cloudfoundry/routing-api/blob/master/docs/api_docs.md#list-tcp-routes)
