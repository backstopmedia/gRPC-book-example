# Archiver Client

Sends files to the archive service (running at `localhost:8080`) to be archived in _compressed.zip_.

## Running this example

With the archive server running (run `make run svc=archiver` from the streaming directory), run the following:

```
pip install --user pipenv
pipenv install grpcio==1.11.0
pipenv run python main.py Pipfile Pipfile.lock
```

You should have a new file in this directory called _compressed.zip_ containing the two files passed on the CLI.
