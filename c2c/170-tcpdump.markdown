---
layout: single
title: Watch your packets with tcpdump
permalink: /c2c/tcpdump
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and called appA and appB
- There are no c2c network policies

### What?
Sometimes a user comes to us and says "container to container networking is
broken! AppA can't talk to AppB". After making sure that they have c2c policies
set, the next thing you might do is use tcpdump.

Tcpdump is a CLI tool that allows you to inspect all of the traffic flowing
through your container.

In this story we are going to look at the packets being sent from AppA to AppB.
Then we'll watch the packets being sent in response!

## How

📝 **Curl appB from appA**
1. Get the overlay IPs of appA and appB
1. Continually try to curl appB from appA
   {% include codeHeader.html %}
   ```bash
   watch -n 15 curl -sS APP_A_ROUTE/proxy/APP_B_OVERLAY_IP:8080
   ```

📝 **Look at those packets**
1. In another terminal, ssh onto the Diego Cell where appA is running and
   become root
1. Run `tcpdump`.  Ahhhhh too much information! ctrl+c! ctrl+c!  On a Diego
   Cell there are many packets being sent around, and tcpdump gives information
   about ALL OF THEM. We need to figure out a way to filter this overwhelming
   stream of information.
1.  Filter by packets where the source IP is APP_A_OVERLAY_IP and where the
    destination IP is APP_B_OVERLAY_IP.
   {% include codeHeader.html %}
    ```bash
    tcpdump -n src APP_A_OVERLAY_IP and dst APP_B_OVERLAY_IP
    ```

    You should see something like:
    ```
    $ tcpdump -n src 10.255.77.3 and dst 10.255.77.4
    tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
    listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
    ```

    ...and nothing else. Where are those packets?
    Notice that tcpdump is looking for packets listening on the eth0 interface. That's not where overlay packets go!

1. Look for packets on any interface
   {% include codeHeader.html %}
    ```bash
    tcpdump -n src APP_A_OVERLAY_IP and dst APP_B_OVERLAY_IP -i any
    ```
    Hey! Those are packets!

    Record the packets you see here from one curl.

    If appB was successfully responding, then you should also see packets being
    sent in the opposite direction.

1. See that no packets are being sent from AppB to AppA
   {% include codeHeader.html %}
    ```bash
    tcpdump -n src APP_B_OVERLAY_IP and dst APP_A_OVERLAY_IP -i any
    ```

🤔 **Add c2c policy**
1. Add c2c policy to allow traffic from appA to appB (`cf add-network-policy --help`)
1. Continually try to curl appB from appA

📝 **Look at those packets**
1. Look for packets from appA to appB
   {% include codeHeader.html %}
   ```bash
   tcpdump -n src APP_A_OVERLAY_IP and dst APP_B_OVERLAY_IP -i any
   ```
   Record the packets you see here from one curl.
   * ❓How are these packets different from before?

1. Look for packets from appB to appA
   {% include codeHeader.html %}
   ```bash
   tcpdump -n src APP_B_OVERLAY_IP and dst APP_A_OVERLAY_IP -i any
   ```

## Expected Result

You should see packets being sent in response from appB to appA. You should see `200 OK`.

## Resources
* [tcpdump man page](https://www.tcpdump.org/manpages/tcpdump.1.html)
* [helpful common tcpdump commands](https://www.rationallyparanoid.com/articles/tcpdump.html)
* [debugging non-c2c traffic in CF](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/troubleshooting.md#debugging-non-c2c-packets)
