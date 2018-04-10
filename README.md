# ONAP Docker Image Builder

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

    $ cd onap_builder
    $ export GOPATH="$(pwd)"
    $ cd src/github.com/electrocucaracha/image_builder
    $ dep ensure
    $ go install
    $ export PATH=$PATH:$(go env GOPATH)/bin

## Execution

    $ image_builder --config-file $GOPATH/globals.yml

## Contributing

Bug reports and patches are most welcome.

## License

Apache-2.0
