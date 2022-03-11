---
layout: single
title: Logging
permalink: /c2c/logging
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
- There are no c2c policies

## What

When you use c2c policies you have the option of logging every time that one
app attempts to send traffic to another app.

In this story you will learn (1) how to turn on this feature, (2) how to look
at the logs, and (3) how this feature is implemented (hint: iptables).

## How

🤔 **Turn on logging**
1. Follow [these
   instructions](https://github.com/cloudfoundry/silk-release/blob/77ecbb775780d74c5a8b6e87c5554dab375a9235/docs/traffic_logging.md#traffic-logging)
   to enable c2c policy logging _AND_ ASG logging. (There is a bug where you
   need to enable both in order for c2c logging to work. Maybe you could be the
   one who fixes it?)

📝 **Look at logs**
1. In one terminal, Ssh onto the Diego Cell where appB is running and become
   root.
1. Watch the kern.logs (kern stands for kernel, as in the linux kernel).
```
tail -f kern.log
```
1. In another terminal, curl appB from appA. You should see log line similar to the one below.
  ```
  2019-04-18T18:01:05.306552+00:00 localhost kernel: [22246.987902]
  DENY_C2C_f13ffea0-9d0d-4ee9-                   <----- DENY means that the traffic was blocked by iptables rules. The GUID here is the beginning of the instance GUID that we have seen before.
  IN=s-010255077003                              <----- The interface of the source app (This seems backwards given that this is IN, I'm not sure why this is).
  OUT=s-010255077004                             <----- The interface of the destination app (This seems backwards given that this is OUT, I'm not sure why this is).
  MAC=aa:aa:0a:ff:4d:03:ee:ee:0a:ff:4d:03:08:00  <----- This is a combination of the mac address of the source app's container network interface and the mac address of the source Diego Cell's VTEP.
  SRC=10.255.77.3                                <----- The overlay IP of the source app
  DST=10.255.77.4                                <----- The overlay IP of the destination app
  LEN=60 TOS=0x00 PREC=0x00 TTL=63 ID=12504 DF PROTO=TCP SPT=39304 DPT=8080 WINDOW=27400 RES=0x00 SYN URGP=0
  ```

1. Add network policy from appA to appB (`cf add-network-policy --help`)
1. Curl appB from appA again. You should see log line similar to the one below (it's very similar to the one above).
```
2019-04-18T18:21:00.670494+00:00 localhost kernel: [23442.243710]
OK_0002_c7de6123-d906-4c65-9               <----- OK means that the traffic was allowed by iptables rules.
IN=s-010255077003
OUT=s-010255077004
MAC=aa:aa:0a:ff:4d:03:ee:ee:0a:ff:4d:03:08:00
SRC=10.255.77.3
DST=10.255.77.4
LEN=60 TOS=0x00 PREC=0x00 TTL=63 ID=35333 DF PROTO=TCP SPT=41330 DPT=8080 WINDOW=27400 RES=0x00 SYN URGP=0
MARK=0x2                    <----- Successful logs also include the mark of the source app.
```

🤔 **Find the implementation**
1. Look at the iptables rules on the Diego Cell
1. Find the rule that logs c2c traffic.

## Expected Result

You (1) turned on this feature, (2) looked at the logs, and (3) saw how this
feature is implemented with iptables rules.

## Resources
* [Traffic Logging with Silk Release](https://github.com/cloudfoundry/silk-release/blob/77ecbb775780d74c5a8b6e87c5554dab375a9235/docs/traffic_logging.md#traffic-logging)
