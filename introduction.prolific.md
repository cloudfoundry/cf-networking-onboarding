Networking Program Onboarding Week [a note for facilitators]

Hi there, facilitators!

This is a reminder to read the **[facilitation documentation](FACILITATION.md)** on GitHub.

It's also a short task list of things you might want to do before Onboarding Week starts. After you've completed the tasks you need to, remove this chore from the backlog. (I cannot make it self-destruct. Sad!)

- [ ] Allocate participants for the week.
- [ ] Designate workstation(s).
- [ ] Get one "floater" laptop per pair to facilitate independent docs reading.
- [ ] Create and populate private Tracker Project(s). Invite participants.
- [ ] Add morning standup to participant calendars so they know where to be Monday morning.
- [ ] Schedule boxes and lines whiteboarding sessions
- [ ] Add Retro to participant calendars (ideally EOD Fri)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction

---

Introduction to the Networking Program Onboarding

WELCOME!

### What?
Networking Onboarding Week is a self-paced, guided exploration of the Cloud Foundry Networking Program components, embarked upon with other Cloud Foundry members. What we hope you get out of it:

1. A coherent, if cursory, overview of a complicated product.
1. Empathy for the customer who uses that product.
1. Knowledge of the breadth of work that the networking program does
1. Knowledge of how to debug networking components when things don't go as planned
1. Experience teaching. Unless you joined Pivotal at the same time as your pair and have been arm-in-arm ever since, you have different backgrounds. Moreover, you each have something valuable to offer to the other. Make sure you share it.

### How?
1. Take your time.
1. Be conscious of your pair's progress. Check in frequently to make sure you're both getting the most out of the material.
1. Read each story completely. You'll feel very silly if you struggle for hours with a problem that turns out to have been addressed in plain English....in a part of the story that you skipped.
1. If something in a story is missing / is wrong / could be improved, create an issue or a PR in the **[networking onboarding repo](https://github.com/pivotal/cf-networking-program-onboarding)**.
1. Seriously, take your time. How often are you paid just to learn?

**Pro Tip:** As you begin stories, click the "Start" button. If you feel confident about their content when you finish, click "Finish", "Deliver", and "Accept". If you still have questions on the material, leave it in the delivered state (i.e. with the "Accept"/"Reject" buttons showing) and decide at the end how you would like to follow up on your questions.

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction

---
How to Read Onboarding Stories

### Rules

1. Read all of the words in the story.
1. No really, read all the words.
1. When possible, I will tell you how to name things. For example, push an app and name is appA. That app will be referred to as appA for the rest of the story.
1. When I cannot tell you how to name things, because they are out of my control, they will be in **CAPS_WITH_UNDERSCORES**. For example, when you create routes. Routes contain a domain name that depends on your environment. I cannot control this, so I will refer to that route as APPA_ROUTE, or something similar.
1. Resources listed at the bottom of the stories are there if you get stuck or if you are interested in going deeper. You are not expected to read, or even open, every resource. However you are expected to...
1. ...read the whole story. Please.

### Emoji Key

- 📝 indicates a story or set of instructions where the commands will be provided to you. Lots of copying and pasting.
- 🤔 indicates a story or set of instructions where the commands will NOT be provided to you. Previous stories should set up for success for these exercises. More thinking and digging through docs for these.
- ❓This indicates a non-rhetorical question. I really want you to think about the question. Discuss the answer with your pair and *record an answer in the comments of the story*. We will go over the questions and answers during the next standup.

- [ ] Read this story closely

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction

---
Set Up Your Environment

## How
1. Grab a Networking OSS dev environment that is relatively up-to-date. You will be using this environment for the whole week.

1. Create and target an org and space where you will be doing all of your work. if you are working on a Networking Team machine, you hould have the script `cf_seed` in your workstation to do this for you. You can find this script [here](https://github.com/cloudfoundry/networking-workspace/blob/0925ac5f6d1ca214e2dd871a6e36662fbd8a74b3/shared.bash#L184-L188).

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction

---
Meet Proxy!

## Assumption

- You have a CF deployed

## What

In these onboarding stories you will be using the proxy app a lot. Basically every story, so let's get familiar with it. Proxy is little golang app that is surprisingly powerful. In this story we are going to test out some of its functions.

Often with container to container (c2c) networking stories, you need an app to make a request to another app. Or you want to time how long DNS resolution (turning a URL into an IP address) is taking. You could do this with `cf ssh` and then running `curl` or `dig` (and we will in later stories!), but the proxy app was created so we didn't have to do that. It has different endpoints that will ...proxy... traffic to a given destination and that will do DNS resolution of a URL you give it, among other things.
Using proxy gives a better mirror how users use our product.

Let's check out proxy's power.

## How

📝 **Push a proxy app**

1. Clone the [cf-networking-release repo](https://github.com/cloudfoundry/cf-networking-release)
   ```
git clone https://github.com/cloudfoundry/cf-networking-release
   ```
1. Go to the proxy app
   ```
cd ~/workspace/cf-networking-release/src/example-apps/proxy
   ```
1. Push the app and name it appA
 ```
cf push appA
 ```

🤔 **Use the proxy app**

Sadly, there are not docs around all of the endpoints the proxy app has. (Maybe you will be the one write these docs? Please?)
The best way to look at all the endpoints is to go to [the code](https://github.com/cloudfoundry/cf-networking-release/blob/develop/src/example-apps/proxy/main.go#L14-L26).

When you push an app, an HTTP route is automatically created. Let's call this route PROXY_ROUTE.

1. Use the `/dig/URL_TO_DIG` endpoint to do DNS resolution for google.com.
1. Use the `/digudp/URL_TO_DIG` endpoint to do a DNS resolution  for google.com over udp. Dig usually uses tcp. This is a great way to test if udp traffic is working. (What are tcp and udp? Check out the resource below!)
1. Use the `/proxy/URL` endpoint to send traffic to neopets.com.
1. Use at least two more endpoints.

### Expected Outcome

Now you know the power of proxy!

## Resources
[tcp vs udp](https://www.vpnmentor.com/blog/tcp-vs-udp/)

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction
---

East/West What? Directional Networking Terms

## What

Network traffic can go in many directions, and there are many jargony ways to refer to them. Let's define them.

**Ingress Traffic** is traffic that originates outside the local network that is transmitted to somewhere inside the local network. Remember ingress traffic is coming *in*.

**Egress Traffic** is traffic that originates inside the local network that is transmitted somewhere outside of the local network. Remember egress traffic is *e*xiting the local network.

```
 INGRESS TRAFFIC EXAMPLE          EGRESS TRAFFIC EXAMPLE

         +-----+                          +-----+
         |     |                          |     |
     +---+     +------+               +---+     +------+
   +-+                |             +-+                |
 +-+   The Internet   ++          +-+   The Internet   ++
 |                   +-+          |                   +-+
 +-------------------+            +-------------------+

               V                                ^
               |                                ^
               |                                |
+----------------------+         +----------------------+
| Container    |       |         | Container    |       |
|              |       |         |              |       |
| +-------+    |       |         | +-------+    |       |
| |       |    |       |         | |       |    |       |
| | MyApp | <--+       |         | | MyApp | >--+       |
| |       |            |         | |       |            |
| +-------+            |         | +-------+            |
+----------------------+         +----------------------+

```

**North/South** traffic is any communication between two different networks.  Both Ingress and Egress are examples of North/South traffic.

**East/West** traffic is any communication within one network.

```
                  EAST/WEST TRAFFIC EXAMPLE

+-------------------------------------------------------------+
| My Local Network                                            |
|                                                             |
|                                                             |
|                                                             |
| +----------------------+         +----------------------+   |
| | Container1           |         | Container2           |   |
| |                      |         |                      |   |
| | +-------+            |         | +-------+            |   |
| | |       |            |         | |       |            |   |
| | | MyApp | +----------------->  | | MyApp |            |   |
| | |   1   |            |         | |   2   |            |   |
| | +-------+            |         | +-------+            |   |
| +----------------------+         +----------------------+   |
|                                                             |
+-------------------------------------------------------------+

```

## Questions

❓How would you use ingress, egress, north/south, and east/west to describe the following situations:
- You visit neopets in your browser.
- Your pair ssh-es onto your computer.
- You set up a local netcat server and send traffic to it from your terminal.

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: introduction
L: questions
---

[RELEASE] Introduction ⇧
L: introduction
