FROM alpine
LABEL maintainer="clemenko@docker.com", GITREPO="https://github.com/clemenko/dockerfiles/tree/master/demo_flask"
RUN apk -U upgrade && apk add --no-cache curl py2-pip  &&\
    pip install --no-cache-dir --upgrade pip &&\
    pip install --no-cache-dir flask redis pymongo &&\
    rm -rf /var/cache/apk/*
WORKDIR /code
ADD . /code
EXPOSE 5000
HEALTHCHECK CMD curl -f http://localhost:5000/healthz || exit 1
CMD ["python", "app.py"]
