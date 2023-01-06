FROM alpine:3.17
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY shelly-openmetrics-exporter /
EXPOSE 54901
