package pagygo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginate_DefaultPagination(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5}
	config := Config{}

	paginator := New(items, config)
	result := paginator.Paginate()

	assert.Equal(t, 1, result.Meta.CurrentPage)
	assert.Equal(t, 25, result.Meta.PerPage)
	assert.Equal(t, 1, result.Meta.TotalPagesCount)
	assert.Equal(t, 5, result.Meta.TotalItemsCount)
	assert.True(t, result.Meta.IsFirstPage)
	assert.True(t, result.Meta.IsLastPage)
	assert.Nil(t, result.Meta.PrevPage)
	assert.Nil(t, result.Meta.NextPage)
}

func TestPaginate_WithCustomPerPage(t *testing.T) {
	items := make([]interface{}, 50)
	config := Config{}

	paginator := New(items, config).PerPage(10)
	result := paginator.Paginate()

	assert.Equal(t, 10, result.Meta.PerPage)
	assert.Equal(t, 5, result.Meta.TotalPagesCount)
	assert.Equal(t, 50, result.Meta.TotalItemsCount)
	assert.True(t, result.Meta.IsFirstPage)
	assert.False(t, result.Meta.IsLastPage)
	assert.Nil(t, result.Meta.PrevPage)
	assert.NotNil(t, result.Meta.NextPage)
}

func TestPaginate_WithCustomPage(t *testing.T) {
	items := make([]interface{}, 50)
	config := Config{}

	paginator := New(items, config).PerPage(10).Page(3)
	result := paginator.Paginate()

	assert.Equal(t, 3, result.Meta.CurrentPage)
	assert.Equal(t, 10, result.Meta.PerPage)
	assert.Equal(t, 5, result.Meta.TotalPagesCount)
	assert.Equal(t, 50, result.Meta.TotalItemsCount)
	assert.False(t, result.Meta.IsFirstPage)
	assert.False(t, result.Meta.IsLastPage)
	assert.NotNil(t, result.Meta.PrevPage)
	assert.NotNil(t, result.Meta.NextPage)
}

func TestPaginate_WithFilter(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5}
	config := Config{}

	paginator := New(items, config).FilterBy(func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	result := paginator.Paginate()

	assert.Equal(t, 2, len(result.Items))
	assert.Equal(t, 2, result.Items[0].(int))
	assert.Equal(t, 4, result.Items[1].(int))
	assert.Equal(t, 2, result.Meta.TotalItemsCount)
}

func TestPaginate_WithOrder(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5}
	config := Config{}

	paginator := New(items, config).OrderBy(func(items []interface{}) []interface{} {
		for i := len(items)/2 - 1; i >= 0; i-- {
			opp := len(items) - 1 - i
			items[i], items[opp] = items[opp], items[i]
		}
		return items
	})
	result := paginator.Paginate()

	assert.Equal(t, 5, result.Items[0].(int))
	assert.Equal(t, 4, result.Items[1].(int))
	assert.Equal(t, 3, result.Items[2].(int))
	assert.Equal(t, 2, result.Items[3].(int))
	assert.Equal(t, 1, result.Items[4].(int))
}

func TestPaginate_EmptyItems(t *testing.T) {
	var items []interface{}
	config := Config{}

	paginator := New(items, config)
	result := paginator.Paginate()

	assert.Equal(t, 0, len(result.Items))
	assert.Equal(t, 0, result.Meta.TotalItemsCount)
	assert.Equal(t, 1, result.Meta.TotalPagesCount)
	assert.True(t, result.Meta.IsFirstPage)
	assert.True(t, result.Meta.IsLastPage)
}

func TestPaginate_InvalidPage(t *testing.T) {
	items := make([]interface{}, 10) // 10 items
	for i := 0; i < 10; i++ {
		items[i] = i + 1
	}

	config := Config{}
	paginator := New(items, config).PerPage(2).Page(6)
	result := paginator.Paginate()

	assert.Equal(t, 5, result.Meta.TotalPagesCount)
	assert.Equal(t, 6, result.Meta.CurrentPage)
	assert.Equal(t, 0, len(result.Items))
	assert.False(t, result.Meta.IsLastPage)
	assert.False(t, result.Meta.IsFirstPage)
	assert.Nil(t, result.Meta.NextPage)
	assert.NotNil(t, result.Meta.PrevPage)
}

func TestPaginate_MaxPerPageExceeded(t *testing.T) {
	items := make([]interface{}, 20) // 20 items
	for i := 0; i < 20; i++ {
		items[i] = i + 1
	}

	maxPerPage := 5
	config := Config{
		MaxPerPage: &maxPerPage,
	}

	paginator := New(items, config).PerPage(10).Page(1)
	result := paginator.Paginate()

	assert.Equal(t, 5, result.Meta.PerPage)
	assert.Equal(t, 4, result.Meta.TotalPagesCount)
	assert.Equal(t, 5, len(result.Items))
	assert.Equal(t, 1, result.Items[0].(int))
	assert.Equal(t, 5, result.Items[4].(int))
	assert.False(t, result.Meta.IsLastPage)
	assert.True(t, result.Meta.IsFirstPage)
	assert.NotNil(t, result.Meta.NextPage)
	assert.Nil(t, result.Meta.PrevPage)
}
