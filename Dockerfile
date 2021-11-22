FROM golang:1.17.3

WORKDIR /usr/src/app

RUN go get -d -v github.com/go-yaml/yaml
RUN go get -d -v github.com/influxdata/influxdb1-client

COPY . .

RUN go build -o /promtoinflux

CMD [ "/promtoinflux" ]