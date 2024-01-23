package usecase

import (
	"context"
	"github.com/Dostonlv/task/config"
	"github.com/Dostonlv/task/internal/blogs"
	"github.com/Dostonlv/task/internal/models"
	"github.com/Dostonlv/task/pkg/logger"
	"github.com/Dostonlv/task/pkg/utils"
	"github.com/google/uuid"
)

// Blogs UseCase
type blogsUC struct {
	cfg       *config.Config
	blogsRepo blogs.Repository
	logger    logger.Logger
}

// blogs UseCase constructor
func NewBlogsUseCase(cfg *config.Config, blogsRepo blogs.Repository, logger logger.Logger) blogs.UseCase {
	return &blogsUC{
		cfg:       cfg,
		blogsRepo: blogsRepo,
		logger:    logger,
	}
}

// Create blogs
func (b *blogsUC) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	return b.blogsRepo.Create(ctx, blog)
}

// Update blogs
func (b *blogsUC) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	updatedBlog, err := b.blogsRepo.Update(ctx, blog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

// Delete blogs
func (b *blogsUC) Delete(ctx context.Context, blogID uuid.UUID) error {
	if err := b.blogsRepo.Delete(ctx, blogID); err != nil {
		return err
	}
	return nil
}

// GetByID blogs
func (b *blogsUC) GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error) {
	return b.blogsRepo.GetByID(ctx, blogID)
}

// GetAll blogs
func (b *blogsUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogList, error) {
	return b.blogsRepo.GetAll(ctx, title, query)
}
