FROM alpine:3.23
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY shelly-openmetrics-exporter /
EXPOSE 54901
