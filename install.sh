#!/bin/bash
set -euo pipefail

arch=$(uname -m)
if [ "$arch" = "x86_64" ]; then
  url="https://github.com/wangyuling93/backtrace-jxyd/releases/latest/download/backtrace-linux-amd64.tar.gz"
else
  url="https://github.com/wangyuling93/backtrace-jxyd/releases/latest/download/backtrace-linux-arm64.tar.gz"
fi

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
cd "$tmp"

if command -v wget >/dev/null 2>&1; then
  wget -q -O backtrace.tar.gz "$url"
else
  curl -fsSL -o backtrace.tar.gz "$url"
fi

tar --warning=no-unknown-keyword -xf backtrace.tar.gz 2>/dev/null || tar -xf backtrace.tar.gz
chmod +x backtrace

# raw ICMP sockets need root (or CAP_NET_RAW)
if [ "$(id -u)" -eq 0 ]; then
  ./backtrace
else
  sudo ./backtrace
fi
