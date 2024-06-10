package postgres

import (
	"context"
	"database/sql"
	"strings"
	"user/api/models"
	"user/pkg"

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
	INSERT INTO users (
		full_name, 
		nick_name, 
		photo, 
		birthday, 
		location, 
		created_by
	) 
	VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		ST_SetSRID(ST_MakePoint($5, $6), 4326), 
		$7) 
	RETURNING id;`

	row := u.db.QueryRow(ctx, query, user.FullName, user.NickName, user.Photo, user.Birthday, user.Location.Latitude, user.Location.Longitude, user.CreatedBy)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepo) CreateMany(ctx context.Context, users models.AddUsers) error {

	query := `
	INSERT INTO users (
		full_name, 
		nick_name,
		photo,
		birthday,
		location,
		created_by
	) 
	VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		ST_SetSRID(ST_MakePoint($5, $6), 4326), 
		$7
	);`

	for i := 0; i < len(users.Users); i++ {
		_, err := u.db.Exec(ctx, query, users.Users[i].FullName, users.Users[i].NickName, users.Users[i].Photo, users.Users[i].Birthday, users.Users[i].Location.Latitude, users.Users[i].Location.Longitude, users.Users[i].CreatedBy)

		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userRepo) Update(ctx context.Context, user models.UpdateUser) (int64, error) {
	var id int64
	row := u.db.QueryRow(ctx, `
	UPDATE 
		users
	SET
		full_name = $2, 
		nick_name = $3, 
		photo = $4, 
		birthday = $5, 
		location = ST_SetSRID(ST_MakePoint($6, $7), 4326),
		updated_at = NOW(), 
		updated_by = $8
	WHERE 
    	id = $1
	RETURNING id;`, user.Id, user.FullName, user.NickName, user.Photo, user.Birthday, user.Location.Latitude, user.Location.Longitude, user.UpdatedBy)

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepo) UpdateMany(ctx context.Context, users models.UpdateUsers) error {
	query := `
	UPDATE 
		users
	SET
		full_name = $2, 
		nick_name = $3, 
		photo = $4, 
		birthday = $5, 
		location = ST_SetSRID(ST_MakePoint($6, $7), 4326),
		updated_at = NOW(), 
		updated_by = $8
	WHERE 
    	id = $1;`

	for i := 0; i < len(users.Users); i++ {
		_, err := u.db.Exec(ctx, query, users.Users[i].Id, users.Users[i].FullName, users.Users[i].NickName, users.Users[i].Photo, users.Users[i].Birthday, users.Users[i].Location.Latitude, users.Users[i].Location.Longitude, users.Users[i].UpdatedBy)

		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, user models.DeleteUser) error {
	_, err := u.db.Exec(ctx, `
	UPDATE 
		users 
	SET
		deleted_at = NOW(),
		deleted_by = $2
	WHERE 
		id = $1;`, user.Id, user.DeletedBy)

	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) DeleteMany(ctx context.Context, users models.DeleteUsers) error {

	query := `
	UPDATE 
		users 
	SET
		deleted_at = NOW(),
		deleted_by = $2
	WHERE id = $1;`

	for i := 0; i < len(users.Users); i++ {
		_, err := u.db.Exec(ctx, query, users.Users[i].Id, users.Users[i].DeletedBy)

		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userRepo) GetById(ctx context.Context, id int64) (models.GetUser, error) {

	var (
		user                                                       models.GetUser
		nickName, photo, birthday, deletedAt, updatedBy, deletedBy sql.NullString
	)

	row := u.db.QueryRow(ctx, `
	SELECT
		full_name,
		nick_name,
		photo,
		TO_CHAR(birthday, 'YYYY-MM-DD') AS birthday,
		ST_X(location) AS longitude, 
		ST_Y(location) AS latitude, 
		TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
		TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS') AS deleted_at,
		TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
		created_by,
		updated_by,
		deleted_by
	FROM
		users
	WHERE id = $1;`, id)

	err := row.Scan(
		&user.FullName,
		&nickName,
		&photo,
		&birthday,
		&user.Location.Latitude,
		&user.Location.Longitude,
		&user.CreatedAt,
		&deletedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&updatedBy,
		&deletedBy,
	)

	user.NickName = pkg.NullStringToString(nickName)
	user.Photo = pkg.NullStringToString(photo)
	user.Birthday = pkg.NullStringToString(birthday)
	user.DeletedAt = pkg.NullStringToString(deletedAt)
	user.UpdatedBy = pkg.NullStringToString(updatedBy)
	user.DeletedBy = pkg.NullStringToString(deletedBy)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepo) GetList(ctx context.Context, req models.GetListRequest) (models.GetListResponse, error) {
	resp := models.GetListResponse{}

	orderBy := ""
	offset := (req.Page - 1) * req.Limit

	if len(req.Sort) > 0 {
		orderBy += ` ORDER BY ` + strings.Join(req.Sort, ", ") + ` `
	}
	query := `
	SELECT
		full_name,
		nick_name,
		photo,
		TO_CHAR(birthday, 'YYYY-MM-DD') AS birthday,
		ST_X(location) AS longitude, 
		ST_Y(location) AS latitude, 
		TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
		TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
	FROM
		users
	WHERE deleted_at IS NULL 
	` + orderBy + `
	OFFSET
		$1
	LIMIT
		$2;`

	rows, err := u.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			user                      models.UserList
			nickName, photo, birthday sql.NullString
		)

		err := rows.Scan(
			&user.FullName,
			&nickName,
			&photo,
			&birthday,
			&user.Location.Latitude,
			&user.Location.Longitude,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		user.NickName = pkg.NullStringToString(nickName)
		user.Photo = pkg.NullStringToString(photo)
		user.Birthday = pkg.NullStringToString(birthday)

		resp.Users = append(resp.Users, user)
	}

	defer rows.Close()

	query = `
	WITH ordered_users AS (
		SELECT full_name
		FROM users
		WHERE deleted_at IS NULL
		` + orderBy + `
	)
	SELECT COUNT(*)
	FROM ordered_users;`

	row := u.db.QueryRow(ctx, query)

	err = row.Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
