FROM debian:stable-slim AS runner

WORKDIR /app

# Copy the native binary
COPY dist/hifi-webui ./hifi-webui

# Mark executable
RUN chmod +x ./hifi-webui

EXPOSE 3000

CMD ["./hifi-webui"]
