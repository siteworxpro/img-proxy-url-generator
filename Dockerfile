FROM siteworxpro/golang:1.24.3 AS build

WORKDIR /app

ADD . .

ENV GOPRIVATE=git.siteworxpro.com
ENV GOPROXY=direct
ENV CGO_ENABLED=0

RUN go mod tidy && go build -o imgproxy .

FROM alpine:latest AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=build /app/imgproxy /app/imgproxy

RUN  adduser -u 1001 -g appuser appuser -D && \
    chown -R appuser:appuser /app

USER 1001

ENTRYPOINT ["/app/imgproxy", "server"]