package pagygo

import (
	"math"
	"sync"
)

type (
	Context struct {
		items    []interface{}
		page     int
		perPage  int
		orderBy  func([]interface{}) []interface{}
		filterBy func(interface{}) bool
		config   Config
		mutex    sync.RWMutex
	}
	Pagygo interface {
		Page(int) Pagygo
		PerPage(int) Pagygo
		OrderBy(callback func([]interface{}) []interface{}) Pagygo
		FilterBy(callback func(interface{}) bool) Pagygo
		Paginate() Paginated
	}
	Paginated struct {
		Items []interface{}
		Meta  Meta
	}
	Config struct {
		DefaultPerPage  *int
		DefaultPage     *int
		MaxPerPage      *int
		DefaultOrderBy  *func([]interface{}) []interface{}
		DefaultFilterBy *func(interface{}) bool
	}
	Meta struct {
		CurrentPage     int
		PrevPage        *int
		NextPage        *int
		PerPage         int
		TotalPagesCount int
		TotalItemsCount int
		IsFirstPage     bool
		IsLastPage      bool
	}
)

func New(items []interface{}, config Config) Pagygo {
	defaultPerPage := 25
	defaultPage := 1
	var defaultOrderBy func([]interface{}) []interface{}
	var defaultFilterBy func(interface{}) bool

	if config.DefaultPerPage != nil {
		defaultPerPage = *config.DefaultPerPage
	}

	if config.DefaultPage != nil {
		defaultPage = *config.DefaultPage
	}

	if config.DefaultOrderBy != nil {
		defaultOrderBy = *config.DefaultOrderBy
	}

	if config.DefaultFilterBy != nil {
		defaultFilterBy = *config.DefaultFilterBy
	}

	return &Context{
		items:    items,
		page:     defaultPage,
		perPage:  defaultPerPage,
		orderBy:  defaultOrderBy,
		filterBy: defaultFilterBy,
		config:   config,
	}
}

func (c *Context) Page(page int) Pagygo {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if page < 1 {
		page = 1
	}

	c.page = page

	return c
}

func (c *Context) PerPage(perPage int) Pagygo {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if perPage < 1 {
		perPage = 1
	}

	if c.config.MaxPerPage != nil && perPage > *c.config.MaxPerPage {
		perPage = *c.config.MaxPerPage
	}

	c.perPage = perPage

	return c
}

func (c *Context) OrderBy(callback func([]interface{}) []interface{}) Pagygo {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.orderBy = callback

	return c
}

func (c *Context) FilterBy(callback func(interface{}) bool) Pagygo {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.filterBy = callback

	return c
}

func (c *Context) Paginate() Paginated {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var currentItems []interface{}

	if c.filterBy != nil {
		for _, item := range c.items {
			if c.filterBy(item) {
				currentItems = append(currentItems, item)
			}
		}
	} else {
		currentItems = c.items
	}

	if c.orderBy != nil {
		currentItems = c.orderBy(currentItems)
	}

	totalItemsCount := len(currentItems)
	var totalPagesCount int
	var prevPage *int = nil
	var nextPage *int = nil

	if totalItemsCount == 0 {
		totalPagesCount = 1
	} else {
		totalPagesCount = int(math.Ceil(float64(totalItemsCount) / float64(c.perPage)))
	}

	if (c.page-1) >= 1 && (c.page-1) <= totalPagesCount {
		var prevPageValue = c.page - 1
		prevPage = &prevPageValue
	}

	if (c.page+1) > 1 && (c.page+1) <= totalPagesCount {
		var nextPageValue = c.page + 1
		nextPage = &nextPageValue
	}

	isFirstPage := c.page == 1
	isLastPage := c.page == totalPagesCount || totalItemsCount == 0

	meta := Meta{
		TotalItemsCount: totalItemsCount,
		CurrentPage:     c.page,
		PerPage:         c.perPage,
		TotalPagesCount: totalPagesCount,
		PrevPage:        prevPage,
		NextPage:        nextPage,
		IsFirstPage:     isFirstPage,
		IsLastPage:      isLastPage,
	}

	startIndex := (c.page - 1) * c.perPage
	endIndex := startIndex + c.perPage

	if startIndex > totalItemsCount {
		startIndex = totalItemsCount
	}

	if endIndex > totalItemsCount {
		endIndex = totalItemsCount
	}

	currentItems = currentItems[startIndex:endIndex]

	return Paginated{
		Meta:  meta,
		Items: currentItems,
	}
}
