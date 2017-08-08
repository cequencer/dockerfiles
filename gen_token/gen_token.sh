#!/bin/sh

LICENSE_FILE="${LICENSE_FILE:-/data/docker_subscription.lic}"

# generate values
GEN_TOKEN_OUTPUT="$(/gen_token "${LICENSE_FILE}")"

# generate values
KEY="$(echo "${GEN_TOKEN_OUTPUT}" | jq -r '.["X-DOCKER-KEY-ID"]')"
TOKEN="$(echo "${GEN_TOKEN_OUTPUT}" | jq -r '.["X-DOCKER-TOKEN"]')"
TIMESTAMP="$(echo "${GEN_TOKEN_OUTPUT}" | jq -r '.["X-DOCKER-TIMESTAMP"]')"

# sleep 5 second to make sure we allow the time to be valid
sleep 5

DB_URL="$(curl -s -X GET \
  "https://license.enterprise.docker.com/v1/dss/cve-db-updates/0?schema=2" \
  -H "X-DOCKER-KEY-ID: ${KEY}" \
  -H "X-DOCKER-TOKEN: ${TOKEN}" \
  -H "X-DOCKER-TIMESTAMP: ${TIMESTAMP}" | jq -r .urls[])"

CVE_DATE="$(date +%F)"

curl "${DB_URL}" -o /data/docker_scanning_database."${CVE_DATE}".tar
