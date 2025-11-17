FROM oven/bun:latest AS builder

WORKDIR /app

RUN corepack enable pnpm

COPY . .

RUN pnpm install

RUN pnpm run build


FROM oven/bun:latest AS runner

WORKDIR /app

COPY dist/ ./hifi-webui

EXPOSE 3000

CMD ["./hifi-webui"]
