package user

import (
	"context"
	"stranger-danger/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	temporary_dev_secret_key = "devKey"
)

type service struct {
	Repository
	timeout time.Duration
}

type UserJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	idString := strconv.Itoa(int(user.ID))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserJWTClaims{
		ID:       idString,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    idString,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(temporary_dev_secret_key))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, Username: user.Username, ID: idString}, nil
}
