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
func (serv *serviceSeller) Update(sellID int, domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)
	domain.Password = hashedPassword
	result, err := serv.sellerRepository.Update(sellID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceSeller) SellerByID(id int) (Domain, error) {

	result, err := serv.sellerRepository.SellerByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceSeller) Delete(id int) (string, error) {

	result, err := serv.sellerRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
