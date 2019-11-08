package resolver

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Add the resolvers for your queries and mutations here
// Also create a new file in this package to create resolvers of individual types
type DBStruct struct {
	Db *gorm.DB
}

type Resolver struct {
	DBStruct *DBStruct
}

type userRegisterInput struct {
	Email         string
	FullName      string
	Password      string
	ContactNumber string
}

func (r *Resolver) Register(ctx context.Context, args userRegisterInput) (*UserResolver, error) {

	user, err := r.DBStruct.register(ctx, args.Email, args.FullName, args.Password, args.ContactNumber)
	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db:   r.DBStruct,
		user: *user,
	}

	return &s, nil
}

type userLoginInput struct {
	Email    string
	Password string
}

func (r *Resolver) Login(ctx context.Context, args userLoginInput) (*UserResolver, error) {

	user, err := r.DBStruct.login(ctx, args.Email, args.Password)

	if err != nil {
		return nil, err
	}

	s := UserResolver{
		db:   r.DBStruct,
		user: *user,
	}

	return &s, nil
}
