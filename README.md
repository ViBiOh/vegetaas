# vegetaas

[![Build](https://github.com/ViBiOh/vegetaas/workflows/Build/badge.svg)](https://github.com/ViBiOh/vegetaas/actions)
[![codecov](https://codecov.io/gh/ViBiOh/vegetaas/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/vegetaas)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_vegetaas&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_vegetaas)

An HTTP wrapper around [vegeta](https://github.com/tsenart/vegeta).

Thanks to [tsenart](https://github.com/tsenart) for the awesome `vegeta`.

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/vegetaas/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/vegetaas](https://hub.docker.com/r/vibioh/vegetaas/tags).

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

You'll find a Kubernetes exemple in the [`infra/`](infra) folder, using my [`app chart`](https://github.com/ViBiOh/charts/tree/main/app)

## Usage

Query parameters:

- `url` url to request
- `rps` request per second, as integer
- `duration` duration, in Golang format

## CI

Following variables are required for CI:

|            Name            |           Purpose           |
| :------------------------: | :-------------------------: |
|      **DOCKER_USER**       | for publishing Docker image |
|      **DOCKER_PASS**       | for publishing Docker image |
| **SCRIPTS_NO_INTERACTIVE** |  for running scripts in CI  |

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of vegetaas:
  -address string
        [server] Listen address {VEGETAAS_ADDRESS}
  -cert string
        [server] Certificate file {VEGETAAS_CERT}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {VEGETAAS_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {VEGETAAS_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {VEGETAAS_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {VEGETAAS_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {VEGETAAS_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {VEGETAAS_CSP} (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options {VEGETAAS_FRAME_OPTIONS} (default "deny")
  -graceDuration duration
        [http] Grace duration when SIGTERM received {VEGETAAS_GRACE_DURATION} (default 30s)
  -hsts
        [owasp] Indicate Strict Transport Security {VEGETAAS_HSTS} (default true)
  -idleTimeout duration
        [server] Idle Timeout {VEGETAAS_IDLE_TIMEOUT} (default 2m0s)
  -key string
        [server] Key file {VEGETAAS_KEY}
  -loggerJson
        [logger] Log format as JSON {VEGETAAS_LOGGER_JSON}
  -loggerLevel string
        [logger] Logger level {VEGETAAS_LOGGER_LEVEL} (default "INFO")
  -loggerLevelKey string
        [logger] Key for level in JSON {VEGETAAS_LOGGER_LEVEL_KEY} (default "level")
  -loggerMessageKey string
        [logger] Key for message in JSON {VEGETAAS_LOGGER_MESSAGE_KEY} (default "message")
  -loggerTimeKey string
        [logger] Key for timestamp in JSON {VEGETAAS_LOGGER_TIME_KEY} (default "time")
  -okStatus int
        [http] Healthy HTTP Status code {VEGETAAS_OK_STATUS} (default 204)
  -port uint
        [server] Listen port (0 to disable) {VEGETAAS_PORT} (default 1080)
  -prometheusAddress string
        [prometheus] Listen address {VEGETAAS_PROMETHEUS_ADDRESS}
  -prometheusCert string
        [prometheus] Certificate file {VEGETAAS_PROMETHEUS_CERT}
  -prometheusGzip
        [prometheus] Enable gzip compression of metrics output {VEGETAAS_PROMETHEUS_GZIP}
  -prometheusIdleTimeout duration
        [prometheus] Idle Timeout {VEGETAAS_PROMETHEUS_IDLE_TIMEOUT} (default 10s)
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {VEGETAAS_PROMETHEUS_IGNORE}
  -prometheusKey string
        [prometheus] Key file {VEGETAAS_PROMETHEUS_KEY}
  -prometheusPort uint
        [prometheus] Listen port (0 to disable) {VEGETAAS_PROMETHEUS_PORT} (default 9090)
  -prometheusReadTimeout duration
        [prometheus] Read Timeout {VEGETAAS_PROMETHEUS_READ_TIMEOUT} (default 5s)
  -prometheusShutdownTimeout duration
        [prometheus] Shutdown Timeout {VEGETAAS_PROMETHEUS_SHUTDOWN_TIMEOUT} (default 5s)
  -prometheusWriteTimeout duration
        [prometheus] Write Timeout {VEGETAAS_PROMETHEUS_WRITE_TIMEOUT} (default 10s)
  -readTimeout duration
        [server] Read Timeout {VEGETAAS_READ_TIMEOUT} (default 5s)
  -shutdownTimeout duration
        [server] Shutdown Timeout {VEGETAAS_SHUTDOWN_TIMEOUT} (default 10s)
  -tracerRate string
        [tracer] Jaeger sample rate, 'always', 'never' or a float value {VEGETAAS_TRACER_RATE} (default "always")
  -tracerURL string
        [tracer] Jaeger endpoint URL (e.g. http://jaeger:14268/api/traces) {VEGETAAS_TRACER_URL}
  -url string
        [alcotest] URL to check {VEGETAAS_URL}
  -userAgent string
        [alcotest] User-Agent for check {VEGETAAS_USER_AGENT} (default "Alcotest")
  -writeTimeout duration
        [server] Write Timeout {VEGETAAS_WRITE_TIMEOUT} (default 10s)
```
