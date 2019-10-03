Introducing Route Registrar

⚠️⚠️⚠️
**WARNING: The Route Registrar stories are very new. ** You will probably find some typos. You will probably find some things that aren't as clear as they could be.
Please open a [PR or an issue](https://github.com/pivotal/cf-networking-program-onboarding) when you find these problems.
⚠️⚠️⚠️

## What?

Cloud Foundry components (like CAPI, Diego, and Policy Server) are deployed into a private network. This means they can only be accessed within their network. But what if you want to make an endpoint to a CF component available outside of Cloud Foundry?

**Route Registrar** makes Cloud Foundry components available outside of Cloud Foundry. These routes get registered with the GoRouter, just like app routes.
Or, if you like analogies:  ```Route Registrar:CF Components::cf map-route:CF Apps```

In this Route Registrar series of stories you are going to create your own instance group to run a HTTP server. Then you are going to use route registrar to map a route to the server.


## Vocab 💬

**Off-Platform** - Anything not in a Cloud Foundry deployment. This term is usually used when talking about where traffic originates. For example, traffic from Wendy (the end user) is off-platform traffic. Your local machine is off platform.

**On-Platform** - Anything that is within a Cloud Foundry deployment. This term is usually used when talking about where traffic originates. For example, when CAPI sends information to Diego, this is on-platform traffic.

**App Routes** - These are routes that resolve to a CF app.

**Component Routes** - These are routes that resolve to a bosh component. For example, uaa.beanie.c2c.cf-app.com is a component route that resolves to UAA.

## Notes
- This set of stories uses the instance group `my-http-server`, which you will create. It is handy to use this VM with nearly nothing on it so that there is much less traffic coming/going to it. However, all of this work _could_ be done on any VM.

L: route-registrar
---

User Workflow - Life Without Route Registrar

## Assumptions
- You have an OSS CF deployed

## What?

In this story you are going to run your own HTTP server on a new bosh instance group _without_ creating a route via Route Registrar. This story will show why Route Registrar is needed.

## How?

📝 **Create your own instance group with no route registrar routes**

0. Deploy your own instance group by adding something like this to your bosh manifest

```
instance_groups:
- name: my-http-server
  azs:
  - z1
  instances: 2       # <------------ Make sure you have two instances for load balancing
  jobs:
  - name: route_registrar
    properties:
      route_registrar:
        routes: []   # <------------ No routes to start with
    release: routing
  networks:
  - name: default
  stemcell: default
  update:
    serial: true
  vm_type: minimal
```

📝 **Run an HTTP server on your new VMs**
0. Copy the http server code from [this gist](https://gist.github.com/ameowlia/2768de0c1d857a9981ed2df9809de6a9) onto your local machine.
0. Look at the file. It is a small go program that starts an HTTP server on port 9994 that responds to any request with a friendly hello and the mac address of the machine responding.
0.  Compile and copy this file onto the new instance group VMs.
 ```
 GOOS=linux go build go-server.go                                   # <----- compile the golang server
 bosh scp go-server my-http-server:/tmp/go-server                   # <----- copy the compiled server to both instances of my-http-server
 bosh ssh my-http-server -c "sudo mv /tmp/go-server /bin/go-server" # <----- move the compiled server to the /bin/ directory on both instances of my-http-server
 ```
0.  In one terminal, ssh onto `my-http-server/0`, become root, and run the server.
0.  In a second terminal, ssh onto `my-http-server/1`, become root, and run the server.

📝 **Try to hit the http sever from your local machine**

0.  In a third terminal from your local machine, run `bosh is` and record the IPs for both instances of my-http-server. Let's call these MY_HTTP_SERVER_0_IP and MY_HTTP_SERVER_1_IP.

0. Still in the third terminal, try to `curl MY_HTTP_SERVER_0_IP:9994` .

0. Try to `curl MY_HTTP_SERVER_1_IP:9994`.

❓What happens? Why can't you reach these endpoints?

📝 **Try to hit the http sever from within the private CF network**

0. In the third terminal, bosh ssh onto any VM other than my-http-server.

0. Try to `curl MY_HTTP_SERVER_0_IP:9994`.

0. Try to `curl MY_HTTP_SERVER_1_IP:9994`.

❓Why can you reach these endpoints?


### Expected Results

MY_HTTP_SERVER_0_IP and MY_HTTP_SERVER_1_IP are both within CF's private network. This means that those IPs are only accessible from within the private network. You should be able to hit the HTTP server from any other VM in your CF deployment. You should not be able to hit the HTTP server from your local machine.

L:route-registrar
---

User Workflow - Life With Route Registrar

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have two instances of my-http-server deployed from the previous story.
- There is an HTTP server on both instances of my-http-server from the previous story.

## Recorded values from previous stories
```
MY_HTTP_SERVER_0_IP=<value>
MY_HTTP_SERVER_1_IP=<value>
```
## What?

Let's get that HTTP server accessible from off-platform! In this story you are going to add a route to my-http-server with route registrar. You will learn how route registrar load balances requests.

## How?

📝 **Update the instance group to include route registrar routes**

1. Run `cf domains` to find out the SYSTEM_DOMAIN for your deployment
2. Update your bosh manifest to add routes via route registrar. Redeploy.

```
instance_groups:
-  name: my-http-server
  azs:
  - z1
  instances: 2
  jobs:
  - name: route_registrar
    properties:
      route_registrar:
        routes:
        - name: meow-route               # <<< Add this new stuff to routes
          port: 9994                     # <<<
          registration_interval: 10s     # <<<
          uris:                          # <<<
          - meow.SYSTEM_DOMAIN           # <<< Make sure to replace SYSTEM_DOMAIN. You can also replace meow if you want. But why would you?
    release: routing
  networks:
  - name: default
  stemcell: default
  update:
    serial: true
  vm_type: minimal
```

🤔**Run the HTTP server on both instances of my-http-server**
1. Look back at the story `Life Without Route Registrar` if you need help with this.

🤔**Hit the route**
 1. Curl the route you created (with no port) from your local terminal.

You should see something like...
```Hello from machine with mac address 42:01:0a:00:01:14```

❓Curl the route a couple more times. Does the mac address change? Why or why not?


### Expected results
You should see the mac address load balance evenly between the two instances of my-http-server. If you are only seeing one mac address, you might not have both servers running successfully.

L:route-registrar
---

Route Registrar - Find Those Routes!

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have my-http-server deployed with route registrar routes setup from the previous story

## What?

In this story you are going to follow how route registrar registers routes (say that 5 times fast).

In the HTTP Routes section you learned how the Route Emitter on the Diego Cell repeatedly sends route registration messages to NATS. Then GoRouter subscribes to those NATS messages and then populates its route table. The same thing happens with component routes. But instead of the Route Emitter emitting routes (teehehe), it's Route Registrar that is emitting routes repeatedly.

## How?

🤔 **Look at NATS**

0. Subscribe to the the NATS messages for your component route from the my-http-server VM.
 - You can find the NATS username, password, and host on the my-http-server VM at `/var/vcap/jobs/route_registrar/config/registrar_settings.json`
 - See the story `Route Propagation - Part 3 - Route Emitter and NATS` if you help.

 ❓How do the component route NATS messages compare to the app route NATS messages?


🤔**Look at the routes table**
0. Bosh ssh onto the router VM.
0. Look at the GoRouter routes table and find your component route.
 - See the story `Route Propagation - Part 4 - GoRouter` if you need a reminder on how to this.

 ❓How does the component route compare to the app routes?


### Bonus Question

❓❓❓So if GoRouter routes off-platform users to other components, how do off-platform users route to GoRouter?!?!


### Expected Result
There are (almost) no differences between the app routes and the component routes. GoRouter does not know the difference between them and treats them the same.

L:route-registrar
---

User Workflow - Instance Specific Routes

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have two instances of my-http-server deployed. (See story `Life Without Route Registrar` for help if needed).
- There is an HTTP server on both instances of my-http-server. (See story `Life Without Route Registrar` for help if needed).

## What?

In the `Life With Route Registrar` story you got the component route working and load balancing between two instances of my-http-server. But what if you want to be able to target a specific instance?

In this story you are going to create instance specific component routes.

## How

🤔**Make and use instance specific routes**

0. Update the [prepend_instance_index property](https://github.com/cloudfoundry/routing-release/blob/develop/jobs/route_registrar/spec#L95-L96) in your bosh manifest to turn on instance specific routing.

0. Redeploy

0. Use the new routes!

0. Prove that you are hitting only one instance and that you can choose which instance you are hitting.


🤔**Check the gorouter routing table**

0. Look at the gorouter routes table and find your instance component routes.

 ❓How do these routes differ from the route you saw in the `Life With Route Registrar` story?

L:route-registrar
---

[RELEASE] Route Registrar ⇧
L: route-registrar
