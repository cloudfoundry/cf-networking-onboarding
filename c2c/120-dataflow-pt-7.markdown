---
layout: single
title: Dataflow Part 7 - Overlay Leases and ARP
permalink: /c2c/dataflow-pt-7
sidebar:
  title: "Container-to-container Networking"
  nav: sidebar-c2c
---

## Assumptions
- You have a CF deployed with at least 2 diego cells
- You have two
  [proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)
  apps pushed and named appA and appB (the fewer apps you have deployed the
  better)

## Review
This track of stories is going to go through the steps (listed below) that were
covered in the dataflow overview.  The steps and diagram will be at the top of
each story in case you need to orient yourself. Higher quality diagram
[here](https://storage.googleapis.com/cf-networking-onboarding-images/c2c-data-plane.png).

![c2c traffic
flow](https://storage.googleapis.com/cf-networking-onboarding-images/overlay-underlay-silk-network.png)

1. AppB (on Diego Cell 1) makes a request to AppA's overlay IP address (on
   Diego Cell 2). This packet is called the overlay packet (aka the c2c
   packet).
1. The packet exits the app container through the veth interface.
1. The overlay packet is marked with a ...mark... that is unique to the source
   app.
1. Because the packet is an overlay packet, it is sent to the silk-vtep
   interface on the Diego Cell. This interface is a VXLAN interface.
1. **The overlay packet is encapsulated inside of an underlay packet. This
   underlay packet is addressed to the underlay IP of the Diego Cell where the
   destination app is located (appA in this case).    <------- CURRENT STORY**
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

In the last story we learned that the VTEP is responsible for encapsulating the
overlay packet and sending it to the correct Diego Cell. But how does it get
that information?

Each CF Deployment is given a [range of
IPs](https://github.com/cloudfoundry/silk-release/blob/develop/jobs/silk-controller/spec#L30-L33)
to act as the overlay IPs. By default the overlay range is `10.255.0.0/16`.
This is CIDR (pronounced cider) notation for the IPs that range from
10.255.0.0 to 10.255.255.255. Every Diego Cell is given a subset of this range
to give to the apps.

The Silk Controller is a database backed process that keeps track of which
subnet of IPs is leased out to which Diego Cell. Every Diego Cell has a process
running called the Silk Daemon, which constantly polls the Silk Controller and
asks: "what overlay ranges map to what Diego Cells"? The Silk Daemon then
writes that information to the Address Resolution Protocol (ARP) table.

ARP is at layer 2 in the OSI (Open System Interconnection) model. If you
haven't heard of OSI, read
[this](https://www.webopedia.com/quick_ref/OSI_Layers.asp) as a primer. The
main bit, is that at layer 2 [MAC
addresses](https://whatismyipaddress.com/mac-address) are used to address data.
Because of this, in the ARP table and in the silk database, the Diego Cells are
referenced by their MAC addresses and not their underlay IPs. If you aren't
familiar with OSI or MAC addresses, read those links before continuing on.

In this story, you are going to look at the Silk Controller database to see
what overlay subnet is leased to which Diego Cells. Then you'll going to look
at the ARP table.

## How

🤔 **Look at silk controller database**

1. Get the mysql username and password for the silk controller db. Download
   your bosh manifest and see what is set for the properties:
   `silk-controller.properties.database.name` and
   `silk-controller.properties.database.password`. You might need to look up
   these values in credhub. This will differ based on your deployment.
1. Bosh ssh onto the database VM and become root.
1. Login to your database. For a pxc-mysql deployment, it looks like this
   ```bash
   /var/vcap/packages/pxc/bin/mysql -u network_connectivity -p -h sql-db.service.cf.internal -D DATABASE_NAME
   ```
1. Run the following query to look at all of the subnets.
   ```
   mysql> select * from subnets;
   +----+---------------+---------------------------+-----------------------------+------------------+
   | id | underlay_ip   | overlay_subnet            | overlay_hwaddr              | last_renewed_at  |
   +----+---------------+---------------------------+-----------------------------+------------------+
   |  0 | DIEGO_CELL_IP | DIEGO_CELL_OVERLAY_SUBNET | DIEGO_CELL_VTEP_MAC_ADDRESS | TIMESTAMP        |
   |  1 | 10.0.1.12     | 10.255.77.0/24            | ee:ee:0a:ff:4d:00           | 1555520514       | <--- Let's call these values: DIEGO_CELL_0_IP, DIEGO_CELL_0_OVERLAY_SUBNET, DIEGO_CELL_0_VTEP_MAC_ADDRESS
   |  2 | 10.0.1.13     | 10.255.82.0/24            | ee:ee:0a:ff:52:00           | 1555520512       | <--- Let's call these values: DIEGO_CELL_1_IP, DIEGO_CELL_1_OVERLAY_SUBNET, DIEGO_CELL_1_VTEP_MAC_ADDRESS
   |  3 | 10.0.1.17     | 10.255.0.225/32           | ee:ee:0a:ff:00:e1           | 1555520513       | <--- This is an istio router, which is also on the overlay. Each istio router gets one overlay IP. You might or might not have these.
   |  4 | 10.0.1.18     | 10.255.0.160/32           | ee:ee:0a:ff:00:a0           | 1555520515       | <--- Another istio router.
   +----+---------------+---------------------------+-----------------------------+------------------+
   ```

📝**Look at the arp table**

1. Ssh onto the Diego Cell with DIEGO_CELL_0_IP as the underlay IP and become root.
1. Find the mac address of the silk-vtep interface
   ```bash
   ip link show silk-vtep # <---- Should match DIEGO_CELL_0_VTEP_MAC_ADDRESS
   ```

1. Look at the ARP table. This is how the VXLAN VTEP knows where each Diego Cell is located.
   ```bash
   arp
   ```

   You should see something like this. The output is split up so we can look at it
   section by section. Yours might be in a different order.

   ⬇️ This is the entry for the other Diego Cell. This should match DIEGO_CELL_1_OVERLAY_SUBNET and DIEGO_CELL_1_VTEP_MAC_ADDRESS.
   ```
   Address                  HWtype  HWaddress           Flags Mask            Iface
   10.255.82.0              ether   ee:ee:0a:ff:52:00   CM                    silk-vtep
   ```

   ⬇️ This is the entry for the Istio Routes. You might or might not have these. The MAC address and overlay ips should match the database.
   ```
   Address                  HWtype  HWaddress           Flags Mask            Iface
   10.255.0.225             ether   ee:ee:0a:ff:00:e1   CM                    silk-vtep
   10.255.0.160             ether   ee:ee:0a:ff:00:a0   CM                    silk-vtep
   ```

   ⬇️ This is the entries for the 2 apps running on this cell. The addresses match the overlay IPs of the two apps.
   ```
   Address                  HWtype  HWaddress           Flags Mask            Iface
   10.255.77.3              ether   ee:ee:0a:ff:4d:03   CM                    s-010255077003
   10.255.77.4              ether   ee:ee:0a:ff:4d:04   CM                    s-010255077004
   ```

   ⬇️ This is the entry for the eth0 interface.
   ```
   Address                  HWtype  HWaddress           Flags Mask            Iface
   10.0.0.1                 ether   42:01:0a:00:00:01   C                     eth0
   ```

## Resources
* [The Layers of Networking - OSI](https://www.webopedia.com/quick_ref/OSI_Layers.asp)

