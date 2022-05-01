---
layout: single
title: Request an External URL
permalink: /bosh-dns/external
sidebar:
  title: "Bosh DNS"
  nav: sidebar-bosh-dns
---

## Assumptions
- You have 2 my-http-server instances deployed (see the story [Life Without
  Route Registrar](../route-registrar/life-without-rr) for setup)

## What

In this story you are going to follow the DNS request for an external URL.
Surprisingly (maybe) this will still involve Bosh DNS, even though the request
is for an external URL.

## How

📝 **See that your VM already has Bosh DNS**
By default Bosh DNS is on every VM in a OSS Cloud Foundry deployment.
1. Bosh ssh onto one of the my-http-server VMs and become root.
1.  Run `monit summary`. You should see a process called bosh-dns running.

📝 **Look at the DNS servers**

1. Bosh ssh onto one of the my-http-server VMs.
1. Look at the /etc/resolv.conf file. This file contains the IPs for the DNS
   servers used for all DNS lookups.

   The file should look something like this.
{% include codeHeader.html %}
   ```bash
   cat /etc/resolv.conf
   ```

   ```
   # This file was automatically updated by bosh-dns
   nameserver 169.254.0.2          <-------------- record this value as BOSH_DNS_IP

   nameserver 169.254.169.254      <-------------- record this value as NON_BOSH_DNS_IP
   search c.cf-container-networking-gcp.internal google.internal
   ```

📝 **Find Bosh DNS running**

So you have a value for BOSH_DNS_IP, but who do you _know_ this is the Bosh DNS IP?

1. Use netstat to see what IP the Bosh DNS process is bound to.

{% include codeHeader.html %}
   ```bash
   netstat -tulpn
   ```

   ```
   Active Internet connections (only servers)
   Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
   tcp        0      0 169.254.0.2:53          0.0.0.0:*               LISTEN      3227/bosh-dns
   tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      743/sshd
   tcp        0      0 127.0.0.1:53080         0.0.0.0:*               LISTEN      3227/bosh-dns
   tcp        0      0 127.0.0.1:2822          0.0.0.0:*               LISTEN      3300/monit
   tcp        0      0 127.0.0.1:2825          0.0.0.0:*               LISTEN      613/bosh-agent
   udp        0      0 169.254.0.2:53          0.0.0.0:*                           3227/bosh-dns
   ```

📝 **Do a non-Bosh DNS lookup**

1. Use dig to do a DNS request for any non-CF url.

{% include codeHeader.html %}
   ```bash
   dig neopets.com
   ```

   ```
   ; <<>> DiG 9.10.3-P4-Ubuntu <<>> neopets.com
   ;; global options: +cmd
   ;; Got answer:
   ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 35487
   ;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

   ;; OPT PSEUDOSECTION:
   ; EDNS: version: 0, flags:; udp: 512
   ;; QUESTION SECTION:
   ;neopets.com.     IN  A

   ;; ANSWER SECTION:
   neopets.com.    3599  IN  A 23.96.35.235

   ;; Query time: 49 msec
   ;; SERVER: 169.254.0.2#53(169.254.0.2)
   ;; WHEN: Thu Oct 03 18:15:20 UTC 2019
   ;; MSG SIZE  rcvd: 67
   ```

1. Let's go through and try to understand this dig request.

| **snippet from dig response** |  **meaning** |
| -- | -- |
|`ANSWER: 1` |This means that the DNS request successfully found an IP for the url. If an IP was not found, it would be `ANSWER: 0`|
|`23.96.35.235` |This is the IP for the neopets.com. Try it in your browser! |
|`SERVER: 169.254.0.2#53(169.254.0.2)` |This means that the DNS server that handled this request is at IP 169.254.0.2 and port 53 (this is the standard port for DNS requests). |

Do you recognize that server IP? That's the BOSH_DNS_IP that you recorded earlier!

📝 **Look at logs**

1. Look at the bosh DNS logs. You should see something like...

{% include codeHeader.html %}
   ```bash
   tail -f /var/vcap/sys/log/bosh-dns/bosh_dns*
   ```

   ```
   [ForwardHandler] 2019/10/03 18:15:20
   INFO - handlers.ForwardHandler Request [1]
   [neopets.com.] 0 [recursor=169.254.169.254:53] 49064000ns
   ```
1. Do you recognize that recursor IP? That's the NON_BOSH_DNS_IP you recorded earlier!

📝 **Tell dig what DNS server to use**

1. Try digging the external URL again, but this time force dig to use the
BOSH_DNS_IP as the DNS server
* ❓ Does this succeed? Why or why not?

1. Try digging the external URL again, but this time force dig to use the
NON_BOSH_DNS_IP as the DNS server
* ❓ Does this succeed? Why or why not?

## Expected Outcomes
Bosh DNS only knows information about Bosh DNS routes. For any other URL
(neopets, for example) asks a different DNS server. Both the Bosh DNS server
and the non-Bosh DNS server can (with recursion) resolve the external URL.

## Helpful Commands

**Do a DNS lookup**
```
dig URL [@SERVER_IP]

# for example
dig neopets.com @169.254.4.4
```
