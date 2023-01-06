# shelly-openmetrics-exporter

This is an openmetrics exporter written in Go.
It's `/probe` and `/metrics` endpoint report in OpenTracing format.

## Features

Use this tool to fetch power readings from your Shellys via Prometheus (or any OpenTracing compatible agent).
It supports multiple power lines (like they are available on `Shelly 3EM` or `Shelly Plus 2PM`).

This exporter is compatible with the _First Generation Shelly Devices API_ and the _Second Generation Shelly Devices API_.
Authentication is not supported.

## Usage

Run this program and point prometheus to it.
Use `target` to define which shelly to query.
You may use an IP or a domain name.
You may configure a port like this `name:port`.

If your Shelly is password protected,
you may use the `username` and `password` arguments to pass credentials.

## TODO

- [x] detect shelly generation
- [x] temperature
- [x] goreleaser
- [ ] tests
- [ ] rename shelly-openmetrics-exporter
- [x] authentication

## License

This software is licensed [under MIT license](/LICENSE).
