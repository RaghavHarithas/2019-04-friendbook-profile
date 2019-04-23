package profile

import (
	"log"

	db "github.com/rajch/2019-04-friendbook-profile/database/connection"
)

const (
	GET_PROFILE_BY_EMAIL = `SELECT * FROM profile WHERE email=?`
	INSERT_PROFILE       = `INSERT INTO profile (email, username, firstname, lastname, gender, date_of_birth, mobile_number, city) VALUES (:email, :username, :firstname, :lastname,:gener,:date_of_birth,:mobile_number,:city)`
	UPDATE_PROFILE       = `UPDATE profile SET username=:username, firstname=:firstname, lastname=:lastname, gender=:gender, date_of_birth=:date_of_birth, mobile_number=:mobile_number, city=:city`
)

type Profile struct {
	Email        string  `db:"email"`
	UserName     *string `db:"username,omitempty"`
	FirstName    *string `db:"firstname,omitempty"`
	LastName     *string `db:"lastname,omitempty"`
	Gender       *string `db:"gender,omitempty"`
	DateOfBirth  *string `db:"date_of_birth,omitempty"`
	MobileNumber *string `db:"mobile_number, omitempty"`
	City         *string `db:"city,omitempty"`
	Visibility   *string `db:"visibility,omitempty"`
	Status       *string `db:"status,omitempty"`
}

func GetProfileByEmail(email string) (*Profile, error) {
	profile := Profile{}
	err := db.Pool.Get(&profile, GET_PROFILE_BY_EMAIL, email)
	if err != nil {
		log.Printf("database error: %s", err)
		return nil, err
	}

	return &profile, nil
}

func CreateOrUpdateProfile(profile *Profile) error {
	_, err := GetProfileByEmail(profile.Email)
	if err != nil {
		_, err := db.Pool.NamedExec(INSERT_PROFILE, profile)
		return err
	}
	_, err = db.Pool.NamedExec(UPDATE_PROFILE, profile)

	return err
}
