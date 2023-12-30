package models

import "mime/multipart"

type InventoryResponse struct {
	ProductID int
	Stock     int
}

type InventoryUpdate struct {
	Productid int `json:"product_id"`
	Stock     int `json:"stock"`
}

type Inventories struct {
	ID              uint    `json:"id"`
	CategoryID      int     `json:"category_id"`
	ProductName     string  `json:"product_name"`
	Size            string  `json:"size"`
	Stock           int     `json:"stock"`
	IfPresentAtCart bool    `json:"if_present_at_cart"`
	Price           float64 `json:"price"`
	Image           string  `json:"product_image"`
}

type AddInventories struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type EditInventoryDetails struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
	Size       string  `json:"size"`
}
type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"products_id"`
}

type InventoryDetails struct {
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}

type Cart struct {
	InventoryID int `json:"products_id"`
	Quantity    int `json:"quantity"`
}

type ImageUp struct {
	InventoryID int `json:"inventory_id"`
	Url         []Url
}
type Url struct {
	Url *multipart.FileHeader `json:"url"`
}
type Image struct {
	Image []string `json:"image"`
}
