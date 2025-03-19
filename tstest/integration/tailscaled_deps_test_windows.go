// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by gen_deps.go; DO NOT EDIT.

package integration

import (
	// And depend on a bunch of tailscaled innards, for Go's test caching.
	// Otherwise cmd/go never sees that we depend on these packages'
	// transitive deps when we run "go install tailscaled" in a child
	// process and can cache a prior success when a dependency changes.
	_ "github.com/dblohm7/wingoes/com"
	_ "github.com/tailscale/wireguard-go/tun"
	_ "golang.org/x/sys/windows"
	_ "golang.org/x/sys/windows/svc"
	_ "golang.org/x/sys/windows/svc/eventlog"
	_ "golang.org/x/sys/windows/svc/mgr"
	_ "golang.zx2c4.com/wintun"
	_ "golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
	_ "tailscale.com/client/local"
	_ "tailscale.com/cmd/tailscaled/childproc"
	_ "tailscale.com/control/controlclient"
	_ "tailscale.com/derp/derphttp"
	_ "tailscale.com/drive/driveimpl"
	_ "tailscale.com/envknob"
	_ "tailscale.com/feature/condregister"
	_ "tailscale.com/health"
	_ "tailscale.com/hostinfo"
	_ "tailscale.com/ipn"
	_ "tailscale.com/ipn/auditlog"
	_ "tailscale.com/ipn/conffile"
	_ "tailscale.com/ipn/desktop"
	_ "tailscale.com/ipn/ipnlocal"
	_ "tailscale.com/ipn/ipnserver"
	_ "tailscale.com/ipn/store"
	_ "tailscale.com/logpolicy"
	_ "tailscale.com/logtail"
	_ "tailscale.com/logtail/backoff"
	_ "tailscale.com/net/dns"
	_ "tailscale.com/net/dnsfallback"
	_ "tailscale.com/net/netmon"
	_ "tailscale.com/net/netns"
	_ "tailscale.com/net/proxymux"
	_ "tailscale.com/net/socks5"
	_ "tailscale.com/net/tsdial"
	_ "tailscale.com/net/tshttpproxy"
	_ "tailscale.com/net/tstun"
	_ "tailscale.com/paths"
	_ "tailscale.com/safesocket"
	_ "tailscale.com/syncs"
	_ "tailscale.com/tailcfg"
	_ "tailscale.com/tsd"
	_ "tailscale.com/tsweb/varz"
	_ "tailscale.com/types/flagtype"
	_ "tailscale.com/types/key"
	_ "tailscale.com/types/logger"
	_ "tailscale.com/types/logid"
	_ "tailscale.com/util/clientmetric"
	_ "tailscale.com/util/eventbus"
	_ "tailscale.com/util/multierr"
	_ "tailscale.com/util/osdiag"
	_ "tailscale.com/util/osshare"
	_ "tailscale.com/util/syspolicy"
	_ "tailscale.com/util/winutil"
	_ "tailscale.com/util/winutil/gp"
	_ "tailscale.com/version"
	_ "tailscale.com/version/distro"
	_ "tailscale.com/wf"
	_ "tailscale.com/wgengine"
	_ "tailscale.com/wgengine/netstack"
	_ "tailscale.com/wgengine/router"
)
