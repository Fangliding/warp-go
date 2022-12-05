# Go 语言实现的第三方 WARP 客户端

[中文文档](https://gitlab.com/ProjectWARP/warp-go/-/blob/master/README.zh_CN.md) | [English Document](https://gitlab.com/ProjectWARP/warp-go/-/blob/master/README.md)

~~建议点个 Star~~ :D

## 命令行参数

配置文件

默认为 `warp.conf`

```
--config <config file name>
```

注册 WARP

默认保存配置文件到`warp.conf` 可通过`--config`参数修改位置

```
--register
```

注册 WARP 并 升级到 WARP+

```
--register --license <WARP+ license>
```

注册 WARP Team

Team 配置文件需要使用特别的方法获取,

请访问 [Warp Team Api](https://warp-team-api.herokuapp.com) 获取 `Token` 作为 `Team Config`

```
--register --team-config <WARP Team Config>
```

注册 WARP 并 自定义设备名

```
--register --device-name <Device Name>
```

升级到 WARP+

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

```
--update --license <WARP+ license>
```

修改设备名

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

```
--update --device-name <Device Name>
```

升级到 WARP+ 并 修改设备名

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

```
--update --license <WARP+ license> --device-name <Device Name>
```

移除 WARP 设备 并 删除配置文件

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

```
--remove
```

前台运行

_Linux/Darwin/FreeBSD 默认以 daemon 形式启动 即后台进程_

```
--foreground
```

生成 `WireGuard` 配置文件

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

注意: 生成完成后不要使用 `--remove` 卸载, 并保存好您的 `warp-go` 和 `WireGuard` 配置文件

```
--export-wireguard <File Name>
```

生成 `Sing-Box` Socks 配置文件

默认加载配置文件为`warp.conf` 可通过`--config`参数修改

注意: 生成完成后不要使用 `--remove` 卸载, 并保存好您的 `warp-go` 和 `Sing-Box` 配置文件

> Socks 监听地址为 127.0.0.1:2000

```
--export-singbox <File Name>
```

打印帮助信息

```
-h
```

打印版本号和版权信息

```
-v
```

## 用法

1.先向 WARP 服务器进行注册 如果有 WARP+ License 可以带上 `--license` 参数

```
warp-go --register
```

2.启动 Warp-Go

```
warp-go --foreground
```

## 注意事项

- OpenVZ, LXC, Docker 容器请开放 Tun 权限
- 以 Daemon 运行时如果先关闭程序请使用 `kill -15 <PID>` (模拟 Ctrl+C 正常关闭程序)
- 建议以 systemd 服务的形式启动本程序

## 配置文件解析

配置文件使用 ini 的格式, 与 WireGuard 配置文件略有不同

- 账户部分

此部分注册时会自动生成, 请勿修改

```
[Account]
Device     = <Device ID>
PrivateKey = <WireGuard Private Key>
Token      = <Cloudflare API Token>
Team       = <Boolean>
```

- Peer 部分

此部分注册时会自动生成

`PublicKey` 字段无效, 仅供参考

`Endpoint` 字段仅接受 `IP:端口` 的格式 不支持域名

`Endpoint6` 字段无效, 仅供参考

`KeepAlive` 字段就是 WireGuard 的 `PersistentKeepalive` 字段, 用于 NAT 后保持 UDP 会话活跃, 单位: 秒

`AllowedIPs` 字段用于连接成功以后自动添加路由表, 默认生成的配置没有此项, 您可以配合 PostUp 脚本自动添加路由表, 也可以使用本字段

_不填写 `AllowedIPs` 字段就相当于 `Table=off`_

```
[Peer]
PublicKey = <Warp Endpoint Public Key>
Endpoint  = <Warp Endpoint>
Endpoint6 = <Warp Endpoint V6>
KeepAlive = 30

#AllowedIPs = 0.0.0.0/0, ::/0
```

- Script 部分

此部分注册时不会自动生成, 需要手动追加到配置文件中, 与 `WireGuard` 配置一样

`PreUp` 字段用于初始化 `WireGuard-Go` 以前执行的命令行

`PostUp` 字段用于连接成功以后执行的命令行

`PostDown` 字段用于程序退出前执行的命令行

```
[Script]
#PreUp  = <Command>
#PostUP = <Command>
#PostDown = <Command>
```

## 版权声明

- Cloudflare 与 `ProjectWARP` 项目组拥有本项目的最终解释权
- 使用本项目即代表您同意 Cloudflare WARP 用户协议, 并承担一切违规使用的责任与后果
- 本项目基于 `MIT` 协议开源, 为了避免 DMCA 以及 滥用, 本储存库不存放任何获取 `Reserved` 字段相关源代码
- 如果您愿意加入 `ProjectWARP` 项目组, 可以 Email 到 `coiaprant@gmail.com` , 并附上您的 GitLab 用户名, 我们将在审核您的个人主页之后给予答复 (如果您在其他平台有相关储存库也请一并附上链接)

- 本仓库贡献维护者 [@CoiaPrant](https://gitlab.com/CoiaPrant)
