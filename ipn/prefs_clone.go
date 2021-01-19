// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by tailscale.com/cmd/cloner -type Prefs; DO NOT EDIT.

package ipn

import (
	"inet.af/netaddr"
	"tailscale.com/control/controlclient"
	"tailscale.com/types/preftype"
)

// Clone makes a deep copy of Prefs.
// The result aliases no memory with the original.
func (src *Prefs) Clone() *Prefs {
	if src == nil {
		return nil
	}
	dst := new(Prefs)
	*dst = *src
	dst.AdvertiseTags = append(src.AdvertiseTags[:0:0], src.AdvertiseTags...)
	dst.AdvertiseRoutes = append(src.AdvertiseRoutes[:0:0], src.AdvertiseRoutes...)
	if dst.Persist != nil {
		dst.Persist = new(controlclient.Persist)
		*dst.Persist = *src.Persist
	}
	return dst
}

// A compilation failure here means this code must be regenerated, with command:
//   tailscale.com/cmd/cloner -type Prefs
var _PrefsNeedsRegeneration = Prefs(struct {
	ControlURL            string
	RouteAll              bool
	AllowSingleHosts      bool
	CorpDNS               bool
	WantRunning           bool
	ShieldsUp             bool
	RouteAllVia           string
	AdvertiseTags         []string
	Hostname              string
	OSVersion             string
	DeviceModel           string
	NotepadURLs           bool
	ForceDaemon           bool
	AdvertiseRoutes       []netaddr.IPPrefix
	AdvertiseDefaultRoute bool
	NoSNAT                bool
	NetfilterMode         preftype.NetfilterMode
	Persist               *controlclient.Persist
}{})
