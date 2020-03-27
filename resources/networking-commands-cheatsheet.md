# Networking Commands Cheatsheet

Here are some commands that you learned during networking onboarding. If you find any typos or if there are any commands you would like to add, please PR!

## arp
ARP stands for Address Resolution Protocol. The ARP table keeps track of which IP addresses map to which hardware addresses. The arp CLI is used to view the ARP table.

```
# sample command
arp

# sample output
Address                  HWtype  HWaddress           Flags Mask            Iface
10.255.73.0              ether   ee:ee:0a:ff:49:00   CM                    silk-vtep
10.255.235.3             ether   ee:ee:0a:ff:eb:03   CM                    s-010255235003
10.0.0.1                 ether   42:01:0a:00:00:01   C                     eth0
```

## dig
The dig CLI does DNS lookups.

```
# sample command
dig URL [@SERVER_IP]
dig neopets.com @169.254.4.4

# sample output
 ; <<>> DiG 9.10.3-P4-Ubuntu <<>> neopets.com
 ;; global options: +cmd
 ;; Got answer:
 ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 35487
 ;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

 ;; OPT PSEUDOSECTION:
 ; EDNS: version: 0, flags:; udp: 512
 ;; QUESTION SECTION:
 ;neopets.com.			IN	A

 ;; ANSWER SECTION:
 neopets.com.		3599	IN	A	23.96.35.235

 ;; Query time: 49 msec
 ;; SERVER: 169.254.0.2#53(169.254.0.2)
 ;; WHEN: Thu Oct 03 18:15:20 UTC 2019
 ;; MSG SIZE  rcvd: 67
```

| **snippet from dig response** |  **meaning** |
| -- | -- |
|`ANSWER: 1` |This means that the DNS request successfully found an IP for the url. If an IP was not found, it would be `ANSWER: 0`|
|`23.96.35.235` |This is the IP for the neopets.com. Try it in your browser! |
|`SERVER: 169.254.0.2#53(169.254.0.2)` |This means that the DNS server that handled this request is at IP 169.254.0.2 and port 53 (this is the standard port for DNS requests). |


## ifconfig
The ifconfig CLI is used to look at networking interfaces.

```
# sample command
ifconfig

# sample output
eth0      Link encap:Ethernet  HWaddr 42:01:0a:00:01:0c
          inet addr:10.0.1.12  Bcast:10.0.1.12  Mask:255.255.255.255
          UP BROADCAST RUNNING MULTICAST  MTU:1460  Metric:1
          RX packets:72072478 errors:0 dropped:0 overruns:0 frame:0
          TX packets:69556380 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:13157676674 (13.1 GB)  TX bytes:43292132622 (43.2 GB)

lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:179614177 errors:0 dropped:0 overruns:0 frame:0
          TX packets:179614177 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:75826807392 (75.8 GB)  TX bytes:75826807392 (75.8 GB)

s-010255235003 Link encap:Ethernet  HWaddr aa:aa:0a:ff:eb:03
          inet addr:169.254.0.1  Bcast:0.0.0.0  Mask:255.255.255.255
          UP BROADCAST RUNNING NOARP MULTICAST  MTU:1410  Metric:1
          RX packets:2178 errors:0 dropped:0 overruns:0 frame:0
          TX packets:3093 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:442683 (442.6 KB)  TX bytes:772499 (772.4 KB)

silk-vtep Link encap:Ethernet  HWaddr ee:ee:0a:ff:eb:00
          inet addr:10.255.235.0  Bcast:0.0.0.0  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1410  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
```

## ip
The ip CLI can "show / manipulate routing, devices, policy routing and tunnels".

```
# sample commands for ip netns

# list networking namespaces
ip netns

# create networking namespace
ip netns add NETWORK_NAMESPACE_NAME

# execute a command in a particular networking namespace
ip netns exec NETWORK_NAMESPACE_NAME COMMAND
```

```
# list all networking interfaces
ip link list 

# sample output
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1460 qdisc mq state UP mode DEFAULT group default qlen 1000
    link/ether 42:01:0a:00:01:0c brd ff:ff:ff:ff:ff:ff
42652: silk-vtep: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1410 qdisc noqueue state UNKNOWN mode DEFAULT group default
    link/ether ee:ee:0a:ff:eb:00 brd ff:ff:ff:ff:ff:ff
42656: s-010255235003@if42655: <BROADCAST,MULTICAST,NOARP,UP,LOWER_UP> mtu 1410 qdisc noqueue state UP mode DEFAULT group default
    link/ether aa:aa:0a:ff:eb:03 brd ff:ff:ff:ff:ff:ff link-netnsid 0
```

```
# list routes in route table
ip route

# sample output (yup, it's not pretty and there are no headers)
default via 10.0.0.1 dev eth0  proto dhcp  src 10.0.1.12  metric 1024
default via 10.0.0.1 dev eth0  proto dhcp  metric 1024
10.0.0.1 dev eth0  proto dhcp  scope link  src 10.0.1.12  metric 1024
10.0.0.1 dev eth0  proto dhcp  metric 1024
10.255.0.0/16 dev silk-vtep  proto kernel  scope link  src 10.255.235.0
10.255.73.0/24 via 10.255.73.0 dev silk-vtep  src 10.255.235.0
10.255.235.3 dev s-010255235003  proto kernel  scope link  src 169.254.0.1
```

## iptables
The iptables CLI lets users view and modify the iptables rules in a networking namespace.

```
# basic commands
iptables -S        # display rules in iptables-save format
iptables -L        # display rules in table format
iptables -A RULE   # add a rule
iptables -D RULE   # delete a rule
iptables -X CHAIN  # delete a chain
```

## netstat
Netstat is a tool that can show information about network connections, routing tables, and network interface statistics.

```
# sample command
netstat -tulp
# -t  <---- show tcp sockets
# -u  <---- show udp sockets
# -l  <---- display listening sockets
# -p  <---- display PID/program name for sockets
```



## route
The route CLI lets users list and modify routing tables

```
# list the routing table
route -n

# sample output
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.0.0.1        0.0.0.0         UG    1024   0        0 eth0
0.0.0.0         10.0.0.1        0.0.0.0         UG    1024   0        0 eth0
10.0.0.1        0.0.0.0         255.255.255.255 UH    1024   0        0 eth0
10.0.0.1        0.0.0.0         255.255.255.255 UH    1024   0        0 eth0
```



## tcpdump
The tcpdump CLI lets users see a dump of the current network packets on a network.

```
# to get packets
tcpdump

# to get packets from a particular source
tcpdump -n src SOURCE_IP 

# to get packets from a particular destination
tcpdump -n dst DESTINATION_IP 

# to get packets from any interface (by default it only shows eth0)
tcpdump -i any

# to get very specific packets
tcpdump -n src SOURCE_IP and dst DESTINATION_IP -i any


# sample output
20:55:33.304757 IP vm-5fcab44e-1a4d-44b4-4a75-98553f9fa353.c.cf-container-networking-gcp.internal.35448 > 28848a08-abcf-4866-a55f-f8e6b96a5a6a.diego-api.default.cf.bosh.ssh: Flags [.], ack 323112, win 1419, options [nop,nop,TS val 2687832119 ecr 3036103886], length 0
```
