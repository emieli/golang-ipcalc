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
		address, network, broadcast string
	)

	if prefix.Addr().Is4() {
		address, network, broadcast = prefix4(prefix)
	}

	err = templates.Prefix(address, network, broadcast).Render(r.Context(), w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func prefix4(prefix netip.Prefix) (address string, network string, broadcast string) {

	address = prefix.Addr().String()
	hostBits := 32 - uint32(prefix.Bits())
	addressesInSubnet := uint32(math.Pow(2, float64(hostBits)))

	// // Netmask
	// netmaskAsInt := uint32(math.MaxUint32-addressesInSubnet) + 1
	// netmaskAsSlice := new(bytes.Buffer)
	// err := binary.Write(netmaskAsSlice, binary.BigEndian, netmaskAsInt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("netmask: %08b\n", netmaskAsSlice.Bytes())

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
