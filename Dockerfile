FROM siteworxpro/golang:1.24.0 AS build

WORKDIR /app

ADD . .

ENV GOPRIVATE=git.s.int
ENV GOPROXY=direct
ENV CGO_ENABLED=0

RUN go mod tidy && go build -o imgproxy .

FROM alpine:latest AS runtime

EXPOSE 9000

WORKDIR /app

COPY --from=build /app/imgproxy /app/imgproxy

RUN  adduser -u 1001 -g appuser appuser -D && \
    chown -R appuser:appuser /app

USER 1001

# docker buildx build --push --sbom=true --provenance=true --platform linux/amd64,linux/arm64 -t siteworxpro/img-proxy-url-generator:v1.4.0-grpc .

ENTRYPOINT ["/app/imgproxy", "grpc"]