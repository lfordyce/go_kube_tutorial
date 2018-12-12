# Build stage
#FROM golang:1.10.2-alpine3.7 AS build
FROM golang:1.10.2-alpine3.7
# Support CGO and SSL
#RUN apk --no-cache add gcc g++ make ca-certificates
#WORKDIR /go/src/github.comcast.com/lfordy200/k8_test

#COPY Gopkg.lock Gopkg.toml ./
#COPY webservice webservice
#COPY vendor vendor
#COPY shapeutil shapeutil

#COPY . .

# Compile them
#RUN go install ./...

#RUN go get -v ./...
#RUN go install -v ./...

#CMD ["app"]
#ENTRYPOINT ["app", "-f=7", "-s=9"]

# Run the app
#CMD [ "app" ]

#EXPOSE 8080

#FROM alpine:3.7
#WORKDIR /usr/bin
#COPY --from=build /go/bin .

ADD ./main.go /go/src/github.comcast.com/lfordy200/k8_test/main.go

ADD ./webservice /go/src/github.comcast.com/lfordy200/k8_test/webservice

ADD ./vendor /go/src/github.comcast.com/lfordy200/k8_test/vendor

RUN set -ex && \
  cd /go/src/github.comcast.com/lfordy200/k8_test && \
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./k8_test /usr/bin/k8_test

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "k8_test" ]