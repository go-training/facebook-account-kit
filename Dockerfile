FROM golang:1.11-alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /facebook-account-kit -tags netgo -ldflags '-w -extldflags "-static"' .

### Put the binary onto Heroku image
FROM plugins/base:multiarch
LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"
EXPOSE 8080
COPY --from=server_builder /app/templates /templates
COPY --from=server_builder /app/images /images
COPY --from=server_builder /facebook-account-kit /facebook-account-kit
CMD ["/facebook-account-kit"]
