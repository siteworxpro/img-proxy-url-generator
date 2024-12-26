FROM siteworxpro/golang:1.23.4 AS build

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

ENTRYPOINT ["/app/imgproxy", "grpc"]