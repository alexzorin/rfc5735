package rfc5735

import "testing"

var reserved = []string{
	"192.168.1.1",
	"127.0.0.1",
	"10.0.0.1",
	"255.255.255.255",
}

var notReserved = []string{
	"203.12.160.35",
	"8.8.4.4",
	"193.168.1.1",
}

func TestReserved(t *testing.T) {
	for _, v := range reserved {
		if !IsReservedString(v) {
			t.Fatalf("%s should have been reserved", v)
		}
	}
}

func TestNotReserved(t *testing.T) {
	for _, v := range notReserved {
		if IsReservedString(v) {
			t.Fatalf("%s should not have been reserved", v)
		}
	}
}
