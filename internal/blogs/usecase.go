package blogs

import (
	"context"

	"github.com/Dostonlv/task/internal/models"
	"github.com/Dostonlv/task/pkg/utils"
	"github.com/google/uuid"
)

// blogs use case
type UseCase interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Update(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, blogID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogList, error)
}
