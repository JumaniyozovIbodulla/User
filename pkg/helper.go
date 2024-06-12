package pkg

import (
	"database/sql"
	"errors"
	"time"
	"user/api/models"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func ValidateBirthday(date string) error {
	const layout = "2006-01-02"

	parsedDate, err := time.Parse(layout, date)

	if err != nil {
		return err
	}

	if parsedDate.Year() > time.Now().Year() {
		return errors.New("year is not valid")
	}

	return nil
}


func ValidateManyBirthday(users models.AddUsers) error {
	const layout = "2006-01-02"

	for i := 0; i < len(users.Users); i++ {
		parsedDate, err := time.Parse(layout, users.Users[i].Birthday)

		if err != nil {
			return err
		}

		if parsedDate.Year() > time.Now().Year() {
			return errors.New("year is not valid")
		}
	}
	return nil
}

func ValidateUpdateManyBirthday(users models.UpdateUsers) error {
	const layout = "2006-01-02"

	for i := 0; i < len(users.Users); i++ {
		parsedDate, err := time.Parse(layout, users.Users[i].Birthday)

		if err != nil {
			return err
		}

		if parsedDate.Year() > time.Now().Year() {
			return errors.New("year is not valid")
		}
	}
	return nil
}