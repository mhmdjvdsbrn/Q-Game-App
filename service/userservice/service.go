package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"q-game-app/entity"
	"q-game-app/pkg/phonenumber"
	"regexp"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	auth AuthGenerator
	repo Repository
}

// New now requires AuthGenerator
func New(repo Repository, auth AuthGenerator) *Service {
	return &Service{
		auth: auth,
		repo: repo,
	}
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User struct {
		ID          uint   `json:"id"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	} `json:"user"`
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, errors.New("invalid phone number")
	}
	//check uniq phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number %s is not unique", req.PhoneNumber)
		}
	}
	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is too short")
	}

	//TODO -> check password with regex
	//validate passsword
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password is too short, minimum is 8")
	}
	if err := isValidPassword(req.Password); err != nil {
		return RegisterResponse{}, err
	}
	//create user to db
	createdUser, err := s.repo.RegisterUser(entity.User{Name: req.Name, PhoneNumber: req.PhoneNumber, Password: getMD5Hash(req.Password)})
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	//return user
	return RegisterResponse{
		User: struct {
			ID          uint   `json:"id"`
			PhoneNumber string `json:"phone_number"`
			Name        string `json:"name"`
		}{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		},
	}, nil
}

var (
	lowerRegex   = regexp.MustCompile(`[a-z]`)
	upperRegex   = regexp.MustCompile(`[A-Z]`)
	digitRegex   = regexp.MustCompile(`\d`)
	specialRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func isValidPassword(password string) error {

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if !lowerRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one lowercase letter (a-z)")
	}

	if !upperRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter (A-Z)")
	}

	if !digitRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one digit (0-9)")
	}

	if !specialRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one special character (!@#$%%^&* etc.)")
	}

	return nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	//check phone number existence
	// get user by phone number
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("phone number or password%s is wrong", req.PhoneNumber)
	}

	//compare user.password with the req.password
	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("usernameeeee or password is wrong")

		if err != nil {
			return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
		}
	}
	//return ok
	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)

	}
	return LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}
