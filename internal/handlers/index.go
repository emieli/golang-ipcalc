package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"ipcalc/internal/templates"
	"log"
	"math"
	"net/http"
	"net/netip"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := templates.Hello("Emil").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func Prefix(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	input := r.PostForm.Get("prefix")
	input = strings.TrimSpace(input)
	prefix, err := netip.ParsePrefix(input)
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid prefix %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var (
		address, first, last string
	)

	if prefix.Addr().Is4() {
		address, first, last = prefix4(prefix)
		err = templates.Prefix4(address, first, last).Render(r.Context(), w)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}

	if prefix.Addr().Is6() {
		address, first, last = prefix6(prefix)
		err = templates.Prefix6(address, first, last).Render(r.Context(), w)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}

func prefix4(prefix netip.Prefix) (address string, network string, broadcast string) {

	address = prefix.Addr().String()
	hostBits := 32 - uint32(prefix.Bits())
	addressesInSubnet := uint32(math.Pow(2, float64(hostBits)))

	// Network address
	// We shift the rightmost bits right and then left again.
	// This resets them to all zeroes, giving us the network address.
	addressAsInt := binary.BigEndian.Uint32(prefix.Addr().AsSlice())
	networkAsInt := addressAsInt >> hostBits << hostBits
	networkAsSlice := new(bytes.Buffer)
	err := binary.Write(networkAsSlice, binary.BigEndian, networkAsInt)
	if err != nil {
		panic(err)
	}
	networkAddr := netip.AddrFrom4([4]byte(networkAsSlice.Bytes()))
	network = networkAddr.String()

	// Broadcast address
	broadcastAsInt := networkAsInt + addressesInSubnet - 1
	broadcastAsSlice := new(bytes.Buffer)
	err = binary.Write(broadcastAsSlice, binary.BigEndian, broadcastAsInt)
	if err != nil {
		panic(err)
	}
	broadcastAddr := netip.AddrFrom4([4]byte(broadcastAsSlice.Bytes()))
	broadcast = broadcastAddr.String()

	return
}

// Find first and last address in subnet.
// IPv6 addresses are 128 bits long, but largest int in Golang is uint64, so we have to split
// the address into a top and bottom part that are processed individually.
func prefix6(prefix netip.Prefix) (address string, first string, last string) {

	address = prefix.Addr().String()

	top := binary.BigEndian.Uint64(prefix.Addr().AsSlice()[0:8])
	bottom := binary.BigEndian.Uint64(prefix.Addr().AsSlice()[8:16])

	// First address
	bytes := make([]byte, 0, 16)
	var networkAddr netip.Addr

	if prefix.Bits() > 64 {
		bitDiff := 128 - prefix.Bits()
		networkAddrAsInt := bottom >> uint64(bitDiff) << uint64(bitDiff)
		bytes = binary.BigEndian.AppendUint64(bytes, top)
		bytes = binary.BigEndian.AppendUint64(bytes, networkAddrAsInt)
		networkAddr, _ = netip.AddrFromSlice(bytes)
	} else {
		bitDiff := 64 - prefix.Bits()
		networkAddrAsInt := top >> uint64(bitDiff) << uint64(bitDiff)
		bytes = binary.BigEndian.AppendUint64(bytes, networkAddrAsInt)
		bytes = binary.BigEndian.AppendUint64(bytes, 0)
		networkAddr, _ = netip.AddrFromSlice(bytes)
	}
	first = networkAddr.String()

	// Last address
	bytes = make([]byte, 0, 16)
	top = binary.BigEndian.Uint64(networkAddr.AsSlice()[0:8])
	bottom = binary.BigEndian.Uint64(networkAddr.AsSlice()[8:16])
	var lastAddr netip.Addr

	if prefix.Bits() > 64 {
		hostBits := 128 - prefix.Bits()
		addressesInSubnet := uint64(math.Pow(2, float64(hostBits)))
		lastAddressAsInt := bottom + addressesInSubnet - 1

		bytes = binary.BigEndian.AppendUint64(bytes, top)
		bytes = binary.BigEndian.AppendUint64(bytes, lastAddressAsInt)
		lastAddr, _ = netip.AddrFromSlice(bytes)
	} else {
		hostBits := 64 - prefix.Bits()
		addressesInSubnet := uint64(math.Pow(2, float64(hostBits)))
		lastAddressAsInt := top + addressesInSubnet - 1

		bytes = binary.BigEndian.AppendUint64(bytes, lastAddressAsInt)
		bytes = binary.BigEndian.AppendUint64(bytes, math.MaxUint64)
		lastAddr, _ = netip.AddrFromSlice(bytes)
	}
	last = lastAddr.String()

	return
}
