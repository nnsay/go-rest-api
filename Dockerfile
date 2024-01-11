FROM golang:alpine
RUN apk add --update git
RUN mkdir -p ${GOPATH}/go-rest-api
WORKDIR ${GOPATH}/go-rest-api/
COPY rest-api.go go.mod go.sum ${GOPATH}/go-rest-api/
RUN go build rest-api.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/go-rest-api/rest-api .
ENV APP_PORT 8080
EXPOSE $APP_PORT
CMD [ "./rest-api" ]
