FROM siteworxpro/golang:latest AS build

WORKDIR /app

ADD . .

ENV GOPRIVATE=git.s.int
ENV GOPROXY=direct
ENV CGO_ENABLED=0

RUN go mod tidy && go build -o imgproxy .

FROM alpine AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=build /app/imgproxy /app/imgproxy

ENTRYPOINT ["/app/imgproxy", "server"]