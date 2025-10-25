package service

import (
	"context"

	"user-api/internal/models"
	"user-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error)
	GetUser(ctx context.Context, id int32) (*models.UserResponse, error)
	ListUsers(ctx context.Context) ([]*models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.Create(ctx, req.Name, req.DOB)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) ListUsers(ctx context.Context) ([]*models.UserResponse, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*models.UserResponse
	for _, user := range users {
		response := user.ToResponse()
		responses = append(responses, &response)
	}

	return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.Update(ctx, id, req.Name, req.DOB)
	if err != nil {
		return nil, err
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.userRepo.Delete(ctx, id)
}
