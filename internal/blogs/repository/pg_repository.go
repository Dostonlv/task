package repository

import (
	"context"
	"database/sql"
	"github.com/Dostonlv/task/pkg/utils"

	"github.com/Dostonlv/task/internal/blogs"
	"github.com/Dostonlv/task/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type blogsRepo struct {
	db *sqlx.DB
}

// NewBlogsRepository constructor
func NewBlogsRepository(db *sqlx.DB) blogs.Repository {
	return &blogsRepo{db: db}
}

// Create blogs
func (b *blogsRepo) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	newUUID := uuid.New()
	result := &models.Blog{}
	query := `INSERT INTO blogs (id,title,content) VALUES ($1,$2,$3) RETURNING *`
	if err := b.db.QueryRowxContext(ctx, query, newUUID, &blog.Title, &blog.Content).StructScan(result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.Create.QueryRowxContext")
	}
	return result, nil
}

// Update blogs
func (b *blogsRepo) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	result := &models.Blog{}
	query := `UPDATE blogs SET title=$1,content=$2 WHERE id=$3 RETURNING *`
	if err := b.db.QueryRowxContext(ctx, query, &blog.Title, &blog.Content, &blog.ID).StructScan(result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.Update.QueryRowxContext")
	}
	return result, nil
}

// Delete  blogs
func (b *blogsRepo) Delete(ctx context.Context, blogID uuid.UUID) error {
	query := `DELETE FROM blogs WHERE id=$1`
	result, err := b.db.ExecContext(ctx, query, blogID)
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "blogsRepo.Delete.RowsAffected")
	}
	return nil

}

// GetByID blogs
func (b *blogsRepo) GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error) {
	result := &models.Blog{}
	query := `SELECT id, title, content, created_at FROM blogs WHERE id=$1`
	if err := b.db.QueryRowxContext(ctx, query, blogID).StructScan(result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetByID.QueryRowxContext")
	}
	return result, nil
}

// GetAll blogs
func (b *blogsRepo) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogList, error) {
	var (
		totalCount    int
		getTotalCount = `SELECT COUNT(id) FROM blogs WHERE 1=1`
		getAllBlogs   = `SELECT id, title, content, created_at FROM blogs WHERE 1=1`
	)
	if title != "" {
		getTotalCount += " AND title ILIKE '%" + title + "%';"
		getAllBlogs += ` AND title ILIKE '%` + title + `%'`
	}
	getAllBlogs += ` ORDER BY created_at  OFFSET $1  LIMIT $2;`
	if err := b.db.QueryRowContext(ctx, getTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryRowContext")
	}
	if totalCount == 0 {
		return &models.BlogList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Blogs:      make([]*models.Blog, 0),
		}, nil
	}
	rows, err := b.db.QueryxContext(ctx, getAllBlogs, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	blogList := make([]*models.Blog, 0, query.GetSize())
	for rows.Next() {
		blog := &models.Blog{}
		if err := rows.StructScan(blog); err != nil {
			return nil, errors.Wrap(err, "blogsRepo.GetAll.StructScan")
		}
		blogList = append(blogList, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.rows.Err")
	}

	return &models.BlogList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Blogs:      blogList,
	}, nil

}
