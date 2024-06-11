package postgres

import (
	"context"
	"testing"
	"user/api/models"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	userRepo := NewUser(db)

	reqUser := models.AddUser{
		FullName: faker.Name(),
		NickName: faker.Name(),
		Photo:    faker.URL(),
		Birthday: "2000-10-10",
		Location: models.Tolocation{
			Latitude:  faker.Latitude(),
			Longitude: faker.Longitude(),
		},

		CreatedBy: faker.Name(),
	}

	id, err := userRepo.Create(context.Background(), reqUser)
	if assert.NoError(t, err) {
		createdUser, err := userRepo.GetById(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqUser.FullName, createdUser.FullName)
			assert.Equal(t, reqUser.NickName, createdUser.NickName)
			assert.Equal(t, reqUser.Photo, createdUser.Photo)
			assert.Equal(t, reqUser.Birthday, createdUser.Birthday)
			assert.Equal(t, reqUser.Location, createdUser.Location)
			assert.Equal(t, reqUser.CreatedBy, createdUser.CreatedBy)

		} else {
			return
		}
	}
}


func TestCreateMany(t *testing.T) {
	userRepo := NewUser(db)
	resp, err := userRepo.GetList(context.Background(), models.GetListRequest{})
	var length int64

	if assert.NoError(t, err) {
		length = resp.Count
	}

	reqUsers := models.AddUsers{
        Users: []models.AddUser{
            {
                FullName: faker.Name(),
                NickName: faker.Name(),
                Photo:    faker.URL(),
                Birthday: "2001-10-10",
                Location: models.Tolocation{
                    Latitude:  faker.Latitude(),
                    Longitude: faker.Longitude(),
                },
                CreatedBy: faker.Name(),
            },
            {
                FullName: faker.Name(),
                NickName: faker.Name(),
                Photo:    faker.URL(),
                Birthday: "2002-10-10",
                Location: models.Tolocation{
                    Latitude:  faker.Latitude(),
                    Longitude: faker.Longitude(),
                },
                CreatedBy: faker.Name(),
            },
        },
    }

	err = userRepo.CreateMany(context.Background(), reqUsers)
	if assert.NoError(t, err) {
		resp, err := userRepo.GetList(context.Background(), models.GetListRequest{})
		if assert.NoError(t, err) {
			assert.Equal(t, length + 2, resp.Count)	
		} else {
			return
		}
	} 
}

func TestUpdate(t *testing.T) {
	userRepo := NewUser(db)

	reqUser := models.UpdateUser{
		Id:         20,
		FullName:  faker.Name(),
		NickName:   faker.Name(),
		Photo:      faker.Word(),
		Birthday: "2000-10-12",
		Location:      models.Tolocation{
			Latitude:  faker.Latitude(),
			Longitude: faker.Longitude(),
		},
		UpdatedBy:      faker.Name(),
	}

	id, err := userRepo.Update(context.Background(), reqUser)
	if assert.NoError(t, err) {
		updatedUser, err := userRepo.GetById(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqUser.FullName, updatedUser.FullName)
			assert.Equal(t, reqUser.NickName, updatedUser.NickName)
			assert.Equal(t, reqUser.Photo, updatedUser.Photo)
			assert.Equal(t, reqUser.Birthday, updatedUser.Birthday)
			assert.Equal(t, reqUser.Location, updatedUser.Location)
			assert.Equal(t, reqUser.UpdatedBy, updatedUser.UpdatedBy)
		} else {
			return
		}
	}
}


func TestUpdateMany(t *testing.T) {
	userRepo := NewUser(db)

	reqUser := models.UpdateUsers{
		Users: []models.UpdateUser{
			{
				Id:         20,
				FullName:  faker.Name(),
				NickName:   faker.Name(),
				Photo:      faker.Word(),
				Birthday: "2000-11-12",
				Location:      models.Tolocation{
					Latitude:  faker.Latitude(),
					Longitude: faker.Longitude(),
				},
				UpdatedBy:      faker.Name(),
			},
			{
				Id:         21,
				FullName:  faker.Name(),
				NickName:   faker.Name(),
				Photo:      faker.Word(),
				Birthday: "2000-10-12",
				Location:      models.Tolocation{
					Latitude:  faker.Latitude(),
					Longitude: faker.Longitude(),
				},
				UpdatedBy:      faker.Name(),
			},
		},
	}

	err := userRepo.UpdateMany(context.Background(), reqUser)
	if assert.NoError(t, err) {
		return
	}
}

func TestDelete(t *testing.T) {
	userRepo := NewUser(db)

	reqUser := models.AddUser{
		FullName: faker.Name(),
		NickName: faker.Name(),
		Photo:  faker.Name(),
		Birthday: "2001-01-12",
		Location: models.Tolocation{
			Latitude:  faker.Latitude(),
			Longitude: faker.Longitude(),
		},
		CreatedBy:      faker.Name(),
	}

	id, err := userRepo.Create(context.Background(), reqUser)
	if assert.NoError(t, err) {
		err := userRepo.Delete(context.Background(), models.DeleteUser{
			Id: id,
			DeletedBy: "me",
		})
		if assert.NoError(t, err) {
			return
		}
	}
}


func TestGetList(t *testing.T) {
	userRepo := NewUser(db)

	reqUser := models.AddUser{
		FullName: faker.Name(),
		NickName: faker.Name(),
		Photo:    faker.URL(),
		Birthday: "2000-10-10",
		Location: models.Tolocation{
			Latitude:  faker.Latitude(),
			Longitude: faker.Longitude(),
		},

		CreatedBy: faker.Name(),
	}

	response, err := userRepo.GetList(context.Background(), models.GetListRequest{
		Page: 1,
		Limit: 10,
	})
	if assert.NoError(t, err) {
		count := response.Count

		_, err := userRepo.Create(context.Background(), reqUser)

		if assert.NoError(t, err) {
			testResponse, err := userRepo.GetList(context.Background(), models.GetListRequest{
				Page: 1,
				Limit: 10,
			})
			if assert.NoError(t, err) {
				testCount := testResponse.Count
				assert.Equal(t, count+1, testCount)
			} else {
				return
			}
		}
	}
}
