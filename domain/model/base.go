package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"ID" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdateAt  time.Time `json:"updated_at" valid:"-"`
}
