# ONAP Docker Image Builder


[![Build Status](https://travis-ci.org/electrocucaracha/onap_builder.png)][1]
[![Go Report Card](https://goreportcard.com/badge/github.com/electrocucaracha/onap_builder
)][2]
[![GoDoc](https://godoc.org/github.com/electrocucaracha/onap_builder?status.svg)][3]

This golang project allows to generate ONAP Docker images from source code. Its 
main goal is centralize the process to generate images as well as having better
control of the creation process.

## Requirements

| Name   | Version |
|--------|---------|
| go     | +1.10.1 |
| docker | +18.03  |
| mvn    | +3.5.3  |
| git    | +2.14.3 |
| java   | +9.0.4  |

## Installation

    $ go get github.com/electrocucaracha/onap_builder

## Execution

    $ image_builder --config-file configs/globals.yml

## Contributing

Bug reports and patches are most welcome.

## License

Apache-2.0

[1]: https://travis-ci.org/electrocucaracha/onap_builder
[2]: https://goreportcard.com/report/github.com/electrocucaracha/onap_builder
[3]: https://godoc.org/github.com/electrocucaracha/onap_builder
