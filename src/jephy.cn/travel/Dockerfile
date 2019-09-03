#docker build -f Dockerfile -t travel .
#docker run -d -it --rm=true -p 88:80 travel
FROM jephy-base
#FROM centos:latest

EXPOSE 80

RUN mkdir -p /opt/apps/travel/bin
WORKDIR /opt/apps/travel

COPY ./nginx /opt/apps/travel/nginx
COPY ./web /opt/apps/travel/web
COPY ./docker /opt/apps/travel/docker
COPY ./bin /opt/apps/travel/bin

COPY ./api/etc/travel_api.toml /opt/apps/travel/etc/travel_api.toml

#RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
#RUN yum makecache -y
#RUN yum install -y epel-release
#RUN yum install -y nginx
#RUN yum install -y supervisor
#RUN yum install -y nodejs

RUN cd /opt/apps/travel/web && npm install

#CMD ["sh", "/opt/apps/travel/docker/start.sh"]
ENTRYPOINT ["sh", "/opt/apps/travel/docker/start.sh"]
