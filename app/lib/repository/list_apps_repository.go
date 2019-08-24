package repository
//
//import (
//	"github.com/Masterminds/squirrel"
//	"github.com/jinzhu/gorm"
//	"github.com/jmoiron/sqlx"
//
//	"git.jasonraimondi.com/jason/jasontest/app/models"
//)
//
//type ListAppsRepository struct {
//	queryBuilder squirrel.StatementBuilderType
//	dbx          *gorm.DB
//}
//
//func (s *ListAppsRepository) ForPhoto(photoId string, currentPage int64, itemsPerPage int64) (*Paginator, error) {
//	query := s.queryBuilder.Select().
//		From("apps").
//		Join("photo_app ON photo_app.app_id=apps.id").
//		Where(squirrel.Eq{
//			"photo_app.photo_id": photoId,
//		})
//	totalCount, err := TotalCountForQuery(s.dbx, query)
//	query = query.Column("apps.*")
//	if err != nil {
//		return nil, err
//	}
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
//		p := models.App{}
//		err = rows.StructScan(&p)
//		if err != nil {
//			continue
//		}
//		results = append(results, p)
//	}
//	// @todo env var for base url
//	baseURL := "http://localhost:1323/photos/" + photoId + "/apps"
//	return NewPaginator(baseURL, totalCount, itemsPerPage, currentPage, results)
//}
