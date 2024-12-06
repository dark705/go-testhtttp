# Simple HTTP Server with metrics, health check and so on

## Github

Can be found on: https://github.com/dark705/go-testhtttp

## Docker

Can be found on: https://hub.docker.com/r/dark705/go-testhttp 

## API endpoints

## Test endpoints

* /test - return "Ok"
* /host - return hostname

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