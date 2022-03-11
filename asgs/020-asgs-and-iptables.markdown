---
layout: single
title: ASGs and Iptables
permalink: /asgs/asgs-and-iptables
sidebar:
  title: "Application Security Groups (ASGs)"
  nav: sidebar-asgs
---

## Assumptions
- You have a CF deployed
- You have completed the iptables-primer track
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and named appA (the fewer apps you have deployed the better)
- You have the public_networks security group bound to all running and staging
  containers

## What
In the first story you learned what ASGs *are*, but how are they implemented
under-the-hood?  (Hint: iptables)

Information about ASGs are stored in CAPI's database. When a container create
is triggered (for example by a cf push), CAPI passes ASG information along to
Diego (container scheduler). Diego then passes the ASG information along to
Garden (the container creator), which passes it along to the container
networking components. After this long journey, the container networking
components turn this ASG information into iptables rules on the Diego Cells.

Iptables rules are a series of rules that network packets match against to
decide whether that traffic is allowed or not. This is one way firewalls can be
implemented.  ASGs are only concerned with egress traffic from CF apps to
something external to the CF foundation (usually the internet).

The iptables man page is extremely helpful as a reference! Consult it if you
have any questions throughout this story.

Let's investigate what iptables rules are created for the public_networks ASG.

## How

📝 **Find the iptables rule for the public_networks ASG**
1. Look at the ASG public_networks
  ```
  cf security-group public_networks
  ```
  Note which IP ranges are specified. Save these rules somewhere accessible for
  future reference.

1. Ssh onto the Diego Cell where appA is running and become root.
  ```
  # First determine the IP of the Diego Cell
  cf ssh appA  -c "env | grep CF_INSTANCE_IP"
  ```
  ```
  # Look at bosh output to see which Diego Cell has that IP
  bosh instances
  ```
  ```
  # Ssh on
  bosh ssh diego-cell/DIEGO_CELL_INSTANCE_GUID
  ```

1. List all of the iptables rules on the filter table on the Diego Cell
  ```
  iptables -t filter -S
  ```

  Yep. It's a lot. Take a deep breath. Read it line by line. I promise it will
  start to become comprehensible.

  Remember those IP ranges that are in the public_networks security group?  If
  you search for one of them (ie. `0.0.0.0-9.255.255.255`) you should be able
  to find an iptables rule that looks something like this:

  ```
  # If you have logging enabled
  -A netout--772fbbd5-862a-4b3d-7 -m iprange --dst-range 0.0.0.0-9.255.255.255 -g netout--772fbbd5-862a-4--log
  # If you don't have logging enabled
  -A netout--772fbbd5-862a-4b3d-7 -m iprange --dst-range 0.0.0.0-9.255.255.255 -j ACCEPT
  ```

Copy the iptables rule that you found on the Diego Cell that looks like the
line above. It will be necessary to reference it soon.

🤔 **Decipher the iptables rule for the public_networks ASG**
When you run `iptables -S` the iptables rules are displayed in the format that
they were created in. For example, if you prepended `iptables` to the line
above and ran (DON'T! this is hypothetical) `iptables -A
netout--772fbbd5-862a-4b3d-7 -m iprange ...` it would add that iptables rule.

With that in mind, use `iptables --help` to define what the following flags
mean and record them somewhere to reference later...

```
-A                                 ...
-m                                 ...
--dst-range                        ...
-g                                 ...
-j                                 ...
```

1. What is the name of the chain that the above iptables rule is appended to?
   Let's call this CHAIN-NAME.
1. What conditions need to be met for a packet to match this rule?
1. If logging is enabled: If packets meet the above conditions, what chain of
   rules will they jump to and evaluate against next? Let's call this
   JUMP-TO-CHAIN-NAME.
1. Look at the rules only in the chains from the questions above.
  ```
  iptables --list CHAIN-NAME
  iptables --list JUMP-TO-CHAIN-NAME # <---- only if logging is enabled
  ```
Packets stop matching against iptables rules when they reach either the ACCEPT
target(let the packet flow through!), the DROP target (stop that packet in its
tracks!), or the REJECT target (nicely tell the sender that this destination is
not available to them) targets.

🤔 **Read iptables rules to predict behavior**
1. Pretend there is a packet trying to hit the IP 23.96.32.148 via tcp. Where will it hit ACCEPT or DROP?
1. Pretend there is a packet trying to hit the IP 172.20.0.3 via tcp. Where will it hit ACCEPT or DROP?
1. ssh onto your app and try to curl 23.96.32.148. Did it succeed or fail? Did it match your expectations?
1. ssh onto your app and try to curl 172.20.0.3. Did it succeed or fail? Did it match your expectations?

## Expected Result
You should be able to find the iptables rules for the public_networks and be
able to follow the flow to either ACCEPT or DROP.

Iptables are hard! Hopefully they have been demystified a little bit.  Spend
some extra time looking at all the iptables rules on a Diego Cell. Research
what you don't understand. Ask lots of questions.

## Resources
**iptables**
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)
* [Julia Evans iptables basics](https://jvns.ca/blog/2017/06/07/iptables-basics/)
* [iptables primer](https://danielmiessler.com/study/iptables/)

**ASGs**
* [Application Security Groups Documentation](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html)
* [Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)
* ["Taking Security to the Next Level—Application Security Groups" by Abby Kearns](https://blog.pivotal.io/pivotal-cloud-foundry/products/taking-security-to-the-next-level-application-security-groups)
* ["Making sense of Cloud Foundry security group declarations" by Sadique Ali](https://sdqali.in/blog/2015/05/21/making-sense-of-cloud-foundry-security-group-declarations/)
