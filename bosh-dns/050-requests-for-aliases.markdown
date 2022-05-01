---
layout: single
title: Request a Bosh DNS Alias
permalink: /bosh-dns/request-alias
sidebar:
  title: "Bosh DNS"
  nav: sidebar-bosh-dns
---
## Assumption
- You have my-http-server deployed with an alias setup from the previous story

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
HTTP_SERVER_ALIAS=<value>
```

## What
In this story you are going to look at what happens under the hood when you do a DNS request for HTTP_SERVER_ALIAS.

## How

📝 **Do a DNS lookup for your alias**
1. Bosh ssh onto any Cloud Foundry VM
1. Use dig to do a DNS request for your alias.

{% include codeHeader.html %}
   ```bash
   dig HTTP_SERVER_ALIAS
   ```

   ```
   ; <<>> DiG 9.10.3-P4-Ubuntu <<>> meow.meow
   ;; global options: +cmd
   ;; Got answer:
   ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 28967
   ;; flags: qr aa rd ra; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 0

   ;; QUESTION SECTION:
   ;meow.meow.      IN  A

   ;; ANSWER SECTION:
   meow.meow.   0 IN  A 10.0.1.23   # <------------ This should match the IP of one of the my-http-server VMs
   meow.meow.    0 IN  A 10.0.1.22   # <------------ This should match the IP of the other my-http-server VM

   ;; Query time: 2 msec
   ;; SERVER: 169.254.0.2#53(169.254.0.2)        # <------------ This should match the BOSH_DNS_IP
   ;; WHEN: Thu Oct 03 20:45:06 UTC 2019
   ;; MSG SIZE  rcvd: 83
   ```

📝 **Look at logs**

1. Look at the bosh-dns logs on the machine you did the dig on in the steps above

{% include codeHeader.html %}
   ```bash
   tail -f /var/vcap/sys/log/bosh-dns/bosh_dns*
   ```

   ```
   [RequestLoggerHandler] 2019/10/03 20:49:43
   INFO - handlers.DiscoveryHandler Request [1]
   [amelia.meow.] 0 160000ns                     # <------------ Note, there is no recursor
   ```

   * ❓ Remember how with neopets.com there was a recursor in the logs? Based
     on what you know about recursors, why do you think there is no recursor
     listed in this log line?

📝 **Tell dig what DNS server to use**

1. Try digging your alias again, but this time force dig to use the BOSH_DNS_IP
   as the DNS server
* ❓ Does this succeed? Why or why not?

1. Try digging your alias again, but this time force dig to use the
   NON_BOSH_DNS_IP as the DNS server
* ❓ Does this succeed? Why or why not?

## Expected Results
The Bosh DNS server knows how to recurse to the non-Bosh DNS server. However,
the non-Bosh DNS server does not recurse to the Bosh DNS server. Because of
this, the non-Bosh DNS server will not be able to resolve HTTP_SERVER_ALIAS.

## Helpful Command

**Do a DNS lookup**
```
dig URL [@SERVER_IP]

# for example
dig neopets.com
# OR
dig neopets.com @169.254.4.4
```
