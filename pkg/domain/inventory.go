// package domain

// type Inventories struct {
// 	ID          uint     `json:"id" gorm:"unique;not null"`
// 	CategoryID  int      `json:"category_id"`
// 	Category    Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
// 	ProductName string   `json:"product_name"`
// 	//	Image       string   `json:"image"`
// 	Size  string  `json:"size" gorm:"size:5;default:'M6';check:size IN ('W4', 'M6', 'W4', 'M8', 'W10')"`
// 	Stock int     `json:"stock"`
// 	Price float64 `json:"price"`
// }

// type Category struct {
// 	ID       uint   `json:"id" gorm:"unique;not null"`
// 	Category string `json:"category"`
// }

package domain

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
}

type Inventories struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size" gorm:"size:W4;default:'W6';Check:size IN ('W4', 'W6', 'M6', 'W10', 'M10');"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
}
