# Database Client

Sends a query to the database service's `Search` method.

## Running this example

With the database server running (run `make run svc=database` from the streaming directory), run the following:

```
ruby client.pb searchTerm
```

You should have a new file in this directory called _compressed.zip_ containing the two files passed on the CLI.

