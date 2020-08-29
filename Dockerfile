FROM golang:1.14 AS build-env
#disable crosscompiling
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . /go/src/go-api
WORKDIR /go/src/go-api
RUN ls -la
RUN go build -o /go/bin/go-api cmd/main.go


FROM alpine
EXPOSE 80
EXPOSE $PORT
COPY --from=build-env /go/bin/go-api /app/
WORKDIR /app
CMD ["./go-api"]