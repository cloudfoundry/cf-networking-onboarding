---
layout: single
title: Network Namespaces
permalink: /c2c/dataflow-pt-1
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells

## Review
This part of the c2c module is a set of stories that goes through the steps
(listed below) that were covered in the dataflow overview.  The steps and
diagram will be at the top of each story in case you need to orient yourself.
Higher quality diagram
[here](https://storage.googleapis.com/cf-networking-onboarding-images/c2c-data-plane.png).

![c2c traffic flow](https://storage.googleapis.com/cf-networking-onboarding-images/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on
   Diego Cell 2). This packet is called the overlay packet (aka the c2c
   packet).
1. **The packet exits the app container through the veth interface. <-------
   CURRENT STORY**
1. The overlay packet is marked with a ...mark... that is unique to the source
   app.
1. Because the packet is an overlay packet, it is sent to the silk-vtep
   interface on the Diego Cell. This interface is a VXLAN interface.
1. The overlay packet is encapsulated inside of an underlay packet. This
   underlay packet is addressed to underlay IP of the Diego Cell where the
   destination app is located (appA in this case).
1. The underlay packet exits the cell.
1. The packet then travels over the physical underlay network to the correct
   Diego Cell.
1. The packet arrives to the correct Diego Cell
1. The underlay packet is decapsulated. Now it's just the overlay packet again.
1. Iptables rules check that appA is allowed to talk to appB based on the mark
   on the overlay packet.
1. If traffic is allowed, the overlay network directs the traffic to the
   correct place.

## What
Each CF app runs in a container, but what *is* a container? A container is a
collection of **namespaces** and **cgroups**.

**Namespaces** "partition kernel resources such that one set of processes sees
one set of resources while another set of processes sees a different set of
resources" (thanks [wiki](https://en.wikipedia.org/wiki/Linux_namespaces)).
There are different types of namespaces. For example, the mount namespace lets
processes see different file trees. Another example (most related to this
onboarding) is the network namespace. The network namespace isolates network
interfaces (we'll get into what those are in the next story).

**Cgroups** (pronounced cee-groups) are resource limits. Cgroups let you say:
"these processes can only use 1G of memory".

Most important in this onboarding context is the networking namespace.
Container networking components are responsible for setting up the networking
namespace for each app.

Even the bosh processes are run in containers! In CF this is done with [Bosh
Process Manager (BPM)](https://github.com/cloudfoundry/bpm-release).

In this story you'll play around with BPM containers, network namespaces, and
even make a network namespace for yourself.

## How

1. Read Julia Evan's post: [What even is a
   container](https://jvns.ca/blog/2016/10/10/what-even-is-a-container/)

📝 **Look at a container**
1. Ssh onto a Diego Cell and become root.
1. Look at the network interfaces (again, we'll go deeper in the next story).
{% include codeHeader.html %}
   ```bash
   ifconfig
   ```
1. Inspect the directories that the root user has access to. For example, look at all the log files.
{% include codeHeader.html %}
   ```bash
   ls /var/vcap/sys/log
   ```
1. Get into the BPM container for the vxlan-policy-agent.
{% include codeHeader.html %}
   ```bash
   bpm list
   ```
{% include codeHeader.html %}
   ```bash
   bpm shell vxlan-policy-agent
   ```
1. Look at the network interfaces. How do they compare to the host vm network interfaces?
1. Look at the log files you can access. How do they compare to the files accessible to root user?
  * ❓Based on this information, does BPM create a [mount
    namespace](https://medium.com/@teddyking/linux-namespaces-850489d3ccf) for
    the vxlan-policy-agent container?
  * ❓Based on this information, does BPM create a network namespace for the
    vxlan-policy-agent container?
1. Exit the container.

📝 **Make your own network namespace**

1. Still on a Diego Cell as root, create your own network namespace called meow.
{% include codeHeader.html %}
   ```bash
   ip netns add meow
   ```
1. List all of the networking namespaces
{% include codeHeader.html %}
   ```bash
   ip netns
   ```
You should only see meow. Hmmm. You might think you would see the other
networking namespaces for all the apps on this cell. (I certainly thought so
when I first tried this.) You'll learn how to view an app's networking
namespace one day, I promise.

1. Curl google.com from the Diego Cell. See that it works!

1. Curl google.com from inside of your networking namespace
{% include codeHeader.html %}
   ```bash
   ip netns exec meow curl google.com
   ```

What? It doesn't work!? You should see `curl: (6) Could not resolve host: google.com`. Try another URL. They will all fail.

### Expected Outcome
The meow networking namespace can't send any traffic out of the container. By
default, network namespaces are completely isolated and have no network
interfaces.  In the next story you'll explore network interfaces. You'll learn
why the meow namespace needs one in order for you to curl google.com.

## Resources
* [ip netns man page](http://man7.org/linux/man-pages/man8/ip-netns.8.html)
* [linux network namespaces/veth/route table blog](https://tanzu.vmware.com/developer/blog/a-container-is-a-linux-namespace-and-networking-basics/)  #replaced a dead link
* [network namespaces blog](https://blogs.igalia.com/dpino/2016/04/10/network-namespaces/)
* [interface explanations](https://www.computerhope.com/unix/uifconfi.htm)
* [linux namespaces overview](https://medium.com/@teddyking/linux-namespaces-850489d3ccf)
