package sellers

import (
	"rulzmotoshop/app/middleware"
	"rulzmotoshop/business"
	"rulzmotoshop/helpers/encrypt"
	"time"
)

type serviceSeller struct {
	sellerRepository Repository
	contextTimeout   time.Duration
	jwtAuth          *middleware.ConfigJWT
}

func NewServiceSeller(repoSeller Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &serviceSeller{
		sellerRepository: repoSeller,
		contextTimeout:   timeout,
		jwtAuth:          jwtauth,
	}
}

func (serv *serviceSeller) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.sellerRepository.Register(domain)

	if result == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceSeller) Login(email, password string) (Domain, error) {

	result, err := serv.sellerRepository.Login(email, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "seller")

	return result, nil
}
