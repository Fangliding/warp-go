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

Register for WARP and upgrade to WARP+

```
--register --license <WARP+ license>
```

Running in the foreground

_Linux/Darwin/FreeBSD starts as a daemon by default, that is, a background process_

```
--foreground
```

print help information

```
-h
```

Print version number and copyright information

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
```

- Peer section

This section is automatically generated when you register

`Endpoint` field only accepts the format `IP:port`. Domain names are not supported

`Endpoint6` field is invalid, for informational purposes only

`KeepAlive` field is the `PersistentKeepalive` field of WireGuard, used to keep the UDP session active after NAT

`AllowedIPs` field is used to automatically add the routing table after the connection is successful. The default generated configuration does not have this option. You can automatically add the routing table with the PostUPJ script, or you can use this field (due to the parser problem of `go-ini` does not support Multiple lines of `AllowedIPs` use `,` to separate multiple CIDR entries)

```
[Peer]
Endpoint  = <Warp Endpoint>
Endpoint6 = <Warp Endpoint V6>
KeepAlive = 30

#AllowedIPs = 0.0.0.0/0, ::/0
```

- Script section

Will this part be automatically generated during registration? It needs to be manually appended to the configuration file

`PreUp` field is used to initialize the command line previously executed by `WireGuard-Go` (due to the parser problem of `go-ini` does not support multi-line `PreUp`, please separate multiple command lines with `;`, or save as one script)

`PostUp` field is used to execute the command line after the connection is successful (due to the parser problem of `go-ini`, multi-line `PostUp` is not supported, please use `;` to separate multiple command lines, or save it as a script)

`PostDown` field is used for the command line executed before the program exits (due to the parser problem of `go-ini`, multi-line `PostDown` is not supported, please use `;` to separate multiple command lines, or save as a script)

```
[Script]
#PreUp = <Command>
#PostUP = <Command>
#PostDown = <Command>
```

## Copyright

- Cloudflare and the `ProjectWARP` project team have the final interpretation right of this project
- By using this project, you agree to the Cloudflare WARP User Agreement and assume all responsibilities and consequences for illegal use
- This project is open source based on the `MIT` protocol. In order to avoid DMCA, this repository does not store any related source code. If you need it, you can Email to `coiaprant@gmail.com` to get the source code
- If you are willing to join the `ProjectWARP` project group, you can Email to `coiaprant@gmail.com` with your GitLab username, we will reply after reviewing your profile (if you have related information on other platforms Please include a link to the repository as well)
- Brother Yong Copy Copy (kkkyg) Please consciously stay away from this project

- Contributing maintainer of this repository [@CoiaPrant](https://gitlab.com/CoiaPrant)

*Translate by Google* :D
