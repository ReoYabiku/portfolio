package repository

import "portfolio/app/entity"

type Mysql interface {
	GetBiobraphies() ([]entity.Biography, error)
	GetAchievements() ([]entity.Achievement, error)
}
