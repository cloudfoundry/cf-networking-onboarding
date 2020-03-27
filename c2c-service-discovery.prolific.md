User Flow: Container to Container Networking with Service Discovery

## Assumptions
- You have a CF deployed with silk release
- You have appA  talking to appB via c2c networking and policy (see previous story "User Flow Container to Container Networking")

## What

In the previous story you were able to use c2c networking to get appA to talk directly to appB using the overlay IP for appB (explanation of what overlay is will come soon, I swear).

But IPs are ugly and URLs are pretty. Also, what happens when you restage an app? Also, how do I load balance across many instances when I can only use IPs?

In order to fix these problems, we implemented Service Discovery, which is apart of cf-networking-release. Service Discovery is also sometimes called app-sd. Service Discovery is a fancy way of saying we handle the URL -> IP translation for internal routes. Now appA can "discover" where the "service" appB is, without having to know the IP.

Service Discovery is implemented using "internal routes" these routes will *only* work from one CF app to another. They will not be accessible from clients outside of CF.

## How

1. Start off where you left off from the previous story "User Flow Container to Container Networking"). You should have appA talking to appB via an overlay IP using `watch  "curl CF_INSTANCE_INTERNAL_IP:8080"` inside of the appA container in one terminal. 

1. In another terminal, run `cf restart appB`
   Predictably, the curl from appA to appB fails when appB is stopped. But it should come back when appB starts running again, right? ...Right? WHY IS IT STILL FAILING?

1. Recheck the overlay IP for appB `cf ssh appB -c "env | grep CF_INSTANCE_INTERNAL_IP"`
   What?! It moved! Use this new overlay IP to curl appB from appA and see that it still works. Lesson learned: IPs suck. Let's use Service Discovery instead.

1. When you create a route, any route, you have to supply a domain. To create an internal route, it must use an internal domain. We'll get into why in another story. For now, run `cf domains` and see that you should have a domain (or two) that is labeled `internal`.  Note the name of an internal domain that DOES NOT CONTAIN THE WORD "istio". You probably have the internal domain "apps.internal". Let's use that. (If you don't have a non istio internal domain, follow the resource at the bottom of this story to add a custom internal domain).

1. Using `cf map-route`, create and map a route for appB that uses the domain "apps.internal". May I suggest the route, appB.apps.internal?

1. In the terminal that is in the container for appA, run `watch  "curl appB.apps.internal:8080"`.

1. Restart appB.

1. Scale appB. Can you tell which instance you are hitting?

### Expected Result
Now that you are using internal routes to communicate via c2c, it shouldn't matter that appB is restarted. As long as appB is running, appA should be able to access it thanks to Service Discovery. When there are multiple instances of appB, the internal route will automatically load balance between all of the instances. 

## Resources

[internal domain docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains)
[servide discovery docs](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: c2c-service-discovery
L: user-workflow

---
User Flow: Create your own Internal Domain

## Assumptions
- You have a CF deployed with silk release
- You have appA  talking to appB via c2c networking and policy (see previous story "User Flow: Container to Container Networking with Service Discovery")

## What

In the previous story you created an internal route for appA to talk to appB using the domain "apps.internal", but what if we wanted to create our own internal domain (may I suggest meow.meow.meow)?

## How

1. Start off where you left off from the previous story "User Flow Container to Container Networking"). You should have appA talking to appB via an overlay IP using `watch  "curl appB.apps.internal:8080"` inside of the appA container in one terminal. 

1. In another terminal, create a new internal domain `cf create-shared-domain meow.meow.meow --internal`
   Check that it worked
   ```
$ cf domains
Getting domains in org o as admin...
name                         status   type   details
meow.meow.meow               shared          internal
   ```

1. Using `cf map-route`, create and map a route for appB that uses our new internal domain "meow.meow.meow". May I suggest the route, appB.meow.meow.meow?

1. In the terminal that is in the container for appA, use this new internal route to curl appB `watch "curl -sS appB.meow.meow.meow:8080"`
    What? "Could not resolve host"???? Why doesn't it work like our other internal route? Unlike other domains, internal domains require one more step in order for them to work.

1. Download the manifest for your CF. Look at the property `internal_domains` on the `bosh_dns_adapter` job. It probably looks like this:
   ```
internal_domains:
      - apps.internal.
   ```

   So unfortunately, there is a deploy time dependency for internal domains. I know, this makes me sad too. Let's dig in why this is.

   Warning, Architecture Description Ahead (please follow along with the diagram below): When an app makes ANY network request that requires DNS lookup (that is, any request to a URL, not an IP) , the DNS lookup first hits Bosh DNS. Bosh DNS then checks to see if the domain of the URL being requested matches any of the internal domains that it knows about from the Bosh DNS Adapter `internal_domains` property. There is no reason why this couldn't be dynamic, Bosh DNS Adapter *could* make an API call to CAPI to figure out what the up-to-date internal domains are. But it doesn't, so it's not dynamic. Then the Bosh DNS Adapter calls out to the Service Discovery Controller, which keeps track of what internal route maps to what overlay IP, very similar to the routes table in the GoRouter.


   ![href](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/architecture-diagram.png?raw=true)

1. Add our internal domain meow.meow.meow to the bosh manifest.

1. Redeploy your environment.

1. In the terminal that is in the container for appA, use this new internal route to curl appB `watch "curl -sS appB.meow.meow.meow:8080"`

### Expected Result

appA should be able to successfully reach appB using the internal route with our brand new internal domain.

## Resources

https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains
https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: c2c-service-discovery
L: user-workflow
---

Break things with Internal Domains!

## Assumptions

## What
In the previous story ("User Flow: Create your own Internal Domain") we talked about how Bosh DNS redirects all DNS lookups where the request matches an internal domain to the Bosh DNS Adapter. In this story we are going to exploit this.

   ![href](https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/architecture-diagram.png?raw=true)

## How

😇 **Pretend you are innocent user1**
1. Push [proxy app](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) and call it appA.

1. Make sure appA has an http route, let's call it APPA_ROUTE.

1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA to neopets.com
 ```
 watch "curl -sS APPA_ROUTE/proxy/neopets.com"
 ```
 You should get back some html for neopets! Fun! :D

😈 **Now pretend you are malicious user2**
As a malicious actor, you know that appA is sending traffic to neopets.com. You want to break their app and make it so appA can't reach neopets. You can do this by shadowing the neopets.com domain with an internal domain. 

1. Create the internal domain `neopets.com` (look at `cf create-shared-domain --help` if you don't remember how)

  Instead of adding this new internal domain to the bosh manifest and redeploying, we are going to hack this in on the Diego Cell. This is a great, fast, (and dangerous) debugging technique. It should be used with heavy caution, but it is often the fastest way to change a bosh property if things are really bad and need to be fixed immediately.

1. Ssh onto the Diego Cell where appA is running and become root.

1. You need to find what config file holds the information you want to change. The config on the VMs does not directly match the bosh manifest. To find any config for any bosh job in CF go to `/var/vcap/jobs/JOB_NAME`. There are many files there for Bosh DNS Adapter. In order to figure out exactly what you need to change, look at [Bosh DNS Job in the release code](https://github.com/cloudfoundry/cf-networking-release/tree/develop/jobs/bosh-dns-adapter) and look where the `internal_domains` property is used. [hint](https://github.com/cloudfoundry/cf-networking-release/blob/develop/jobs/bosh-dns-adapter/templates/handlers.json.erb#L11). [hint](https://github.com/cloudfoundry/cf-networking-release/blob/develop/jobs/bosh-dns-adapter/spec#L10).

1. Edit the correct config file and add an entry for our new domain `neopets.com`. It should look exactly like `apps.internal` except for the name.

1. Now you'll need to restart the Bosh DNS Adapter process so that it will run with our new config. Linux Bosh VMs use monit as a process manager. Run `monit summary` to see all of the processes running on this VM. Restart the Bosh DNS Adapter by running `monit restart bosh-dns-adapter`. Keep running `monit summary` until the Bosh DNS Adapter is successfully running again. If it fails to start, then you probably made a syntax error in the config file. Look at the logs and fix the error.

😇 **Back to pretending you are innocent user1**
1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA to neopets.com
   `watch "curl -sS APPA_ROUTE/proxy/neopets.com"`
   It should still show neopets. Why isn't it broken yet!?


😈 **Now pretend you are malicious user2**
You restarted the Bosh DNS Adapter, but look at the diagram again. It's actually Bosh DNS _not_ Bosh DNS adapter that does the hairpinning for internal domains.
You had to restart Bosh DNS Adapter process so it would run with the new config file, but you also need to restart the Bosh DNS process, so it can get these new values from the Bosh DNS Adapter.
1. Restart Bosh DNS
 ```
 monit restart bosh-dns
 ```

😇 **Back to pretending you are innocent user1**
1. From your terminal, use appA's `/proxy` endpoint to send traffic from appA to neopets.com
   `watch "curl -sS APPA_ROUTE/proxy/neopets.com"`
   Where did neopets go???


### Expected Result
AppA should no longer be able to access neopets. :(

## Questions

1. Neopets is a silly example. What is a worse example that customers could run into?
1. What permissions does a user require to exploit this?
1. Would you consider this a security concern?
1. Would your assessment change if internal domains were dynamic and didn't require setting a bosh property at deploy time? Why or why not?
1. What will happen to your config changes if you redeployed your environment?

## Resources
https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md#internal-domains
https://github.com/cloudfoundry/cf-networking-release/blob/develop/docs/app-sd.md

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: c2c-service-discovery

---

[RELEASE] Container to Container Networking with Service Discovery ⇧
L: c2c-service-discovery
