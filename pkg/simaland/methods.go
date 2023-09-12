package simaland

import (
	"fmt"
)

// Authors - Возвращает список авторов
//
// page [int] - номер страницы
// count [int] - кол-во элементов (enum: )
func (c *Client) ListAuthors(page, count int) ([]Author, error) {
	res := []Author{}
	err := c.get("author", &res,
		SetQuery("page", fmt.Sprint(page)),
		SetQuery("perPage", fmt.Sprint(count)),
	)
	return res, err
}

// GetAuthor - Возвращает одного автора
func (c *Client) GetAuthor(id int) (*Author, error) {
	var res *Author
	err := c.get(fmt.Sprintf("author/%v", id), &res)
	return res, err
}

func (c *Client) ListDeliveryAddress() (interface{}, error) {
	var res interface{}
	err := c.get("user-delivery-address", &res)
	return res, err
}

// MARK: - Category
// Методы для работы со списком категорий

func (c *Client) ListCategory(page int, count int) (*ListResponse[Category], error) {
	var res *ListResponse[Category]
	err := c.get("category", res,
		SetCount(count),
		SetPage(page),
	)
	return res, err
}

// MARK: - Item
// Методы для работы со списком товаров

// ListItems - Производит листинг товаров
func (c *Client) ListItems(page int, count int, category int) (*ListResponse[Item], error) {
	var res *ListResponse[Item]
	err := c.get("item", res,
		SetCount(count),
		SetPage(page),
		SetQuery("category_id", fmt.Sprint(category)),
	)
	return res, err
}

// GetItem - Получает товар из Simaland
func (c *Client) GetItem(id int) (*Item, error) {
	res := &Item{}
	err := c.get(fmt.Sprintf("item/%v", id), res)
	return res, err
}
