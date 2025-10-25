FROM node:25 AS builder
WORKDIR /usr/src/housebook
COPY . .
RUN npm install
RUN npm run lint && npm run build

# TODO: switch to Caddy in production
FROM node:25 AS runner
WORKDIR /usr/src/housebook
COPY --from=builder /usr/src/housebook .
EXPOSE 5173
CMD ["npm", "run", "dev"]
