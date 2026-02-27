# checkport

[![go][go-src]][go-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

A small tool, written in Go, to chekc if a port is available or not.

## Install

```bash
go install github.com/ewilan-riviere/checkport@latest
```

## Usage

Just check a port:

```bash
checkport 3000
```

Use `sudo` to know which process uses this port:

```bash
sudo checkport 3000
```

## Test

```sh
go test -v
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.25&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/checkport/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://packagist.org/packages/ewilan-riviere/checkport
[license-src]: https://img.shields.io/github/license/ewilan-riviere/checkport.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/checkport/blob/main/LICENSE
