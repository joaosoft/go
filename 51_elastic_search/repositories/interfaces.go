package repositories

import "github.com/joaosoft/golang-learn/51_elastic_search/domain"

type IRepository interface {
	CreateIndex(index string, mapping map[string]interface{}) error
	DeleteIndex(index string) error
	Insert(data []domain.Something) error
}
