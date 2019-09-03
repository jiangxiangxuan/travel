#!/bin/bash

mkdir /data/logs -p
mkdir /etc/supervisor.d -p

cp -rf /opt/apps/travel/nginx/travel_include.conf /etc/nginx/conf.d/travel_include.conf
cp -rf /opt/apps/travel/nginx/nginx.conf /etc/nginx/nginx.conf
cp -rf /opt/apps/travel/docker/supervisord/*  /etc/supervisor.d/
cp -rf /opt/apps/travel/etc/travel_api.toml /etc/travel_api.toml

#service nginx restart
/usr/sbin/nginx
/usr/bin/supervisord -c /etc/supervisord.conf
