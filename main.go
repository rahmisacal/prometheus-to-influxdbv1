package main

import (
	"./configuration"
	"./domain"
	"./infrastructure"
	"encoding/json"
)

func main() {

	var (
		conf = configuration.Configuration()
		data = domain.Json{}
		urls [4]string
	)

	for i := 0; i < len(conf.Clusters); i++ {

		urls[0] = conf.Clusters[i].Cpu
		urls[1] = conf.Clusters[i].Memory
		urls[2] = conf.Clusters[i].Network
		urls[3] = conf.Clusters[i].Disk
		token := conf.Clusters[i].Token
		clusterName := conf.Clusters[i].Name

		for _, url := range urls {
			_ = json.Unmarshal(infrastructure.GetMetric(url, token, clusterName), &data) // metrics is gotten from prometheus with using http client

			var result = data.Data.Result //returned json from the prometheus mapping to result
			for _, metrics := range result {
				go infrastructure.Influxdb(
					metrics.Metric.Name,
					metrics.Metric.Instance,
					metrics.Value[1],
				)
				defer func() {
					go infrastructure.Influxdb(metrics.Metric.Name, metrics.Metric.Instance, metrics.Value[1])
				}()
			}
		}
	}
}