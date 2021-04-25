FROM alpine:3.13
ADD ./ping-probe /usr/local/bin/ping-probe
ENTRYPOINT ["ping-probe", "--config", "/etc/config/config.prod.yaml"]