package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

type Result struct {
	i int
	s string
}

var (
	rIp = []string{
		"219.141.136.12", "221.179.155.161", "202.106.50.1",
		"202.96.209.133", "211.136.112.200", "210.22.97.1",
		"58.60.188.222", "120.196.165.24", "210.21.196.6",
		"61.139.2.69", "211.137.96.205", "119.6.6.6",
		"36.111.200.100", "39.134.254.6", "42.48.16.100",
		"202.101.224.69", "211.141.90.68", "220.248.192.12",
	}
	rName = []string{
		"北京电信", "北京移动", "北京联通",
		"上海电信", "上海移动", "上海联通",
		"广州电信", "广州移动", "广州联通",
		"成都电信", "成都移动", "成都联通",
		"湖南电信", "湖南移动", "湖南联通",
		"江西电信", "江西移动", "江西联通",
	}
	ca = []color.Attribute{color.FgHiYellow, color.FgHiMagenta, color.FgHiBlue, color.FgHiGreen, color.FgHiCyan, color.FgHiRed, color.FgHiWhite}
	m  = map[string]string{
		"AS4134":     "电信163 [普通线路]",
		"CN2GT":      "电信CN2GT[混合线路]",
		"AS4809":     "电信CN2 [优质线路]",
		"AS4837":     "联通4837[普通线路]",
		"AS9929":     "联通9929[优质线路]",
		"AS9808":     "移动CMI [普通线路]",
		"AS58453":    "移动CMI [普通线路]",
		"AS58807":    "移动CMIN2[优质线路]",
	}
)

func trace(ch chan Result, i int) {
	hops, err := Trace(net.ParseIP(rIp[i]))
	if err != nil {
		s := fmt.Sprintf("%v %-15s %v", rName[i], rIp[i], err)
		ch <- Result{i, s}
		return
	}

	key := classifyPath(hops)
	if key == "" {
		s := fmt.Sprintf("%v %-15s %v", rName[i], rIp[i], "测试超时")
		ch <- Result{i, s}
		return
	}

	as := m[key]
	var c *color.Color
	switch {
	case strings.Contains(as, "[优质线路]"):
		c = color.New(color.FgHiGreen).Add(color.Bold)
	case strings.Contains(as, "[混合线路]"):
		c = color.New(color.FgHiCyan).Add(color.Bold)
	default:
		c = color.New(color.FgHiYellow).Add(color.Bold)
	}
	s := fmt.Sprintf("%v %-15s %-23s", rName[i], rIp[i], c.Sprint(as))
	ch <- Result{i, s}
}

// classifyPath 扫描整条路径上的可识别 ASN，再按优先级判定。
// 避免「首个前缀命中即 return」：例如先经 AS4134(218.30) 再经 59.43，
// 旧逻辑会漏掉 163、误报纯 CN2；现改为 4134+4809 → CN2GT。
func classifyPath(hops []*Hop) string {
	seen := map[string]bool{}
	var order []string
	for _, h := range hops {
		for _, n := range h.Nodes {
			asn := ipAsn(n.IP.String())
			if asn == "" || seen[asn] {
				continue
			}
			seen[asn] = true
			order = append(order, asn)
		}
	}
	if len(order) == 0 {
		return ""
	}

	has := func(a string) bool { return seen[a] }

	// 精品线优先（与常见 backtrace 系工具一致）
	if has("AS9929") {
		return "AS9929"
	}
	if has("AS58807") {
		return "AS58807"
	}
	// 电信：163 入境后再进 CN2 → GT；仅 CN2 → GIA；仅 163 → 163
	if has("AS4809") && has("AS4134") {
		return "CN2GT"
	}
	if has("AS4809") {
		return "AS4809"
	}
	if has("AS4134") {
		return "AS4134"
	}
	if has("AS4837") {
		return "AS4837"
	}
	if has("AS9808") {
		return "AS9808"
	}
	if has("AS58453") {
		return "AS58453"
	}
	return order[0]
}

func ipAsn(ip string) string {
	if isInIPRanges(ip) {
		return "AS58807"
	}

	switch {
	case strings.HasPrefix(ip, "59.43"):
		return "AS4809"
	// 电信 163：国际关口常见段（不止 202.97；如 218.30 出境跳）
	case strings.HasPrefix(ip, "202.97"), strings.HasPrefix(ip, "218.30"):
		return "AS4134"
	case strings.HasPrefix(ip, "218.105"), strings.HasPrefix(ip, "210.51"):
		return "AS9929"
	case strings.HasPrefix(ip, "219.158"):
		return "AS4837"
	case strings.HasPrefix(ip, "223.118"), strings.HasPrefix(ip, "223.119"), strings.HasPrefix(ip, "223.120"), strings.HasPrefix(ip, "223.121"):
		return "AS58453"
	default:
		return ""
	}
}

func isInIPRanges(ip string) bool {
	ipRanges := []string{
		"223.119.8.0/21",
		"223.119.32.0/24",
		"223.119.34.0/24",
		"223.119.35.0/24",
		"223.119.36.0/24",
		"223.119.37.0/24",
		"223.119.100.0/24",
		"223.120.128.0/17",
		"223.120.134.0/23",
		"223.120.138.0/23",
		"223.120.158.0/23",
		"223.120.164.0/22",
		"223.120.168.0/22",
		"223.120.172.0/22",
		"223.120.174.0/23",
		"223.120.184.0/22",
		"223.120.188.0/22",
		"223.120.192.0/23",
		"223.120.200.0/23",
		"223.120.210.0/23",
		"223.120.212.0/23",
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, ipRange := range ipRanges {
		_, subnet, err := net.ParseCIDR(ipRange)
		if err != nil {
			continue
		}

		if subnet.Contains(parsedIP) {
			return true
		}
	}

	return false
}
