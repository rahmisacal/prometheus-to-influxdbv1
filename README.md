# prometheus-influxdb exporter
A Simple Application Using Golang

## About : prometheus-influxdb exporter
The App is used golang http client, authenticated to clusters with openshift cluster's token which define in secret at openshift-authentication namespace, received cpu,memory,network and disk usage metrics of the cluster which define with environment variables finally this application is sent this metrics to influxdb every 3 minutes

- golang 1.17.3
- influxdb-client-v1(influxdb v1.7.9)

## Author:
- team-openshift-admin

## Usage

 - Set environment deployment.yaml (cpu, memory, disk, network, name, token)
 - sure that access to influxdb:port and target k8s cluster from sourceIp