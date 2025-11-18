FROM oven/bun:latest AS runner

WORKDIR /app

COPY dist/hifi-webui ./hifi-webui

EXPOSE 3000

CMD ["./hifi-webui"]
