FROM alpine:3.17
ENTRYPOINT ["/shelly-prometheus-exporter"]
COPY shelly-prometheus-exporter /
