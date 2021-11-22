# prometheus-influxdb exporter
A Simple Application Using Golang

## About : prometheus-influxdb exporter
The App is used golang http client, authenticated to clusters with openshift cluster's token which define in secret at openshift-authentication namespace, received cpu,memory,network and disk usage metrics of all clusters which define in config.yaml from prometheus http api finally this application is sent this metrics to influxdb every 3 minutes

- golang 1.17.3
- influxdb-client-v1(influxdb v1.7.9)

## Author:
- team-openshift-admin