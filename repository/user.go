package repository

import (
	"context"

	"github.com/racsoJosu/rest-ws/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error 
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository){
	implementation = repository

}

func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user);
}

func GetUserById(ctx context.Context, id string) (*models.User, error ){
	return implementation.GetUserById(ctx, id);
}
func GetUserByEmail(ctx context.Context, email string) (*models.User, error ){
	return implementation.GetUserByEmail(ctx, email);
}
func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post);
}

func Close() error{
	return implementation.Close()
}