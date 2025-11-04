package cdn

import (
	"context"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/entity"
	"github.com/fancyshenmue/golang-gin-sample/internal/domain/repository"
)

// VerizonUseCase handles Verizon Edge CDN operations
type VerizonUseCase struct {
	repo repository.VerizonRepository
}

// NewVerizonUseCase creates a new Verizon use case
func NewVerizonUseCase(repo repository.VerizonRepository) *VerizonUseCase {
	return &VerizonUseCase{
		repo: repo,
	}
}

// GetOrigin retrieves an origin by ID
func (uc *VerizonUseCase) GetOrigin(ctx context.Context, originID int) (*entity.CDNOrigin, error) {
	return uc.repo.GetOrigin(ctx, originID)
}

// AddOrigin adds a new origin
func (uc *VerizonUseCase) AddOrigin(ctx context.Context, origin *entity.CDNOrigin) (*entity.CDNOrigin, error) {
	return uc.repo.AddOrigin(ctx, origin)
}

// DeleteOrigin deletes an origin
func (uc *VerizonUseCase) DeleteOrigin(ctx context.Context, originID int) error {
	return uc.repo.DeleteOrigin(ctx, originID)
}

// ListCNAMEs retrieves all CNAMEs
func (uc *VerizonUseCase) ListCNAMEs(ctx context.Context) ([]*entity.CDNCNAME, error) {
	return uc.repo.ListCNAMEs(ctx)
}

// GetCNAME retrieves a CNAME by ID
func (uc *VerizonUseCase) GetCNAME(ctx context.Context, cnameID int) (*entity.CDNCNAME, error) {
	return uc.repo.GetCNAME(ctx, cnameID)
}

// AddCNAME adds a new CNAME
func (uc *VerizonUseCase) AddCNAME(ctx context.Context, cname *entity.CDNCNAME) (*entity.CDNCNAME, error) {
	return uc.repo.AddCNAME(ctx, cname)
}

// DeleteCNAME deletes a CNAME
func (uc *VerizonUseCase) DeleteCNAME(ctx context.Context, cnameID int) error {
	return uc.repo.DeleteCNAME(ctx, cnameID)
}

// PurgeCache purges cache for URLs
func (uc *VerizonUseCase) PurgeCache(ctx context.Context, urls []string) error {
	return uc.repo.PurgeCache(ctx, urls)
}
