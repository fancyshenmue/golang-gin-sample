package repository

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
)

// CloudflareRepository defines the interface for Cloudflare operations
type CloudflareRepository interface {
	// Zone operations
	ListZones(ctx context.Context) ([]*entity.CDNZone, error)
	GetZone(ctx context.Context, zoneID string) (*entity.CDNZone, error)
	CreateZone(ctx context.Context, name string) (*entity.CDNZone, error)
	DeleteZone(ctx context.Context, zoneID string) error

	// DNS operations
	ListDNSRecords(ctx context.Context, zoneID string) ([]*entity.DNSRecord, error)
	GetDNSRecord(ctx context.Context, zoneID, recordID string) (*entity.DNSRecord, error)
	CreateDNSRecord(ctx context.Context, record *entity.DNSRecord) (*entity.DNSRecord, error)
	DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error

	// Cache operations
	PurgeCache(ctx context.Context, purge *entity.CachePurge) error
}

// VerizonRepository defines the interface for Verizon Edge operations
type VerizonRepository interface {
	// Origin operations
	GetOrigin(ctx context.Context, originID int) (*entity.CDNOrigin, error)
	AddOrigin(ctx context.Context, origin *entity.CDNOrigin) (*entity.CDNOrigin, error)
	DeleteOrigin(ctx context.Context, originID int) error

	// CNAME operations
	ListCNAMEs(ctx context.Context) ([]*entity.CDNCNAME, error)
	GetCNAME(ctx context.Context, cnameID int) (*entity.CDNCNAME, error)
	AddCNAME(ctx context.Context, cname *entity.CDNCNAME) (*entity.CDNCNAME, error)
	DeleteCNAME(ctx context.Context, cnameID int) error

	// Cache operations
	PurgeCache(ctx context.Context, urls []string) error
}

// GreyCDNRepository defines the interface for GreyCDN operations
type GreyCDNRepository interface {
	ListSites(ctx context.Context) ([]string, error)
	PurgeBySite(ctx context.Context, site string) error
	PurgeBySiteURI(ctx context.Context, site, uri string) error
}
