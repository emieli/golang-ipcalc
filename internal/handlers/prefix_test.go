package handlers

import (
	"net/netip"
	"testing"
)

func TestPrefix4(t *testing.T) {
	prefix, err := netip.ParsePrefix("1.2.3.4/24")
	if err != nil {
		t.Errorf("parse prefix: %q", err)
	}
	gotAddress, gotNetwork, gotBroadcast, gotHosts := prefix4(prefix)

	expectedAddress := "1.2.3.4"
	expectedNetwork := "1.2.3.0"
	expectedBroadcast := "1.2.3.255"
	expectedHosts := "254"

	if gotAddress != expectedAddress {
		t.Errorf("expected %q but got %q", expectedAddress, gotAddress)
	}
	if gotNetwork != expectedNetwork {
		t.Errorf("expected %q but got %q", expectedNetwork, gotNetwork)
	}
	if gotBroadcast != expectedBroadcast {
		t.Errorf("expected %q but got %q", expectedBroadcast, gotBroadcast)
	}
	if gotHosts != expectedHosts {
		t.Errorf("expected %q but got %q", gotHosts, expectedHosts)
	}
}

func TestPrefix6(t *testing.T) {
	prefix, err := netip.ParsePrefix("2001::/33")
	if err != nil {
		t.Errorf("parse prefix: %q", err)
	}
	gotAddress, gotFirst, gotLast, got64s := prefix6(prefix)

	expectedAddress := "2001::"
	expectedFirst := "2001::"
	expectedLast := "2001:0:7fff:ffff:ffff:ffff:ffff:ffff"
	expected64s := "2B"

	if gotAddress != expectedAddress {
		t.Errorf("expected %q but got %q", expectedAddress, gotAddress)
	}
	if gotFirst != expectedFirst {
		t.Errorf("expected %q but got %q", expectedFirst, gotFirst)
	}
	if gotLast != expectedLast {
		t.Errorf("expected %q but got %q", expectedLast, gotLast)
	}
	if got64s != expected64s {
		t.Errorf("expected %q but got %q", got64s, expected64s)
	}
}
