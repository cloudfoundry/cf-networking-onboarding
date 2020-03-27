User Workflow: Application Security Groups

## Assumptions
- You have a CF deployed
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and named appA

## What?
Application Security Groups (ASGs) are a collection of egress (outbound) rules that specify the protocols, ports, and IP ranges where applications can send traffic. ASGs define rules that *allow* traffic. They are a whitelist, not a blacklist.
Diego Cells use these ASGs to filter and log outbound network traffic.

When applications are staging, there need to be ASGs permissive enough to allow download particular resources (for example: ruby gems for ruby apps).
After an application is running, devs often want the ASGs to be more restrictive and secure. To distinguish between these two different security
requirements, administrators can define different security groups for *staging* containers versus *running* containers.

To provide granular control when securing egress app traffic, an administrator can also assign security groups to apply across a CF deployment, or to specific spaces or orgs within a foundation.

### How?

📝 **Look at the defaults**
1. Look at all the security group CLI commands. The commands can be confusing; familiarize yourself with all of the commands available.
 ```
cf help -a | grep security-group
 ```
1. As admin view the list of security groups `cf security-groups`

 In most default OSS deployments there will be two ASGs: `public_networks` and `dns`. These default ASGs are bound (aka applied) to the entire foundation.

1. View the rules for the `public_networks` security group.
 ```
 cf security-group public_networks
 ```

The public_network ASG allows egress traffic to access the entire public internet, via every protocol.

🤔 **Using Running ASGs**
Because the wide open `public_networks` security group is bound to all running and staging contains for the entire foundation, your app should be able to connect to any website on the internet. Let's test this.
1. Ssh onto appA (`cf ssh --help`)
1. Curl www.neopets.com. Success!
1. Unbind the public_networks **running** security-group.
1. When you bind/unbind ASGs you will see this helpful tip `TIP: Changes will not apply to existing running applications until they are restarted.` So restart your app!
1. Ssh onto appA again
1. Curl www.neopets.com.

❓ What happened? Why did it fail?

🤔 **Using Staging ASGs**
1. Unbind the public_networks **staging** security-group.
1. Push a new proxy app, and name it appB.

❓ What happened? Why did it fail?

**Reset your ASGs**
1. Rebind public_networks to both running and staging containers.

### Expected Result
When you have public_networks bound to all staging and running containers your apps can access the entire internet!
When public_networks is not bound to running containers then your running apps cannot access the internet.
When public_networks is not bound to staging containers, then the staging container is not able to access the internet to install godep (for go apps) and other staging requirements, so `cf push` will fail.

#### Note:
If you're working with PCF Dev, you should see three security groups. One is named `all_pcfdev` which opens all egress traffic. Because of this security group, any other group is redundant.

## Resources
[Application Security Groups Documentation](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html)
[Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)
["Taking Security to the Next Level—Application Security Groups" by Abby Kearns](https://blog.pivotal.io/pivotal-cloud-foundry/products/taking-security-to-the-next-level-application-security-groups)
["Making sense of Cloud Foundry security group declarations" by Sadique Ali](https://sdqali.in/blog/2015/05/21/making-sense-of-cloud-foundry-security-group-declarations/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: asgs
L: user-workflow
L: questions
---
Iptables Primer

Please go to the section labeled "iptables-primer" and complete those stories before moving on.

L: asgs
---

ASGs and Iptables

## Assumptions
- You have a CF deployed
- You have completed the iptables-primer track
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and named appA (the fewer apps you have deployed the better)
- You have the public_networks security group bound to all running and staging containers

## What?
In the previous story "User Workflow: Application Security Groups" you learned what ASGs *are*, but how are they implemented under the hood?
(Hint: iptables)

Information about ASGs are stored in CAPI's database. When a container create is triggered (for example by a cf push), CAPI passes ASG information along to Diego (container scheduler). Diego then passes the ASG information along to Garden (the container creator), which passes it along to the container networking components. After this long journey, the container networking components turn this ASG information into iptables rules on the Diego Cells.

Iptables rules are a series of rules that network packets match against to decide whether that traffic is allowed or not. This is one way firewalls can be implemented.
ASGs are only concerned with egress traffic from CF apps to something external to the CF foundation (usually the internet).

The iptables man page is extremely helpful as a reference! Consult it if you have any questions throughout this story.

Let's investigate what iptables rules are created for the public_networks ASG.

## How?

📝 **Find the iptables rule for the public_networks ASG**
1. Look at the ASG public_networks
 ```
 cf security-group public_networks
 ```
 Note which IP ranges are specified. Save these rules somewhere accessible for future reference.
1. Ssh onto the Diego Cell where appA is running and become root.
 ```
 # First determine the IP of the Diego Cell
cf ssh appA  -c "env | grep CF_INSTANCE_IP"

 # Look at bosh output to see which Diego Cell has that IP
bosh instances

 # Ssh on
bosh ssh diego-cell/DIEGO_CELL_INSTANCE_GUID
 ```
1. List all of the iptables rules on the filter table on the Diego Cell
 ```
 iptables -t filter -S
 ```

Yep. It's a lot. Take a deep breath. Read it line by line. I promise it will start to become comprehensible.

Remember those IP ranges that are in the public_networks security group?
If you search for one of them (ie. `0.0.0.0-9.255.255.255`) you should be able to find an iptables rule that looks something like this:

```
# If you have logging enabled
-A netout--772fbbd5-862a-4b3d-7 -m iprange --dst-range 0.0.0.0-9.255.255.255 -g netout--772fbbd5-862a-4--log

# If you don't have logging enabled
-A netout--772fbbd5-862a-4b3d-7 -m iprange --dst-range 0.0.0.0-9.255.255.255 -j ACCEPT
```

Copy the iptables rule that you found on the Diego Cell that looks like the line above. Paste it in the story below.
It will be necessary to reference it soon.

```
# Put your iptables rule...
# HERE
```

🤔 **Decipher the iptables rule for the public_networks ASG**
When you run `iptables -S` the iptables rules are displayed in the format that they were created in. For example, if you prepended `iptables` to the line above and ran (DON'T! this is hypothetical) `iptables -A netout--772fbbd5-862a-4b3d-7 -m iprange ...` it would add that iptables rule. 

With that in mind, use `iptables --help` to define what the following flags mean and record it below...

```
-A                                 ...
-m                                 ...
--dst-range                        ...
-g                                 ...
-j                                 ...
```

1. What is the name of the chain that the above iptables rule is appended to? Let's call this CHAIN-NAME.
1. What conditions need to be met for a packet to match this rule?
1. If logging is enabled: If packets meet the above conditions, what chain of rules will they jump to and evaluate against next? Let's call this JUMP-TO-CHAIN-NAME.

1. Look at the rules only in the chains from the questions above.
 ```
iptables --list CHAIN-NAME
iptables --list JUMP-TO-CHAIN-NAME # <---- only if logging is enabled
 ```

Packets stop matching against iptables rules when they reach either the ACCEPT target(let the packet flow through!), the DROP target (stop that packet in its tracks!), or the REJECT target (nicely tell the sender that this destination is not available to them) targets.

🤔 **Read iptables rules to predict behavior**
1. Pretend there is a packet trying to hit the IP 23.96.32.148 via tcp. Where will it hit ACCEPT or DROP?
1. Pretend there is a packet trying to hit the IP 172.20.0.3 via tcp. Where will it hit ACCEPT or DROP?
1. ssh onto your app and try to curl 23.96.32.148. Did it succeed or fail? Did it match your expectations?
1. ssh onto your app and try to curl 172.20.0.3. Did it succeed or fail? Did it match your expectations?


### Expected Result
You should be able to find the iptables rules for the public_networks and be able to follow the flow to either ACCEPT or DROP.

Iptables are hard! Hopefully they have been demystified a little bit.
Spend some extra time looking at all the iptables rules on a Diego Cell. Research what you don't understand. Ask lots of questions.

## Resources
**iptables**
[iptables man page](http://ipset.netfilter.org/iptables.man.html)
[Julia Evans iptables basics](https://jvns.ca/blog/2017/06/07/iptables-basics/)
[Aidan's iptables ppt](https://docs.google.com/presentation/d/1qLkNu633yLHP5_S_OqOIIBETJpW3erk2QuGSVo71_oY/edit#slide=id.p)
[iptables primer](https://danielmiessler.com/study/iptables/)

**ASGs**
[Application Security Groups Documentation](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html)
[Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)
["Taking Security to the Next Level—Application Security Groups" by Abby Kearns](https://blog.pivotal.io/pivotal-cloud-foundry/products/taking-security-to-the-next-level-application-security-groups)
["Making sense of Cloud Foundry security group declarations" by Sadique Ali](https://sdqali.in/blog/2015/05/21/making-sense-of-cloud-foundry-security-group-declarations/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: asgs

---

[RELEASE] Application Security Groups (ASGs) ⇧
L: asgs
