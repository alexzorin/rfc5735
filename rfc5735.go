package rfc5735

import "net"

var reservedCidrs = []string{
	"0.0.0.0/8",
	"10.0.0.0/8",
	"127.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.0.2.0/24",
	"192.88.99.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
	"198.51.100.0/24",
	"203.0.113.0/24",
	"224.0.0.0/4",
	"240.0.0.0/4",
}

var reservedNets []*net.IPNet

func init() {
	var v *net.IPNet
	var e error
	for _, cidr := range reservedCidrs {
		_, v, e = net.ParseCIDR(cidr)
		if e != nil {
			panic(e)
		}
		reservedNets = append(reservedNets, v)
	}
	reservedCidrs = nil
}

// Only accepts a regular formatted IPv4 adddress, i.e.:
//
// 	192.168.1.1
//
// CIDR notation (192.168.1.1/32) will not work.
//
// Will return false if the address is not in a valid format.
func IsReservedString(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return false
	}
	return IsReserved(ip)
}

func IsReserved(ip net.IP) bool {
	for _, v := range reservedNets {
		if v.Contains(ip) {
			return true
		}
	}
	return false
}
