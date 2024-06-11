package postgres

import (
	"context"
	"testing"
	"user/api/models"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
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


func TestCreateUserMany(t *testing.T) {
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

// func TestUpdateStudent(t *testing.T) {
// 	studentRepo := NewStudent(db)

// 	reqStudent := models.Student{
// 		Id:         "3eccd9c5-c3ed-460d-9b2e-e10b600c3a0e",
// 		FirstName:  faker.Name(),
// 		LastName:   faker.Word(),
// 		Age:        12,
// 		ExternalId: faker.ID,
// 		Phone:      faker.Phonenumber(),
// 		Email:      faker.Email(),
// 	}

// 	id, err := studentRepo.Update(context.Background(), reqStudent)
// 	if assert.NoError(t, err) {
// 		createdStudent, err := studentRepo.GetStudent(context.Background(), id)
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
// 			assert.Equal(t, reqStudent.Age, createdStudent.Age)
// 			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
// 		} else {
// 			return
// 		}
// 	}
// }

// func TestDeleteStudent(t *testing.T) {
// 	studentRepo := NewStudent(db)

// 	reqStudent := models.AddStudent{
// 		FirstName: faker.Name(),
// 		Age:       10,
// 		LastName:  faker.Word(),
// 	}

// 	id, err := studentRepo.Create(context.Background(), reqStudent)
// 	if assert.NoError(t, err) {
// 		err := studentRepo.Delete(context.Background(), id)
// 		if assert.NoError(t, err) {
// 			return
// 		}
// 	}
// }

// func TestGetAllStudent(t *testing.T) {
// 	studentRepo := NewStudent(db)
// 	reqStudent := models.AddStudent{
// 		FirstName: faker.Name(),
// 		Age:       10,
// 		LastName:  faker.Word(),
// 	}

// 	response, err := studentRepo.GetAll(context.Background(), models.GetAllStudentsRequest{})
// 	if assert.NoError(t, err) {
// 		count := response.Count

// 		_, err := studentRepo.Create(context.Background(), reqStudent)

// 		if assert.NoError(t, err) {
// 			testResponse, err := studentRepo.GetAll(context.Background(), models.GetAllStudentsRequest{})
// 			if assert.NoError(t, err) {
// 				testCount := testResponse.Count
// 				assert.Equal(t, count+1, testCount)
// 			} else {
// 				return
// 			}
// 		}
// 	}
// }
