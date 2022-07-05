#!/usr/bin/env sh

set -e
set -x 

if [[ ! -z "${METROGALAXY_API_ADDRESS}" ]]; then
  sed -i "s/metrogalaxy-api/$METROGALAXY_API_ADDRESS/g" /etc/nginx/conf.d/default.conf
fi

nginx -g 'daemon off;'