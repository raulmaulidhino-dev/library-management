package dto

import (
	"base-gin/app/domain/dao"
)

type BookReq struct {
	ID           uint      `json:"-"`
	Title         string    `json:"title" binding:"required,min=6,max=56"`
	Subtitle     string    `json:"subtitle" binding:"omitempty,min=4,max=64"`
	AuthorID      uint      `json:"author_id" binding:"required"`
	PublisherID   uint      `json:"publisher_id" binding:"required"`
}

func (o *BookReq) ToEntity() dao.Book {
	return dao.Book{
        Title:        o.Title,
        Subtitle:     o.Subtitle,
        AuthorID:     o.AuthorID,
        PublisherID:   o.PublisherID,
    }
}

type BookResp struct {
	ID       int    `json:"id"`
	Title     string `json:"title"`
	Subtitle string `json:"subtitle"`
	AuthorName string `json:"author_name"`
	PublisherName string `json:"publisher_name"`
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = int(item.ID)
	o.Title = item.Title
	o.Subtitle = item.Subtitle
	o.AuthorName = item.Author.FullName
	o.PublisherName = item.Publisher.Name
}