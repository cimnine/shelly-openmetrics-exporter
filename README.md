# shelly-prometheus-exporter

This is a prometheus exporter written in Go.
It has no other dependencies but the prometheus client libraries.
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

## TODO

- [x] detect shelly generation
- [x] temperature
- [ ] goreleaser
- [ ] tests
- [x] authentication

## License

This software is licensed [under MIT license](/LICENSE).
