---
layout: single
title: Life Without Route Registrar
permalink: /route-registrar/life-without-rr
sidebar:
  title: "Route Registrar"
  nav: sidebar-route-registrar
---

## What

In this story you are going to run your own HTTP server on a new bosh instance
group _without_ creating a route via Route Registrar. This story will show why
Route Registrar is needed.

## How?

📝 **Create your own instance group with no route registrar routes**

1. Deploy your own instance group by adding something like this to your bosh manifest

```
instance_groups:
- azs:
  - z1
  instances: 2       # <------------ Make sure you have two instances for load balancing
  jobs:
  - name: route_registrar
    properties:
      nats:
        tls:
          client_cert: ((nats_client_cert.certificate))
          client_key: ((nats_client_cert.private_key))
          enabled: true
      route_registrar:
        routes: []   # <------------ No routes to start with
    release: routing
  name: my-http-server
  networks:
  - name: default
  stemcell: default
  update:
    serial: true
  vm_type: minimal
```

📝 **Run an HTTP server on your new VMs**
1. Copy the http server code from
   [this gist](https://gist.github.com/ameowlia/2768de0c1d857a9981ed2df9809de6a9)
   onto your local machine.
1. Look at the file. It is a small go program that starts an HTTP server on
   port 9994 that responds to any request with a friendly hello and the mac
   address of the machine responding.
1. Compile and copy this file onto the new instance group VMs.
   ```
   GOOS=linux go build go-server.go                                   # <----- compile the golang server
   bosh scp go-server my-http-server:/tmp/go-server                   # <----- copy the compiled server to both instances of my-http-server
   bosh ssh my-http-server -c "sudo mv /tmp/go-server /bin/go-server" # <----- move the compiled server to the /bin/ directory on both instances of my-http-server
   ```
1.  In one terminal, ssh onto `my-http-server/0`, become root, and run the server.
1.  In a second terminal, ssh onto `my-http-server/1`, become root, and run the server.

📝 **Try to hit the http server from your local machine**

1. In a third terminal from your local machine, run `bosh is` and record the
   IPs for both instances of my-http-server. Let's call these
   MY_HTTP_SERVER_0_IP and MY_HTTP_SERVER_1_IP.

1. Still in the third terminal, try to `curl MY_HTTP_SERVER_0_IP:9994` .

1. Try to `curl MY_HTTP_SERVER_1_IP:9994`.
* ❓ What happens? Why can't you reach these endpoints?

📝 **Try to hit the http server from within the private CF network**

1. In the third terminal, bosh ssh onto any VM other than my-http-server.

1. Try to `curl MY_HTTP_SERVER_0_IP:9994`.

1. Try to `curl MY_HTTP_SERVER_1_IP:9994`.
* ❓ Why can you reach these endpoints?

## Expected Results

MY_HTTP_SERVER_0_IP and MY_HTTP_SERVER_1_IP are both within CF's private
network. This means that those IPs are only accessible from within the private
network. You should be able to hit the HTTP server from any other VM in your CF
deployment. You should not be able to hit the HTTP server from your local
machine.
