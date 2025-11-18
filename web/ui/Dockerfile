FROM oven/bun:latest AS builder

WORKDIR /app

COPY . .

RUN bun install --frozen-lockfile

RUN bun run build


# FROM oven/bun:latest AS runner

# WORKDIR /app

# COPY dist/hifi-webui ./hifi-webui

# EXPOSE 3000

# CMD ["./hifi-webui"]
