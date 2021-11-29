package items

import (
	"rulzmotoshop/business"
)

type serviceItem struct {
	itemRepository Repository
}

func NewServiceItem(repoItem Repository) Service {
	return &serviceItem{
		itemRepository: repoItem,
	}
}

func (serv *serviceItem) AllItem() ([]Domain, error) {

	result, err := serv.itemRepository.AllItem()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceItem) Create(sellID int, domain *Domain) (Domain, error) {

	result, err := serv.itemRepository.Create(sellID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceItem) Update(sellID int, itID int, domain *Domain) (Domain, error) {

	result, err := serv.itemRepository.Update(sellID, itID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceItem) Delete(sellID int, id int) (string, error) {

	result, err := serv.itemRepository.Delete(sellID, id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}

func (serv *serviceItem) MyItemBySeller(sellID int) ([]Domain, error) {

	result, err := serv.itemRepository.MyItemBySeller(sellID)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceItem) ItemByID(id int) (Domain, error) {

	result, err := serv.itemRepository.ItemByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceItem) ItemByIdSeller(orgzID int) ([]Domain, error) {

	result, err := serv.itemRepository.ItemByIdSeller(orgzID)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
