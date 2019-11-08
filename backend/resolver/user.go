package resolver

import (
	"context"
	"errors"
	"strings"
	"time"

	"../models"
)

func (db *DBStruct) register(ctx context.Context, email string, name string, password string, contact string) (*models.UserDetail, error) {

	var user models.UserDetail

	// tempDb := models.GetDB()

	if !db.Db.Where(&models.UserDetail{Email: strings.TrimSpace(email)}).Find(&models.UserDetail{}).RecordNotFound() {
		err := errors.New("Email already exists")
		return nil, err
	}

	user = models.UserDetail{
		Email:         email,
		FullName:      name,
		Password:      HASH(password),
		ContactNumber: contact,
	}
	err := db.Db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DBStruct) login(ctx context.Context, email string, password string) (*models.UserDetail, error) {
	var user models.UserDetail

	// tempDb := models.GetDB()
	dbResponse := db.Db.Model(&models.UserDetail{}).Where(&models.UserDetail{Email: strings.TrimSpace(email)}).First(&user)

	if dbResponse.RecordNotFound() {
		err := errors.New("User not registered")
		return nil, err
	}

	if user.Password != HASH(password) {
		err := errors.New("Invalid Credentials")
		return nil, err
	}

	user.UserToken = tokenGenerator(password)

	err := db.Db.Model(&user).Updates(models.UserDetail{UserToken: user.UserToken}).Error

	if err != nil {
		err := errors.New("Unable to login Please Try Again")
		return nil, err
	}

	return &user, nil
}

type UserResolver struct {
	db   *DBStruct
	user models.UserDetail
}

func (u *UserResolver) UserId(ctx context.Context) *int32 {
	return &u.user.UserId
}

func (u *UserResolver) Email(ctx context.Context) *string {
	return &u.user.Email
}
func (u *UserResolver) FullName(ctx context.Context) *string {
	return &u.user.FullName
}
func (u *UserResolver) UserToken(ctx context.Context) *string {
	return &u.user.UserToken
}
func (u *UserResolver) Password(ctx context.Context) *string {
	pwd := "****"
	return &pwd
}
func (u *UserResolver) ContactNumber(ctx context.Context) *string {
	return &u.user.ContactNumber
}

func (u *UserResolver) CreatedAt(ctx context.Context) *string {
	timeString := u.user.CreatedAt.String()
	return &timeString
}

func (u *UserResolver) UpdatedAt(ctx context.Context) *string {
	timeString := u.user.UpdatedAt.String()
	return &timeString
}

func tokenGenerator(password string) string {
	var salt = "**bINgeBoNd**"
	t := time.Now()
	return HASH(password + salt + t.String())
}
