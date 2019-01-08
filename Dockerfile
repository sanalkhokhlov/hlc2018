FROM golang:1.11-alpine3.7 AS build

WORKDIR /go/src/bitbucket.org/sLn/hlc2018
ADD . /go/src/bitbucket.org/sLn/hlc2018

RUN go build -o /bin/app main.go

FROM alpine:3.7
COPY --from=build /bin/app /app

EXPOSE 80
CMD ["/app"]