FROM nginx:latest

COPY docker/nginx/conf.d/application.template /etc/nginx/conf.d/default.conf

EXPOSE 81

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]