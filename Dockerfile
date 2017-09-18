FROM alpine:stable

RUN apk add --no-cache ca-certificates

ENV HELM_CHART_PUBLISHER_VERSION=v0.2.3

RUN curl -L https://github.com/luizbafilho/helm-chart-publisher/releases/download/${HELM_CHART_PUBLISHER_VERSION}/helm-chart-publisher_linux-amd64 -o /usr/local/bin/helm-chart-publisher
RUN chmod +x /usr/local/bin/helm-chart-publisher

CMD ["/helm-chart-publisher", "-c", "/etc/helm-chart-publisher/config.yaml"]

