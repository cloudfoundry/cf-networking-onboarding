---
layout: single
title: Can CF Apps Use Bosh DNS?
permalink: /bosh-dns/cf-apps
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
In the previous stories you learned how to set up and use a Bosh DNS alias to
curl the golang http server from other Bosh VM. But what about CF apps? Can
apps use Bosh DNS aliases?

## How

🤔 **Try your Bosh DNS alias from an app**

1. Push any app
1. Cf ssh onto that app
1. Curl HTTP_SERVER_ALIAS:9994
* ❓ Does it work? Why or why not? Are iptables involved? (hint: yes)

🤔 **Make your Bosh DNS alias available from an app**

1. Update your ASGs (application security groups) to make the alias available to your app.

## Expected Result

Most likely, the default security groups for your CF deployment do not give
your apps access to any private IP ranges. If you update the ASGs to allow apps
access to private IPs (specifically, the IPs for the my-http-server VMs) your
app should be able to use the alias.

If it is not working:
- make sure you restarted your app after applying the security group
- wait a couple minutes. The DNS records seem to take a little bit to propagate
  throughout the CF deployment.
