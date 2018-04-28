# nodejs example

Simple example gRPC server.

* Install dependencies with `npm i`
* Run server & gateway with `npm start`
* Test server with `npm test`

Additionally, you can install [grpcnode](https://www.npmjs.com/package/grpcnode), globally with `npm i grpcnode -g`.

Then you can do this to get a list of all gRPCs:

```sh
grpcnode client ls -I ../../proto sfapi.proto
```

You can do so much more!

I like to use [jq](https://stedolan.github.io/jq/) to pull stuff out.

Here is an example of getting a single person id off the top:
```sh
grpcnode client run -I ../../proto sfapi.proto  -c '/sfapi.v1.Starfriends/ListPeople()' | jq '.people[0].id'
```

Let's use it to get that person (Luke):
```sh
grpcnode client run -I ../../proto sfapi.proto  -c '/sfapi.v1.Starfriends/GetPerson({"id": "SyAbJp35ViM"})'
```

You can also watch for streaming ship events:

```sh
grpcnode client run -I ../../proto sfapi.proto -c '/sfapi.v1.Starfriends/ListStarshipActions()'
```