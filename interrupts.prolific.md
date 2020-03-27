Introduction to Interrupt Stories

🎉 Congratulations!
You made it through all of the main networking onboarding stories. Now, if you have time left over, it is time to put your skills to the test.
The next stories include actual interrupt questions (edited for formatting and spelling). It is your job to answer these questions.

🍀 Good luck

L: interrupt
---

Interrupt: why don't ASGs work for Container to Container Networking?

## Interrupt

Hi folks,

What's the mechanism that prevents Diego hosted apps from hitting other diego-cells directly? I'm looking at the default ASG and it's not explicitly omitting my 10.0.x.y pcf-deployment network, yet i cannot curl from a cf-ssh session in an AI to another Diego cell?

```
name: default_security_group
         rules:
         - destination: 0.0.0.0-169.253.255.255
           protocol: all
         - destination: 169.255.0.0-255.255.255.255
           protocol: all
```

Note: I'm *very* pleased to see this restriction is in place but I thought we were depending on ASGs here. I've been asked by Voya to export some configuration for their compliance team to document that this isolation/control is in place. - but the default ASG is misleading because it, in theory, would allow a call to my diego cell running on 10.0.0.105 in my lab.

## Restate the question
_Sometimes interrupts misuse networking or CF terms. It's your job to get to the bottom of what they are asking for. In 2 sentences or less, restate their question. Put your interpretation below this text._

```
PUT YOUR INTERPRETATION HERE
```

## Your answer
_Now that we're on the same page about what is being asked, what is the answer? Provide your answer below this text._

```
PUT YOUR ANSWER HERE
```

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: interrupt
---
Interrupt: silk, VXLAN, and limitations

## Interrupt

Hello,

I've started reading about the container to container networking and I've noticed that the "batteries included pack"
the Silk CNI plugin and the VXLAN Policy Agent use VXLAN.
In the [description](https://docs.cloudfoundry.org/concepts/understand-cf-networking.html ) of the VXLAN Policy agent it is written that the outbound traffic is tagged with the unique identifier of the source app using the VXLAN Group-Based Policy (GBP) header. Checking the [VXLAN GBP header documentation](https://tools.ietf.org/html/draft-smith-vxlan-group-policy-05) I've read that the "Group Policy ID" field is 16bit long which means 2^16= 64k unique source apps can be tagged. The VXLAN GBP header also has the VXLAN Network Identifier (VNI) field. I've also read in the [VXLAN RFC](https://tools.ietf.org/html/rfc7348) in section 4. VXLAN that "Only VMs within the same VXLAN segment can communicate with each other. Each VXLAN segment is identified through a 24-bit segment ID, termed the 'VXLAN Network Identifier (VNI)'."

Three things are not clear to me:
1. Does this setup use only one VXLAN? The docs talk about one overlay network to which all the Diego cells are connected.
2. Does the limitation that the VXLAN GBP header brings, means that only 64k unique source apps can be used for container to container networking per cf deployment?
3. Can we somehow overcome the limitation brought by the "Group Policy ID" field's size of 16bit in the VXLAN GBP by using the current setup and adding another VXLAN overlay or we have to re-implement all of the swappable components?

## Restate the question
_Sometimes interrupts misuse networking or CF terms. It's your job to get to the bottom of what they are asking for. Restate their question in a few sentences. Put your interpretation below this text._

```
PUT YOUR INTERPRETATION HERE
```

## Your answer
_Now that we're on the same page about what is being asked, what is the answer? Provide your answer below this text._

```
PUT YOUR ANSWER HERE
```

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: interrupt
---

Interrupt: how to limit or block traffic

## Interrupt

Hello networking people, I wondered is there a way to limit outgoing traffic of an app? or blocking all udp traffic? how about rate limiting?

## Restate the question
_Sometimes interrupts misuse networking or CF terms. It's your job to get to the bottom of what they are asking for. Restate their question in a few sentences. Put your interpretation below this text._

```
PUT YOUR INTERPRETATION HERE
```

## Your answer
_Now that we're on the same page about what is being asked, what is the answer? Provide your answer below this text._

```
PUT YOUR ANSWER HERE
```

🙏 _If this story needs to be updated: please, please, PLEASE submit a PR. Amelia will be eternally grateful. How? Go to [this repo](https://github.com/pivotal/cf-networking-program-onboarding). Search for the phrase you want to edit. Make the fix!_

L: interrupt
---

[RELEASE] Interrupts ⇧
L: interrupt
