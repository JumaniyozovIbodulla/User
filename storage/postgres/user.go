package postgres

import (
	"context"
	"user/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) userRepo {
	return userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, user models.AddUser) (int64, error) {

	var id int64

	query := `
	INSERT INTO
		users (full_name, nick_name, photo, birthday, location, created_by) 
	VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7) RETURNING id;`

	row := u.db.QueryRow(ctx, query, user.FullName, user.NickName, user.Photo, user.Birthday, user.Location.Latitude, user.Location.Longitude, user.CreatedBy)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepo) CreateUsers(ctx context.Context, users models.AddUsers) error {

	query := `INSERT INTO
					users (full_name, nick_name, photo, birthday, location, created_by) VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6), 4326), $7);`

	for i := 0; i < len(users.Users); i++ {
		_, err := u.db.Exec(ctx, query, users.Users[i].FullName, users.Users[i].NickName, users.Users[i].Photo, users.Users[i].Birthday, users.Users[i].Location.Latitude, users.Users[i].Location.Longitude, users.Users[i].CreatedBy)

		if err != nil {
			return err
		}
	}
	return nil
}
