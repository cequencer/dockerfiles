FROM alpine
MAINTAINER jesse.lovelace@gmail.com

ENV PYTHON_VERSION=2.7.12-r0
ENV PY_PIP_VERSION=8.1.2-r0
ENV SUPERVISOR_VERSION=3.3.0

RUN apk update && apk add -u python=$PYTHON_VERSION py-pip=$PY_PIP_VERSION ca-certificates
RUN pip install supervisor==$SUPERVISOR_VERSION

COPY logserver.ini /etc/
COPY build/logserver /usr/bin/
COPY client_secret.json /

RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/supervisord.conf

EXPOSE 80:80

CMD ["/usr/bin/supervisord"]
