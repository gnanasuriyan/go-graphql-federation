package services

import (
	"context"

	"github.com/gnanasuriyan/go-graphql/user-micro-service/graph/model"
	"github.com/google/wire"
)

type IUserService interface {
	FindUserByID(ctx context.Context, IDUser int) (*model.User, error)
	CreateUser(ctx context.Context, name string) (*model.User, error)
	ListUsers(ctx context.Context) (*model.UserList, error)
	Me(ctx context.Context) (*model.User, error)
}

type UserServiceOptions struct {
	// include user service dependencies
}

type UserService struct {
	*UserServiceOptions
}

var NewUserService = wire.NewSet(wire.Struct(new(UserServiceOptions), "*"), InitUserService, wire.Bind(new(IUserService), new(*UserService)))

var userList []model.User

func InitUserService(options *UserServiceOptions) (*UserService, error) {
	service := &UserService{UserServiceOptions: options}
	// populate some user data
	userList = append(userList, model.User{ID: 1, Username: "User 1"})
	return service, nil
}

func (us *UserService) FindUserByID(ctx context.Context, IDUser int) (*model.User, error) {
	for _, user := range userList {
		if user.ID == IDUser {
			return &user, nil
		}
	}
	return nil, nil
}

func (us *UserService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	user := model.User{
		Username: name,
		ID:       len(userList) + 1,
	}
	userList = append(userList, user)
	return &user, nil
}

func (us *UserService) ListUsers(ctx context.Context) (*model.UserList, error) {
	response := &model.UserList{
		Total: len(userList),
		Data:  []*model.User{},
	}
	for _, user := range userList {
		tempUser := user
		response.Data = append(response.Data, &tempUser)
	}
	return response, nil
}

func (us *UserService) Me(ctx context.Context) (*model.User, error) {
	return us.FindUserByID(ctx, 1)
}
