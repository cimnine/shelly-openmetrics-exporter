FROM alpine:3.21
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY shelly-openmetrics-exporter /
EXPOSE 54901
