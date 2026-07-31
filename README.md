# 三网回程路由测试（含江西 / 赣州）

基于 [zhanghanyun/backtrace](https://github.com/zhanghanyun/backtrace) / [ludashi2020/backtrace](https://github.com/ludashi2020/backtrace)，在原有省市节点下方增加江西移动与赣州电信/联通。

新增目标：

| 名称 | IP | 备注 |
| --- | --- | --- |
| 江西移动 | `211.141.90.68` | 江西移动省内专用 DNS（南昌） |
| 赣州电信 | `218.87.136.7` | |
| 赣州联通 | `220.248.192.12` | |

暂无单独「赣州移动」节点：公开常用地址与省网 DNS 相同，没有可靠的赣州本地探测 IP。

## 使用

```shell
curl https://cdn.jsdelivr.net/gh/wangyuling93/backtrace-jxyd@main/install.sh -sSf | sh
```

## 示例输出

```text
项目地址：github.com/wangyuling93/backtrace-jxyd
正在测试三网回程路由...
北京电信 219.141.136.12  电信CN2 [优质线路]
...
湖南移动 39.134.254.6    电信CN2 [优质线路]
江西移动 211.141.90.68   电信CN2 [优质线路]
赣州电信 218.87.136.7    电信CN2 [优质线路]
赣州联通 220.248.192.12  电信CN2 [优质线路]
```
