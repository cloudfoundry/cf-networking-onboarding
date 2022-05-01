---
layout: single
title: \"bypass\" c2c policies
permalink: /c2c/bypass
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and named appA and appB
- You have at least 2 Diego Cells
- There are no c2c network policies

## What

Before container networking and the overlay network, apps could talk to each
other via the backend IP and port.

The old workflow was:
1. Create an ASG that allows apps to send traffic to the Diego Cell IPs
1. Bind that ASG to appA
1. AppA discovers appB's backend IP and port.
1. AppA sends traffic to appB.

Turns out that old workflow still (kind of) works. Even worse, some CF users
still rely on this behavior! Fun!

In this story you are going to see when this workflow does and doesn't work and
why.

## How
🤔 **Get Diego Cell IPs**
1. Get and record the following variables:
```
DIEGO_CELL_1_IP = ...
DIEGO_CELL_2_IP = ...
```

🤔 **Create a wide open ASG**
1. Create an ASG that allows traffic to DIEGO_CELL_1_IP and DIEGO_CELL_2_IP
1. Bind that ASG to appA.
1. Restart appA

🤔 **Setup**
1. Scale appB to 2 instances.
1. Make sure that you have the following configuration.
- Diego Cell 1: 1 instance of appA, 1 instance of appB (appB_1)
- Diego Cell 2: 1 instance of appB, this instance will be called appB_2

🤔 **Get environment variables**
1. Get and record the following variables:
```
APP_B_1_BACKEND_PORT = ...
APP_B_2_BACKEND_PORT = ...
```
hint: you can ssh onto a specific instance of an app, by passing the `-i` flag (`cf ssh --help`).

📝 **Bypass c2c rules and route integrity**
1. Ssh onto appA
1. See if you can access appB
{% include codeHeader.html %}
   ```bash
   curl DIEGO_CELL_1_IP:APP_B_1_BACKEND_PORT
   ```

1. See if you can access appB
{% include codeHeader.html %}
   ```bash
   curl DIEGO_CELL_2_IP:APP_B_2_BACKEND_PORT
   ```

## Expected Result
AppA should not be able to access appB_1. AppA should be able to access appB_2.

## ❓ Questions

* Why can appA access apps on other Diego Cells, but not on its own? (Hint:
  look at the iptables rules and review the diagrams in _Route Propagation -
  Part 5 - DNAT Rules_)
* Do you think this is a security concern?
* Do you think we should "fix" this? What would you suggest?
