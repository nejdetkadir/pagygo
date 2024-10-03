package main

import (
	"github.com/nejdetkadir/pagygo"
	"log"
)

func main() {
	// Define a dataset to paginate
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Defaults
	defaultPerPage := 3
	defaultPage := 1
	maxPerPage := 5

	// Create default pagination configuration
	config := pagygo.Config{
		DefaultPerPage:  &defaultPerPage, // Default items per page
		DefaultPage:     &defaultPage,    // Default starting page
		MaxPerPage:      &maxPerPage,     // Maximum items per page
		DefaultOrderBy:  nil,             // No default order
		DefaultFilterBy: nil,             // No default filter
	}

	// Initialize Pagygo with items and config
	paginator := pagygo.New(items, config)

	// Example 1: Default pagination (page 1, 3 items per page)
	log.Println("Example 1: Default Pagination")
	result := paginator.Paginate()
	logPagination(result)

	// Example 2: Custom page and items per page
	log.Println("\nExample 2: Custom Page and Items Per Page")
	result = paginator.Page(2).PerPage(5).Paginate()
	logPagination(result)

	// Example 3: Filtering items (only even numbers)
	log.Println("\nExample 3: Filtering Items (Even Numbers)")
	result = paginator.FilterBy(func(item interface{}) bool {
		return item.(int)%2 == 0
	}).Paginate()
	logPagination(result)

	// Example 4: Custom sorting (reverse order)
	log.Println("\nExample 4: Custom Sorting (Reverse Order)")
	result = paginator.OrderBy(func(items []interface{}) []interface{} {
		for i := len(items)/2 - 1; i >= 0; i-- {
			opp := len(items) - 1 - i
			items[i], items[opp] = items[opp], items[i]
		}
		return items
	}).Paginate()
	logPagination(result)

	// Example 5: Custom page, filtering, and sorting
	log.Println("\nExample 5: Custom Page with Filtering and Sorting")
	result = paginator.Page(2).PerPage(3).
		FilterBy(func(item interface{}) bool { return item.(int) > 3 }).
		OrderBy(func(items []interface{}) []interface{} { return items }).
		Paginate()
	logPagination(result)
}

// Helper function to log the results
func logPagination(result pagygo.Paginated) {
	log.Printf("Current Page: %d", result.Meta.CurrentPage)
	log.Printf("Items on Page: %v", result.Items)
	log.Printf("Total Items: %d", result.Meta.TotalItemsCount)
	log.Printf("Items Per Page: %d", result.Meta.PerPage)
	log.Printf("Total Pages: %d", result.Meta.TotalPagesCount)
	log.Printf("Is First Page: %v", result.Meta.IsFirstPage)
	log.Printf("Is Last Page: %v", result.Meta.IsLastPage)
	if result.Meta.PrevPage != nil {
		log.Printf("Previous Page: %d", *result.Meta.PrevPage)
	}
	if result.Meta.NextPage != nil {
		log.Printf("Next Page: %d", *result.Meta.NextPage)
	}
}
