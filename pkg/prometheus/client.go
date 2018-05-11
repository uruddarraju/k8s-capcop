package prometheus

import (
	"net"
	"net/http"
	"time"

	"fmt"
	log "github.com/Sirupsen/logrus"
	promapi "github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"golang.org/x/net/context"
)

var DefaultRoundTripper http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

type Client struct {
	Client v1.API
}

type ResultRecord struct {
	Vectors map[string]string
	Value   float64
}

type ContainerRecord struct {
	Namespace string
	Labels    map[string]string
	PodName   string
}

func NewClient(endpoint string) (*Client, error) {
	cfg := promapi.Config{
		RoundTripper: DefaultRoundTripper,
		Address:      endpoint,
	}

	httpClient, err := promapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	promClient := v1.NewAPI(httpClient)

	return &Client{Client: &promClient}, nil
}

func (c *Client) GetAllContainersUsageLessThan(usageInBytes float64) []ContainerRecord {

	var containers []ContainerRecord

	ctx, _ := context.WithCancel(context.Background())

	result, err := c.Client.Query(ctx, fmt.Sprintf("sum(rate(container_network_transmit_bytes_total[1d])) by (pod_name, namespace) < %.6f", usageInBytes), time.Now())
	if err != nil {
		log.Errorf("error during query: %s", err.Error())
		return nil
	}

	vec := result.(model.Vector)
	for _, val := range vec {
		fmt.Println(val.Metric[model.LabelName("namespace")])
		cr := ContainerRecord{
			Namespace: string(val.Metric[model.LabelName("namespace")]),
			PodName:   string(val.Metric[model.LabelName("pod_name")]),
		}

		containers = append(containers, cr)
	}

	return containers
}
