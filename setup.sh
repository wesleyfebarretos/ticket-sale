#!/bin/bash

source .env

sudo chmod go-w ./infra/elastic_stack/filebeat/filebeat.yml
sudo chown 0:0 ./infra/elastic_stack/filebeat/filebeat.yml

sudo chmod go-w ./infra/elastic_stack/metricbeat/metricbeat.yml
sudo chown 0:0 ./infra/elastic_stack/metricbeat/metricbeat.yml

docker compose up -d
