---
layout: single
title: Write a Firewall
permalink: /iptables/write-a-firewall
sidebar:
  title: "Iptables"
  nav: sidebar-iptables
---

## What

Make a super basic firewall for your docker container. This (extremely practical) firewall will only let egress traffic exit if it is going to neopets.com.

## How

🤔 **Make your own rule**
1. Make your own chain.
1. Attach rule to that chain that accepts traffic if it is sent to ip 23.96.35.235 (neopets!) port 80 using tcp.
1. Attach a rule to that chain that drops all other traffic.
1. Add a jump rule to either the OUTPUT, FORWARD, or INPUT chains so that the traffic exiting the docker container will hit your custom chain.
1. Curl google.com. Does it fail?
1. Curl 23.96.35.235:80. Does it succeed?
1. Curl http://neopets.com. Does it fail or succeed? Why?
1. Practice deleting chains and rules: delete all of the rules and chains that you created.

## ❓ Question
Why didn't curling http://neopets.com work?

## Expected Result
Hopefully you realize by now that iptables rules are very powerful and very fun :D

## Extra Credit
Use iptables rules to make it so you can curl neopets.com, but not google.com

## Resources
* [iptables man page](http://ipset.netfilter.org/iptables.man.html)
* [Julia Evans iptables basics](https://jvns.ca/blog/2017/06/07/iptables-basics/)
* [iptables.info - great resource linked to in Julia's blog, but only available
  with the way back
  machine](https://web.archive.org/web/20180310124055/http://www.iptables.info/en/iptables-contents.html)
* [iptables primer](https://danielmiessler.com/study/iptables/)

