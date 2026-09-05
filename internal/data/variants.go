//variants.go defines the Variant struct and operations (Insert, GetByImageID) for the three generated outputs a worker will produce from each image

package data

import (
	"context"
	"database/sql"
	"time"
)

type Variant struct {
	ID             string    `json:"-"`
	ImageID        string    `json:"-"`
	Name           string    `json:"name"`
	StoredFilename string    `json:"-"`
	Width          string    `json:"width"`
	Height         string    `json:"height"`
	SizeBytes      int64     `json:"-"`
	URL            string    `json:"url,omitempty"`
	CreatedAt      time.Time `json:"-"`
}

type VariantModel struct {
	DB *sql.DB
}

func (m VariantModel) Insert(v *Variant) error {
	query := `
			INSERT INTO variants (image_id, name, stored_filename, width, height, size_bytes) VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, v.ImageID, v.Name, v.StoredFilename, v.Width, v.Height, v.SizeBytes).Scan(&v.ID, &v.CreatedAt)
}

func (m VariantModel) GetByImageID(imageID string) ([]Variant, error) {
	query := `
			SELECT id, image_id, name, stored_filename, width, height, size_bytes, created_at
			FROM variants WHERE image_id = $1 ORDER BY name`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, imageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []Variant
	for rows.Next() {
		var v Variant
		if err := rows.Scan(&v.ID, &v.ImageID, &v.Name, &v.StoredFilename, &v.Width, &v.Height, &v.CreatedAt); err != nil {
			return nil, err
		}
		variants = append(variants, v)
	}
	return variants, rows.Err()
}
