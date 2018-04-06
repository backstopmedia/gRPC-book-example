# gRPC-book-example

Based loosely on models from [swapi](https://swapi.co/).

## Getting Started

* Install Go
* Clone this repo into your $GOPATH
* `make setup`

Lint proto with:

```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) gwihlidal/protoc $(pwd)/proto/swapi.proto -I$(pwd) --lint_out=.
```
