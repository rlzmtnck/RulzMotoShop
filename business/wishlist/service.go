package wishlist

import (
	"rulzmotoshop/business"
)

type serviceTrans struct {
	transRepository Repository
}

func NewServiceWish(repoTrans Repository) Service {
	return &serviceTrans{
		transRepository: repoTrans,
	}
}

func (serv *serviceTrans) Wish(userID int, domain *Domain) (Domain, error) {

	result, err := serv.transRepository.Wish(userID, domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return result, nil
}

func (serv *serviceTrans) GetAllUserWish(userID int) ([]Domain, error) {

	result, err := serv.transRepository.GetAllUserWish(userID)

	if err != nil {
		return []Domain{}, err
	}

	return result, err
}
