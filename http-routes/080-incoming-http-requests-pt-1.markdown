---
layout: single
title: Incoming HTTP Requests Part 1 - See what's listening with netstat
permalink: /http-routes/incoming-http-requests-pt-1
sidebar:
  title: "HTTP Routes"
  nav: sidebar-http-routes
---

## Assumptions
- You have a CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
Netstat is a tool that can show information about network connections, routing
tables, and network interface statistics.  In the previous story we saw that
the GoRouter sent traffic for `APP_A_ROUTE` to
`DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT`.  Let's use netstat to see what is
listening at on the Diego Cell and specifically at
`DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT`.

## How
📝 **Look at open ports on a Diego Cell**
1. Ssh onto the Diego Cell where appA is deployed and become root.
1. Use netstat to look at open ports
 ```
 netstat -ntulp
 # -n  <---- show raw ip addresses and socket numbers
 # -t  <---- show tcp sockets
 # -u  <---- show udp sockets
 # -l  <---- display listening sockets
 # -p  <---- display PID/program name for sockets
 ```
  You should recognize the program names in the far right column. Most of them
  are the long running cf component processes.

1. Find the local address for the Route Emitter. What port is it running on?
   Does that match what is in the [spec
   file](https://github.com/cloudfoundry/diego-release/blob/develop/jobs/route_emitter/spec)?

1. Search for `DIEGO_CELL_ENVOY_PORT` in the output. Can you find it?

## Expected Result
You won't see the `DIEGO_CELL_ENVOY_PORT` anywhere in the netstat output because
nothing is *actually* running there.  But if there's nothing running there, how
does the traffic reach the app? Would you believe that iptables are involved?
Check out the next story to learn more :)
