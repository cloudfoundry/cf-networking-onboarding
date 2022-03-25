# See it live!

Check it out: https://cloudfoundry.github.io/cf-networking-onboarding

# What is CF Networking Onboarding?
CF Networking Onboarding is a guided exploration of the Cloud Foundry Networking
components.

[Amelia](https://github.com/ameowlia) created the first batch of the onboarding
stories in early 2019. She wanted an easy way to onboard new team members to the
Cloud Foundry Networking Team.

Though she initially built it for R&D engineers joining her team, it turned out
that there was a large demand for currated ways to learn more about the
inner-workings of Cloud Foundry. Over 75 people -including R&D engineers,
support engineers, platform architects, and project managers- have gone through
these stories as a week long class.

These stories have now been converted to this website to make the content
available to anyone.

# How to contribute

1. Get everything running locally.
2. Make some changes. Edit a story for clarity or even add your own module!
3. Submit a PR!

## Run locally
```
git clone git@github.com:cloudfoundry/cf-networking-onboarding.git
cd cf-networking-onboarding
bundle exec jekyll serve --incremental
```

## Understanding where things are located
```
$ tree # output edited to show the most important dirs and files.
.
├── _data
│   ├── navigation.yml # This is where you edit the sidebars and top menu
│   └── ui-text.yml
├── _site # DON'T touch anything in here! This is generated!
├── about.markdown # the about page in the main menu
├── asgs # edit asg stories here
│   ├── 000-intro.markdown
│   ├── ...
├── bosh-dns # edit bosh dns stories here
│   ├── 000-intro.markdown
│   ├── ...
├── c2c # edit c2c stories here
│   ├── 000-intro.markdown
│   ├── ...
├── curriculum.markdown # the curriculum page in the main menu
├── get-started # edit the get started stories here
│   ├── 000-intro.markdown
│   ├── ...
├── http-routes # edit the http route stories here
│   ├── 000-landing-page.markdown
│   ├── ...
├── index.markdown # the index page
├── interrupts # edit the interrupt stories here
│   ├── 000-intro.markdown
│   ├── ...
├── iptables-primer # edit the iptables stories here
│   ├── 000-intro.markdown
│   ├── ...
├── route-integrity # edit the route integrity stories here
│   ├── 000-intro.markdown
│   ├── ...
├── route-registrar # edit the route registrar stories here
│   ├── 000-intro.markdown
│   ├── ...
├── route-services # edit the route services stories here
│   ├── 000-introduction.markdown
│   ├── ...
├── service-discovery # edit the service discovery stories here
│   ├── 000-intro.markdown
│   ├── ...
└── tcp-routes # edit the tcp route stories here
    ├── 000-intro.markdown
│   ├── ...

```

