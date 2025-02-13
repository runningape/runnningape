package category

import "github.com/runningape/goblog/app/models"

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}
