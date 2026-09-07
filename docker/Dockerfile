FROM alpine:3.19
LABEL org.opencontainers.image.title="Antybrowser CLI" \
      org.opencontainers.image.description="CLI for Antybrowser anti-detect browser - manage profiles, proxies, extensions, and automations from the terminal" \
      org.opencontainers.image.url="https://antybrowser.com" \
      org.opencontainers.image.source="https://github.com/antybrowser/antybrowser-cli" \
      org.opencontainers.image.documentation="https://docs.antybrowser.com" \
      org.opencontainers.image.vendor="Antybrowser.com" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.keywords="antybrowser,cli,browser,automation,proxy,fingerprint,multi-account,anti-detect,scraping,privacy"
RUN apk add --no-cache curl && \
    curl -fsSL https://github.com/antybrowser/antybrowser-cli/releases/download/cli-v1.0.0/antybrowser-linux-amd64.tar.gz | tar xz -C /usr/local/bin && \
    apk del curl
ENTRYPOINT ["antybrowser"]
