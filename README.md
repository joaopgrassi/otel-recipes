# otel-recipes


## Starting a local Collector

Requirements:
- Docker and Compose

At the root of the repo, execute:

```shell
docker-compose up
```

Traces can be viewed:

- Console
- Jaeger: http://localhost:16686

If you don't want to see the Traces in the console or don't want
to block the terminal, you can start compose in detached mode:
`docker-compose up -d`.
To stop all containers, just run `docker-compose stop`.