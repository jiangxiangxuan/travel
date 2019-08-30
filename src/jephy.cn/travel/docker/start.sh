#!/bin/bash

mkdir /data/logs -p

ln -s /opt/apps/travel/nginx/include.conf /etc/nginx/vhost/travel_include.conf
cp -rf /opt/apps/travel/nginx/nginx.conf /etc/nginx/nginx.conf
cp -rf /opt/apps/travel/docker/supervisord/*  /etc/supervisord/

service nginx restart
/usr/bin/supervisord
