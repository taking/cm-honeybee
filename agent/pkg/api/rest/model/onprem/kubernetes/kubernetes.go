package kubernetes

import "time"

type Kubernetes struct {
	Nodes     []Node                 `json:"nodes"`
	Workloads map[string]interface{} `json:"workloads"`
}

type Node struct {
	Name      interface{} `json:"name,omitempty"`
	Labels    interface{} `json:"labels,omitempty"`
	Addresses interface{} `json:"addresses,omitempty"`
	NodeInfo  interface{} `json:"nodeinfo,omitempty"`
}

type Helm struct {
	Repo    []Repo    `json:"repo"`
	Release []Release `json:"release"`
}

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Release struct {
	Name             string    `json:"name"`
	Namespace        string    `json:"namespace"`
	Revision         int       `json:"revision"`
	Updated          time.Time `json:"updated"`
	Status           string    `json:"status"`
	ChartNameVersion string    `json:"chart"`
	AppVersion       string    `json:"app_version"`
}