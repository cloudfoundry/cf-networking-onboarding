# Facilitation Cheat Sheet

This doc is intended to make facilitation of CF Networking Onboarding easier. This doc contains copy-and-pastable messages for slack and for calendar invites.

## Slack Messages

### Polling for a date

```
✨ 🌍 I want to hold an [TIMEZONE] friendly CF Networking Onboarding week!

I am going to hold this onboarding on the date that works best for everyone.

👉👉👉 If you are interested in joining, vote with emoji which dates are best for you.
1️⃣ [DATE 1]
2️⃣ [DATE 2]
3️⃣ [DATE 3]
4️⃣ [DATE 4]


🧵 See more details about this week in the thread.
```

### What is it

```
💡 What is it?

During this onboarding individuals take a week away from their teams and go through a series of stories to learn about the the networking and routing domain of CF. The goals of this onboarding is to introduce people to the breadth of networking and routing features, dive deep into how these features are implemented "under-the-hood", and how to debug these features.
✅ Features Covered: Application Security Groups (ASGs), HTTP Routes, Route Integrity, Route Services , TCP Routes, Container to Container Networking, Service Discovery for Container to Container Networking
✅ Topics Covered: Iptables rules, Network namespaces, Network interfaces, Routes table, Overlay vs Underlay, VXLAN, ARP table, Debugging with tcpdump

❓ Who can participate?
This onboarding is useful for anyone who spends their time debugging Cloud Foundry or needs to understand Cloud Foundry deeply. This includes: R&D eng, PAs, customer engineers, PMs, and probably others! This onboarding assumes a working knowledge about bosh so it helps to have at least 6 months of CF experience under your belt before participating.

❓ A full week?? It is super hard for me to take that much time away. Can I attend other meetings?
Yes. A full week. If you need to step away for a meeting here or there, that is okay. But you will not have time to do your normal job during this onboarding.

❓ What does [TIMEZONE] friendly mean?
Onboarders work together and often pair. In order for this to work, it helps if all participants work in near-ish timezones. This onboarding session is intended for [TIMEZONE] participants. If you do not work [TIMEZONE] hours, feel free to DM me and I will place you on a list to be notified about future onboardings for your timezone.
```

### Advertising, once a date is decided

```
🙋‍♀️ How to claim a spot
1. Make sure you are available [DATE]
2. Get approval from your manager to attend.
3. DM [FACILITATOR] to claim your spot.
  * tell them you have manager approval.
  * tell them what timezone + hours you typically work.
  * tell them if you are interested in being a facilitator for future onboardings.

Due date: Let [FACILITATOR] know you can attend by [ONE WEEK BEFORE ONBOARDING STARTS].
```

### Welcoming onborders to slack channel one week ahead of start

```
Welcome [ONBOARDERS]

🎉 You have signed up for CF Networking Onboarding for the week of [DATE]. I wanted to get in touch with everyone and give you an idea about what to expect that week.

💬 For starters, your facilitator, [FACILITATOR],  will continue working their "day job" but will check in a few times a day to answer questions and make sure everyone is going well. They will also be there for stand ups, and will lead some sessions on the architecture of some of the Cloud Foundry networking and routing systems. You should have just gotten these calendar invites.

👷‍♀️👩‍⚕️🧟‍♀️ This is a group with different backgrounds and skills. The stories can get a bit technical, but by and large you'll be going through actual user journeys and then jumping in to debug them -- HTTP Routes, TCP Routes, container-to-container networking policies, route integrity, etc. Please help each other out as much as possible, write down what you learn, and think about how difficulties you have with the stories could be feedback for this onboarding or for teams that build the products you're using.

🍐 There is a lot of timezone overlap in this group! We will chat as a group on Friday at the "Welcome to CF Networking Onboarding" event about how we want to handle pairing.

🛠 This onboarding is a full week of work. You should expect to not being doing any other team work during the week, but if you need to jump out for a meeting here or there that is okay.

👉 Action Items for you
* Think about if you want to pair during the week or work solo.
* Do the [pre-work](https://cloudfoundry.github.io/cf-networking-onboarding/recommended-reading/)!
* Otherwise, if you have any questions, feel free to reply here.
```

### Events for onboarding, share one week ahead

```
🗓 Events for onboarding

👋 Friday [DATE]: Welcome to CF Networking Onboarding
When: [TIME]

☀️ Everyday: Standup
When: [TIME]

📈 Tuesday [DATE]: Whiteboarding Session 1
When: [TIME]
Miro: [LINK]

📈 Wednesday [DATE]: Iptables in TAS
When: [TIME]
slide deck: [LINK]

📈 Thursday [DATE]: Whiteboarding Session 2
When: [TIME]
Miro: [LINK]

💬 Friday [DATE]: Retro
When: [TIME]
Retro: [LINK], [PASSWORD]

✏️ Friday [DATE]: Feedback Time
When: [TIME]
Form: [LINK]

🎥 All meetings will be held in [ZOOM LINK]
```

### First day of onboarding message

```
☀️  Welcome to your first day of cf networking onboarding!!

👇👇👇  Steps to start your day 👇👇👇

💻 1. Get onto your workstation
* You are welcome to work locally from your own workstation.
* However, [FACILITATOR] has made a remote workstations for all of you to use.
* This workstation is already set up for working on this onboarding and will have all of the correct CLIs installed.

⭐️ 2. Start going through stories!
* The onboarding is located [here](https://cloudfoundry.github.io/cf-networking-onboarding/).
* Start with the "get started" module.

🤔 3 . If you get stuck...
* Getting stuck is part of learning.
* These stories are designed to make you think, not to give you all the answers.
* Read some docs, try some things out!
* Ask you fellow onboarders for hints.
* If you REALLY get stuck you can always ping [FACILITATOR].

⏰ 4. Don't forget standup!
Everyday: [TIME]
Zoom: [LINK]
```

## Calendar Invites

### Week long CF Networking Onboarding Event
Send as soon as people confirm that they can join.

> This invite is just a way to block your calendar to remind you that you have this onboarding. 
>
> One week before the onboarding I will send out more specific invites for all the events of the week.

### Welcome to CF Networking Onboarding
> Let's meet each other and chat about expectations for next week.
> 
> Slides: [LINK]


### CF Networking Onboarding - Standup
> Our daily standup! This is a chance to check in with everyone, talk about blockers, answer the questions from the stories, etc.


### CF Networking Onboarding - Whiteboarding Session 1
> It's time to review what you have learned about HTTP Routes.
>
> We will spend the hour working together to diagram the routing control plane and data plane. You all will be leading this effort, with the help of your facilitator to keep you on track.
> 
> This is a great time to clarify anything you feel iffy on or to dig down and ask more questions.
>
> Miro: [MIRO LINK]


### CF Networking Onboarding - IPTables in CF
> We are going to talk about iptables in TAS!
>
> Slides: [LINK]

### CF Networking Onboarding - Whiteboarding Session 2
> It's time to review what you have learned about Container-to-Container Networking.
> 
> We will spend the hour working together to diagram the networking control plane and data plane. You all will be leading this effort, with the help of your facilitator to keep you on track.
>
> This is a great time to clarify anything you feel iffy on or to dig down and ask more questions.
>
> Miro: [MIRO LINK]

### CF Networking Onboarding - Retro
> Let's talk about how it went this week.

### CF Networking Onboarding - Feedback Time
> This is not a meeting. I am blocking off time on your calendar for you to fill out this feedback form: [LINK].
