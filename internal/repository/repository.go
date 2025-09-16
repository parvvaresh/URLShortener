package repository

import "github.com/jmoiron/sqlx"

type URLRepository struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(code, url string) error {
	_, err := r.db.Exec("INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", code, url)
	return err
}

func (r *URLRepository) Get(code string) (string, error) {
	var url string
	err := r.db.Get(&url, "SELECT original_url FROM urls WHERE short_code = $1", code)
	return url, err
}
