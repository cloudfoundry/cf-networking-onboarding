---
layout: single
title: Break Things
permalink: /service-discovery/break-things
sidebar:
  title: "C2C Service Discovery"
  nav: sidebar-service-discovery
---

## What
In the previous story we talked about how Bosh DNS redirects all DNS lookups
where the request matches an internal domain to the Bosh DNS Adapter. In this
story we are going to exploit this.

![href](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/architecture-diagram.png?raw=true)

## How

😇 **Pretend you are innocent user1**
1. Push [proxy
   app](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
   and call it appA.

1. Make sure appA has an http route, let's call it APPA_ROUTE.

1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA
   to neopets.com
   ```bash
   watch "curl -sS APPA_ROUTE/proxy/neopets.com"
   ```
You should get back some html for neopets! Fun! :D

😈 **Now pretend you are malicious user2**

  As a malicious actor, you know that appA is sending traffic to neopets.com.
  You want to break their app and make it so appA can't reach neopets. You can
  do this by shadowing the neopets.com domain with an internal domain.

1. Create the internal domain `neopets.com` (look at `cf create-shared-domain
   --help` if you don't remember how)

  Instead of adding this new internal domain to the bosh manifest and
  redeploying, we are going to hack this in on the Diego Cell. This is a great,
  fast, (and dangerous) debugging technique. It should be used with heavy
  caution, but it is often the fastest way to change a bosh property if things
  are really bad and need to be fixed immediately, or if you are just experimenting and impatient.

1. Ssh onto the Diego Cell where appA is running and become root.

1. You need to find what config file holds the information you want to change.
   The config on the VMs does not directly match the bosh manifest. To find any
   config for any bosh job in CF go to `/var/vcap/jobs/JOB_NAME`. There are
   many files there for Bosh DNS Adapter. In order to figure out exactly what
   you need to change, look at [Bosh DNS Job in the release
   code](https://github.com/cloudfoundry/cf-networking-release/tree/develop/jobs/bosh-dns-adapter)
   and look where the `internal_domains` property is used.
   [hint](https://github.com/cloudfoundry/cf-networking-release/blob/develop/jobs/bosh-dns-adapter/templates/handlers.json.erb#L11).
   [hint](https://github.com/cloudfoundry/cf-networking-release/blob/develop/jobs/bosh-dns-adapter/spec#L10).

1. Edit the correct config file and add an entry for our new domain
   `neopets.com`. It should look exactly like `apps.internal` except for the
   name.

1. Now you'll need to restart the Bosh DNS Adapter process so that it will run
   with our new config. Linux Bosh VMs use monit as a process manager. Run
   `monit summary` to see all of the processes running on this VM. Restart the
   Bosh DNS Adapter by running `monit restart bosh-dns-adapter`. Keep running
   `monit summary` until the Bosh DNS Adapter is successfully running again. If
   it fails to start, then you probably made a syntax error in the config file.
   Look at the logs and fix the error.

😇 **Back to pretending you are innocent user1**
1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA
   to neopets.com `watch "curl -sS APPA_ROUTE/proxy/neopets.com"` It should
   still show neopets. Why isn't it broken yet!?

😈 **Now pretend you are malicious user2**

You restarted the Bosh DNS Adapter, but look at the diagram again. It's
actually Bosh DNS _not_ Bosh DNS adapter that does the hairpinning for internal
domains.  You had to restart Bosh DNS Adapter process so it would run with the
new config file, but you also need to restart the Bosh DNS process, so it can
get these new values from the Bosh DNS Adapter.

1. Restart Bosh DNS
   ```bash
   monit restart bosh-dns
   ```

😇 **Back to pretending you are innocent user1**
1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA
   to neopets.com
   ```bash
   watch "curl -sS APPA_ROUTE/proxy/neopets.com"
   ```
   Where did neopets go???

## Expected Result
AppA should no longer be able to access neopets. :(

## ❓ Questions

1. Neopets is a silly example. What is a worse example that customers could run
   into?
1. What permissions does a user require to exploit this?
1. Would you consider this a security concern?
1. Would your assessment change if internal domains were dynamic and didn't
   require setting a bosh property at deploy time? Why or why not?
1. What will happen to your config changes if you redeployed your environment?

## Resource
* [configuring internal domain
  docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains)

