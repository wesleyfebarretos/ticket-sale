#!/bin/bash

source .env

sudo chmod go-w ./internal/infra/elastic_stack/filebeat/filebeat.yml
sudo chown 0:0 ./internal/infra/elastic_stack/filebeat/filebeat.yml

sudo chmod go-w ./internal/infra/elastic_stack/metricbeat/metricbeat.yml
sudo chown 0:0 ./internal/infra/elastic_stack/metricbeat/metricbeat.yml

sudo chmod go-w ./internal/infra/elastic_stack/logstash/logstash.conf
sudo chown 0:0 ./internal/infra/elastic_stack/logstash/logstash.conf

docker compose up -d
