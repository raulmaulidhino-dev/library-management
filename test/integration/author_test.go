package integration_test

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Author_Success(t *testing.T) {
	req := dto.AuthorCreateReq{
		FullName: util.RandomStringAlpha(8),
		Gender: util.RandomGender(),
		BirthDate: util.RandomBirthDate().Format("2006-01-02"),
	}

	w := doTest("POST", server.RootAuthor, req, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 201, w.Code)

	var resp dto.SuccessResponse[dto.AuthorCreateResp]
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if assert.Nil(t, err) {
		data := resp.Data
		assert.Greater(t, data.ID, 0)
		assert.Equal(t, req.FullName, data.FullName)
		assert.Equal(t, req.Gender, data.Gender)

		item, _ := authorRepo.GetByID(uint(data.ID))
		if assert.NotNil(t, item) {
			assert.Equal(t, req.FullName, item.FullName)

			var gender string
			if *item.Gender ==  domain.GenderMale {
				gender = "m"
			} else {
				gender = "f"
			}

			assert.Equal(t, req.Gender, gender)
		}
	}
}