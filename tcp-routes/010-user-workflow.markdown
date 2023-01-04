---
layout: single
title: Using TCP Routes
permalink: /tcp-routes/user-workflow
sidebar:
  title: "TCP Routes"
  nav: sidebar-tcp-routes
---

## Assumptions
- You have a CF deployed
- You have one TCP Router deployed. (Check by running `bosh vms` and looking
  for a VM called tcp-router). If you have more than one that's okay, but the
  steps written below assume that there is only one.
- This story assumes that you are using GCP and have access to the account
  where your CF is deployed. You can still complete this story if this is not
  true, but any differences will be left as an exercise to the reader.

## What
GoRouter only handles incoming HTTP traffic.  All of the routes that we have
been talking about so far are *HTTP* routes (though we often drop the HTTP in
conversation and just call them routes). If you want to send og TCP traffic,
then you are going to need to set up a TCP route.

There is a parallel system for TCP routes similar to HTTP routes:
- An HTTP client connects to an HTTP route on an HTTP domain, though an HTTP
  load balancer, which sends traffic to HTTP Routers (GoRouters).
- TCP client connects to a TCP route on a TCP domain, through a TCP load
  balancer, which sends traffic to TCP Routers.

   ```
    +-----------+         +------------------+        |HTTP Router|        +------+
    |HTTP Client| ----->  |HTTP Load Balancer| -----> |(GoRouter) | -----> |      |
    +-----------+         +------------------+        +-----------+        |      |
                                                                           | App  |
    +-----------+         +------------------+        +-----------+        |      |
    |TCP Client | ----->  |TCP Load Balancer | -----> |TCP Router | -----> |      |
    +-----------+         +------------------+        +-----------+        +------+
   ```

Let's create a TCP Route and send traffic to it!

## How

📝 **Prep in Google Cloud Console**
1. Make sure that you have DNS set up properly for the TCP Router.  In Google
   Cloud Console, go to the Zone Details for your env (Network Services -->
   Cloud DNS --> <your-env>-zone)
2. You should find a domain that starts with `tcp`, let's call this TCP_DOMAIN.
   This domain should have an IP next to it, let's call this
   TCP_LOAD_BALANCER_IP.

   ![example TCP domain DNS on GCP](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/example-tcp-domain-dns.png)

3. In Google Cloud Console, find the Load Balancer with the ip
   TCP_LOAD_BALANCER_IP. (Network Services --> Load balancing) Here you will be
   able to see all of the VMs that the load balancer ...balances load between.
   In the example below, and most likely in your case, there is only one TCP
   Router deployed, so there will only be one VM listed.

   ![example TCP load balancer on GCP](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/example-tcp-load-balancer.png)

4. Click the VM instance that the TCP load balancer sends traffic to. Find the
   VM's internal IP. Let's call this TCP_ROUTER_IP.  ![example TCP router vm on GCP](https://storage.googleapis.com/cf-networking-onboarding-images-owned-by-ameowlia/example-tcp-router-details.png)
5. In the terminal, check that the TCP_ROUTER_IP matches the IP that bosh
   reports for the TCP Router. It comes full circle!

📝 **Push a TCP server app**
1. Push [this tcp listener app](https://github.com/cloudfoundry/cf-acceptance-tests/tree/master/assets/tcp-listener)
   with no HTTP route.  This app is listening for TCP traffic on port 8080 and
   it logs all of the messages sent to it.
   ```bash
   cf push tcp-app --no-route
   ```

📝 **Create a TCP Route**
1. See that you have a default-tcp router group. Router Groups are used to reserve ports for tcp routes.
   ```bash
   cf router-groups
   ```
1. Create a shared TCP domain
   ```bash
   cf create-shared-domain TCP_DOMAIN --router-group default-tcp
   ```
1. See that `cf map-route --help` has different usage instructions for TCP routes and HTTP routes.
1. Create a route with the TCP domain and map it to tcp-app, let's call this TCP_ROUTE:TCP_PORT.
   ```bash
   cf map-route tcp-app TCP_DOMAIN --random-port
   ```

📝 **Test with curl**

Curl sends traffic via HTTP(S), but because HTTP is built on top of TCP, we can
still use curl to test out TCP route.
1. In one terminal, run `cf logs tcp-app`
1. In another terminal `curl TCP_ROUTE:TCP_PORT`
1. See the HTTP headers show up in your logs

🤔 **Test with netcat**

Netcat is a helpful utility for reading and writing traffic over TCP or UDP.
You will use netcat to send tcp traffic.
1. In one terminal, run `cf logs tcp-app`
1. In another terminal, run a docker container and get some helpful tools
   ```
   docker run --privileged -it ubuntu bin/bash
   apt-get update -y
   apt-get install netcat -y
   ```
1. Use `nc -h` to look at the help text and figure out how to connect to TCP_ROUTE:TCP_PORT.
1. If you successfully open a connection, it will hold it open so you can type anything. Mash some keys and press enter.
1. See your key mashes in the app logs. You sent that via TCP!

Don't delete your TCP app/route/domain yet! You'll need them in the next stories.

## Extra Credit
What happens if you try to send TCP traffic to an HTTP route? Why can you send
HTTP traffic (kind of) over TCP, but not the other way around?

## Resources

* [Basic Golang TCP Server](https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go)
* [CF Docs - configure TCP domain](https://docs.cloudfoundry.org/adminguide/enabling-tcp-routing.html#-configure-cf-with-your-tcp-domain)
* [CF Docs - HTTP vs TCP routes](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html#-http-vs.-tcp-routes)
* [CF Docs - create TCP routes](https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html#-create-a-tcp-route-with-a-port)
* [Sending tcp traffic via netcat](https://askubuntu.com/questions/443227/sending-a-simple-tcp-message-using-netcat)
* [netcat fun! by julia evans](https://jvns.ca/blog/2013/10/01/day-2-netcat-fun/)
* [Did you brew install nc on your mac and it broke bosh? yup.](https://github.com/cloudfoundry/bosh-cli/pull/403)
