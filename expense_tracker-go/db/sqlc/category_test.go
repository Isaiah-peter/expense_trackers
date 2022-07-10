package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateNewCategory() (Category, error) {
	arg := CreateCategoryParams{
		UserID: 1,
		Icon:   "how",
		Name:   "Book",
	}

	category, err := testQueries.CreateCategory(ctx, arg)

	return category, err
}

func TestCreateCategory(t *testing.T) {
	category, err := CreateNewCategory()

	arg := CreateCategoryParams{
		UserID: 1,
		Icon:   "how",
		Name:   "Book",
	}

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Icon, category.Icon)
	require.Equal(t, arg.Name, category.Name)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdatedAt)
}

func TestGetCategory(t *testing.T) {
	category, err := CreateNewCategory()

	require.NoError(t, err)

	result, err := testQueries.GetCategory(ctx, category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, category.Icon, result.Icon)
	require.Equal(t, category.Name, result.Name)
	require.Equal(t, category.UserID, result.UserID)
	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)
	require.NotZero(t, result.UpdatedAt)
}

func TestUpdateCategory(t *testing.T) {
	category, err := CreateNewCategory()

	require.NoError(t, err)
	require.NotEmpty(t, category)

	arg := UpdateCategoryParams{
		ID: category.ID,
		Name: "good",
		Icon: "feel",
		UpdatedAt: time.Now(),
	}

	err = testQueries.UpdateCategory(ctx, arg)
	require.NoError(t, err)
}

func TestListCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateNewCategory()
	}

	arg := ListCategoriesParams{
		Limit: 5,
		Offset: 5,
	}

	categorys, err := testQueries.ListCategories(ctx, arg)
	require.NoError(t, err)

	for _, category := range(categorys) {
		require.NotEmpty(t, category)
	}
}

func TestDeleteCategory(t *testing.T) {
	category, err := CreateNewCategory()

	require.NoError(t, err)
	err = testQueries.DeleteCategory(ctx, category.ID)
	require.NoError(t, err)
}

func TestListCategoryByUserID(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateNewCategory()
	}

	arg := ListCategoryByUserIdParams{
		Limit: 5,
		Offset: 5,
		UserID: 1,
	}

	categorys, err := testQueries.ListCategoryByUserId(ctx, arg)
	require.NoError(t, err)

	for _, category := range(categorys) {
		require.NotEmpty(t, category)
	}
}
