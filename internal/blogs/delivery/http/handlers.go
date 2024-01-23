package http

import (
	"github.com/Dostonlv/task/config"
	"github.com/Dostonlv/task/internal/blogs"
	"github.com/Dostonlv/task/internal/models"
	"github.com/Dostonlv/task/pkg/httpErrors"
	"github.com/Dostonlv/task/pkg/logger"
	"github.com/Dostonlv/task/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// blogs handlers
type blogsHandlers struct {
	cfg     *config.Config
	blogsUC blogs.UseCase
	logger  logger.Logger
}

// NewBlogsHandlers blog handlers constructor
func NewBlogsHandlers(cfg *config.Config, blogsUC blogs.UseCase, logger logger.Logger) blogs.Handlers {
	return &blogsHandlers{cfg: cfg, blogsUC: blogsUC, logger: logger}
}

// Create
// @Summary Create blog
// @Description Create blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param body body models.BlogSwagger true "blog"
// @Success 201 {object} models.Blog
// @Failure 500 {object}  httpErrors.RestErr
// @Router /blogs [POST]
func (h *blogsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		blog := &models.Blog{}
		if err := utils.SanitizeRequest(c, blog); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}
		createdBlog, err := h.blogsUC.Create(c.Request().Context(), blog)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, createdBlog)
	}
}

// Update
// @Summary Update
// @Description Update blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.BlogSwagger true "body"
// @Success 200 {object} models.Blog
// @Failure 500 {object}  httpErrors.RestErr
// @Router /blogs/{id} [PUT]
func (h *blogsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		blog := &models.Blog{}
		if err = utils.SanitizeRequest(c, blog); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		updatedBlog, err := h.blogsUC.Update(c.Request().Context(), &models.Blog{
			ID:      blogsID,
			Title:   blog.Title,
			Content: blog.Content,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, updatedBlog)
	}
}

// Delete
// @Summary Delete
// @Description Delete blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object}  httpErrors.RestErr
// @Router /blogs/{id} [DELETE]
func (h *blogsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		if err = h.blogsUC.Delete(c.Request().Context(), blogsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.NoContent(http.StatusNoContent)
	}
}

// GetByID
// @Summary GetByID
// @Description Getting blog by id
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Blog
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [GET]
func (h *blogsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		blog, err := h.blogsUC.GetByID(c.Request().Context(), blogsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, blog)
	}
}

// GetAll
// @Summary GetAll
// @Description Get all blogs with pagination and search
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param title query string false "title"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {object} models.BlogList
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs [GET]
func (h *blogsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		gp, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		blogList, err := h.blogsUC.GetAll(c.Request().Context(), c.QueryParam("title"), gp)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, blogList)
	}
}
