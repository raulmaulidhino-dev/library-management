package dto

import "base-gin/app/domain/dao"

type PublishersCreateReq struct {
	Name string `json:"name" binding"required,min=6,max=48"`
	City string `json:"city" binding"required,min=2,max=32"`
}

func (o PublishersCreateReq) ToEntity() dao.Publishers {
	return dao.Publishers {
		Name: o.Name,
		City: o.City,
	}
}

type PublishersCreateResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (v *PublishersCreateResp) FromEntity(item *dao.Publishers) {
	v.ID = int(item.ID)
    v.Name = item.Name
    v.City = item.City
}