FROM redis:latest

ENV ALLOW_EMPTY_PASSWORD=yes

COPY redis.conf /usr/local/etc/redis/redis.conf

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
