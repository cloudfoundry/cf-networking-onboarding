## Contributing
Thank you for taking the time to contribute back to this project. Please read the style guide and check your work before submitting a PR.

### Small Changes to Existing Stories
If you want make small changes (better wording, fixing typos etc), it is probably easiest change it straight in the code.

### New Stories && New Epics
Stories are grouped in `.prolific.md` files based on which track of work they are in. For example, all ASG stories are in the `asg.prolific.md` file.

If you want to create a new track of work, you will need to create a new prolific file and make sure to reference it in the [generate script](./generate-tracker-csv.go).

All stories in the same track should have a matching label and the last story in the prolific file should be a release marker.

When creating a new story it can be easiest to develop it first in tracker, make sure that it looks pretty, and then copy and paste the change into your text editor. Look at this example [prolific file](./example.prolific.md) for help on how to format prolific files. And look at the [prolific docs](https://www.pivotaltracker.com/integrations/prolific).

### Check your work
Regardless of how large your change is. Please run the `./build` script before making a PR to make sure you didn't break anything.

### Style Guide

#### Story Sections 
Stories should be built around the following sections. Not all are always required.
- **Assumptions** - what should be in place already? What apps should be pushed?
- **What** - what, in English, is going to be done in this story? Explain the context for this story.
- **How** - steps for how to complete the goal
- **Expected Result** - what is the endstate for the participant? Should the apps be communicating successfully? Should the participant be able to describe the OSI networking layers?
- **Questions** - questions for deeper learning. 
- **Extra Credit** - Suggestions for further experiments or reading that the participant could do to learn more.
- **Resources** - links to resources that would be helpful if the participant got stuck or for further reading if they are interested in a subject. Always optional. 

#### Variable Names
If you can provide a name for something, then do it. For example: app names, space names, org names, etc. Refer to this name without backticks.
 - Yes: Push an app and name it appA. Ssh onto appA. 
 - No: Push an app and ssh onto it. (Bad: did not provide name.)
 - No: Push an app and name is `appA`. Ssh onto `appA`. (Bad: used backticks.)
 
If you can't provide a name for something because it is environment related or dynamically created, then give it an all caps name. For example: routes, domains, IPs can't be named by a story writer.
- Yes: Map an HTTP route to appA. This route will be referred to as APP_A_ROUTE.
- No: Map an HTTP route to appA called meow.my-domain.com. (Bad: cannot predict domain names.)
- No: Map an HTTP route to appA. (Bad: no name given.)
- No: Map an HTTP route to appA and call it appa-route-thing (Bad: wrong format for name.) 
 
#### Capitalization
Capitalize things that are meant to be capitalized. For example: URL, NATs, CATs, IaaS. 

Capitalize words at the beginning of the sentence even if it's awkward. 
- Yes: Iptables rules are fun.
- No: c2c stands for container to container networking. (Bad: not capitalized at beginning of sentence)
 
#### Emoji Guide
- 📝 indicates a set of instructions where the commands will be provided to the participant. Lots of copying and pasting.
- 🤔 indicates a set of instructions where the commands will NOT be provided to the participant. Previous stories should set up for success for these exercises. More thinking and digging through docs for these.
- ❓This indicates a non-rhetorical question. The answers should be recorded in the comments and later reviewed with the onboarding facilitator.

#### Labels 
Each story is given one label that matches the track of work that they are apart of. 
Other labels are: 
- deploy: this means the participant will need to do a bosh deploy for this story.

#### Formatting Code 
Formatting code in trackers stories is hard. It's even harder if you want it to look good. 

Try to keep code segments in their own line. 

- Yes: 
```
tcpdump -n src APP_A_OVERLAY_IP and dst APP_B_OVERLAY_IP
```

- No: Run `tcpdump -n src APP_A_OVERLAY_IP and dst APP_B_OVERLAY_IP` (Bad: all on one line)

#### Point of View 
Write from the 2nd person point of view.
- Yes: In this story you will learn about iptables rules. 
- No: In this story we will learn about iptables rules. (Bad: 3rd person POV)


