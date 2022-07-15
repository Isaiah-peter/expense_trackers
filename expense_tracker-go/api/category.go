package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/Isaiah-peter/expense_tracker/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	UserID int32  `json:"user_id" binding:"required"`
	Icon   string `json:"icon"`
	Name   string `json:"name" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		UserID: req.UserID,
		Icon:   req.Icon,
		Name:   req.Name,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) GetCategory(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) listAllCategory(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit:  req.PageSize,
		Offset: (req.Offset - 1) * req.PageSize,
	}

	categories, err := server.store.ListCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (server *Server) listAllCategoryByUserID(ctx *gin.Context) {
	var req listUserRequest
	var id userupdateID

	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoryByUserIdParams{
		UserID: id.ID,
		Limit:  req.PageSize,
		Offset: (req.Offset - 1) * req.PageSize,
	}

	catergories, err := server.store.ListCategoryByUserId(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, catergories)
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var req createCategoryRequest
	var id userupdateID

	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, id.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.Icon == "" {
		req.Icon = category.Icon
	}

	if req.UserID == 0 {
		req.UserID = category.UserID
	}

	if req.Name == "" {
		req.Name = category.Name
	}

	arg := db.UpdateCategoryParams{
		ID:        id.ID,
		Icon:      req.Icon,
		Name:      req.Name,
		UpdatedAt: time.Now(),
	}

	if err := server.store.UpdateCategory(ctx, arg); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"updates": "Successfully updated"})
	}
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category successfully deleted"})
}
