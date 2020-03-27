Introducing Bosh DNS

⚠️⚠️⚠️
**WARNING: The Bosh DNS stories are very new. ** You will probably find some typos. You will probably find some things that aren't as clear as they could be.
Please open a [PR or an issue](https://github.com/pivotal/cf-networking-program-onboarding) when you find these problems.
⚠️⚠️⚠️

## Assumptions
- You have completed the Route Registrar stories

## What?

As you learned in the last section, Route Registrar is used to create routes that make CF components available to off-platform users. You _could_ use these routes for talking from one CF component to another, _but you should not_. Using Route Registrar in place of internal component routes (1) exposes CF components unnecessarily to the big bad internet and (2) adds extra hops through the load balancer and GoRouter, which cause latency.

So what do you do when one component wants to talk to another?

You use Bosh DNS!

In this Bosh DNS story series you are going to add a custom URL for an HTTP server and trace the DNS requests for Bosh DNS routes and non-Bosh DNS routes.

## Vocab

**DNS** - stands for **D**omain **N**ame **S**ystem. DNS translates domain names (urls) to IP addresses. You can do a DNS request in the terminal by using the command line tool `dig`. With dig you can learn that neopets.com resolves to 23.96.35.235.

**Bosh DNS** - provides native DNS Support for deployments (in our case, Cloud Foundry). It gives the bosh release creator (Rahcel, that's (probably) you!) an easy way to reference other VMs with load balancing. For example, you could set up Bosh DNS such that cc.cf.internal resolves to the IP for the VM where the Cloud Controller lives. This Bosh DNS server is typically deployed on every Bosh deployed machine in a deployment (that is, every Cloud Foundry VM) using a Bosh addon.

**Bosh DNS Routes** - routes available via Bosh DNS that route to Cloud Foundry components.

**Alias** - a custom Bosh DNS Route. For example you can make the alias meow.meow send traffic to the api VM. Fun!

**External URLs** - non-CF urls. Anything on the internet. Example: neopets.com.

### Notes
- This set of stories uses the instance group `my-http-server` that was created in the Route Registrar set of stories. It is handy to use this VM with nearly nothing on it so that there is much less traffic on it.
- The networking program does not maintain Bosh DNS. However, DNS resolution is an important part of the networking traffic flow, so it is important to understand how it works.


### Links
- [Bosh DNS Docs](https://bosh.io/docs/dns/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Bosh DNS Reading

## What?

Read some docs!

1. Read the first two sections ("Architecture" and "Types of DNS addresses") thoroughly
1. (At least) skim the rest

➡️ [Bosh DNS Docs](https://bosh.io/docs/dns/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Bosh DNS -  DNS Requests for External URLs

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have 2 my-http-server instances deployed (see the story `User Workflow - Life Without Route Registrar` for setup)

## What?

In this story you are going to follow the DNS request for an external URL. Surprisingly (maybe) this will still involve Bosh DNS, even though the request is for an external URL.

## How?

📝**See that your VM already has Bosh DNS**
By default Bosh DNS is on every VM in a OSS Cloud Foundry deployment.
1. Bosh ssh onto one of the my-http-server VMs and become root.
1.  Run `monit summary`. You should see a process called bosh-dns running.

📝**Look at the DNS servers**

1. Bosh ssh onto one of the my-http-server VMs.
1. Look at the /etc/resolv.conf file. This file contains the IPs for the DNS servers used for all DNS lookups.

The file should look something like this.
```
$cat /etc/resolv.conf

# This file was automatically updated by bosh-dns
nameserver 169.254.0.2          <-------------- record this value as BOSH_DNS_IP

nameserver 169.254.169.254      <-------------- record this value as NON_BOSH_DNS_IP
search c.cf-container-networking-gcp.internal google.internal
```

📝**Find Bosh DNS running**

So you have a value for BOSH_DNS_IP, but who do you _know_ this is the Bosh DNS IP?

1. Use netstat to see what IP the Bosh DNS process is bound to.

 ```
 $ netstat -tulpn

 Active Internet connections (only servers)
 Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
 tcp        0      0 169.254.0.2:53          0.0.0.0:*               LISTEN      3227/bosh-dns
 tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      743/sshd
 tcp        0      0 127.0.0.1:53080         0.0.0.0:*               LISTEN      3227/bosh-dns
 tcp        0      0 127.0.0.1:2822          0.0.0.0:*               LISTEN      3300/monit
 tcp        0      0 127.0.0.1:2825          0.0.0.0:*               LISTEN      613/bosh-agent
 udp        0      0 169.254.0.2:53          0.0.0.0:*                           3227/bosh-dns
 ```

📝**Do a non-Bosh DNS lookup**

1. Use dig to do a DNS request for any non-CF url.

 ```
 $ dig neopets.com

 ; <<>> DiG 9.10.3-P4-Ubuntu <<>> neopets.com
 ;; global options: +cmd
 ;; Got answer:
 ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 35487
 ;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

 ;; OPT PSEUDOSECTION:
 ; EDNS: version: 0, flags:; udp: 512
 ;; QUESTION SECTION:
 ;neopets.com.			IN	A

 ;; ANSWER SECTION:
 neopets.com.		3599	IN	A	23.96.35.235

 ;; Query time: 49 msec
 ;; SERVER: 169.254.0.2#53(169.254.0.2)
 ;; WHEN: Thu Oct 03 18:15:20 UTC 2019
 ;; MSG SIZE  rcvd: 67
 ```

1. Let's go through and try to understand this dig request.

| **snippet from dig response** |  **meaning** |
| -- | -- |
|`ANSWER: 1` |This means that the DNS request successfully found an IP for the url. If an IP was not found, it would be `ANSWER: 0`|
|`23.96.35.235` |This is the IP for the neopets.com. Try it in your browser! |
|`SERVER: 169.254.0.2#53(169.254.0.2)` |This means that the DNS server that handled this request is at IP 169.254.0.2 and port 53 (this is the standard port for DNS requests). |

Do you recognize that server IP? That's the BOSH_DNS_IP that you recorded earlier!

📝**Look at logs**

1. Look at the bosh DNS logs. You should see something like...

 ```
$ tail -f /var/vcap/sys/log/bosh-dns/bosh_dns*

 [ForwardHandler] 2019/10/03 18:15:20
 INFO - handlers.ForwardHandler Request [1]
 [neopets.com.] 0 [recursor=169.254.169.254:53] 49064000ns
 ```
1. Do you recognize that recursor IP? That's the NON_BOSH_DNS_IP you recorded earlier!

❓What is a recursor?

📝**Tell dig what DNS server to use**

1. Try digging the external URL again, but this time force dig to use the BOSH_DNS_IP as the DNS server
❓Does this succeed? Why or why not?

1. Try digging the external URL again, but this time force dig to use the NON_BOSH_DNS_IP as the DNS server
❓Does this succeed? Why or why not?


### Expected Outcomes
Bosh DNS only knows information about Bosh DNS routes. For any other URL (neopets, for example) asks a different DNS server. Both the Bosh DNS server and the non-Bosh DNS server can (with recursion) resolve the external URL.

### Helpful Commands

**Do a DNS lookup**
```
dig URL [@SERVER_IP]

# for example
dig neopets.com @169.254.4.4
```

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Add a Custom Bosh DNS Alias

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have 2 my-http-server instances deployed (see the story `User Workflow - Life Without Route Registrar` for setup)

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
```

## What?
In this story you are going to add your own fun alias for your go HTTP server.

## How?

📝**Add your own alias**

1. Update your manifest to include a Bosh DNS alias. This alias could be added for any job on the instance group.
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

1. Make sure the go server is running on both of your my-http-server VMs. See the story ` Life Without Route Registrar` for help with this step.

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

Your new alias should only be accessible within Cloud Foundry and not from your local machine.

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Bosh DNS Records Table

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have my-http-server deployed with an alias setup from the previous story

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
HTTP_SERVER_ALIAS=<value>
```

## What?
So how does Bosh DNS work? How does it figure out what IP to send traffic to?

## How?

📝**Find your alias in the Bosh DNS records table**

1. Bosh ssh onto any VM in your CF deployment.
1. Look at the Bosh DNS records table. (You might need to install jq)
 ```
 cat /var/vcap/instance/dns/records.json | jq .
 ```
1. Find the data about HTTP_SERVER_ALIAS.

❓What data is present?
❓If you delete this file does Bosh DNS still work on this machine?
❓What is updating this file? (hint: read the [Bosh DNS Docs](https://bosh.io/docs/dns/))


### Links
- [Bosh DNS Docs](https://bosh.io/docs/dns/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Bosh DNS - DNS Requests for Bosh DNS Aliases

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have my-http-server deployed with an alias setup from the previous story

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
HTTP_SERVER_ALIAS=<value>
```

## What?
In this story you are going to look at what happens under the hood when you do a DNS request for HTTP_SERVER_ALIAS.

## How?

📝**Do a DNS lookup for your alias**
1. Bosh ssh onto any Cloud Foundry VM
1. Use dig to do a DNS request for your alias.

 ```
 $ dig HTTP_SERVER_ALIAS

 ; <<>> DiG 9.10.3-P4-Ubuntu <<>> meow.meow
 ;; global options: +cmd
 ;; Got answer:
 ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 28967
 ;; flags: qr aa rd ra; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 0

 ;; QUESTION SECTION:
 ;meow.meow.			IN	A

 ;; ANSWER SECTION:
 meow.meow.		0	IN	A	10.0.1.23   # <------------ This should match the IP of one of the my-http-server VMs
 meow.meow.		0	IN	A	10.0.1.22   # <------------ This should match the IP of the other my-http-server VM

 ;; Query time: 2 msec
 ;; SERVER: 169.254.0.2#53(169.254.0.2)        # <------------ This should match the BOSH_DNS_IP
 ;; WHEN: Thu Oct 03 20:45:06 UTC 2019
 ;; MSG SIZE  rcvd: 83
 ```

📝**Look at logs**

1. Look at the bosh-dns logs on the machine you did the dig on in the steps above

 ```
 $ tail -f /var/vcap/sys/log/bosh-dns/bosh_dns*

 [RequestLoggerHandler] 2019/10/03 20:49:43
 INFO - handlers.DiscoveryHandler Request [1]
 [amelia.meow.] 0 160000ns                     # <------------ Note, there is no recursor
 ```

❓ Remember how with neopets.com there was a recursor in the logs? Based on what you know about recursors, why do you think there is no recursor listed in this log line?


📝**Tell dig what DNS server to use**

1. Try digging your alias again, but this time force dig to use the BOSH_DNS_IP as the DNS server
❓Does this succeed? Why or why not?

1. Try digging your alias again, but this time force dig to use the NON_BOSH_DNS_IP as the DNS server
❓Does this succeed? Why or why not?



### Expected Results
The Bosh DNS server knows how to recurse to the non-Bosh DNS server. However, the non-Bosh DNS server does not recurse to the Bosh DNS server. Because of this, the non-Bosh DNS server will not be able to resolve HTTP_SERVER_ALIAS.

### Helpful Commands

**Do a DNS lookup**
```
dig URL [@SERVER_IP]

# for example
dig neopets.com
# OR
dig neopets.com @169.254.4.4
```

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

Can CF Apps Use Bosh DNS?

## Assumptions
- You have a CF deployed
- You have done the other stories in this section
- You have my-http-server deployed with an alias setup from the previous story

## Recorded values from previous stories
```
BOSH_DNS_IP=<value>
NON_BOSH_DNS_IP=<value>
HTTP_SERVER_ALIAS=<value>
```

## What?
In the previous stories you learned how to set up and use a Bosh DNS alias to curl the golang http server from other Bosh VM. But what about CF apps? Can apps use Bosh DNS aliases?

## How?

🤔**Try your Bosh DNS alias from an app**

1. Push any app
1. Cf ssh onto that app
1. Curl HTTP_SERVER_ALIAS:9994

❓Does it work? Why or why not? Are iptables involved? (hint: yes)

🤔**Make your Bosh DNS alias available from an app**

1. Update your ASGs (application security groups) to make the alias available to your app.

## Expected Result

Most likely, the default security groups for your CF deployment do not give your apps access to any private IP ranges. If you update the ASGs to allow apps access to private IPs (specifically, the IPs for the my-http-server VMs) your app should be able to use the alias.

If it is not working:
- make sure you restarted your app after applying the security group
- wait a couple minutes. The DNS records seem to take a little bit to propagate throughout the CF deployment.

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L:bosh-dns
---

[RELEASE] Bosh DNS ⇧
L:bosh-dns
