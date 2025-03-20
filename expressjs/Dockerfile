FROM node:lts AS builder
WORKDIR /build

COPY . .

RUN npm install

ENV NODE_ENV=production

EXPOSE 3000

CMD [ "node", "index.js" ]
