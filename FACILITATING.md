## Facilitating Networking Program Onboarding Week

### Overview
Networking Program Onboarding Week exists to provide Networking Program engineers dedicated time to explore the networking components of Cloud Foundry in a self-paced learning environment.

Either way, the idea is to give people who want more experience with networking and routing in CF the space to work through common problems alongside someone who has the same questions as they do, coming away with knowledge drawn from figuring something out themselves (something that is often lacking when paired with an experienced dev).

Networking Program Onboarding Week is intended to be a **facilitated** experience. We want participants to get "productively lost" together, not wander in aimless hair-tearing frustration. 

### Preparation Suggestions
* Allocate participants for the week. Try to get an even number of participants so that everyone always has a pair.
* Designate workstation(s). If they have been recently reimaged, run a workstation script on them. I like the [Networking Workspace](https://github.com/cloudfoundry/networking-workspace). Let Amelia know if you need access to this repo.
* Setup remote pairing hardware for remote participants (if applicable). Get headsets. Make sure zoom is downloaded. 
* Create and populate Tracker Project(s). Invite participants to all trackers.
* A month before, send out an email to all participants, reminding them to clear their schedules for the week.  Schedule day-long meetings to help avoid people scheduling more meetings for that time.
```
Howdy folks!
If you're getting this email, it means that you're scheduled to participate in Networking Program Onboarding soon (the week of __________________).

During the week you will be pairing full time with each other, working on stories exploring actual user journeys and workflows that are central to the CF networking ecosystem.

There will be more communication closer to the start of Networking Program Onboarding, so don't worry too much about the details of preparation just yet.  You will be pairing full time, however, so your teams should expect you to be away for the majority of the week. If you generally have a lot of meetings we recommend cancelling or moving as many of them as possible, but if you need to step out for a meeting here or there, that shouldn't be a problem.

If this week doesn't work for you, let us know!  

Thanks!
The Networking Program Onboarding Staff
```
* The week before, send out an email to all participants, reminding them that they are signed up for onboarding week and setting expectations about how the week will go. Here is a template that can be used as a starting point:

```
Howdy folks!

If you're getting this email, it means that you're scheduled to participate in Networking Program  Onboarding next week (the week of __________________). I wanted to get in touch with everyone and give you an idea about what to expect next week.

For starters, _____________ and _____________ will be your facilitators. They'll continue working their "day jobs," but they'll check in a few times a day to answer questions and make sure everyone is going well. They will also be there for morning stand ups, and will even lead some sessions on the architecture of some of the Cloud Foundry networking and routing systems.

This is a group with different backgrounds and skills. The stories can get a bit technical, but by and large you'll be going through actual user journeys and then jumping in to debug them -- HTTP Routes, TCP Routes, container-to-container networking policies, route integrity, etc. Please help each other out as much as possible, write down what you learn, and think about how difficulties you have with the stories could be feedback for us as facilitators or for teams that build the products you're using.

I also wanted to point out that you will be pairing full time with each other. Your teams should expect you to be away for the majority of the week, but if you need to jump out for a meeting here or there, just let your pair know.

Action Items for you:
1. If you have any issues with the scheduling, you can always sign up for a different week, but please try to find someone else to take your spot.
2. Let your anchor and PM know that you'll be away next week for Networking Program Onboarding. Feel free to loop me in if there are any concerns about that.
3. On Monday morning, please meet ______________. We'll have pairing stations ready to go for you. We'll start with an Onboarding orientation, we'll set you up with pre-populated Pivotal Tracker projects with stories for you to work on, and we'll send you on your way.

Otherwise, if you have any questions, feel free to reply to this email.

Thanks!
The Networking Program  Onboarding Staff
```
* A week before, have a 15 minute meeting with all the participants to make sure everyone is on the same page with expectations and to answer any questions.
* A week before, create a slack channel and invite everyone. Add yourself and @ameowlia (if you want me, anyway) as interrupts.
* A day before, set up the dev environments that the pairs will be using 
* A day before, post all relavent links and information in slack:
```
## First Day 
Where: 
When: 

## Links
:pair: Piar.ist
Link: 
Username: 
Password: 

:retro: Retro
Link: 
Team:
Password:

:tracker: Trackers
<Dev-Env-1> - 
<Dev-Env-2>  - 
<Dev-Env-3>  - 

Targeting Environments
<information on how to bosh target and cf target>

```

### Calendar Invites
* Hold standup every day. Make sure it's on the calenar, this way participants know when to start! This is especially important for the first day.
* Whiteboarding sessions on Tuesday and Thursday afternoons. These should be one hour long and make sure that you get a conference room with a good whiteboard (if you are not remote). If you are remote then set up a miro board.
* Retro on Friday. Prepare the retro board ahead of time so participants can add items to it all week.
* One team lunch. Make sure you get approval for budget.
* Feedback time on Friday. Block off 30 minutes for people to fill out the [feedback form](https://docs.google.com/forms/d/e/1FAIpQLSc4_cMKZh283_gM4T6BRkny9YXpLYhxrY9mP1tMv4y0SvkrGQ/viewform?usp=sf_link).

### Daily Theme Suggestions
* Be conscious of your pair's needs. Support one another. Teach one another.
* When you learn something new, tell your pair. Get in the habit of being honest about ignorance and progress.
* Learn to be comfortable with the idea of breaking something. If you're afraid to break your environment, you'll have a harder time learning and exploring.
* This is your week. Pursue the topics that interest you.
* How will you apply what you've learned this week to the future?
