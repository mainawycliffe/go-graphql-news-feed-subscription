# This dockerfile uses vendor directory instead of fetching dependancies remotely
# first run "go mod vendor" to create vendors directory before building with this
# dockerfile

#build stage
FROM golang:alpine AS builder
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app
COPY . .
WORKDIR /app/graphql/server
RUN apk add --no-cache git
RUN go build -v -o app -mod=vendor

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/graphql/server/app /app
RUN mkdir -p images
ENTRYPOINT ./app
EXPOSE 8080
