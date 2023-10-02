FROM node:20.7.0-alpine3.18
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build
EXPOSE 4000
ENTRYPOINT [ "yarn", "serve" ]