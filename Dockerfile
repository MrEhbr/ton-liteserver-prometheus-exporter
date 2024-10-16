FROM            alpine:3

COPY            ton-liteserver-prometheus-exporter /bin/

ENTRYPOINT      ["/bin/ton-liteserver-prometheus-exporter"]
