---
layout: single
title: Why don't ASGs work for C2C?
permalink: /interrupts/asgs-and-c2c
sidebar:
  title: "Interrupts"
  nav: sidebar-interrupts
---

## Interrupt

_Hi folks,_

_What's the mechanism that prevents Diego hosted apps from hitting other
diego-cells directly? I'm looking at the default ASG and it's not explicitly
omitting my 10.0.x.y deployment network, yet i cannot curl from a cf-ssh
session in an AI to another Diego cell?_

```
name: default_security_group
         rules:
         - destination: 0.0.0.0-169.253.255.255
           protocol: all
         - destination: 169.255.0.0-255.255.255.255
           protocol: all
```

_Note: I'm *very* pleased to see this restriction is in place but I thought we
were depending on ASGs here. I've been asked to export some configuration for
their compliance team to document that this isolation/control is in place. -
but the default ASG is misleading because it, in theory, would allow a call to
my diego cell running on 10.0.0.105 in my lab._

## Restate the question
Sometimes interrupts misuse networking or CF terms. It's your job to get to
the bottom of what they are asking for. In 2 sentences or less, restate their
question.

## Answer
Now that we're on the same page about what is being asked, what is the answer?

