# 三网回程路由测试（含江西 / 赣州）

基于 [zhanghanyun/backtrace](https://github.com/zhanghanyun/backtrace) / [ludashi2020/backtrace](https://github.com/ludashi2020/backtrace)，在原有省市节点下方增加赣州电信、江西移动、江西联通。每组按 **电信 → 移动 → 联通** 排列。

新增目标：

| 名称 | IP | 备注 |
| --- | --- | --- |
| 赣州电信 | `218.87.136.7` | 赣州本地（江西理工大学） |
| 江西移动 | `211.141.90.68` | 江西移动省内专用 DNS（南昌） |
| 江西联通 | `220.248.192.12` | 江西联通省网 DNS（南昌） |

## 使用

```shell
curl https://cdn.jsdelivr.net/gh/wangyuling93/backtrace-jxyd@main/install.sh -sSf | sh
```

## 示例输出

```text
项目地址：github.com/wangyuling93/backtrace-jxyd
正在测试三网回程路由...
北京电信 219.141.136.12  电信CN2 [优质线路]
北京移动 221.179.155.161 电信CN2 [优质线路]
北京联通 202.106.50.1    电信CN2 [优质线路]
...
湖南电信 36.111.200.100  电信CN2 [优质线路]
湖南移动 39.134.254.6    电信CN2 [优质线路]
湖南联通 42.48.16.100    电信CN2 [优质线路]
赣州电信 218.87.136.7    电信CN2 [优质线路]
江西移动 211.141.90.68   电信CN2 [优质线路]
江西联通 220.248.192.12  电信CN2 [优质线路]
```
