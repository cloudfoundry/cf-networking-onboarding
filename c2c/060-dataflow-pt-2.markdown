---
layout: single
title: Dataflow Part 2 - Network Interfaces
permalink: /c2c/dataflow-pt-2
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have CF deployed with at least 2 diego cells
- You have the meow network namespace created
- You have one [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy) app pushed and named appA (the fewer apps you have deployed the better)

## Review

This track of stories is going to go through the steps (listed below) that were covered in the dataflow overview.
The steps and diagram will be at the top of each story in case you need to orient yourself. Higher quality diagram [here](https://storage.googleapis.com/cf-networking-onboarding-images/c2c-data-plane.png).

![c2c traffic flow](https://storage.googleapis.com/cf-networking-onboarding-images/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on Diego Cell 2). This packet is called the overlay packet (aka the c2c packet).
1. **The packet exits the app container through the veth interface. <------- CURRENT STORY**
1. The overlay packet is marked with a ...mark... that is unique to the source app.
1. Because the packet is an overlay packet, it is sent to the silk-vtep interface on the Diego Cell. This interface is a VXLAN interface.
1. The overlay packet is encapsulated inside of an underlay packet. This underlay packet is addressed to underlay IP of the Diego Cell where the destination app is located (appA in this case).
1. The underlay packet exits the cell.
1. The packet then travels over the physical underlay network to the correct Diego Cell.
1. The packet arrives to the correct Diego Cell
1. The underlay packet is decapsulated. Now it's just the overlay packet again.
1. Iptables rules check that appA is allowed to talk to appB based on the mark on the overlay packet.
1. If traffic is allowed, the overlay network directs the traffic to the correct place.

## What
In this story you are going to look at network interfaces.

A network interface is ... the interface between two different networks, physical or virtual.  You can list network interfaces with `ifconfig` (old way) or `ip link list` (new way).
In order to have packets from a CF app leave an app container, there needs to be a network interface that can send packets elsewhere. In the CF case, we want them to go to the Diego Cell.
In order to have packets leave the Diego Cell, there needs to be a network interface to the underlay network.

Let's look at the interfaces on the Diego cell and in our meow network namespace.

## How
📝 **Look at network interfaces**

1. List all of the network interfaces in the Diego Cell (this output is edited for brevity and clarity)
{% include codeHeader.html %}
   ```bash
   ip link
   ```
   returns
   ```
   1: lo                 <------------- The loopback network interface that lets the system communicate with itself over localhost.
   2: eth0               <------------- A ethernet interface. Traffic goes here to leave the Diego Cell.
   1555: silk-vtep       <------------- A VXLAN overlay network interface. Overlay packets go here to be encapsulated in an  underlay packet before exiting the Diego Cell.
   1559: s-010255096003@if1558: link-netnsid 0   <-------------  The interface that links to the network namespace with id 0. The name is `s-CONTAINER-IP`. This is the veth interface.
   There will be one of these network interfaces per app on the Diego Cell.
   ```

2. Now list all of the networking interfaces in the meow networking namespace
{% include codeHeader.html %}
   ```bash
   ip netns exec meow ip link
   ```
   returns
   ```
   1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
   link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
   ```

There is only the loopback interface. The loopback interface is a special
virtual network interface that a computer uses to communicate with itself. The
standard domain name for the address is localhost.  No wonder you can't curl
google.com! There is no network interface for packets that want to leave the
network.

The solution is to create a veth (virtual ethernet) pair. A veth pair consists
of two virtual ethernet interfaces. One is placed in the host network
namespace, the other in the meow network namespace. The veth pair acts like a
bridge between the network namespace and the host.  You already saw one side of
the veth pair for the proxy app when you ran `ip link` inside of the Diego
Cell.  Let's look at the other half of the veth pair.

You might be wondering: why can't each networking namespace connect directly to `eth0`? "One
of the consequences of network namespaces is that only one interface can be
assigned to a namespace at a time. If the root namespace owns eth0, which
provides access to the external world, only programs within the root namespace
could reach the Internet. "This explanation comes the extra extra credit link
about [making your own veth
pair](https://blogs.igalia.com/dpino/2016/04/10/network-namespaces/).

🤔 **Look at network interfaces inside the proxy app container**
1. Ssh proxy app.
1. List all of the network interfaces.

### Expected Result
You should see an eth0 interface inside of the proxy app container. This is how
traffic exits the app container.

## Extra Credit
Look at the
[code](https://github.com/cloudfoundry/silk/blob/master/cni/lib/pair_creator.go)
and
[tests](https://github.com/cloudfoundry/silk/blob/master/cni/lib/pair_creator_test.go)
in silk where veth pairs are set up.

## Extra Extra Credit
📝 **Make your own veth pair**

Follow [these
instructions](https://blogs.igalia.com/dpino/2016/04/10/network-namespaces/) to
create a veth pair to connect the meow network namespace. If successful, you
will be able to curl google.com.

## Resources
* [interface explanations](https://www.computerhope.com/unix/uifconfi.htm)
* [linux network namespaces/veth/route table
  blog](https://devinpractice.com/2016/09/29/linux-network-namespace/)
