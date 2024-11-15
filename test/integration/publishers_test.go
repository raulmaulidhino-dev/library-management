package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Success(t *testing.T) {
	req := dto.PublishersCreateReq{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}

	w := doTest("POST", server.RootPublishers, req, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 201, w.Code)

	var resp dto.SuccessResponse[dto.PublishersCreateResp]
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if assert.Nil(t, err) {
		data := resp.Data
		assert.Greater(t, data.ID, 0)
		assert.Equal(t, req.Name, data.Name)
		assert.Equal(t, req.City, data.City)

		item, _ := publishersRepo.GetByID(uint(data.ID))
		if assert.NotNil(t, item) {
			assert.Equal(t, req.Name, item.Name)
            assert.Equal(t, req.City, item.City)
		}
	}
}

func TestPublishers_Create_ErrorUnauthorized(t *testing.T) {
	req := dto.PublishersCreateReq{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}

	w := doTest("POST", server.RootPublishers, req, "asdf")
	assert.Equal(t, 401, w.Code)
}