# VPN DNS Changer [![Quality assurance](https://github.com/mishamyrt/vpn-dns/actions/workflows/qa.yaml/badge.svg)](https://github.com/mishamyrt/vpn-dns/actions/workflows/qa.yaml) [![Maintainability](https://api.codeclimate.com/v1/badges/0feb5c97955ba991b140/maintainability)](https://codeclimate.com/github/mishamyrt/vpn-dns/maintainability)

The service that changes DNS settings when connecting to a VPN. Solves some problems on macOS Ventura.

## Installing

To install automatically, run the command.

```sh
curl -s https://raw.githubusercontent.com/mishamyrt/vpn-dns/main/scripts/install_latest.py | python3
```

## Configuration

Before you start, you must create a configuration file:

```sh
mkdir ~/.config/vpn-dns
vi ~/.config/vpn-dns/config.yaml
```

The configuration consists of the following keys:

* `interface` — network interface name. The macbook is likely to have `Wi-Fi`.
* `VPNs` — list of VPN connection settings.
    * `name` — name of the VPN connection. The exact name can be seen in the output of the `scutil --nc list` command (what is written in "quotes").
    * `servers` — list of DNS that will be set if the connection is active.
* `fallback_servers` — list of DNS that will be set if none of the VPN connections listed are active.

If several connections are active, the DNS lists will summarise. Priority corresponds to the order in the file: higher priority is higher.

An example can be seen in the file [basic-config.yaml](./testdata/basic-config.yaml).

## Usage

Commands are available to control the application:

* `start` — starts the application in the background.
* `stop` — stops the background application.
* `run` — runs the application in the current process without daemonization.
* `autostart` — controls the automatic start-up of the application.

### Basic

```sh
vpn-dns start
```

### Custom configuration

You can pass the path to the configuration file in the `--config` argument.

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

## Development

The repository is set up with auxiliary [Makefile](./Makefile) targets:

* `build` — compile the project.
* `clear` — remove all build artifacts.
* `test` — run unit tests.
* `lint` — check code style. Requires `golangci-lint` and `revive` to be installed.
* `coverage` — generate and open test coverage HTML report.
* `install` — installs the built version on the system. To work correctly, it should be run with `sudo`.

## License

[GPL-3.0](./LICENSE).
