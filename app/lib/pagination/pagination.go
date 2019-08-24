package pagination

import (
	paginator "github.com/pilagod/gorm-cursor-paginator"
)

type PagingQuery struct {
	Cursor paginator.Cursor
	Limit  *int
	Order  *string
}

func GetPhotoPaginator(q PagingQuery) *paginator.Paginator {
	p := paginator.New()

	p.SetKeys("CreatedAt", "ID") // [defualt: "ID"] (supports multiple keys, and order of keys matters)

	if q.Cursor.After != nil {
		p.SetAfterCursor(*q.Cursor.After) // [default: nil]
	}

	if q.Cursor.Before != nil {
		p.SetBeforeCursor(*q.Cursor.Before) // [default: nil]
	}

	if q.Limit != nil {
		p.SetLimit(*q.Limit) // [default: 10]
	}

	if q.Order != nil && *q.Order == "asc" {
		p.SetOrder(paginator.ASC) // [default: paginator.DESC]
	}
	return p
}
