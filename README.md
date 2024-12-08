# Simple HTTP Server with metrics, health check and so on

## Github

Can be found on: https://github.com/dark705/go-testhtttp

## Docker

Can be found on: https://hub.docker.com/r/dark705/go-testhttp

## API endpoints

### Test endpoints

* /test - return "Ok"
* /host - return hostname

### Kubernetes endpoint probes

* /kuber/startup - Startup probe. See Environment.
* /kuber/live - Live probe. See Environment.
* /kuber/ready - Ready probe. See Environment.

## Configuration

### Environment

* VERSION - Version of application. Default:"version_not_set"
* LEVEL - LogLevel. Default: "info". Possible values:

    - "debug"
    - "info"
    - "warning"
    - "error"
    - "fatal"

* HTTP_PORT - HTTP port of application.Default: "8000"`
* HTTP_REQUEST_HEADER_MAX_SIZE - Maximum HTTP request header size in bites. Default: "10000"
* HTTP_REQUEST_READ_HEADER_TIMEOUT_MILLISECONDS - Maximum time for read HTTP request header in milliseconds. Default: "
  2000"
* PROMETHEUS_PORT - Prometheus port. Default:"9000"

* KUBER_PROBE_START_UP_SECONDS - Time seconds after start, when Startup probe will return Ok. Default:"0"
* KUBER_PROBE_PROBABILITY_LIVE - Probability (from 0 to 100) Live probe return Ok. Default:"100", means - always live.
* KUBER_PROBE_PROBABILITY_READY - Probability (from 0 to 100) Ready probe return Ok. Default:"100", means - always
  ready. 