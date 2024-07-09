FROM node:latest as build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:latest

COPY docker/nginx/conf.d/application.template /etc/nginx/conf.d/default.conf

COPY --from=build /app/dist /app

EXPOSE 81

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]