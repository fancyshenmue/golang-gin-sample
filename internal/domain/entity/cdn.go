package entity

// CDNZone represents a CDN zone/domain
type CDNZone struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// DNSRecord represents a DNS record
type DNSRecord struct {
	ID      string `json:"id"`
	ZoneID  string `json:"zone_id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

// CachePurge represents a cache purge operation
type CachePurge struct {
	ZoneID string   `json:"zone_id"`
	URLs   []string `json:"urls"`
	PurgeAll bool   `json:"purge_all"`
}

// CDNOrigin represents a CDN origin server
type CDNOrigin struct {
	ID              int    `json:"id"`
	CustomerID      int    `json:"customer_id"`
	Name            string `json:"name"`
	Hostname        string `json:"hostname"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	UseIPRestriction bool  `json:"use_ip_restriction"`
}

// CDNCNAME represents a CDN CNAME record
type CDNCNAME struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	OriginID   int    `json:"origin_id"`
	DirPath    string `json:"dir_path"`
	EnableGzip int    `json:"enable_gzip"`
}
