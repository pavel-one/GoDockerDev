FROM golang
ARG DOCKER_UID

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN mkdir /.cache
RUN chown ${DOCKER_UID}:${DOCKER_UID} /.cache

WORKDIR /usr/src/app

CMD ["air"]