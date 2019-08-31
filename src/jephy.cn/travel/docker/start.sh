#!/bin/bash

mkdir /data/logs -p

ln -s /opt/apps/travel/nginx/include.conf /etc/nginx/conf.d/travel_include.conf
cp -rf /opt/apps/travel/nginx/nginx.conf /etc/nginx/nginx.conf
cp -rf /opt/apps/travel/docker/supervisord/*  /etc/supervisord/

cp -rf /opt/apps/travel/etc/travel_api.toml /etc/travel_api.toml

service nginx restart
/usr/bin/supervisord
