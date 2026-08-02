package main

import (
	"net"
	"testing"
)

func hopsFromIPs(ips ...string) []*Hop {
	hops := make([]*Hop, 0, len(ips))
	for i, s := range ips {
		hops = append(hops, &Hop{
			Distance: i + 1,
			Nodes:    []*Node{{IP: net.ParseIP(s)}},
		})
	}
	return hops
}

func TestClassifyPath_CN2GT(t *testing.T) {
	// 对齐用户实测：218.30(AS4134) → 59.43(AS4809)，应标 CN2GT 而非纯 CN2
	hops := hopsFromIPs("45.194.1.1", "10.80.1.1", "218.30.1.1", "59.43.1.1", "111.74.1.1")
	got := classifyPath(hops)
	if got != "CN2GT" {
		t.Fatalf("got %q, want CN2GT", got)
	}
}

func TestClassifyPath_CN2Only(t *testing.T) {
	hops := hopsFromIPs("45.194.1.1", "59.43.1.1", "111.74.1.1")
	got := classifyPath(hops)
	if got != "AS4809" {
		t.Fatalf("got %q, want AS4809", got)
	}
}

func TestClassifyPath_CMIN2(t *testing.T) {
	hops := hopsFromIPs("45.194.1.1", "223.120.128.1", "221.183.1.1")
	got := classifyPath(hops)
	if got != "AS58807" {
		t.Fatalf("got %q, want AS58807", got)
	}
}

func TestClassifyPath_163Only(t *testing.T) {
	hops := hopsFromIPs("45.194.1.1", "202.97.1.1", "111.74.1.1")
	got := classifyPath(hops)
	if got != "AS4134" {
		t.Fatalf("got %q, want AS4134", got)
	}
}
