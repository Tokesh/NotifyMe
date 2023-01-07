package repositories

import (
	"context"
	"fmt"
	"project/source/domain/entity"
	"project/source/infrastructure/postgresql"
)

type Repository struct {
	client postgresql.Client
	//logger *logging.Logger
}

func New(client postgresql.Client) Repository {
	return Repository{
		client: client,
		//logger add TODO
	}
}

func (r *Repository) CreateUser(user entity.User) error {
	// username, user_email, user_password, user_activation_status,status
	q := `
		INSERT INTO users(username, user_email, user_password,
		                          user_activation_status, status) 
		values($1,$2,$3,$4,$5)
	`

	_, err := r.client.Query(context.TODO(), q, user.Username, user.UserEmail, user.Password, user.ActivationStatus, user.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindUserID(user entity.User) entity.User {
	q := `
		SELECT user_id, username from users where username = $1
	`
	err := r.client.QueryRow(context.TODO(), q, user.Username).Scan(&user.UserID, &user.Username)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user
	}
	return user
}

func (r *Repository) FindUserByID(user_id int) (entity.User, error) {
	q := `
		SELECT user_id, username from users where user_id = $1
	`
	user := entity.User{}
	err := r.client.QueryRow(context.TODO(), q, user_id).Scan(&user.UserID, &user.Username)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user, err
	}
	return user, nil
}

func (r *Repository) FindUserPasswordRepo(user entity.User) entity.User {
	q := `
		SELECT user_password from users where username = $1
	`
	//fmt.Println(r.client.QueryRow(context.TODO(), q, user.Username).Scan(user.Password))
	err := r.client.QueryRow(context.TODO(), q, user.Username).Scan(&user.Password)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user
	}
	return user
}
