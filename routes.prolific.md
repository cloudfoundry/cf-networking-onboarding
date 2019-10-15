User Workflow: HTTP Routes

## Assumptions
- You have a CF deployed
- You have two [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) apps pushed, named appA and appB

## What?

**Routes** are the URLs that can be used to access CF apps.
**Route Mappings** are the join table between routes (URLs) and the apps they send traffic to. Apps can have many routes.
And routes can send traffic to many apps. So Route Mappings is a many-to-many mapping.

## How?

📝 **Create a route that maps to two apps**
0. By default `cf push` creates a route. Look at all of the routes for all of your apps.
 ```
 cf apps
 ```
0. Use curl to hit appA.
 ```
 curl APP_A_URL
 ```
 It should respond with something like `{"ListenAddresses":["127.0.0.1","10.255.116.44"],"Port":8080}`
 We'll get into the listen addresses later, but for now the most important thing to know is that the 10.255.X.X address is the overlay IP address. This IP is unique per app instance.
0. Create your own route (`cf map-route --help`) and map it to both appA **AND** appB.
0. Curl your new route over and over again `watch "curl -sS MY-NEW-ROUTE`".

### Expected Result
You have a route that maps to both appA and appB. See that the overlay IP changes, showing that you are routed evenly(ish) between all the apps mapped to the routes.

L: http-routes
L: user-workflow
---

Route Propagation - Part 0.1 - Creating a Route Overview

## What
In the previous story you used the CF CLI to create and map routes. But what happens under the hood to make all of this work? (hint: iptables is involved)
There are two main data flows for routes, (1) when an app dev pushes a new app with a route and (2) when an internet user connects to an app using the route.

Let's focus on the first case. Here is what happens under the hood:

Each step marked with a ✨ will be explained in more detail in a story in this track.
**When an app dev pushes a new app with a route**
1. ✨ The app dev pushes an app with a route using CAPI.
1. CAPI sends this information to Diego.
1. Diego schedules the container create on a specific Diego Cell.
1. Garden creates the container for your app.
1. ✨ Diego deploys a sidecar envoy inside of the app container, which will proxy traffic to your app.
1. ✨ When the container is being set up, iptables rules are created on the Diego Cell to send traffic that is intended for the app to the sidecar proxy.
1. ✨ When the app is created, Diego sends the route information to the Route Emitter. The Route Emitter sends the route information to GoRouter via NATS.
1. ✨ The GoRouter keeps a mapping of routes -> ip:ports in a routes table, which is consulted when someone curls the route.

## How
The following stories will look at how many components (CAPI, Diego BBS, Route Emitter, Nats, GoRouter, DNAT Rules, Envoy) work together to make routes work.

0. 🤔 Step through steps above and follow along on [the HTTP Routing section of this diagram](https://realtimeboard.com/app/board/o9J_kyWPVPM=/).

### Expected Result
You can talk about route propagation at a high level.

## Logistics
In the next few stories, you are going to need to remember values from one story to another, there will be a space provided at the bottom of each story for your to record these values so you can store them.
It can be annoying to scroll up and down in the story as you use the values, so it could be helpful to store these values in a doc outside of tracker.

## Resources for the entire route propagation track
**Capi**
[CAPI API docs](https://apidocs.cloudfoundry.org/7.8.0/)
[CAPI V3 API docs](http://v3-apidocs.cloudfoundry.org/version/3.70.0/index.html)

**Diego**
[cfdot docs](https://github.com/cloudfoundry/cfdot)
[diego design notes](https://github.com/cloudfoundry/diego-design-notes#what-are-all-these-repos-and-what-do-they-do)
[diego bbs API docs](https://github.com/cloudfoundry/bbs/tree/master/doc)

**NATs**
[NATS message bus repo](https://github.com/nats-io/gnatsd)
[NATS ruby gem repo](https://github.com/nats-io/ruby-nats)

**GoRouter**
[GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)
[Detailed Diagram of several Route Related User Flows](https://realtimeboard.com/app/board/o9J_kyWPVPM=/)

**Iptables**
[iptables man page](http://ipset.netfilter.org/iptables.man.html)
[Aidan's iptables in CF ppt](https://docs.google.com/presentation/d/1qLkNu633yLHP5_S_OqOIIBETJpW3erk2QuGSVo71_oY/edit#slide=id.p)

**Route Integrity**
[Route Integrity/Misrouting Docs](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting)

**Envoy**
[What is Envoy?](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy)


L: http-routes
---

Route Propagation - Part 0.2 - HTTP Traffic Overview

## What
In the previous story you went over the first of two main data flows for routes:
(1) when an app dev pushes a new app with a route and (2) when an internet user connects to an app using the route.

Here is what happens under the hood for the second case:

Each step marked with a ✨ will be explained in more detail in a story in this track.
**When an internet user sends traffic to your app**
1. The user visits your route in the browser or curls it via the command line.
1. The traffic first hits a load balancer in front of the CF Foundation.
1. The load balancer sends it to one of the GoRouters.
1. ✨ The GoRouter consults the route table and sends it to the listed IP and port. If Route Integrity is enabled, it sends this traffic via TLS.
1. ✨ The traffic makes its way to the correct Diego Cell, where it hits iptables DNAT rules that reroutes the traffic to the sidecar envoy for the app.
1. ✨ The Envoy terminates the TLS from the GoRouter and then sends the traffic on to the app.

## How
The following stories will look at how many components (CAPI, Diego BBS, Route Emitter, Nats, GoRouter, DNAT Rules, Envoy) work together to make routes work.

0. 🤔 Step through steps above and follow along on [the HTTP Routing section of this diagram](https://realtimeboard.com/app/board/o9J_kyWPVPM=/)

### Expected Result
You can talk about HTTP network traffic flow with fellow CF engineers.

## Logistics
In the next few stories, you are going to need to remember values from one story to another, there will be a space provided at the bottom of each story for your to record these values so you can store them.

## Resources for the entire route propagation track
**Capi**
[CAPI API docs](https://apidocs.cloudfoundry.org/7.8.0/)
[CAPI V3 API docs](http://v3-apidocs.cloudfoundry.org/version/3.70.0/index.html)

**Diego**
[cfdot docs](https://github.com/cloudfoundry/cfdot)
[diego design notes](https://github.com/cloudfoundry/diego-design-notes#what-are-all-these-repos-and-what-do-they-do)
[diego bbs API docs](https://github.com/cloudfoundry/bbs/tree/master/doc)

**NATs**
[NATS message bus repo](https://github.com/nats-io/gnatsd)
[NATS ruby gem repo](https://github.com/nats-io/ruby-nats)

**GoRouter**
[GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)
[Detailed Diagram of several Route Related User Flows](https://realtimeboard.com/app/board/o9J_kyWPVPM=/)

**Iptables**
[iptables man page](http://ipset.netfilter.org/iptables.man.html)
[Aidan's iptables in CF ppt](https://docs.google.com/presentation/d/1qLkNu633yLHP5_S_OqOIIBETJpW3erk2QuGSVo71_oY/edit#slide=id.p)

**Route Integrity**
[Route Integrity/Misrouting Docs](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting)

**Envoy**
[What is Envoy?](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy)


L: http-routes
---

Route Propagation - Part 1 - CAPI

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- I recommend deleting all other apps

## What
The Cloud Controller API (CAPI) maintains the database of all apps, domains, routes, and route mappings.
However CAPI does not keep track of *where* those apps are deployed. Nor does CAPI track the IPs and ports each
route should, well, *route* to. That's the job of the router, often called GoRouter.

CAPI keeps track of the desired state. The user wants a route called MY_ROUTE that sends traffic to appA.
But the user doesn't (shouldn't) care about the logistics needed to make that route happen. That is the responsibility of other components.

Let's look at what information CAPI *does* keep track of.

## How

0. 🤔 Map a route to appA. Let's call this route APP_A_ROUTE. I recommend _deleting_ all other routes.

0. 🤔 Look at the domains, routes, route mappings, and apps in CAPI's database.
    To look at all the domains you can curl CAPI using `cf curl /v2/domains`. Use the [CAPI docs](https://apidocs.cloudfoundry.org/7.8.0/) to figure out the APIs for the other resources.

This is all of the information that CAPI contains about routes. Note there are no IPs anywhere. Note that all of these routes are for CF apps, none of them are for CF components.

### Expected Result
You can view data from CAPI about the route APP_A_ROUTE that you created.

## Recorded Values
Record the following values that you generated or discovered during this story.
```
APP_A_ROUTE=<value>
```

## Resources
[CAPI API docs](https://apidocs.cloudfoundry.org/7.8.0/)


L: http-routes
---

Route Propagation - Part 2 - Diego BBS

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## What

**Diego** is an umbrella term for many components that work together to make CF container orchestration happen. These components are maintained by the Diego team.

**BBS** stands for Bulletin Board System. This is the database that the Diego components use to keep track of DesiredLRPs and ActualLRPs.

**LRPs**, or Long Running Processes, represent work that a client (ie CAPI) wants Diego to keep running indefinitely. Apps are the primary example of LRPs. Diego will try to
keep them running as best it can. When an app stops or fails, it will attempt to restart it and keep it running.

**Desired LRPs** represent what a client (ie CAPI) wants "the world" to look like (for example, how many instances of which apps). They contain no information
about where LRPs should be run, because the user shouldn't care.

**Actual LRPs** represent what Diego is currently *actually* running. Actual LRPs contain information about which Diego Cell the LRP is running on and which port maps to the LRP.

For this story, let's look at the data stored in the BBS and see what information it has about appA. Diego will go on to send this information via the Route Emitter to GoRouter, so GoRouter knows where to send network traffic to.

## How

📝 **Look at actualLRPS**
0. Grab the guid for appA. You'll need it in a moment. Let's call it APP_A_GUID.
 ```
 cf app appA --guid
 ```
0. Ssh onto the Diego Cell vm where appA is running and become root.
0. Use the [cfdot CLI](https://github.com/cloudfoundry/cfdot) to query BBS for actualLRPs. Cfdot is a helpful CLI for using the BBS API.
 It's a great tool for debugging on the Diego Cell.
 ```
 cfdot actual-lrps | jq .
 ```
0. Search through the actual LRPs for APP_A_GUID. It should match the beginning of a process guid. You'll find an entry for each instance of appA that is running.
0. Let's dissect and store the most important information (for us) about appA:
   ```
   {
     "process_guid": "ab2bd185-9d9a-4628-9cd8-626649ec5432-cb50adac-6861-4f03-92e4-9fcc1a204a1e",
     "index": 0,
     "cell_id": "d8d4f5fe-36f2-4f50-8c4a-8df293f6bc5b",
     "address": "10.0.1.12",                  <------ The cell's IP address where this app instance is running, also sometimes called the host IP. Let's call this DIEGO_CELL_IP.
       "ports": [
         {
           "container_port": 8080,            <------ The port the app is listening on inside of its container. 8080 is the default value. Let's call this CONTAINER_APP_PORT.
           "host_port": 61012,                <------ The port on the Diego Cell where traffic to your app is sent to before it is forwarded to the overlay address and the container_port. Let's call this DIEGO_CELL_APP_PORT.
           "container_tls_proxy_port": 61001, <------ The port inside of the app container that envoy is listening on for HTTPS traffic. This is the default value (currently unchangeable). Let's call this CONTAINER_ENVOY_PORT.
           "host_tls_proxy_port": 61014,      <------ The port on the Diego Cell where traffic to your app's envoy sidecar is sent to before it is forwarded to the overlay address and the container_tls_proxy_port. Let's call this DIEGO_CELL_ENVOY_PORT
         },
         {
           "container_port": 2222,            <------ The port exposed on the app container for sshing onto the app container
           "host_port": 61013,                <------ The port on the Diego Cell where ssh traffic to your app container is sent to before it is forwarded to the overlay address and the ssh container_port
           "container_tls_proxy_port": 61002, <------ The ssh port inside of the app container that envoy is listening on for ssh traffic. This is the default value (currently unchangeable).
           "host_tls_proxy_port": 61015       <------ The port on the Diego Cell where ssh traffic to your app's envoy sidecar is sent to before it is forwarded to the overlay address and the ssh container_tls_proxy_port
         }
       ],
     "instance_address": "10.255.116.6",      <------ The overlay IP address of this app instance, let's call this the OVERLAY_IP
     "state": "RUNNING",
      ...
   }
   ```
0. Use the cfdot CLI to query BBS for desiredLRPs.

❓What information is provided for desiredLRPs, but not for actualLRPs?
❓What information is provided for actualLRPs, but not for desiredLRPs?
❓How does this match with the definition of desired and actual LRPs in the "what" section above?

### Expected Result
Get information from BBS about the desiredLRP and actualLRP for appA. Use cfdot CLI to discover the following values and record them.

## Recorded Values
Record the following values that you generated or discovered during this story.
```
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## Resources
[cfdot docs](https://github.com/cloudfoundry/cfdot)
[diego design notes](https://github.com/cloudfoundry/diego-design-notes#what-are-all-these-repos-and-what-do-they-do)


L: http-routes
L: questions
---

Route Propagation - Part 3 - Route Emitter and NATS

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What

There is one Route Emitter per Diego Cell and its job is to... emit routes. According to the ever helpful
[Diego Design Notes](https://github.com/cloudfoundry/diego-design-notes) the Route Emitter "monitors DesiredLRP state
and ActualLRP state via the BBS. When a change is detected, the Route Emitter emits route registration and unregistration
messages to the GoRouter via the NATS message bus." Even when no change is detected, the Route Emitter will periodically
emit the entire routing table as a kind of heartbeat.

For this story, let's look at the messages that the Route Emitter is publishing via NATS. Subscribing to these NATs messages
can be a helpful debugging technique.

## How

📝 **subscribe to NATs messages**
0. Bosh ssh onto the Diego Cell where your app is running and become root
0. Download ruby and the NATS gem
    ```
    apt-get install -y ruby ruby-dev
    gem install nats
    ```
0. Get NATS username, password, and server address
    ```
    cat /var/vcap/jobs/route_emitter/config/route_emitter.json | jq . | grep nats
    ```
0. Use the nats gem to connect to nats: `nats-sub "*.*" -s nats://NATS_USERNAME:NATS_PASSWORD@NATS_ADDRESS`. The `"*.*"` means that you are subscribing to all NATs messages.
    The Route Emitter registers routes every 20 seconds (by default) so that the GoRouter (which subscribes to these messages) has the most up-to-date information about which IPs map to which apps and routes. Depending on how many routes there are, this might be a lot of information.

0. Find the NATs message for APP_A_ROUTE.
 ```
 nats-sub "*.*" -s nats://NATS_USERNAME:NATS_PASSWORD@NATS_ADDRESS` | grep APP_A_ROUTE
 ```
 If you wait, you should see a message that contains information about the route you created. It will look something like this and contain APP_A_ROUTE:
 ```
   [#32] Received on [router.register] :
{
    "host": "10.0.1.12",
    "port": 61012,
    "tls_port": 61014,
    "uris": [
        "proxy.meow.cloche.c2c.cf-app.com"     <--- This should match APP_A_ROUTE
      ],
    "app": "6856799f-aebf-4e2b-81a5-28c74dfb6162",
     "private_instance_id": "a0d2b217-fa7d-4ac1-65a2-7b19",
     "private_instance_index": "0",
    "server_cert_domain_san": "a0d2b217-fa7d-4ac1-65a2-7b19",
    "tags": {
         "component": "route-emitter"
     }
}
 ```

❓Do the values in the NATS message match the values you recorded previously from BBS? Which ones are present? Which ones aren't there?
❓How does it compare to the information in CAPI?


### Expected Result
Inspect NATs messages. Look at what route information is sent to the GoRouter.

## Resources
[NATS message bus repo](https://github.com/nats-io/gnatsd)
[NATS ruby gem repo](https://github.com/nats-io/ruby-nats)

L: http-routes
L: questions
---

Route Propagation - Part 4 - GoRouter

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
So the Route Emitter emits routes via the NATS message Bus. GoRouter subscribes to those messages and keeps a route table that is uses to route network traffic bound for CF apps and CF components.

Let's take a look at that route table.
## How

📝 **look at route table**
0. Bosh ssh onto the router vm and become root.
0. Install jq (a json manipulation and display tool)
 ```
 apt-get install jq
 ```
0. Get the username and password for the routing api
 ```
 head /var/vcap/jobs/gorouter/config/gorouter.yml
 ```
0. Get the routes table
 ```
 curl "http://USERNAME:PASSWORD@localhost:8080/routes" | jq .
 ```
0. Scroll through and look at the routes.
  ❓How does this differ from the route information you saw in CAPI?
   For example, you should see routes for CF components, like UAA and doppler.
   This because the GoRouter is in charge of routing traffic to CF apps *AND* to CF components.
0. Find APP_A_ROUTE in the list of routes. Let's dissect the most important bits.
    ```
    "proxy.meow.cloche.c2c.cf-app.com": [   <------ The name of the route! This should match APP_A_ROUTE
        {
          "address": "10.0.1.12:61014",     <------ This is where GoRouter will send traffic for this route. This should match DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT
          "tls": true                       <------ This means Route Integrity is turned on, so the GoRouter will use send traffic to this app over TLS
        }
      ]
    ```
    See how the traffic is being sent to `10.0.1.12:61014` or DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT?
    This means all traffic is being sent to the sidecar envoy via TLS, this is because route integrity is enabled.
    ❓What port do you think would be listed here if route integrity was not enabled?

### Expected Result
Access the route table on the router vm. Inspect app routes and CF component routes.

See that the GoRouter sends traffic for this route to DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT.
In the next story we will see what what is listening on that port on the cell.

## Resources

[GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)

L: http-routes
L: questions
---

Route Propagation - Part 4.5 - see what's listening with netstat

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
Netstat is tool that can show information about network connections, routing tables, and network interface statistics.
In the previous story we saw that the GoRouter sent traffic for APP_A_ROUTE to DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT.
Let's use netstat to see what is listening at on the Diego Cell and specifically at DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT.

## How
📝 **look at open ports on a Diego Cell**
1. Ssh onto the Diego Cell where appA is deployed and become root.
1. Use netstat to look at open ports
 ```
 netstat -tulp
 # -t  <---- show tcp sockets
 # -u  <---- show udp sockets
 # -l  <---- display listening sockets
 # -p  <---- display PID/program name for sockets
 ```
  You should recognize the program names in the far right column. Most of them are the long running cf component processes.

1. Find the local address for the Route Emitter. What port is it running on? Does that match what is in the [spec file](https://github.com/cloudfoundry/diego-release/blob/develop/jobs/route_emitter/spec)?

1. Search for DIEGO_CELL_ENVOY_PORT in the output. Can you find it?

### Expected Result
You won't see the DIEGO_CELL_ENVOY_PORT anywhere in the netstat output because nothing is *actually* running there.
But if there's nothing running there, how does the traffic reach the app? Would you believe that iptables are involved?
Check out the next story to learn more :)

L: http-routes
---

Route Propagation - Part 5 - DNAT Rules

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track
- It will help if you have completed the "iptables-primer" track, but it is not required.

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
In a previous story you saw that the GoRouter sent traffic for APP_A_ROUTE to DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT.
But in the previous story, you saw nothing was actually listening at that port on the Diego Cell. So how does the network traffic hit the app?
With the help of iptables rules of course! Everything always comes back to iptables rules.

Nothing is actually listening on the Diego Cell at port DIEGO_CELL_ENVOY_PORT. Instead, all packets that are sent there hit iptables rules that then redirect them to... somewhere. Let's find out!

Let's not brute force looking through every iptables rule. Instead, let's reason about what chain and table most likely contain these iptables rules.
Hint, these rules translate ingress network traffic sent from GoRouter.

## How

🤔 **Find those iptables rules**
Aidan Obley made a great diagram showing the different types of traffic in CF and which iptables chains they hit in what order.
We are currently concerned with ingress traffic, which is represented by the orange line.

0. Look at the diagram. Which chain does the ingress traffic hit first?
    ![traffic-flow-through-iptables-on-cf-diagram](https://storage.googleapis.com/cf-networking-onboarding-images/traffic-flow-through-iptables-on-cf.png)

0. Based on the previous diagram, the ingress traffic hits the prerouting chain first. Look at the diagram below and do some research to learn more about the raw, conn_tracking, mangle, and nat tables.
    Which table should contain the rules to redirect our traffic to a new address?
    ![iptables tables and chains diagram](https://storage.googleapis.com/cf-networking-onboarding-images/iptables-tables-and-chains-diagram.png)

    NAT stands for Network Address Translation. That sounds like what we want.  So let's look at iptables rules for the nat table on the prerouting chain.

0. Ssh onto the Diego Cell where your app is running and become root.
0. Run `iptables -S -t nat`
    You should see some custom chains attached to the PREROUTING chain. There will be one custom chain per app running on this Diego Cell.  They will look something like this.
    ```
    -A PREROUTING -j netin--a0d2b217-fa7d-4ac1-65
    -A PREROUTING -j netin--317736ed-70ac-4087-74
    ...
    ```
0. You should also see 4 rules that contain the OVERLAY_IP for appA. If you look closely you'll see that the ports in the iptables rules match the ports we saw when inspecting the actual LRPs.
    Which port represents what?
    ```
    -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61012 -j DNAT --to-destination 10.255.116.6:8080
    -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61013 -j DNAT --to-destination 10.255.116.6:2222
    -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61014 -j DNAT --to-destination 10.255.116.6:61001
    -A netin--a0d2b217-fa7d-4ac1-65 -d 10.0.1.12/32 -p tcp -m tcp --dport 61015 -j DNAT --to-destination 10.255.116.6:61002
    ```

0. For appA, find the rule that will match with the traffic the GoRouter sends to DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT. It should look something like this...
    ![example DNAT rule with explanation](https://storage.googleapis.com/cf-networking-onboarding-images/example-DNAT-rule-with-explanation.png)

    In summary, when the GoRouter sends network traffic to 10.0.1.12:61014 (DIEGO_CELL_IP:DIEGO_CELL_ENVOY_PORT) it gets redirected to 10.255.116.6:61001 (OVERLAY_IP:CONTAINER_ENVOY_PORT).
    But, looking at the information we learned about the actual LRP, the app isn't even listening on 10.255.116.6:61001, envoy is.
    When will the traffic finally reach the app!?!?

### Expected Result
Inspect the iptables rules that DNAT the traffic from the GoRouter and send it to the correct sidecar envoy.

## Resources
[iptables man page](http://ipset.netfilter.org/iptables.man.html)
[Aidan's iptables in CF ppt](https://docs.google.com/presentation/d/1qLkNu633yLHP5_S_OqOIIBETJpW3erk2QuGSVo71_oY/edit#slide=id.p)

L: http-routes
---

Route Propagation - Part 6.1 - Envoy Primer

## Assumptions
- None

## What

Before you go further it will help if you read a quick primer on Envoy.

## How

1. 📚 Read Julia Evan's blog post ["Some Envoy Basics"](https://jvns.ca/blog/2018/10/27/envoy-basics/).
 Unfortunately, the envoy in the docker container referenced in the blog post doesn't work anymore. However, just reading the post is enough to get a nice overview.

L: http-routes
---

Route Propagation - Part 6.2 - Route Integrity and Envoy

## Assumptions
- You have a OSS CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and called appA
- You have one route mapped to appA called APP_A_ROUTE
- You have completed the previous stories in this track

## Recorded values from previous stories
```
APP_A_ROUTE=<value>
APP_A_GUID=<value>
DIEGO_CELL_IP=<value>
CONTAINER_APP_PORT=<value>
DIEGO_CELL_APP_PORT=<value>
CONTAINER_ENVOY_PORT=<value>
DIEGO_CELL_ENVOY_PORT=<value>
OVERLAY_IP=<value>
```

## What
A **proxy** is a process that sits in-between the client and the server and intercepts traffic before forwarding it on to the server. Proxies can add extra functionality, like caching or SSL termination.

In this case, Envoy is a sidecar proxy (Envoy can be other types of proxies too, but forget about that for now). The sidecar Envoy is only present when Route Integrity is turned on (which is done by default).

Route Integrity is when the GoRouter sends all app traffic via TLS. As part of the TLS handshake, the GoRouter validates the certificate's SAN against the ID found in its route table to make sure it is connecting to the intended app instance. This makes communication more secure and prevents stale routes in the route table from causing misrouting, which is a large security concern. Read more about how Route Integrity prevents misrouting [here](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting).

The Envoy sidecar is the process that actually terminates the TLS traffic from the GoRouter making Route Integrity possible. Then the Envoy proxies it onto...can it be? finally?! YES! THE APP!!

Let's look at how the Envoy sidecar is configured to proxy traffic to the app.
## How

📝 **Look at envoy config**
0. Ssh onto AppA
0. Hit the Envoy admin endpoint `curl localhost:61003/admin`
    These are all of the endpoints you can hit. Try `/clusters` what do you see?
0. Run `curl localhost:61003/config_dump`. This gives you all of the information about how the Envoy is configured.
0. Search for the CONTAINER_ENVOY_PORT, in the example it is 61001. This is where the DNAT rules forwarded the traffic to, as we saw in the last story. Find a listener called `listener-8080` that looks similar to the following: 
    ```
     "listeners": [
      {
       "name": "listener-8080",                                                    <---- The name of the listener
       "address": {
        "socket_address": { "address": "0.0.0.0", "port_value": 61001}             <---- This listener is listening on port 61001. That's the CONTAINER_ENVOY_PORT we know and love!
       },
       "filter_chains": [
        {
         "tls_context": { "require_client_certificate": true },                    <---- This means Route Integrity is turned on
         "filters": [
          {
           "name": "envoy.tcp_proxy",
           "config": { "stat_prefix": "0-stats", "cluster": "0-service-cluster" }  <---- This is the name of the cluster where Envoy will forward traffic that is sent to the CONTAINER_ENVOY_PORT, let's call this CLUSTER-NAME
          }
         ]
        }
       ]
      }
     ]
    ```
0. In the same config_dump output, find the cluster, CLUSTER-NAME, that is referenced above. It should look something like this:
    ```
     "clusters": [
      {
       "name": "0-service-cluster",                                          <---- This is the name of the cluster, CLUSTER-NAME
       "hosts": [
        {
         "socket_address": { "address": "10.255.116.6", "port_value": 8080}  <---- This is the port that the app is listening on inside of the container, should match OVERLAY_IP and CONTAINER_APP_PORT
        }
       ]
      }
    ]
    ```

So the traffic gets sent to the OVERLAY_IP:CONTAINER_ENVOY_PORT, then the envoy forwards it on to OVERLAY_IP:CONTAINER_APP_PORT!

We made it! We finally made it to the end! Everything is set up and someone can use that route you made!

### Expected Result
Look at the Envoy's 8080 listener and related cluster and see how network traffic is sent to the app.

## Resources
[Route Integrity/Misrouting Docs](https://docs.cloudfoundry.org/concepts/http-routing.html#-preventing-misrouting)
[What is Envoy?](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy)

L: http-routes
---

Route Propagation - Part 7 - Wrap Up

## Assumptions
- You have completed the previous stories in this track

## What
Let's review!

## How
1. Go to a whiteboard and have one pair diagram what happens when an app dev creates and maps a new route.
1. Go to a whiteboard and have the other pair diagram what happens when a person on the internet makes an HTTP request to a CF route.

### Expected Result
You know everything about routes. (Just kidding.)


L: http-routes

---

[RELEASE] HTTP Routes ⇧
L: http-routes