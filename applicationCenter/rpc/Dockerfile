FROM hub.onlinesign.com/base/alpine:rancher-plus

ADD ./build/dlv /usr/local/bin/dlv
ADD ./build/appCenter-rpc /opt/app/appCenter-rpc
ADD etc/bootstrap.yaml /opt/app/etc/bootstrap.yaml
RUN chmod  a+x /usr/local/bin/dlv

WORKDIR /opt/app

ENTRYPOINT ["/sbin/tini", "--", "/opt/app/appCenter-rpc", "start"]