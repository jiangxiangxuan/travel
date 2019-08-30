FROM centos:latest

EXPOSE 80

RUN mkdir -p /opt/apps/travel/bin
WORKDIR /opt/apps/travel

COPY /opt/apps/travel/src/jephy.cn/travel/nginx /opt/apps/travel/nginx
COPY /opt/apps/travel/src/jephy.cn/travel/web /opt/apps/travel/web
COPY /opt/apps/travel/src/jephy.cn/travel/docker /opt/apps/travel/docker
COPY /opt/apps/travel/bin /opt/apps/travel/bin

COPY /opt/apps/travel/src/jephy.cn/travel/api/etc/travel_api.toml /opt/apps/travel/etc/travel_api.toml

RUN yum install -y nginx
RUN yum install -y epel-release
RUN yum install -y supervisor

CMD sh /opt/apps/travel/docker/start.sh
