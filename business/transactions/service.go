package transactions

import (
	"rulzmotoshop/business"
	randomcode "rulzmotoshop/helpers/rand"
)

type serviceTrans struct {
	transRepository Repository
}

func NewServiceTrans(repoTrans Repository) Service {
	return &serviceTrans{
		transRepository: repoTrans,
	}
}

func (serv *serviceTrans) Trans(userID int, domain *Domain) (Domain, error) {

	randCode, _ := randomcode.GenerateCode(8)

	domain.Trans_code = randCode

	result, err := serv.transRepository.Trans(userID, domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return result, nil
}

func (serv *serviceTrans) GetTransByID(id int) (Domain, error) {

	result, err := serv.transRepository.GetTransByID(id)

	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return result, nil

}

func (serv *serviceTrans) GetAllTrans() ([]Domain, error) {

	result, err := serv.transRepository.GetAllTrans()

	if err != nil {
		return []Domain{}, err
	}

	return result, err
}

func (serv *serviceTrans) GetAllUserTrans(userID int) ([]Domain, error) {

	result, err := serv.transRepository.GetAllUserTrans(userID)

	if err != nil {
		return []Domain{}, err
	}

	return result, err
}
