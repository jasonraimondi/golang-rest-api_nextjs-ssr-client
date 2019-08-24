package repository
//
//import (
//	"github.com/Masterminds/squirrel"
//	"github.com/jinzhu/gorm"
//
//	"git.jasonraimondi.com/jason/jasontest/app/models"
//)
//
//type ListPhotosRepository struct {
//	queryBuilder squirrel.StatementBuilderType
//	dbx          *gorm.DB
//}
//
//func (s *ListPhotosRepository) ForUser(userId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
//	query := s.queryBuilder.
//		Select().
//		From("photos").
//		Where(squirrel.Eq{
//			"user_id": userId,
//		})
//	baseURL := "http://localhost:1323/photos/user/" + userId
//	return s.listPhotos(query, "*", itemsPerPage, currentPage, baseURL)
//}
//
//func (s *ListPhotosRepository) listPhotos(query squirrel.SelectBuilder, sel string, itemsPerPage int64, currentPage int64, baseURL string) (*Paginator, error) {
//	totalCount, err := TotalCountForQuery(s.dbx, query)
//	if err != nil {
//		return nil, err
//	}
//	query = query.Column(sel).OrderBy("created_at DESC")
//	sql, args, err := PaginateQuery(itemsPerPage, currentPage, query)
//	if err != nil {
//		return nil, err
//	}
//	rows, err := s.dbx.Queryx(sql, args...)
//	if err != nil {
//		return nil, err
//	}
//	var results []interface{}
//	for rows.Next() {
//		p := models.Photo{}
//		err = rows.StructScan(&p)
//		if err != nil {
//			continue
//		}
//		results = append(results, p)
//	}
//	// @todo env var for base url
//	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
//}
