# 三网回程路由测试（含江西）

基于 [zhanghanyun/backtrace](https://github.com/zhanghanyun/backtrace) / [ludashi2020/backtrace](https://github.com/ludashi2020/backtrace)，在原有省市节点下方增加江西电信、江西移动、江西联通。每组按 **电信 → 移动 → 联通** 排列。

## 相对上游的改动

- 新增江西三网探针
- **整路径判定**：扫描路径上所有可识别 ASN，不再「首个前缀命中即定论」
- **CN2GT**：同时出现 AS4134 与 AS4809 时标为 `电信CN2GT[混合线路]`（避免把 163→CN2 误报成纯优质 CN2）
- 补充 `218.30.*` 为电信 163 特征前缀
- 明确探测协议为 **ICMP**（与 TCP/UDP 回程可能不同）

新增目标：

| 名称 | IP | 备注 |
| --- | --- | --- |
| 江西电信 | `202.101.224.69` | 江西电信省网 DNS（南昌） |
| 江西移动 | `211.141.90.68` | 江西移动省内专用 DNS（南昌） |
| 江西联通 | `220.248.192.12` | 江西联通省网 DNS（南昌） |

## 使用

```shell
curl https://cdn.jsdelivr.net/gh/wangyuling93/backtrace-jxyd@main/install.sh -sSf | sh
```

需要 root（或 `CAP_NET_RAW`）。重新发布二进制后，`install.sh` 才会拉到新逻辑。

## 协议说明

本工具只用 **ICMP** traceroute。运营商对 ICMP / TCP / UDP 常走不同策略路由，结果可能和 NetQuality 的 TCP/UDP 不一致。若要对比协议，请用：

```shell
bash <(curl -Ls https://cdn.jsdelivr.net/gh/wangyuling93/NetQuality@main/net.sh) -R 江西
```

## 示例输出

```text
项目地址：github.com/wangyuling93/backtrace-jxyd
正在测试三网回程路由（ICMP；与 TCP/UDP 可能不一致）...
北京电信 219.141.136.12  电信CN2 [优质线路]
北京移动 221.179.155.161 移动CMIN2[优质线路]
北京联通 202.106.50.1    移动CMIN2[优质线路]

...

江西电信 202.101.224.69  电信CN2GT[混合线路]
江西移动 211.141.90.68   移动CMIN2[优质线路]
江西联通 220.248.192.12  移动CMIN2[优质线路]
```
