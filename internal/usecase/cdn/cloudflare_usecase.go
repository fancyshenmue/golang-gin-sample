package cdn

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

// CloudflareUseCase handles Cloudflare CDN operations
type CloudflareUseCase struct {
	repo repository.CloudflareRepository
}

// NewCloudflareUseCase creates a new Cloudflare use case
func NewCloudflareUseCase(repo repository.CloudflareRepository) *CloudflareUseCase {
	return &CloudflareUseCase{
		repo: repo,
	}
}

// ListZones retrieves all zones
func (uc *CloudflareUseCase) ListZones(ctx context.Context) ([]*entity.CDNZone, error) {
	return uc.repo.ListZones(ctx)
}

// GetZone retrieves a specific zone
func (uc *CloudflareUseCase) GetZone(ctx context.Context, zoneID string) (*entity.CDNZone, error) {
	return uc.repo.GetZone(ctx, zoneID)
}

// CreateZone creates a new zone
func (uc *CloudflareUseCase) CreateZone(ctx context.Context, name string) (*entity.CDNZone, error) {
	return uc.repo.CreateZone(ctx, name)
}

// DeleteZone deletes a zone
func (uc *CloudflareUseCase) DeleteZone(ctx context.Context, zoneID string) error {
	return uc.repo.DeleteZone(ctx, zoneID)
}

// ListDNSRecords retrieves DNS records for a zone
func (uc *CloudflareUseCase) ListDNSRecords(ctx context.Context, zoneID string) ([]*entity.DNSRecord, error) {
	return uc.repo.ListDNSRecords(ctx, zoneID)
}

// GetDNSRecord retrieves a specific DNS record
func (uc *CloudflareUseCase) GetDNSRecord(ctx context.Context, zoneID, recordID string) (*entity.DNSRecord, error) {
	return uc.repo.GetDNSRecord(ctx, zoneID, recordID)
}

// CreateDNSRecord creates a new DNS record
func (uc *CloudflareUseCase) CreateDNSRecord(ctx context.Context, record *entity.DNSRecord) (*entity.DNSRecord, error) {
	return uc.repo.CreateDNSRecord(ctx, record)
}

// DeleteDNSRecord deletes a DNS record
func (uc *CloudflareUseCase) DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error {
	return uc.repo.DeleteDNSRecord(ctx, zoneID, recordID)
}

// PurgeCache purges cache for a zone
func (uc *CloudflareUseCase) PurgeCache(ctx context.Context, purge *entity.CachePurge) error {
	return uc.repo.PurgeCache(ctx, purge)
}
