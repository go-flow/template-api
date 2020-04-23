package paging

import (
	"fmt"
	"strconv"
)

var (
	// PaginatorPerPageDefault is the amount of results per page
	PaginatorPerPageDefault = 20

	// PaginatorPageKey is the query parameter holding amount of results per page
	PaginatorPageKey = "page"

	// PaginatorPerPageKey is the query parameter holding the amount of results per page
	// to override the default one
	PaginatorPerPageKey = "per_page"
)

// Paginator is a type used to represent the pagination
type Paginator struct {
	// Current page you're on
	Page int `json:"page"`
	// Number of results you want per page
	PerPage int `json:"per_page"`
	// Page * PerPage (ex: 2 * 20, Offset == 40)
	Offset int `json:"offset"`
	// Total potential records matching the query
	TotalEntriesSize int `json:"total_entries_size"`
	// Total records returns, will be <= PerPage
	CurrentEntriesSize int `json:"current_entries_size"`
	// Total pages
	TotalPages int `json:"total_pages"`
	// OrderBy field
	OrderBy string `json:"order_by"`
	// Order Direction
	OrderDir string `json:"order_dir"`
	// Filter
	Filter string `json:"filter"`
}

// PaginationParams is a parameters provider interface to get the pagination params from
type PaginationParams interface {
	Get(key string) string
}

// NewWithDefaults creates Paginator object with default values
func NewWithDefaults() *Paginator {
	return NewPaginator(1, PaginatorPerPageDefault, "", "", "")
}

// NewPaginator returns a new `Paginator` value with the appropriate
// defaults set.
func NewPaginator(page int, perPage int, orderBy, orderDir, filter string) *Paginator {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = PaginatorPerPageDefault
	}
	p := &Paginator{Page: page, PerPage: perPage, OrderBy: orderBy, OrderDir: orderDir, Filter: filter}
	p.Offset = (page - 1) * p.PerPage
	return p
}

// NewPaginatorFromParams takes an interface of type `PaginationParams`,
// the `url.Values` type works great with this interface, and returns
// a new `Paginator` based on the params or `PaginatorPageKey` and
// `PaginatorPerPageKey`. Defaults are `1` for the page and
// PaginatorPerPageDefault for the per page value.
func NewPaginatorFromParams(params PaginationParams) *Paginator {
	page := "1"
	if p := params.Get("page"); p != "" {
		page = p
	}

	perPage := strconv.Itoa(PaginatorPerPageDefault)
	if pp := params.Get("per_page"); pp != "" {
		perPage = pp
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}

	pp, err := strconv.Atoi(perPage)
	if err != nil {
		pp = PaginatorPerPageDefault
	}

	orderBy := ""
	if ob := params.Get("order_by"); ob != "" {
		orderBy = ob
	}

	orderDir := ""
	if od := params.Get("order_dir"); od != "" {
		orderDir = od
	}

	filter := ""
	if f := params.Get("filter"); f != "" {
		filter = f
	}

	return NewPaginator(p, pp, orderBy, orderDir, filter)
}

// Order returns ordering string
func (p *Paginator) Order(defaultOrder string) string {
	if p.OrderBy == "" {
		p.OrderBy = defaultOrder
	}

	if p.OrderDir == "" {
		p.OrderDir = "ASC"
	}

	return fmt.Sprintf("%s %s", p.OrderBy, p.OrderDir)
}
