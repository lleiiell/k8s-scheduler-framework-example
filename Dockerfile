FROM debian:stretch-slim

WORKDIR /

COPY k8s-scheduler-framework-example /usr/local/bin

CMD ["k8s-scheduler-framework-example"]