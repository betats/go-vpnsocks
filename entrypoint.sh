#!/bin/bash

DEFAULT_GW=$(ip route get 1.1.1.1 | grep -o "via [0-9\/\.]*" | sed -e 's/via //')

## for PRIVATE NETWORK request response
ip route add 10.0.0.0/8 via $DEFAULT_GW
ip route add 172.12.0.0/12 via $DEFAULT_GW
ip route add 192.168.0.0/16 via $DEFAULT_GW

if [ -e $ADD_PRIVATE_NETWORK ]; then
	ip router add $ADD_PRIVATE_NETWORK via $DEFAULT_GW
fi

## DNS resolver modify from private resolver
### dns query need to via VPN
echo nameserver $DNS_RESOLVER1 >  /etc/resolv.conf
echo nameserver $DNS_RESOLVER2 >> /etc/resolv.conf

/usr/local/bin/go-vpnsocks
