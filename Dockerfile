FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && mkdir /app \
    && apk add --no-cache bash tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/lib/apt/lists/*
COPY ./build/release/cabservd /app/
#COPY ./build/dlv /app/
COPY ./docker-entrypoint.sh /usr/bin/docker-entrypoint.sh
RUN chmod +x /usr/bin/docker-entrypoint.sh

WORKDIR /app

ENTRYPOINT ["/usr/bin/docker-entrypoint.sh"]