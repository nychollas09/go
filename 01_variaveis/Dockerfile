FROM golang:1.16.5-alpine3.13 AS build  

WORKDIR /go/src/variables

COPY main.go .

RUN go build main.go

FROM scratch

WORKDIR /usr/variables

COPY --from=build /go/src/variables/main .

ENTRYPOINT [ "./main" ]