# Third-party WARP client implemented in Go

[中文文档](https://gitlab.com/ProjectWARP/warp-go/-/blob/master/README.zh_CN.md) | [English Document](https://gitlab.com/ProjectWARP/warp-go/-/blob/master/README.md)

~~Please give me a Star~~ :D

## Command line arguments

Configuration file

The default value is `warp.conf`

```
--config <config file name>
```

Register for WARP

The configuration file is saved to `warp.conf` by default, and the location can be modified by the `--config` parameter

```
--register
```

Register WARP and customize device name

```
--register --device-name <Device Name>
```

Sign up for WARP Team

Team configuration files need to be obtained using a special method

Please visit [Warp Team Api](https://web--public--warp-team-api--coia-mfs4.code.run) to get `Token` as `Team Config`

```
--register --team-config <WARP Team Config>
```

Register for WARP and upgrade to WARP+

```
--register --license <WARP+ license>
```

Upgrade to WARP+

The default loading configuration file is `warp.conf`, which can be modified by the `--config` parameter

```
--update --license <WARP+ license>
```

Change device name

The default loading configuration file is `warp.conf`, which can be modified by the `--config` parameter

```
--update --device-name <Device Name>
```

Upgrade to WARP+ and change device name

The default loading configuration file is `warp.conf`, which can be modified by the `--config` parameter

```
--update --license <WARP+ license> --device-name <Device Name>
```

Reset Private Key

The default loaded configuration file is `warp.conf`, which can be modified by the `--config` parameter

```

--update --reset-key

```

Remove the WARP device and delete the configuration file

The default loading configuration file is `warp.conf`, which can be modified by the `--config` parameter

```
--remove
```

Running in the foreground

_Linux/Darwin/FreeBSD starts as a daemon by default, that is, a background process_

```
--foreground
```

Generate `WireGuard` configuration file

The default loading configuration file is `warp.conf`, which can be modified by the `--config` parameter

WARNING: Do not use `--remove` to uninstall after build, and save your `warp-go` and `WireGuard` configuration files

```
--export-wireguard <File Name>
```

Generate `Sing-Box` Socks configuration file

The default loaded configuration file is `warp.conf`, which can be modified by the `--config` parameter

Note: Do not use `--remove` to uninstall after the build is complete, and save your `warp-go` and `Sing-Box` configuration files

> Socks listening address is 127.0.0.1:2000

```
--export-singbox <File Name>
```

Print help guide

```
-h
```

Print version and copyright

```
-v
```

## Usage

1. Register with the WARP server first. If you have WARP+ License, you can bring the `--license` parameter

```
warp-go --register
```

2. Start Warp-Go

```
warp-go --foreground
```

## Tips

- OpenVZ, LXC, Docker container please open Tun permission
- When running as Daemon, use `kill -15 <PID>` to close the program first (simulate Ctrl+C to close the program normally)
- It is recommended to start this program as a systemd service

## Configuration file

The configuration file uses the ini format, which is slightly different from the WireGuard configuration file

- Account section

This part will be automatically generated during registration, please do not modify it

```
[Account]
Device     = <Device ID>
PrivateKey = <WireGuard Private Key>
Token      = <Cloudflare API Token>
Type       = <free / plus / team>
```

- Device section

This section is automatically generated when you sign up

`Name` TUN device name, default value is `WARP`
`MTU` TUN device MTU, default value is `1280`

```
[Device]
Name = WARP
MTU=1280
```

- Peer section

This section is automatically generated when you register

`PublicKey` field will not be used, for informational purposes only

`Endpoint` field only accepts the format `IP:Port`. Domain names are not supported

`Endpoint6` field will not be used, for informational purposes only

`KeepAlive` field is the `PersistentKeepalive` field of WireGuard, used to keep the UDP session active after NAT

`AllowedIPs` field is used to automatically add the routing table after the connection is successful. The default generated configuration does not have this option. You can automatically add the routing table with the PostUp script, or you can use this field

_Not filling in the `AllowedIPs` field is equivalent to `Table=off`_

```
[Peer]
PublicKey = <Warp Endpoint Public Key>
Endpoint  = <Warp Endpoint>
Endpoint6 = <Warp Endpoint V6>
KeepAlive = 30

#AllowedIPs = 0.0.0.0/0, ::/0
```

- Script section

This part will not automatically generated during registration, It needs to be manually appended to the configuration file, It is same as the `WireGuard` configuration file

`PreUp` field is used to initialize the command line previously executed by `WireGuard-Go`

`PostUp` field is used to execute the command line after the connection is successful

`PostDown` field is used for the command line executed before the program exits

```
[Script]
#PreUp = <Command>
#PostUP = <Command>
#PostDown = <Command>
```

## Copyright

- Cloudflare and the `ProjectWARP` project team have the final interpretation right of this project
- By using this project, you agree to the Cloudflare WARP User Agreement and assume all responsibilities and consequences for illegal use
- This project is open source based on the `MIT` protocol. To avoid DMCA and abuse, this repository does not store any source code related to obtaining the `Reserved` field
- If you are willing to join the `ProjectWARP` project group, you can Email to `coiaprant@gmail.com` with your GitLab username, we will reply after reviewing your profile (if you have related information on other platforms Please include a link to the repository as well)

- Contributing maintainer of this repository [@CoiaPrant](https://gitlab.com/CoiaPrant)

_Translate by Google_ :D

# Donate

- USDT TRC20 Address `TNU2wK4yieGCWUxezgpZhwMHmLnRnXRtmu`
