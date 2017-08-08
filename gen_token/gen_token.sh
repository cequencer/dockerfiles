#!/bin/bash

HEADERS=$(/gen_token /data/docker_subscription.lic |sed -e 's/{"/-H "/g' -e 's/","/" -H "/g' -e 's/}//g' -e "s/\"/'/g")
sleep 1
CVE_URL=$(curl -sk -X GET https://license.enterprise.docker.com/v1/dss/cve-db-updates/0?schema=2 $HEADERS | jq .urls[] )

curl -sk -X GET $CVE_URL -o /data/docker_scanning_database.tar
