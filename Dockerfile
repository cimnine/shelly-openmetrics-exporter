FROM alpine
ENTRYPOINT ["/shelly-prometheus-exporter"]
COPY shelly-prometheus-exporter /
