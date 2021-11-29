package users

import (
	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business"
	"rulzmotoshop/helpers/encrypt"
	"time"
)

type serviceUser struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewServiceUser(repoUser Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &serviceUser{
		userRepository: repoUser,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (serv *serviceUser) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.userRepository.Register(domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceUser) Login(email, password string) (Domain, error) {

	result, err := serv.userRepository.Login(email, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "user")

	return result, nil
}
