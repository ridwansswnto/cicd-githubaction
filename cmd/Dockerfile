FROM golang:1.16-alpine AS build_base
WORKDIR /cmd
COPY go.mod .
COPY . .
RUN go build -v -o hello-app
FROM alpine:3.9
COPY --from=build_base /cmd/hello-app /cmd/hello-app
EXPOSE 8080
CMD ["/cmd/hello-app"]