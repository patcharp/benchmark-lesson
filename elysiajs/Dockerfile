FROM oven/bun AS builder
WORKDIR /build

COPY . .

RUN bun install

ENV NODE_ENV=production
RUN bun --compile \
    --minify-whitespace \
    --minify-syntax \
    --target bun \
    --outfile server \
    ./index.ts

FROM gcr.io/distroless/base AS runner
WORKDIR /app
COPY --from=builder /build/server .

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "/app/server" ]
    