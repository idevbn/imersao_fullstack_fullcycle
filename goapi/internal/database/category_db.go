package database

import (
	"database/sql"

	entity "github.com/idevbn/imersao-fullstack-fullcycle/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{
		db: db,
	}
}

// GetCategories retorna todas as categorias encontradas no banco de dados
func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

// GetCategory recupera uma categoria específica no banco de dados
func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// CreateCategory insere uma categoria no banco de dados
func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)",
		category.ID,
		category.Name,
	)
	if err != nil {
		return "", err
	}

	return category.ID, nil
}
