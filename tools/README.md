# tools

These are quick-hack tools to generate intiial data and proto from [SWAPI](https://swapi.co/). These are here for the purpose of seeing how our initial data was generated.

## you probly shouldn't run them

After these were used, the data and proto were hand-tuned, so they will corrupt the current data & proto, if you run them again.

If you want to ignore this warning (to be fair, git probly has your back) Run these:

```bash
npm i            # install dependencies
npm run data     # get all the data from the API
npm run proto    # generate the gRPC protobuf that matches the API
npm run test     # validate the generated data against the generated protobuf

```