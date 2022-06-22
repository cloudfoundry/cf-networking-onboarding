---
layout: single
title: Dataflow Part 8 - Enforcing Policy
permalink: /c2c/dataflow-pt-8
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and named appA and appB (the fewer apps you have deployed the
  better)

## Review
This track of stories is going to go through the steps (listed below) that were
covered in the dataflow overview.  The steps and diagram will be at the top of
each story in case you need to orient yourself. Higher quality diagram
[here](https://storage.googleapis.com/cf-networking-onboarding-images/c2c-data-plane.png).

![c2c traffic
flow](https://storage.googleapis.com/cf-networking-onboarding-images/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on
   Diego Cell 2). This packet is called the overlay packet (aka the c2c
   packet).
1. The packet exits the app container through the veth interface.
1. The overlay packet is marked with a ...mark... that is unique to the source
   app.
1. Because the packet is an overlay packet, it is sent to the silk-vtep
   interface on the Diego Cell. This interface is a VXLAN interface.
1. The overlay packet is encapsulated inside of an underlay packet. This
   underlay packet is addressed to the underlay IP of the Diego Cell where the
   destination app is located (appA in this case).
1. The underlay packet exits the cell.
1. The packet then travels over the physical underlay network to the correct
   Diego Cell.
1. The packet arrives to the correct Diego Cell
1. The underlay packet is decapsulated. Now it's just the overlay packet again.
1.  **Iptables rules check that appA is allowed to talk to appB based on the
    mark on the overlay packet.   <------- CURRENT STORY**
1. If traffic is allowed, the overlay network directs the traffic to the
   correct place.

## What

In the last story you learned how the VTEP is able to encapsulate overlay
packets and send them to the correct Diego Cell via the underlay network.

Now that the packet has arrived at the correct Diego Cell, how is c2c policy
enforced? (hint: iptables rules)

In this story, you are going to look at the iptables rules on the destination
Diego Cell

## How
🤔 **Setup**
1. Make sure there are no c2c policies. (`cf network-policies --help`)

📝**Look at iptables rules**
1. Ssh onto the Digeo Cell where appB is running and become root.
1. Look at the iptables rules on the filter table
   ```bash
   iptables -S
   ```
   You should see a chain that looks like the following
   ```
   -N vpa--1555607784864807      <----------- Let's call this VPA_CHAIN_NAME
   ```
   VPA stands for the VXLAN Policy Agent. The VPA is a component from silk-release that is responsible for translating the c2c policies in the database into iptables rules on the Diego Cell.
   Any rules related to enforcing c2c policies will be on a custom chain that starts with "vpa".

1. Look at the rules on the vpa chain.
   ```
   iptables -S VPA_CHAIN_NAME
   ```
   There's nothing there. No rules attached. That is because there are no c2c policies.

🤔 **Create c2c policy and look at iptables rules**
1. Create a c2c network policy from appA to appB (`cf add-network-policy --help`). The diagram in the review section shows apps on two different Diego Cell. The same iptables rules will be created and enforced regardless of whether appA and appB are on the same Diego Cell or not. A future story will go into more detail about this. But for this story, just know that it doesn't matter what Diego Cell your apps are running on.
1. On the Diego Cell, look at the rules on the vpa chain again.
   ```
   iptables -S VPA_CHAIN_NAME
   ```
   What! you should get the error.`iptables: No chain/target/match by that name.`.
   When the VPA sees that it needs to update the iptables rules on a Diego Cell, it *doesn't* append new rules to existing chains. Instead, it replaces ALL OF THE RULES. When the VPA replaces all of the rules, the VPA chain get's renamed with a new timestamp.
1. Find the new name of the VPA chain.
1. Look at the rules on that chain.  They should look something like this
   ```
   -N vpa--1555608460726256
   -A vpa--1555608460726256 -s 10.255.77.3/32 -m comment --comment "src:2f978b7f-b3d2-4f50-b08f-776634a6e411" -j MARK --set-xmark 0x2/0xffffffff
   -A vpa--1555608460726256 -d 10.255.77.4/32 -p tcp -m tcp --dport 8080 -m mark --mark 0x2 -m conntrack --ctstate INVALID,NEW,UNTRACKED -j LOG --log-prefix "OK_0002_c7de6123-d906-4c65-9 "
   -A vpa--1555608460726256 -d 10.255.77.4/32 -p tcp -m tcp --dport 8080 -m mark --mark 0x2 -m comment --comment "src:2f978b7f-b3d2-4f50-b08f-776634a6e411_dst:c7de6123-d906-4c65-9717-c5d040568c76" -j ACCEPT
   ```

   Let's look at them line by line.

   ⬇️ This is adding the VPA chain
   ```
   -N vpa--1555608460726256
   ```

   ⬇️ This is about marking outgoing traffic. (We went over this in the story *Container to Container Networking - Part 2 - Marks*). You might not have this rule. You will only have this rule if this app is also the source of a networking policy.
   ```
   -A vpa--1555608460726256 -s 10.255.77.3/32 -m comment --comment "src:2f978b7f-b3d2-4f50-b08f-776634a6e411" -j MARK --set-xmark 0x2/0xffffffff
   ```

   ⬇️ This is about logging when traffic hits this rule. Your deployment might have logging turned off, so this rule might or might not be present. A future story will go more into this.
   ```
   -A vpa--1555608460726256 -d 10.255.77.4/32 -p tcp -m tcp --dport 8080 -m mark --mark 0x2 -m conntrack --ctstate INVALID,NEW,UNTRACKED -j LOG --log-prefix "OK_0002_c7de6123-d906-4c65-9 "
   ```

   ⬇️ THIS IS THE RULES THAT ENFORCES C2C POLICY!!!!
   ```
   -A vpa--1555608460726256 -d 10.255.77.4/32 -p tcp -m tcp --dport 8080 -m mark --mark 0x2 -m comment --comment "src:2f978b7f-b3d2-4f50-b08f-776634a6e411_dst:c7de6123-d906-4c65-9717-c5d040568c76" -j ACCEPT
   ```
   ❓Can you rewrite this rule in sentence form to explain what it is checking?

## Expected Results
You can find and understand the iptables rule that enforces c2c policy.

## Resources
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)

