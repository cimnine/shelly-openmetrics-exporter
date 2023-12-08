FROM alpine:3.19
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY shelly-openmetrics-exporter /
EXPOSE 54901
