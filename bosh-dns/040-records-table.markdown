---
layout: single
title: Records Table
permalink: /bosh-dns/records-table
sidebar:
  title: "Bosh DNS"
  nav: sidebar-bosh-dns
---

## Assumptions
- You have my-http-server deployed with an alias setup from the previous story

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
HTTP_SERVER_ALIAS=<value>
```

## What
So how does Bosh DNS work? How does it figure out what IP to send traffic to?

## How

📝 **Find your alias in the Bosh DNS records table**

1. Bosh ssh onto any VM in your CF deployment.
1. Look at the Bosh DNS records table. (You might need to install jq)
   ```bash
   cat /var/vcap/instance/dns/records.json | jq .
   ```
1. Find the data about HTTP_SERVER_ALIAS.

## ❓ Questions
* What data is present?
* If you delete this file does Bosh DNS still work on this machine?
* What is updating this file? (hint: read the [Bosh DNS
  Docs](https://bosh.io/docs/dns/))

## Resource
* [Bosh DNS Docs](https://bosh.io/docs/dns/)
