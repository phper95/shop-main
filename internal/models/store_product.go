package models

import "shop/pkg/global"

type StoreProduct struct {
	Image        string         `json:"image" valid:"Required;"`
	SliderImage  string         `json:"slider_image" valid:"Required;"`
	StoreName    string         `json:"store_name" valid:"Required;"`
	StoreInfo    string         `json:"store_info" valid:"Required;"`
	Keyword      string         `json:"keyword" valid:"Required;"`
	CateId       int            `json:"cate_id" valid:"Required;"`
	ProductCate  *StoreCategory `json:"product_cate" gorm:"foreignKey:CateId;association_autoupdate:false;association_autocreate:false"`
	Price        float64        `json:"price" valid:"Required;"`
	VipPrice     float64        `json:"vip_price" valid:"Required;"`
	OtPrice      float64        `json:"ot_price" valid:"Required;"`
	Postage      float64        `json:"postage" valid:"Required;"`
	UnitName     string         `json:"unit_name" valid:"Required;"`
	Sort         int16          `json:"sort" valid:"Required;"`
	Sales        int            `json:"sales" valid:"Required;"`
	Stock        int            `json:"stock" valid:"Required;"`
	IsShow       *int8          `json:"is_show" valid:"Required;"`
	IsHot        *int8          `json:"is_hot" valid:"Required;"`
	IsBenefit    *int8          `json:"is_benefit" valid:"Required;"`
	IsBest       *int8          `json:"is_best" valid:"Required;"`
	IsNew        *int8          `json:"is_new" valid:"Required;"`
	Description  string         `json:"description" valid:"Required;"`
	IsPostage    *int8          `json:"is_postage" valid:"Required;"`
	GiveIntegral int            `json:"give_integral" valid:"Required;"`
	Cost         float64        `json:"cost" valid:"Required;"`
	IsGood       *int8          `json:"is_good" valid:"Required;"`
	Ficti        int            `json:"ficti" valid:"Required;"`
	Browse       int            `json:"browse" valid:"Required;"`
	IsSub        *int8          `json:"is_sub" valid:"Required;"`
	TempId       int64          `json:"temp_id" valid:"Required;"`
	SpecType     int8           `json:"spec_type" valid:"Required;"`
	IsIntegral   *int8          `json:"isIntegral" valid:"Required;"`
	Integral     int32          `json:"integral" valid:"Required;"`
	BaseModel
}

//定义商品消息结构
type ProductMsg struct {
	Operation string `json:"operation"`
	*StoreProduct
}

func (StoreProduct) TableName() string {
	return "store_product"
}

func GetProduct(id int64) StoreProduct {
	var product StoreProduct
	global.Db.Where("id =  ?", id).First(&product)

	return product
}

// get all
func GetFrontAllProduct(pageNUm int, pageSize int, maps interface{}, order string) (int64, []StoreProduct) {
	var (
		total int64
		data  []StoreProduct
	)

	global.Db.Model(&StoreProduct{}).Where(maps).Count(&total)
	if order == "" {
		order = "id desc"
	}
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Order(order).Find(&data)

	return total, data
}

// get all
func GetAllProduct(pageNUm int, pageSize int, maps interface{}) (int64, []StoreProduct) {
	var (
		total int64
		data  []StoreProduct
	)

	global.Db.Model(&StoreProduct{}).Where(maps).Count(&total)
	global.Db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("ProductCate").Order("id desc").Find(&data)

	return total, data
}

func GetProductByIDs(maps interface{}) []StoreProduct {
	var (
		data []StoreProduct
	)
	global.Db.Where(maps).Preload("ProductCate").Find(&data)
	return data
}

func AddProduct(m *StoreProduct) error {
	var err error
	if err = global.Db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByProduct(id int64, m *StoreProduct) error {
	var err error
	err = global.Db.Model(&StoreProduct{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func OnSaleByProduct(id int64, status int) (err error) {
	err = global.Db.Model(&StoreProduct{}).Where("id = ?", id).Update("is_show", status).Error
	return
}

func DelByProduct(ids []int64) error {
	var err error
	err = global.Db.Where("id in (?)", ids).Delete(&StoreProduct{}).Error
	if err != nil {
		return err
	}

	return err
}
