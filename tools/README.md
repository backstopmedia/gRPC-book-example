# tools

This is just some quick tooling I made to document how I designed the protos and got the data.

It can totally be deleted once we have all our definitions & data in place.

```bash
npm i            # install dependencies
npm run data     # get all the data from the API
npm run proto    # generate the gRPC protobuf that matches the API
npm run validate # validate the generated data against the generated protobuf

npm run grpc     # start a demo gRPC server
npm run gateway  # start a demo gRPC gateway
npm start        # run both server & gateway
```