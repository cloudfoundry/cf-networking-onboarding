---
layout: single
title: Add a Custom Bosh Alias
permalink: /bosh-dns/custom-alias
sidebar:
  title: "Bosh DNS"
  nav: sidebar-bosh-dns
---

## Assumptions
- You have 2 my-http-server instances deployed (see the story [Life Without
  Route Registrar](../route-registrar/life-without-rr) for setup)

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
```

## What
In this story you are going to add your own fun alias for your go HTTP server.

## How

📝 **Add your own alias**

1. Update your manifest to include a Bosh DNS alias. This alias could be added
   for any job on the instance group.

    ```
    - name: my-http-server
    # ...
    jobs:
    - name: route_registrar
      provides:                             # < ------------ Add this block to add a Bosh DNS alias
        my_custom_link:                     # < ------------
          aliases:                          # < ------------
          - domain: "meow.meow"             # < ------------ Make the domain anything you want :D
            health_filter: "healthy"        # < ------------ Record the domain you choose as HTTP_SERVER_ALIAS
      custom_provider_definitions:          # < ------------
      - name: my_custom_link                # < ------------
        type: my_custom_link_type           # < ------------
    ```

1. Redeploy

1. Make sure the go server is running on both of your my-http-server VMs. See
   the story [Life Without Route Registrar](../route-registrar/life-without-rr)
   for help with this step.

1. Bosh ssh onto any machine _except_ the my-http-server VM.

1. Wait a couple minutes...

1. Try to access your new URL! Success!

    ```
    $ curl HTTP_SERVER_ALIAS:9994

    Hello from machine with mac address 42:01:0a:00:01:16
    ```

1. Try to access your new URL from your local machine.

    ```
    $ curl HTTP_SERVER_ALIAS:9994

    curl: (6) Could not resolve host: meow.meow
    ```

## Expected Results

Your new alias should only be accessible within Cloud Foundry and not from your
local machine.
