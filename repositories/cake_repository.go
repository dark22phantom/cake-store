package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/teddy/cake-store/models"
)

type CakeRepository interface {
	GetListOfCake() ([]models.CakeModel, error)
	GetDetailOfCake(ID int) (models.CakeModel, error)
	UpdateCake(ID int, request models.RequestModels) error
	InsertCake(request models.RequestModels) error
	DeleteCake(ID int) error
}

type cakeRepository struct {
	db *sql.DB
}

func NewCakeRepository(db *sql.DB) CakeRepository {
	return &cakeRepository{
		db: db,
	}
}

func (r *cakeRepository) GetListOfCake() ([]models.CakeModel, error) {
	var cakes models.CakeModel
	var data []models.CakeModel
	query := `
		SELECT id, title, description, rating, image, created_at, updated_at
		FROM cakes 
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&cakes.ID,
			&cakes.Title,
			&cakes.Description,
			&cakes.Rating,
			&cakes.Image,
			&cakes.CreatedAt,
			&cakes.UpdatedAt,
		)
		if err != nil {
			log.Println("error to scan data")
			return data, err
		}
		data = append(data, cakes)
	}

	return data, nil
}

func (r *cakeRepository) GetDetailOfCake(ID int) (models.CakeModel, error) {
	var cake models.CakeModel

	query := `
		SELECT id, title, description, rating, image, created_at, updated_at
		FROM cakes
		WHERE id = ? 
	`
	row := r.db.QueryRow(query, ID)
	err := row.Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return cake, err
	}
	return cake, nil
}

func (r *cakeRepository) UpdateCake(ID int, request models.RequestModels) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	query := `
		UPDATE cakes
		SET title = ?,
			description = ?,
			rating = ?,
			image = ?,
			updated_at = ?
		WHERE id = ?
	`
	updData, err := r.db.Prepare(query)
	defer updData.Close()

	_, err = updData.Exec(
		request.Title,
		request.Description,
		request.Rating,
		request.Image,
		now,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *cakeRepository) InsertCake(request models.RequestModels) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	query := `
		INSERT INTO cakes(
			title,
			description,
			rating,
			image,
			created_at,
			updated_at
		)
		VALUES(?,?,?,?,?,?);
	`
	insData, err := r.db.Prepare(query)
	defer insData.Close()

	_, err = insData.Exec(
		request.Title,
		request.Description,
		request.Rating,
		request.Image,
		now,
		now,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *cakeRepository) DeleteCake(ID int) error {
	query := `
		DELETE FROM cakes
		WHERE id = ?
	`
	delData, err := r.db.Prepare(query)
	defer delData.Close()

	res, err := delData.Exec(ID)
	if err != nil {
		return err
	}

	rowsDel, _ := res.RowsAffected()
	if rowsDel == 0 {
		err = errors.New("Data not found")
		return err
	}
	return nil
}
