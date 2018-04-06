# gRPC-book-example

Based loosely on models from [swapi](https://swapi.co/).

Lint proto with:

```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) gwihlidal/protoc $(pwd)/proto/swapi.proto -I$(pwd) --lint_out=.
```