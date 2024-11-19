package dto

import "base-gin/app/domain/dao"

type PublisherCreateReq struct {
	ID   uint   `json:"-"`
	Name string `json:"name" binding"required,min=6,max=48"`
	City string `json:"city" binding"required,min=2,max=32"`
}

func (o PublisherCreateReq) ToEntity() dao.Publisher {
	return dao.Publisher {
		Name: o.Name,
		City: o.City,
	}
}

type PublisherCreateResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (v *PublisherCreateResp) FromEntity(item *dao.Publisher) {
	v.ID = int(item.ID)
    v.Name = item.Name
    v.City = item.City
}