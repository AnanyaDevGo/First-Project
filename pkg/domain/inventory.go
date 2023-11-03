package domain


type Inventories struct {
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
	Image    string `json:"category_image"`
}
