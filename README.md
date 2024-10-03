![Build and test](https://github.com/nejdetkadir/pagygo/actions/workflows/main.yml/badge.svg?branch=main)
![Go Version](https://img.shields.io/badge/go_version-_1.23.1-007d9c.svg)

![cover](docs/cover.png)

# Pagygo

Pagygo is a flexible pagination library for Golang, designed to handle large or irregular datasets. It supports advanced features such as dynamic filtering, custom sorting, and configurable pagination options to suit a variety of use cases.

## Features
- **Dynamic Filtering:** Filter data based on custom logic before paginating.
- **Custom Sorting:** Define sorting logic using a callback function.
- **Pagination Controls:** Configure the number of items per page and navigate between pages.
- **Thread-Safe Operations:** All operations are safe to use in concurrent environments.
- **Meta Information:** Get detailed pagination metadata like current page, total pages, and more.

## Installation
To install Pagygo, use the following command:

```bash
go get github.com/nejdetkadir/pagygo
```

## Usage
### Initialize Pagygo
Create a new instance of the pagination context and configure the items to paginate, along with custom options if needed.

```go
package main

import (
    "github.com/nejdetkadir/pagygo"
    "fmt"
)

func main() {
    items := []interface{}{1, 2, 3, 4, 5}
    config := pagygo.Config{}

    paginator := pagygo.New(items, config)
    result := paginator.Paginate()

    fmt.Printf("Current Page: %d, Total Items: %d\n", result.Meta.CurrentPage, result.Meta.TotalItemsCount)
}
```

### Customize Pagination
You can specify the current page and the number of items per page.

```go
result := paginator.Page(2).PerPage(10).Paginate()
```

### Filtering Items
Use the FilterBy function to define a filtering logic for the items.

```go
result := paginator.FilterBy(func(item interface{}) bool {
    return item.(int) % 2 == 0
}).Paginate()

fmt.Println(result.Items) // Output: [2 4]
```

### Sorting Items
You can pass a custom sorting function to order the items before pagination.

```go
result := paginator.OrderBy(func(items []interface{}) []interface{} {
    // Reverse the items
    for i := len(items)/2 - 1; i >= 0; i-- {
        opp := len(items) - 1 - i
        items[i], items[opp] = items[opp], items[i]
    }
    return items
}).Paginate()

fmt.Println(result.Items) // Output: [5 4 3 2 1]
```

### Pagination Metadata
After calling Paginate(), you will receive detailed pagination metadata in the Meta struct. The fields include:

- **CurrentPage:** The current page number.
- **PrevPage:** The previous page number (or nil if there is no previous page).
- **NextPage:** The next page number (or nil if there is no next page).
- **PerPage:** The number of items per page.
- **TotalPagesCount:** The total number of pages based on the items and PerPage setting.
- **TotalItemsCount:** The total number of items after filtering (if any filter is applied).
- **IsFirstPage:** A boolean indicating whether the current page is the first page.
- **IsLastPage**: A boolean indicating whether the current page is the last page.

Here is an example of how to access the metadata:

```go
meta := result.Meta
fmt.Printf(
    "Current Page: %d\nTotal Pages: %d\nTotal Items: %d\nIs First Page: %v\nIs Last Page: %v\nPrev Page: %v\nNext Page: %v\nItems Per Page: %d\n",
    meta.CurrentPage,        // Current page number
    meta.TotalPagesCount,    // Total pages available
    meta.TotalItemsCount,    // Total number of items
    meta.IsFirstPage,        // Is this the first page?
    meta.IsLastPage,         // Is this the last page?
    meta.PrevPage,           // Previous page number (nil if no previous page)
    meta.NextPage,           // Next page number (nil if no next page)
    meta.PerPage,            // Number of items per page
)
```

For example, the output might look like this:

```
Current Page: 1
Total Pages: 3
Total Items: 50
Is First Page: true
Is Last Page: false
Prev Page: <nil>
Next Page: 2
Items Per Page: 20
```

### Example with All Features
```go
result := paginator.Page(2).PerPage(5).
    FilterBy(func(item interface{}) bool { return item.(int) > 2 }).
    OrderBy(func(items []interface{}) []interface{} { 
        // Example sorting logic
        return items 
    }).Paginate()

fmt.Println(result.Items) // Paginated, filtered, and sorted items
```

## Example
Check out the [example](example/main.go) for a complete demonstration of pagygo usage.

## Unit Testing
The package includes comprehensive unit tests to ensure correct behavior across various scenarios. To run the tests, use the following:

```b![cover.png](../../../Downloads/cover.png)ash
go test ./...
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/nejdetkadir/pagygo. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [code of conduct](https://github.com/nejdetkadir/pagygo/blob/main/CODE_OF_CONDUCT.md).

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
