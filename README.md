# VPN DNS [![Quality assurance](https://github.com/mishamyrt/vpn-dns/actions/workflows/qa.yaml/badge.svg)](https://github.com/mishamyrt/vpn-dns/actions/workflows/qa.yaml) [![Maintainability](https://api.codeclimate.com/v1/badges/0feb5c97955ba991b140/maintainability)](https://codeclimate.com/github/mishamyrt/vpn-dns/maintainability)

The service that changes DNS servers when connecting to a VPN. Solves connection problems on macOS.

## Installing

To install automatically, run the command.

```sh
curl -s https://raw.githubusercontent.com/mishamyrt/vpn-dns/main/scripts/install_latest.py | python3
```

## Usage

### Configuration

Before you start, you must create a configuration file:

```sh
mkdir ~/.config/vpn-dns
vi ~/.config/vpn-dns/config.yaml
```

Example content:

```yaml
---
interface: Wi-Fi
VPNs:
  - name: AbdtVPN
    servers:
      - 10.129.144.6
fallback_servers:
  - 1.1.1.1
```

### Starting

Commands are available to control the application:

* `start` — Starts the application in the background
* `stop` — Stops the background application
* `run` — Runs the application in the current process

Additionally, you can pass the path to the configuration file in the `--config` argument.

```sh
vpn-dns --config myconfig.yaml start
```

### Automatic startup

To start the application automatically at system startup, run the command:

```sh
vpn-dns autostart enable
```

To disable, run the command:

```sh
vpn-dns autostart disable
```
