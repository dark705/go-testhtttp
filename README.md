# Simple HTTP Server with metrics, health check and so on

## Configuration

### Environment

* VERSION - Version of application. Default:"version_not_set"
* LEVEL - LogLevel. Default: "info". Possible values:

    - "debug"
    - "info"
    - "warning"
    - "error"
    - "fatal"

* HOSTNAME - Hostname. Default: "localhost"
* HTTP_PORT - HTTP port of application.Default: "8000"`
* HTTP_REQUEST_HEADER_MAX_SIZE - Maximum HTTP request header size in bites. Default: "10000"
* HTTP_REQUEST_READ_HEADER_TIMEOUT_MILLISECONDS - Maximum time for read HTTP request header in milliseconds. Default: "
  2000"
