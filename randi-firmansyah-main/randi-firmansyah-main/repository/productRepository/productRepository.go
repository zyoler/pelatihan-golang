package productRepository

import (
	"randi_firmansyah/connections/database"
	"randi_firmansyah/models/productModel"
)

var DB = database.Connected

func FindAll() ([]productModel.Product, error) {
	var products []productModel.Product
	err := DB.Find(&products).Error
	return products, err
}

func FindByID(ID int) (productModel.Product, error) {
	var product productModel.Product
	err := DB.First(&product, ID).Error
	return product, err
}

func Create(product productModel.Product) (productModel.Product, error) {
	err := DB.Create(&product).Error
	return product, err
}

func UpdateV2(product productModel.Product) (productModel.Product, error) {
	err := DB.Save(&product).Error
	return product, err
}

func Update(id int, product productModel.Product) (productModel.Product, error) {
	err := DB.Model(productModel.Product{}).Where("id = ?", id).Updates(product).Error
	return product, err
}

func Delete(product productModel.Product) (productModel.Product, error) {
	err := DB.Delete(&product).Error
	return product, err
}
