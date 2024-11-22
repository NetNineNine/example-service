# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

ARG Version
ARG FullCommit
ARG Date

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X main.version=${Version} -X main.commitSHA=${FullCommit} -X main.buildDate=${Date}" -o /example-service

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11:latest-amd64 AS build-release-stage

WORKDIR /

COPY --from=build-stage /example-service /example-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/example-service"]
