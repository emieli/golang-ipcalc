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
	gotAddress, gotNetwork, gotBroadcast := prefix4(prefix)

	expectedAddress := "1.2.3.4"
	expectedNetwork := "1.2.3.0"
	expectedBroadcast := "1.2.3.255"

	if gotAddress != expectedAddress {
		t.Errorf("expected %q but got %q", expectedAddress, gotAddress)
	}
	if gotNetwork != expectedNetwork {
		t.Errorf("expected %q but got %q", expectedNetwork, gotNetwork)
	}
	if gotBroadcast != expectedBroadcast {
		t.Errorf("expected %q but got %q", expectedBroadcast, gotBroadcast)
	}
}

func TestPrefix6(t *testing.T) {
	prefix, err := netip.ParsePrefix("2001::/3")
	if err != nil {
		t.Errorf("parse prefix: %q", err)
	}
	gotAddress, gotFirst, gotLast := prefix6(prefix)

	expectedAddress := "2001::"
	expectedFirst := "2000::"
	expectedLast := "3fff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"

	if gotAddress != expectedAddress {
		t.Errorf("expected %q but got %q", expectedAddress, gotAddress)
	}
	if gotFirst != expectedFirst {
		t.Errorf("expected %q but got %q", expectedFirst, gotFirst)
	}
	if gotLast != expectedLast {
		t.Errorf("expected %q but got %q", expectedLast, gotLast)
	}
}
