FROM alpine:3.13
ADD ./ping-probe-register /usr/local/bin/ping-probe-register
ENTRYPOINT ["ping-probe-register", "--config", "/etc/config/config.prod.yaml"]