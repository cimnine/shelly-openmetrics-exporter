FROM alpine:3.23
ARG TARGETPLATFORM
ENTRYPOINT ["/shelly-openmetrics-exporter"]
COPY $TARGETPLATFORM/shelly-openmetrics-exporter /
EXPOSE 54901
