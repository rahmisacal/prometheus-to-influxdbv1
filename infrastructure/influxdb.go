package infrastructure

import (
	"../configuration"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	_ "github.com/influxdata/influxdb1-client/v2"
)

func Influxdb(name string, node string, value string) {

	conf := configuration.Configuration()

	host, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.InfluxDb.Url, conf.InfluxDb.Port))
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Application was not connected to influxdb.\n")
	}

	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Influxdb client was not generated successfuly.\n")
	}

	var pts = make([]client.Point, 1)

	const bitSize = 64 // Don't think about it to much. It's just 64 bits.
	floatValue, err := strconv.ParseFloat(value, bitSize)

	pts[0] = client.Point{
		Measurement: conf.InfluxDb.Measurement,
		Tags: map[string]string{
			name: node,
		},
		Fields: map[string]interface{}{
			"value": floatValue,
		},
		Time:      time.Now(),
		Precision: "",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        conf.InfluxDb.Database,
		RetentionPolicy: conf.InfluxDb.RetentionPolicy,
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Metric was not written to influxdb successfully.\n")
	}
	fmt.Printf("INFO : " + time.Now().String() + " " + name + "/" + node + "=" + value + " was sent to influxdb.\n")
}
