---
layout: single
title: Get an Environment
permalink: /get-started/get-an-environment
sidebar:
  title: "Getting Started"
  nav: sidebar-getting-started
---

## Get an environment

The stories in this onboarding are written for a vanilla
[cf-deployment](https://github.com/cloudfoundry/cf-deployment). You will need
bosh creds and CF admin creds.


## Make an org and space to work in

I have a function called `cf_seed` to do this for me.
```
$ type cf_seed
cf_seed is a function
cf_seed ()
{
    cf create-org o;
    cf create-space -o o s;
    cf target -o o -s s
}
```
