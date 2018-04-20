FROM nginx:1.13.9-alpine

RUN rm -rf /etc/nginx/conf.d

COPY nginx.default.conf /etc/nginx/conf.d/default.conf

COPY ./app /usr/share/nginx/html

ADD --chown=775 replace.sh /replace.sh
ADD --chown=775 run.sh /run.sh

EXPOSE 80

CMD ["sh", "/run.sh"]