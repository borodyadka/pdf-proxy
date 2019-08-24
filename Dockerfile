FROM golang:1.12 AS builder

ENV CGO_ENABLED 0
WORKDIR /app
COPY . /app
RUN go build -mod vendor -a -ldflags "-w -s" -installsuffix cgo -o ./main ./cmd/server/main.go


FROM alpine:3.10

WORKDIR /opt
EXPOSE 8080
EXPOSE 8082

# based on https://github.com/Surnet/docker-wkhtmltopdf/blob/19a0f688a4886679858c32ffd7a5a93800a1d30c/Dockerfile-alpine.template
RUN apk add --no-cache \
    curl \
    wkhtmltopdf \
    libstdc++ \
    libx11 \
    libxrender \
    libxext \
    libssl1.1 \
    ca-certificates \
    fontconfig \
    freetype \
    ttf-dejavu \
    ttf-droid \
    ttf-freefont \
    ttf-liberation \
    ttf-ubuntu-font-family \
    && apk add --no-cache --virtual .build-deps msttcorefonts-installer \
    && update-ms-fonts \
    && fc-cache -f \
    && rm -rf /tmp/* \
    && apk del .build-deps

HEALTHCHECK --start-period=5s --interval=5s --timeout=5s --retries=3 \
    CMD curl --fail --silent http://localhost:8080/ping || exit 1

COPY --from=builder /app/main /opt/server
COPY ./entrypoint.sh /opt/entrypoint.sh

ENTRYPOINT ["/opt/entrypoint.sh"]
CMD ["./server"]
