---
layout: single
title: Dataflow Part 3 - Marks
permalink: /c2c/dataflow-pt-3
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have an CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and named appA and appB (the fewer apps you have deployed the
  better)

## Review
This track of stories is going to go through the steps (listed below) that were
covered in the dataflow overview.  The steps and diagram will be at the top of
each story in case you need to orient yourself. Higher quality diagram
[here](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/c2c-data-plane.png).

![c2c traffic flow](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on
   Diego Cell 2). This packet is called the overlay packet (aka the c2c
   packet).
1. The packet exits the app container through the veth interface.
1. **The overlay packet is marked with a ...mark... that is unique to the
   source app.  <------- CURRENT STORY**
1. Because the packet is an overlay packet, it is sent to the silk-vtep
   interface on the Diego Cell. This interface is a VXLAN interface.
1. The overlay packet is encapsulated inside of an underlay packet. This
   underlay packet is addressed to underlay IP of the Diego Cell where the
   destination app is located (appA in this case).
1. The underlay packet exits the cell.
1. The packet then travels over the physical underlay network to the correct
   Diego Cell.
1. The packet arrives to the correct Diego Cell
1. The underlay packet is decapsulated. Now it's just the overlay packet again.
1. Iptables rules check that appA is allowed to talk to appB based on the mark
   on the overlay packet.
1. If traffic is allowed, the overlay network directs the traffic to the
   correct place.

## What
In the last story we found out that packets leave the app container over the
veth pair interface. In this story we are going to look at how overlay packets
are marked.

First off, *why* are packets marked?

Packets are marked with a mark that is unique to the source app. Each instance
of the source app has its packets marked with the same ID. If there is a c2c
policy, then on the destination Diego Cell there are iptables rules that allow
traffic with a certain mark to certain destinations. The policies look like the
following diagram. Using one mark per source app, decreases the number of
iptables rules needed per c2c policy, especially when there are large number of
app instances.

![markful networking policy
diagram](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/diagram-of-silk-networking-policies.png)

If packets weren't marked we *could* use the overlay IP as a unique identifier.
However, this would create the need for many more iptables rules, especially
when there are a large number of app instances.

![markless networking policy
diagram](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/diagram-of-flannel-networking-policies.png)

You will learn more about how the c2c policies check the mark in later stories
(hint: it uses iptables). For now, let's focus on how traffic is marked (hint:
it uses iptables).

Here is an example iptables rule that sets a mark.  ![setmark iptables rule
example](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/set-mark-iptables-rule-example.png)

In this story we are going to find the applicable set-xmark rules for our app
and we're going to find out where that mark value comes from.

## How

📝 **Find set-xmark rules**

1. Delete all network policies. This time you are going to use the networking
   API because former policies from deleted apps can linger in the database,
   but not show up in the CLI.
   ```bash
   cf curl /networking/v1/external/policies > /tmp/policies.json
   cf curl -X POST /networking/v1/external/policies/delete -d @/tmp/policies.json
   # check that they are all deleted
   cf curl /networking/v1/external/policies
   ```

1. Ssh onto the Diego Cell where appA is located and become root.
1. Look for the iptables rules that set marks.
   ```bash
   iptables -S | grep set-xmark
   ```
Nothing! This is because there are no policies. A mark is only allocated to an
app when that app is used in a container to container (c2c) networking policy.

1. In a different terminal, add a c2c policy from appA to appB  (`cf add-network-policy --help`)
1. Back on the Diego Cell, look again for iptables rules that set marks
   ```bash
   iptables -S | grep set-xmark
   ```

You should see something that looks like the colorful example above. Copy and
paste it here.
```
# PASTE THE SET-XMARK RULE HERE
```

The source IP is the overlay IP for your app. The comment is the app guid for
appA. And the mark is...well... where *does* the mark come from?

When a c2c policy is created, the policy server determines if the app has aCancel changes
mark already or not. If the app doesn't have a mark yet, it creates one. Let's
look at all these marks.  The marks are an internal implementation of how c2c
policies work, so they are not exposed on the external API (the API the CLI
uses). But there is also an internal policy server API. The internal policy
server API is the API that other CF components, like the vxlan-policy-agent,
use.

📝 **Look at marks via internal policy server API**

1. You will need certs to call this API, those certs are located on the Diego
   Cell at `/var/vcap/jobs/vxlan-policy-agent/config/certs`
1. Follow the
   [docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/policy-server-internal-api.md)
   for how to list all of the c2c policies (the actual policy server url may vary from the docs. Check `policy_server_url` in `/var/vcap/jobs/vxlan-policy-agent/config/vxlan-policy-agent.json` to get the right one).  You should see something like the
   following. The tag for appA should match the mark you saw in the iptables
   rule.
```
{
  "total_policies": 1,
  "policies": [
    {
      "source": {
        "id": "90ff1b89-a69d-4c77-b1bd-415ae09833ed",  <------- AppA guid
        "tag": "0004"                                  <------- AppA mark, should match what you saw in the iptables rule above
      },
      "destination": {
        "id": "0babce4f-6739-4fc8-8f74-01f11179bfe5",  <------- AppB guid
        "tag": "0005",                                 <------- AppB mark
        "protocol": "tcp",
        "ports": {
          "start": 8080,
          "end": 8080
        }
      }
    }
  ]
}
```

## ❓ Questions
* Hey! AppB has a mark too. Why?
* Marks on packets are limited to 16 bits. How many unique marks is this? Does this give you scaling concerns for c2c networking?

## Expected Outcome
The data about the tag for the source app from the internal policy server API should match the mark in the iptables rule.

## Look at the Code
In the vxlan policy agent (vpa), there is a component called the planner. The planner gets information from the internal policy server API about all of the c2c policies. The planner turns this policy information into proposed iptables rules.

[Here](https://github.com/cloudfoundry/silk-release/blob/0150c154a47770ed98d39453ca75fc1495848fe2/src/code.cloudfoundry.org/vxlan-policy-agent/planner/planner_linux.go#L399-L404)
the VPA goes through all of the source apps and creates mark rules for them

[Here](https://github.com/cloudfoundry/silk-release/blob/0150c154a47770ed98d39453ca75fc1495848fe2/src/code.cloudfoundry.org/lib/rules/rules.go#L108-L113)
is the implementation of *NewMarkSetRule*
