FROM node:24.18.1-slim AS base
WORKDIR /app
COPY package.json ./
RUN npm install --omit=dev
COPY src ./src

FROM node:24.18.1-slim
WORKDIR /app
ENV NODE_ENV=production
COPY --from=base /app/node_modules ./node_modules
COPY --from=base /app/package.json ./package.json
COPY --from=base /app/src ./src
EXPOSE 4000
CMD ["node", "src/server.js"]
