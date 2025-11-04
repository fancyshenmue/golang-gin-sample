package cdn

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

// GreyCDNUseCase handles GreyCDN operations
type GreyCDNUseCase struct {
	repo repository.GreyCDNRepository
}

// NewGreyCDNUseCase creates a new GreyCDN use case
func NewGreyCDNUseCase(repo repository.GreyCDNRepository) *GreyCDNUseCase {
	return &GreyCDNUseCase{
		repo: repo,
	}
}

// ListSites retrieves all sites
func (uc *GreyCDNUseCase) ListSites(ctx context.Context) ([]string, error) {
	return uc.repo.ListSites(ctx)
}

// PurgeBySite purges cache for a site
func (uc *GreyCDNUseCase) PurgeBySite(ctx context.Context, site string) error {
	return uc.repo.PurgeBySite(ctx, site)
}

// PurgeBySiteURI purges cache for a specific URI
func (uc *GreyCDNUseCase) PurgeBySiteURI(ctx context.Context, site, uri string) error {
	return uc.repo.PurgeBySiteURI(ctx, site, uri)
}
