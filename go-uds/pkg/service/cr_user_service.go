package service

import (
	"math"
	"myapp/api"
	"myapp/dto"
	"myapp/pkg/config"
	"myapp/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	AuthenticateUser(identity, password string) (string, error)
	UnAuthenticateUser(user interface{}) error

	GetUsers(page, limit int) (currentPage, lastPage, totalRow int, res []dto.CrUser, err error)
	GetUserById(id int) (*dto.CrUser, error)
	CreateUser(req *dto.CrUser, createdBy string) error
	UpdateUser(req *dto.CrUser, updatedBy string) error
	DeleteUser(id int, deletedBy string) error

	checkPasswordHash(password, hash string) bool
	hashPassword(password string) (string, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (r UserServiceImpl) AuthenticateUser(identity, password string) (string, error) {
	user, err := r.repo.GetUserByEmail(identity)
	if err != nil {
		return "", err
	}

	if user.Status != api.STATUS_ACTIVE {
		return "", err
	}

	if !r.checkPasswordHash(password, user.Password) {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	var permissions []string
	for _, item := range user.Role.RolePermissions {
		permissions = append(permissions, item.Permission.Name)
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["status"] = user.Status

	claims["role"] = map[string]interface{}{
		"id":   user.Role.ID,
		"name": user.Role.Name,
	}

	claims["permissions"] = permissions
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	t, err := token.SignedString([]byte(viper.GetString(config.JwtSecret)))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (r UserServiceImpl) UnAuthenticateUser(user interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r UserServiceImpl) GetUsers(page, limit int) (int, int, int, []dto.CrUser, error) {
	result, totalRow, err := r.repo.GetAllUsers(page, limit)
	if err != nil {
		return 0, 0, 0, nil, err
	}

	lastPage := int(math.Ceil(float64(totalRow) / float64(limit)))

	return page, lastPage, totalRow, result, nil
}

func (r UserServiceImpl) GetUserById(id int) (*dto.CrUser, error) {
	return r.repo.GetUserByID(id)
}

func (r UserServiceImpl) CreateUser(req *dto.CrUser, createdBy string) error {
	if req.Password != "" {
		req.Password, _ = r.hashPassword(req.Password)
	}

	if _, err := r.repo.CreateUser(req, createdBy); err != nil {
		return err
	}

	return nil
}

func (r UserServiceImpl) UpdateUser(req *dto.CrUser, updatedBy string) error {
	if req.Password != "" {
		req.Password, _ = r.hashPassword(req.Password)
	}

	return r.repo.UpdateUser(req, updatedBy)
}

func (r UserServiceImpl) DeleteUser(id int, deletedBy string) error {
	return r.repo.DeleteUser(id, deletedBy)
}

func (r UserServiceImpl) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r UserServiceImpl) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
