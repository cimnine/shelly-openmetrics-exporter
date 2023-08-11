FROM alpine:3.18
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY shelly-openmetrics-exporter /
EXPOSE 54901
