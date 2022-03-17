---
layout: single
title: User Workflow - Application Security Groups
permalink: /asgs/user-workflow
sidebar:
  title: "Application Security Groups (ASGs)"
  nav: sidebar-asgs
---

## Assumptions
- You have a CF deployed
- You have one
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  app pushed and named appA

## What
Application Security Groups (ASGs) are a collection of egress (outbound) rules
that specify the protocols, ports, and IP ranges where applications can send
traffic. ASGs define rules that *allow* traffic. They are a whitelist, not a
blacklist.  Diego Cells use these ASGs to filter and log outbound network
traffic.

When applications are staging, there need to be ASGs permissive enough to allow
download particular resources (for example: ruby gems for ruby apps).  After an
application is running, devs often want the ASGs to be more restrictive and
secure. To distinguish between these two different security requirements,
administrators can define different security groups for *staging* containers
versus *running* containers.

To provide granular control when securing egress app traffic, an administrator
can also assign security groups to apply across a CF deployment, or to specific
spaces or orgs within a foundation.

### How

📝 **Look at the defaults**
1. Look at all the security group CLI commands. The commands can be confusing;
   familiarize yourself with all of the commands available.
  ```
  cf help -a | grep security-group
  ```
1. As admin view the list of security groups `cf security-groups`

In most default OSS deployments there will be two ASGs: `public_networks` and
`dns`. These default ASGs are bound (aka applied) to the entire foundation.

1. View the rules for the `public_networks` security group.
  ```
  cf security-group public_networks
```

The public_network ASG allows egress traffic to access the entire public
internet, via every protocol.

🤔 **Using Running ASGs**
Because the wide open `public_networks` security group is bound to all running
and staging contains for the entire foundation, your app should be able to
connect to any website on the internet. Let's test this.
1. Ssh onto appA (`cf ssh --help`)
1. Curl www.neopets.com. Success!
1. Unbind the public_networks **running** security-group.
1. When you bind/unbind ASGs you will see this helpful tip `TIP: Changes will not apply to existing running applications until they are restarted.` So restart your app!
1. Ssh onto appA again
1. Curl www.neopets.com.
 * ❓ What happened? Why did it fail?

🤔 **Using Staging ASGs**
1. Unbind the public_networks **staging** security-group.
1. Push a new proxy app, and name it appB.
 * ❓ What happened? Why did it fail?

🤔 **Reset your ASGs**
1. Rebind public_networks to both running and staging containers.

## Expected Results
* When you have public_networks bound to all staging and running containers
  your apps can access the entire internet!
* When public_networks is not bound to running containers then your running
  apps cannot access the internet.
* When public_networks is not bound to staging containers, then the staging
  container is not able to access the internet to install godep (for go apps)
  and other staging requirements, so `cf push` will fail.

## Resources
* [Application Security Groups Documentation](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html)
* [Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)
* ["Taking Security to the Next Level—Application Security Groups" by Abby Kearns](https://blog.pivotal.io/pivotal-cloud-foundry/products/taking-security-to-the-next-level-application-security-groups)
* ["Making sense of Cloud Foundry security group declarations" by Sadique Ali](https://sdqali.in/blog/2015/05/21/making-sense-of-cloud-foundry-security-group-declarations/)
