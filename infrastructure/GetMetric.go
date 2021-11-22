package infrastructure

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetMetric(url string, token string, clusterName string) []byte {

	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "GET request was not generated.\n")
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Response Client was not generated.\n")
	}
	defer response.Body.Close()

	jsonBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "response body was not read as json.\n")
	}
	fmt.Printf("INFO : " + time.Now().String() + " " + clusterName + " Metrics was received from prometheus http-api with using http client.\n")
	return jsonBody
}
