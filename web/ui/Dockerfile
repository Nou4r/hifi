FROM oven/bun:1 AS builder

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN bun install -g pnpm

RUN pnpm install

COPY . .

RUN pnpm run build

CMD ["./dist/hifi-webui"]
