---
layout: single
title: Life with Route Registrar
permalink: /route-registrar/user-workflow
sidebar:
  title: "Route Registrar"
  nav: sidebar-route-registrar
---

## Assumptions
- You have a CF deployed.
- You have two instances of my-http-server deployed from the previous story.
- There is an HTTP server on both instances of my-http-server from the previous
  story.

## Recorded values from previous stories
```
MY_HTTP_SERVER_0_IP=<value>
MY_HTTP_SERVER_1_IP=<value>
```
## What

Let's get that HTTP server accessible from off-platform! In this story you are
going to add a route to my-http-server with route registrar. You will learn how
route registrar load balances requests.

## How

📝 **Update the instance group to include route registrar routes**

1. Run `cf domains` to find out the SYSTEM_DOMAIN for your deployment
2. Update your bosh manifest to add routes via route registrar. Redeploy.

```
instance_groups:
- azs:
  - z1
  instances: 2
  jobs:
  - name: route_registrar
    properties:
      nats:
        tls:
          client_cert: ((nats_client_cert.certificate))
          client_key: ((nats_client_cert.private_key))
          enabled: true
      route_registrar:
        routes:
        - name: meow-route               # <<< Add this new stuff to routes
          port: 9994                     # <<<
          registration_interval: 10s     # <<<
          uris:                          # <<<
          - meow.SYSTEM_DOMAIN           # <<< Make sure to replace SYSTEM_DOMAIN. You can also replace meow if you want. But why would you?
    release: routing
  name: my-http-server
  networks:
  - name: default
  stemcell: default
  update:
    serial: true
  vm_type: minimal
```

🤔 **Run the HTTP server on both instances of my-http-server**
1. Look back at the story `Life Without Route Registrar` if you need help with this.

🤔 **Hit the route**
1. Curl the route you created (with no port) from your local terminal.
  You should see something like...
  ```
  Hello from machine with mac address 42:01:0a:00:01:14
  ```
1. Curl the route a couple more times.
* ❓ Does the mac address change? Why or why not?

## Expected results
You should see the mac address load balance evenly between the two instances of
my-http-server. If you are only seeing one mac address, you might not have both
servers running successfully.
