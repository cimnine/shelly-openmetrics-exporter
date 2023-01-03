# shelly-prometheus-exporter

This is a prometheus exporter written in Go.
It has no other dependencies but the prometheus client libraries.
It's `/probe` and `/metrics` endpoint report in OpenTracing format.

## Features

Use this tool to fetch power readings from your shellys via Prometheus (or any OpenTracing compatible agent).
It supports multiple power lines (like they are available on `Shelly 3EM` or `Shelly Plus 2PM`).

This exporter is compatible with the _First Generation Shelly Devices API_ and the _Second Generation Shelly Devices API_.
Authentication is not supported.

## Usage

Run this program and point prometheus to it.
Configure `target` and `shelly_type`:

- `target` is the target shelly.
  You may use an IP or a domain name.
  You may configure a port like this `name:port`.
- `shelly_type` is which generation of shelly you have.
  See https://shelly-api-docs.shelly.cloud/.

## License

This software is licensed [under MIT license](/LICENSE).
