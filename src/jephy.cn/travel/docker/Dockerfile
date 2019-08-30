FROM centos

EXPOSE 80

RUN mkdir -p /opt/apps/travel/bin
WORKDIR /opt/apps/travel

COPY /opt/apps/travel/src/jephy.cn/travel /opt/apps/travel
COPY /opt/apps/travel/bin /opt/apps/travel/bin
RUN yum install nginx -y

CMD sh /opt/apps/travel/docker/start.sh
