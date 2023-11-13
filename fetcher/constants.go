package fetcher

import es8types "github.com/elastic/go-elasticsearch/v8/typedapi/types"

const BANNER = "ElasticSearch Greeting Fetch"

type ElasticGreeting struct {
	Name        string                            `json:"name"`
	ClusterName string                            `json:"cluster_name"`
	ClusterUuid string                            `json:"cluster_uuid"`
	VersionInfo es8types.ElasticsearchVersionInfo `json:"version"`
	Tagline     string                            `json:"tagline"`
}
