---
layout: single
title: Incoming HTTP Requests Part 2 - DNAT Rules
permalink: /http-routes/incoming-http-requests-pt-2
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
- It will help if you have completed the "iptables-primer" track, but it is not
  required.

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
In a previous story you saw that the GoRouter sent traffic for `APP_A_ROUTE` to
`DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT`.  But in the previous story, you saw
nothing was actually listening at that port on the Diego Cell. So how does the
network traffic hit the app?  With the help of iptables rules of course!
Everything always comes back to iptables rules.

Nothing is actually listening on the Diego Cell at port `DIEGO_CELL_ENVOY_PORT`.
Instead, all packets that are sent there hit iptables rules that then redirect
them to... somewhere. Let's find out!

Let's not brute force looking through every iptables rule. Instead, let's
reason about what chain and table most likely contain these iptables rules.
Hint, these rules translate ingress network traffic sent from GoRouter.

## How

🤔 **Find those iptables rules**
Aidan Obley made a great diagram showing the different types of traffic in CF
and which iptables chains they hit in what order.  We are currently concerned
with ingress traffic, which is represented by the orange line.

1. Look at the diagram. Which chain does the ingress traffic hit first?
   ![traffic-flow-through-iptables-on-cf-diagram](https://storage.googleapis.com/cf-networking-onboarding-images/traffic-flow-through-iptables-on-cf.png)

2. Based on the previous diagram, the ingress traffic hits the prerouting chain
   first. Look at the diagram below and do some research to learn more about
   the raw, conn_tracking, mangle, and nat tables.  Which table should contain
   the rules to redirect our traffic to a new address?  ![iptables tables and
   chains
   diagram](https://storage.googleapis.com/cf-networking-onboarding-images/iptables-tables-and-chains-diagram.png)

    NAT stands for Network Address Translation. That sounds like what we want.
    So let's look at iptables rules for the nat table on the prerouting chain.

3. Ssh onto the Diego Cell where your app is running and become root.
4. Run
   ```
   iptables -S -t nat`
   ```
   You should see some custom chains attached to the PREROUTING chain.
   There will be one custom chain per app running on this Diego Cell.
   They will look something like this.
   ```
   -A PREROUTING -j netin--a0d2b217-fa7d-4ac1-65
   -A PREROUTING -j netin--317736ed-70ac-4087-74
   ...
   ```
5. You should also see 4 rules that contain the `OVERLAY_IP` for appA.
   If you look closely you'll see that the ports in the iptables rules match the ports we saw when inspecting the actual LRPs.
   Which port represents what?
   ```
   -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61012 -j DNAT --to-destination 10.255.116.6:8080
   -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61013 -j DNAT --to-destination 10.255.116.6:2222
   -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61014 -j DNAT --to-destination 10.255.116.6:61001
   -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61015 -j DNAT --to-destination 10.255.116.6:61002
   ```

6. For appA, find the rule that will match with the traffic the GoRouter sends to `DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT`.
   It should look something like this...
   ![example DNAT rule with explanation](https://storage.googleapis.com/cf-networking-onboarding-images/example-DNAT-rule-with-explanation.png)

   In summary, when the GoRouter sends network traffic to 10.0.1.12:61014 (`DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT`) 
   it gets redirected to 10.255.116.6:61001 (`OVERLAY_IP:CONTAINER_ENVOY_PORT`).
   But, looking at the information we learned about the actual LRP, the app isn't even listening on 10.255.116.6:61001, envoy is.
   When will the traffic finally reach the app!?!?

## Expected Result
Inspect the iptables rules that DNAT the traffic from the GoRouter and send it
to the correct sidecar envoy.

## Resources
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)
