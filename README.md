goGetent
========

Wrapper on top of [getent](http://en.wikipedia.org/wiki/Getent) in Go. You can use it to reliably access system databases/text files independent of the name service being used.

```
getent is a unix command that helps a user get entries in a number of important text files called databases. This includes the passwd and group databases which store user information – hence getent is a common way to look up user details on Unix. Since getent uses the same name service as the system, getent will show all information, including that gained from network information sources such as LDAP.

The databases it searches in are: ahosts, ahostsv4, ahostsv6, aliases, ethers (Ethernet addresses), group, gshadow, hosts, netgroup, networks, passwd, protocols, rpc, services, and shadow.
```
