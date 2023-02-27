FROM hub.onlinesign.com/base/alpine:rancher-plus

ADD ./build/dlv /usr/local/bin/dlv
ADD ./build/appCenter-api /opt/app/appCenter-api
ADD etc/bootstrap.yaml /opt/app/etc/bootstrap.yaml
RUN chmod  a+x /usr/local/bin/dlv

WORKDIR /opt/app

ENTRYPOINT ["/sbin/tini", "--", "/opt/app/appCenter-api", "start"]