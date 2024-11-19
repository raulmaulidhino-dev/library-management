package dto

import (
	"base-gin/app/domain/dao"
	"base-gin/app/domain"
	"time"
)

type AuthorCreateReq struct {
	ID uint `json:"-`
	FullName string `json:"fullname" binding"required,min=3,max=56`
	Gender string `json:"gender" binding"required,oneof=m f"`
	BirthDate string `json:"birth_date" binding:"required,datetime=2006-01-02"`
}

func (o *AuthorCreateReq) GetGender() domain.TypeGender {
	if o.Gender == "f" {
		return domain.GenderFemale
	}

	return domain.GenderMale
}

func (o *AuthorCreateReq) GetBirthDate() (time.Time, error) {
	return time.Parse("2006-01-02", o.BirthDate)
}

func (o *AuthorCreateReq) ToEntity() dao.Author {
	var gender string = o.Gender
	typeGender := domain.TypeGender(gender)
	genderDomain := &typeGender

	birthDate, _ := o.GetBirthDate()

	return dao.Author {
		FullName: o.FullName,
		Gender: genderDomain,
        BirthDate: &birthDate,
	}
}

type AuthorCreateResp struct {
	ID int `json:"id"`
	FullName string `json:"fullname"`
	Gender string `json:"gender"`
}

func (o *AuthorCreateResp) FromEntity(item *dao.Author) {
	var gender string
	if item.Gender == nil {
		gender = "-"
	} else if *item.Gender == domain.GenderFemale {
		gender = "f"
	} else {
		gender = "m"
	}

	o.FullName = item.FullName
	o.Gender = gender
	o.ID = int(item.ID)
}