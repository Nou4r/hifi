FROM oven/bun:latest AS builder

WORKDIR /app

COPY package.json bun.lockb ./

RUN bun install

COPY . .

RUN bun run build

FROM oven/bun:latest AS runner

WORKDIR /app

COPY --from=builder /app/dist/hifi-webui ./hifi-webui

RUN chmod +x ./hifi-webui

EXPOSE 3000
CMD ["./hifi-webui"]

