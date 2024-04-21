# OTLP test back-end

This folder contains the code for a simple "dummy" OTLP back-end.

The goal of this server is to be the target for OTLP exporters from the
sample applications. The server keeps the OTLP data in-memory which is then queried by the
tests to assert the exported telemetry data.

## Exposed endpoints

The server is exposed via port `4319`.

### OTLP ingest

The server exposes endpoints to ingest OTLP protobuf data via HTTP. The endpoints
follow the paths from the [OTLP specification](https://opentelemetry.io/docs/specs/otlp/).

- `/v1/traces`
- `/v1/metrics`
- `/v1/logs`

### OTLP query

To query data, the server exposes an endpoint `/getotlp`. The following query parameters
can be further specified to filter out data:

- `signal`: Which telemetry signal to get. One of:
  - `traces`
  - `metrics`
  - `logs`
- `servicename`: The `service.name` resource attribute of the sample application. This is
  used by the tests to get the telemetry for the app being tested.

Example query:

```shell
http://localhost:4319/getotlp?signal=trace&servicename=myapp
```
