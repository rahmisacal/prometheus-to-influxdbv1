package domain

type Json struct {
	Status string `json:"status"`
	Data   *Data
}

type Data struct {
	ResultType string `json:"resultType"`
	Result     Result `json:"result"`
}

type Result []struct {
	Metric *Metric `json:"metric"`
	Value  []string `json:"value"`
}

type Metric struct {
	Name      string `json:"__name__"`
	Container string `json:"container"`
	Endpoint  string `json:"endpoint"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
	Service   string `json:"service"`
}

type Configuration struct {
	Clusters []struct {
		Name    string `yaml:"name"`
		Token   string `yaml:"token"`
		Cpu     string `yaml:"node_cpu_utilisation"`
		Memory  string `yaml:"node_memory_utilisation"`
		Network string `yaml:"node_network_receive_bytes_excluding_lo"`
		Disk    string `yaml:"node_disk_io_time_seconds_total"`
	} `yaml:"Clusters"`
	InfluxDb *InfluxDb `yaml:"influxdb"`
}

type InfluxDb struct {
	Url string `yaml:"Url"`
	Port int `yaml:"Port"`
	Measurement string `yaml:"Measurement"`
	Database string `yaml:"Database"`
	RetentionPolicy string `yaml:"RetentionPolicy"`
}