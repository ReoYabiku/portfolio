package infra

import (
	"database/sql"
	"portfolio/app/entity"
	"portfolio/app/repository"
)

type mysql struct {
	db *sql.DB
}

var _ repository.Mysql = &mysql{}

func NewMysql(db *sql.DB) repository.Mysql {
	return &mysql{
		db: db,
	}
}

func (r *mysql) GetBiobraphies() ([]entity.Biography, error) {
	rows, err := r.db.Query("SELECT * FROM biographies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bios := []entity.Biography{}

	for rows.Next() {
		b := &entity.Biography{}
		if err = rows.Scan(&b.ID, &b.Src, &b.Alt, &b.Summary, &b.Title, &b.Description); err != nil {
			return nil, err
		}
		bios = append(bios, *b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bios, err
}

func (r *mysql) GetAchievements() ([]entity.Achievement, error) {
	return nil, nil
}
