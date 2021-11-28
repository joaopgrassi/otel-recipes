FROM node:16.13.0 as builder
WORKDIR /usr/app
COPY package*.json ./
RUN npm install


FROM node:16.13.0-alpine
WORKDIR /usr/app
COPY --from=builder /usr/app ./
COPY . .
EXPOSE 8080
CMD [ "node", "server.js" ]