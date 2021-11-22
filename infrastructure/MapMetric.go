package infrastructure

import (
	"../domain"
	"encoding/json"
	"os"
)

var (
	data = domain.Json{}
	urls [4]string
)

func MapMetric(){

	urls[0] = os.Getenv("cpu")
	urls[1] = os.Getenv("memory")
	urls[2] = os.Getenv("network")
	urls[3] = os.Getenv("disk")

	for _, url := range urls {
		_ = json.Unmarshal(GetMetric(url, os.Getenv("token"), os.Getenv("name")), &data) // metrics is gotten from prometheus with using http client

		var result = data.Data.Result //returned json from the prometheus mapping to result
		for _, metrics := range result {
			Influxdb(
				metrics.Metric.Name,
				metrics.Metric.Instance,
				metrics.Value[1],
			)
		}
	}
}
