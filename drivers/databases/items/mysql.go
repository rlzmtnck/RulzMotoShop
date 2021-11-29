package items

import (
	"rulzmotoshop/business"
	"rulzmotoshop/business/items"

	"gorm.io/gorm"
)

type MysqlItemRepository struct {
	Conn *gorm.DB
}

func NewMysqlItemRepository(conn *gorm.DB) items.Repository {
	return &MysqlItemRepository{
		Conn: conn,
	}
}

func (rep *MysqlItemRepository) Create(sellID int, domain *items.Domain) (items.Domain, error) {

	ev := fromDomain(*domain)

	ev.SellerID = sellID

	result := rep.Conn.Create(&ev)

	if result.Error != nil {
		return items.Domain{}, result.Error
	}

	return toDomain(ev), nil

}

func (rep *MysqlItemRepository) AllItem() ([]items.Domain, error) {

	var item []Items

	result := rep.Conn.Find(&item)

	if result.Error != nil {
		return []items.Domain{}, result.Error
	}

	return toDomainList(item), nil

}

func (rep *MysqlItemRepository) Update(sellID int, itID int, domain *items.Domain) (items.Domain, error) {

	itemUpdate := fromDomain(*domain)

	find := rep.Conn.Where("id = ?", itID).First(&itemUpdate).Error

	if find != nil {
		return items.Domain{}, business.ErrNotFound
	}

	itemUpdate.ID = itID

	result := rep.Conn.Where("seller_id = ?", sellID).Where("id = ?", itID).Updates(&itemUpdate)

	if result.Error != nil {
		return items.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(itemUpdate), nil
}

func (rep *MysqlItemRepository) Delete(orgID int, id int) (string, error) {
	rec := Items{}

	find := rep.Conn.Where("organizer_id = ?", orgID).Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Items has been delete", nil

}

func (rep *MysqlItemRepository) MyItemBySeller(orgID int) ([]items.Domain, error) {
	var item []Items

	result := rep.Conn.Where("seller_id = ?", orgID).Find(&item)

	if result.Error != nil {
		return []items.Domain{}, result.Error
	}

	return toDomainList(item), nil
}

func (rep *MysqlItemRepository) ItemByID(id int) (items.Domain, error) {

	var item Items

	result := rep.Conn.Where("id = ?", id).First(&item)

	if result.Error != nil {
		return items.Domain{}, result.Error
	}

	return toDomain(item), nil
}
func (rep *MysqlItemRepository) ItemByIdSeller(sellsID int) ([]items.Domain, error) {

	var item []Items

	result := rep.Conn.Where("seller_id = ?", sellsID).Find(&item)

	if result.Error != nil {
		return []items.Domain{}, result.Error
	}

	return toDomainList(item), nil
}
