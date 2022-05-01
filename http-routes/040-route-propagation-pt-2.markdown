---
layout: single
title: Route Propagation Part 2 - Diego BBS
permalink: /http-routes/route-propagation-pt-2
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---
## Assumptions
- You have a OSS CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## What

**Diego** is an umbrella term for many components that work together to make CF
container orchestration happen. These jobs are bundled together in
[diego-release](https://github.com/cloudfoundry/diego-release/tree/develop/jobs).

**BBS** stands for Bulletin Board System. This is the database that the Diego
components use to keep track of DesiredLRPs and ActualLRPs.

**LRPs**, or Long Running Processes, represent work that a client (ie Cloud
Controller) wants Diego to keep running indefinitely. Apps are the primary
example of LRPs. Diego will try to keep them running as best it can. When an
app stops or fails, it will attempt to restart it and keep it running.

**Desired LRPs** represent what a client (ie Cloud Controller) wants "the
world" to look like (for example, how many instances of which apps). They
contain no information about where LRPs should be run, because the user
shouldn't care.

**Actual LRPs** represent what Diego is currently *actually* running. Actual
LRPs contain information about which Diego Cell the LRP is running on and which
port maps to the LRP.

For this story, let's look at the data stored in the BBS and see what
information it has about appA. Diego will go on to send this information via
the Route Emitter to GoRouter, so GoRouter knows where to send network traffic
to.

## How

📝 **Look at actualLRPS**
0. Grab the guid for appA. You'll need it in a moment. Let's call it
   APP_A_GUID.
   {% include codeHeader.html %}
   ```bash
   cf app appA --guid
   ```
0. Ssh onto the Diego Cell vm where appA is running and become root. You can
   find where appA is running by running the following command:
   ```bash
   cf curl /v2/apps/<app-guid>/stats
   ```
0. Use the [cfdot CLI](https://github.com/cloudfoundry/cfdot) to query BBS for
   actualLRPs. Cfdot is a helpful CLI for using the BBS API.  It's a great tool
   for debugging on the Diego Cell.
{% include codeHeader.html %}
   ```bash
   cfdot actual-lrps | jq .
   ```
0. Search through the actual LRPs for APP_A_GUID. It should match the beginning
   of a process guid. You'll find an entry for each instance of appA that is
   running.
0. Let's dissect and store the most important information (for us) about appA:
   ```
   {
     "process_guid": "ab2bd185-9d9a-4628-9cd8-626649ec5432-cb50adac-6861-4f03-92e4-9fcc1a204a1e",
     "index": 0,
     "cell_id": "d8d4f5fe-36f2-4f50-8c4a-8df293f6bc5b",
     "address": "10.0.1.12",                  <------ DIEGO_CELL_IP
       "ports": [
         {
           "container_port": 8080,            <------ CONTAINER_APP_PORT
           "host_port": 61012,                <------ DIEGO_CELL_APP_PORT
           "container_tls_proxy_port": 61001, <------ CONTAINER_ENVOY_PORT
           "host_tls_proxy_port": 61014,      <------ DIEGO_CELL_ENVOY_PORT
         },
         {
           "container_port": 2222,            <------ CONTAINER_SSH_PORT
           "host_port": 61013,                <------ DIEGO_CELL_SSH_PORT
           "container_tls_proxy_port": 61002, <------ CONTAINER_ENVOY_SSH_PORT
           "host_tls_proxy_port": 61015       <------ DIEGO_CELL_ENVOY_SSH_PORT
         }
       ],
     "instance_address": "10.255.116.6",      <------ The overlay IP address of this app instance, let's call this the OVERLAY_IP
     "state": "RUNNING",
      ...
   }
   ```
0. Let's define all of these values.
  * 👇 These are important for this module 👇
    * **DIEGO_CELL_IP** - The cell's IP address where this app instance is
      running, also sometimes called the host IP.
    * **CONTAINER_APP_PORT** - The port the app is listening on inside of its
      container. 8080 is the default value.
    * **DIEGO_CELL_APP_PORT** -  The port on the Diego Cell where traffic to your
      app is sent to before it is forwarded to the overlay address and the
      container_port.
    * **CONTAINER_ENVOY_PORT** - The port inside of the app container that envoy
      is listening on for HTTPS traffic. This is the default value (currently
      unchangeable).
    * **DIEGO_CELL_ENVOY_PORT** - The port on the Diego Cell where traffic to
      your app's envoy sidecar is sent to before it is forwarded to the overlay
      address and the container_tls_proxy_port.
  * 👇 These are NOT important for this module 👇
    * **CONTAINER_SSH_PORT** - The port exposed on the app container for sshing
      onto the app container
    * **DIEGO_CELL_SSH_PORT** - The port on the Diego Cell where ssh traffic to
      your app container is sent to before it is forwarded to the overlay address
      and the ssh container_port.
    * **CONTAINER_ENVOY_SSH_PORT** - The ssh port inside of the app container
      that envoy is listening on for ssh traffic. This is the default value
      (currently unchangeable).
    * **DIEGO_CELL_ENVOY_SSH_PORT** - The port on the Diego Cell where ssh
      traffic to your app's envoy sidecar is sent to before it is forwarded to
      the overlay address and the ssh container_tls_proxy_port.
    * **OVERLAY_IP** - The overlay IP address of this app instance.

0. Use the cfdot CLI to query BBS for desiredLRPs.

## ❓ Questions
* What information is provided for desiredLRPs, but not for actualLRPs?
* What information is provided for actualLRPs, but not for desiredLRPs?
* How does this match with the definition of desired and actual LRPs in the "what" section above?

## Expected Result
Get information from BBS about the desiredLRP and actualLRP for appA. Use cfdot
CLI to discover the following values and record them.

## Recorded Values
Record the following values that you generated or discovered during this story.
```
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## Resources
* [cfdot docs](https://github.com/cloudfoundry/cfdot)
* [diego design
  notes](https://github.com/cloudfoundry/diego-design-notes#what-are-all-these-repos-and-what-do-they-do)
