FROM nginx:latest
MAINTAINER tukejonny

# Copy custom nginx.conf
COPY ./conf/nginx.conf /etc/nginx/nginx.conf

# Conf test
RUN nginx -t

# Copy static files
RUN mkdir -p /var/www/public
COPY ./public /var/www/public

EXPOSE 80

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]
