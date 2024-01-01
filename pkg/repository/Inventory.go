package repository

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &inventoryRepository{DB}
}

func (i *inventoryRepository) AddInventory(inventory models.AddInventories, url string) (models.ProductsResponse, error) {
	var count int64
	i.DB.Model(&models.Inventories{}).Where("product_name = ? AND category_id = ?", inventory.ProductName, inventory.CategoryID).Count(&count)
	if count > 0 {

		return models.ProductsResponse{}, errors.New("product already exists in the database")
	}

	if inventory.Stock < 0 || inventory.Price < 0 {
		return models.ProductsResponse{}, errors.New("stock and price cannot be negative")
	}

	query := `
        INSERT INTO inventories (category_id, product_name, size, stock, price)
        VALUES (?, ?, ?, ?, ?)
        RETURNING id,category_id,product_name,size,stock,price
    `
	var productsResponse models.ProductsResponse
	err := i.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price).Scan(&productsResponse).Error
	if err != nil {
		return models.ProductsResponse{}, err
	}

	queryImage := "insert into images (url,inventory_id) values(?,?)"
	imgErr := i.DB.Exec(queryImage, url, productsResponse.ID).Error
	if imgErr != nil {
		return models.ProductsResponse{}, imgErr
	}
	images := []string{}

	rows, err := i.DB.Raw("SELECT  i.url FROM images AS i WHERE i.inventory_id = ?", productsResponse.ID).Rows()
	if err != nil {
		return models.ProductsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		if err := i.DB.ScanRows(rows, &image); err != nil {
			return models.ProductsResponse{}, err
		}
		images = append(images, image)
	}

	err = i.DB.Raw("SELECT * FROM inventories WHERE id = ?", productsResponse.ID).Scan(&productsResponse).Error
	if err != nil {
		return models.ProductsResponse{}, err
	}
	productsResponse.Image = images
	return productsResponse, nil
}
func (prod *inventoryRepository) ImageUploader(inventoryID int, url string) error {
	err := prod.DB.Exec("insert into images (inventory_id,url) values (?,?)", inventoryID, url).Error
	if err != nil {
		return errors.New("error on updating data base")
	}
	return nil
}

func (prod *inventoryRepository) ListProducts(pageList, offset int) ([]models.ProductsResponseDisp, error) {
	product_list := []models.ProductsResp{}

	query := `SELECT
    i.id AS id,
    i.category_id,
    c.category,
    i.product_name,
    i.size,
    i.price,
	i.stock
FROM
    inventories i
left JOIN
    categories c ON i.category_id = c.id
  LIMIT $1 OFFSET $2`
	err := prod.DB.Raw(query, pageList, offset).Scan(&product_list).Error

	if err != nil {
		return []models.ProductsResponseDisp{}, errors.New("error checking Product details")
	}

	var updatedproductDetails []models.ProductsResponseDisp
	var url []string
	for _, p := range product_list {
		err := prod.DB.Raw(`SELECT url FROM Images WHERE inventory_id=?`, p.ID).Scan(&url).Error
		if err != nil {
			return []models.ProductsResponseDisp{}, err
		}

		// Convert models.ProductsResp to models.ProductsResponseDisp
		updatedProduct := models.ProductsResponseDisp{
			ID:          p.ID,
			CategoryID:  p.CategoryID,
			ProductName: p.ProductName,
			Size:        p.Size,
			Price:       p.Price,
			Stock:       p.Stock,
			Image:       url,
		}

		updatedproductDetails = append(updatedproductDetails, updatedProduct)
	}

	return updatedproductDetails, nil

	// return product_list, nil
}

func (db *inventoryRepository) EditInventory(inventory domain.Inventories, id int) (domain.Inventories, error) {

	var modInventory domain.Inventories

	query := "UPDATE inventories SET category_id = ?, product_name = ?, size = ?, stock = ?, price = ? WHERE id = ?"

	if err := db.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price, id).Error; err != nil {
		return domain.Inventories{}, err
	}

	if err := db.DB.First(&modInventory, id).Error; err != nil {
		return domain.Inventories{}, err
	}
	return modInventory, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {

	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integet is not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) CheckInventoryByCatAndName(cat int, prdct string) (bool, error) {
	var count int
	err := i.DB.Raw("select count(*) from inventories where product_name =? and category_id = ?", prdct, cat).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count <= 0 {
		return false, err
	}
	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, stock int) (models.ProductsResponse, error) {

	if i.DB == nil {
		return models.ProductsResponse{}, errors.New("database connection is nil")
	}
	if err := i.DB.Exec("UPDATE inventories SET stock = stock + $1 WHERE id= $2", stock, pid).Error; err != nil {
		return models.ProductsResponse{}, err
	}
	var newdetails models.ProductsResponse
	var newstock int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&newstock).Error; err != nil {
		return models.ProductsResponse{}, err
	}
	newdetails.ID = uint(pid)
	newdetails.Stock = newstock

	return newdetails, nil

}

func (i *inventoryRepository) ShowIndividualProducts(id string) (models.ProductsResponse, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return models.ProductsResponse{}, errors.New("convertion not happened")
	}
	var product models.ProductsResponse
	err := i.DB.Raw(`
	SELECT
		*
		FROM
			inventories
		
		WHERE
			inventories.id = ?
			`, pid).Scan(&product).Error

	if err != nil {
		return models.ProductsResponse{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {

	var k int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=$1", pid).Scan(&k).Error; err != nil {
		return 0, err
	}

	return k, nil
}

func (c *inventoryRepository) FetchProductDetails(productId uint) (models.Inventories, error) {
	var product models.Inventories
	err := c.DB.Raw(`SELECT price,stock FROM inventories WHERE id=?`, productId).Scan(&product).Error
	return product, err
}

func (i *inventoryRepository) GetInventory(prefix string) ([]models.ProductsResponse, error) {
	var productDetails []models.ProductsResponse

	query := `
	SELECT i.*
	FROM inventories i
	LEFT JOIN categories c ON i.category_id = c.id
	WHERE i.product_name ILIKE '%' || $1 || '%'
    OR c.category ILIKE '%' || $1 || '%';

`
	if err := i.DB.Raw(query, prefix).Scan(&productDetails).Error; err != nil {
		return []models.ProductsResponse{}, err
	}

	return productDetails, nil
}

func (i *inventoryRepository) FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error) {
	var product_list []models.ProductsResponse

	query := `SELECT * FROM inventories WHERE category_id = ?`

	if err := i.DB.Raw(query, CategoryIdInt).Scan(&product_list).Error; err != nil {
		return nil, err
	}

	return product_list, nil
}
