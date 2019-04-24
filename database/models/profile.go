package profile

import (
	"log"

	db "github.com/rajch/2019-04-friendbook-profile/database/connection"
)

const (
	GET_PROFILE_BY_EMAIL = `SELECT * FROM profile WHERE email=?`
	INSERT_PROFILE       = `INSERT INTO profile (email, username, firstname, lastname, gender, date_of_birth, mobile_number, city) VALUES (:email, :username, :firstname, :lastname,:gender,:date_of_birth,:mobile_number,:city)`
	UPDATE_PROFILE       = `UPDATE profile SET username=:username, firstname=:firstname, lastname=:lastname, gender=:gender, date_of_birth=:date_of_birth, mobile_number=:mobile_number, city=:city`
	DELETE_PROFILE       = `DELETE FROM profile WHERE email=:email`
)

type Profile struct {
	Email        string  `db:"email" json:"userEmail"`
	UserName     *string `db:"username,omitempty" json:"username,omitempty"`
	FirstName    *string `db:"firstname,omitempty" json:"firstname,omitempty"`
	LastName     *string `db:"lastname,omitempty" json:"lastname,omitempty"`
	Gender       *string `db:"gender,omitempty" json:"gender,omitempty"`
	DateOfBirth  *string `db:"date_of_birth,omitempty" json:"dateOfBirth,omitempty"`
	MobileNumber *string `db:"mobile_number,omitempty" json:"mobilenumber,omitempty"`
	City         *string `db:"city,omitempty" json:"location,omitempty"`
	Visibility   *string `db:"visibility,omitempty" json:"-"`
	Status       *string `db:"status,omitempty" json:"-"`
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
		log.Printf("Database error: %s", err)
		return err
	}
	_, err = db.Pool.NamedExec(UPDATE_PROFILE, profile)
	if err != nil {
		log.Printf("Database error: %s", err)
	}

	return err
}

func DeleteProfile(email string) error {
	profile, err := GetProfileByEmail(email)
	if err != nil {
		log.Printf("Error getting profile in delete profile: %s", err)
		return err
	}

	_, err = db.Pool.NamedExec(DELETE_PROFILE, profile)
	if err != nil {
		log.Printf("Error while deleting: %s", err)
		return err
	}
	return nil
}
