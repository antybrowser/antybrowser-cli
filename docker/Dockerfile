FROM alpine:3.19
RUN apk add --no-cache curl && \
    curl -fsSL https://github.com/antybrowser/antybrowser-cli/releases/download/v1.0.0/antybrowser-linux-amd64.tar.gz | tar xz -C /usr/local/bin && \
    apk del curl
ENTRYPOINT ["antybrowser"]
