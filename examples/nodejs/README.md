# nodejs example

Simple example gRPC server.

* Install dependencies with `npm i`
* Run server & gateway with `npm start`
* Test server with `npm test`

Additionally, you can install [grpcnode](https://www.npmjs.com/package/grpcnode), globally with `npm i grpcnode -g`.

Then you can do this to get a list of all gRPCs:

```sh
grpcnode client ls -I ../../proto swapi.proto
```