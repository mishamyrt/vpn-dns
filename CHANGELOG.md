# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].

## 0.0.7 — Unreleased

### Added

* `status` command.

### Fixed

* Log the output of failed commands instead of faceless errors.

### Changed

* The `fallback_servers` parameter has become optional.

## [0.0.6]

### Added

* Tests🔥:
    * `pkg/login`
    * `pkg/config`
    * `pkg/vpn`
    * `pkg/exec`
    * `pkg/network`
    * `pkg/process`
* `autostart` command help and verbosity.
* `start`  command verbosity.
* **This changelog**.
* Missing comments.

### Changed

* Split packages.
* Improved documentation.

## [0.0.5][]

### Fixed

* Swap autostart messages.
* Overwrite binary on installing.

## [0.0.4][]

### Added

* Autostart control.

## [0.0.3][]

### Added

* Installer script.
* Linting.

### Fixed

* Improve code style.

## [0.0.2][]

### Fixed

* Remove UPX packing (broke the arm64 build, I'll figure it out later).

## [0.0.1][]

Initial release

[keep a changelog]: https://keepachangelog.com/en/1.0.0/

[semantic versioning]: https://semver.org/spec/v2.0.0.html

[0.0.6]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.6

[0.0.5]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.5

[0.0.4]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.4

[0.0.3]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.3

[0.0.2]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.2

[0.0.1]: https://github.com/mishamyrt/vpn-dns/releases/tag/v0.0.1
