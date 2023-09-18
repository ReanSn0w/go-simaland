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

// MARK: - Category
// Методы для работы со списком категорий

func (c *Client) ListCategory(page int, count int) (*ListResponse[Category], error) {
	res := &ListResponse[Category]{}
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
	res := &ListResponse[Item]{}
	err := c.get("item", res,
		SetCount(count),
		SetPage(page),
		SetQuery("category_id", fmt.Sprint(category)),
		SetQuery("expand", "description"),
	)
	return res, err
}

// GetItem - Получает товар из Simaland
func (c *Client) GetItem(id int) (*Item, error) {
	res := &Item{}
	err := c.get(fmt.Sprintf("item/%v", id), res,
		SetQuery("expand", "description"),
	)
	return res, err
}

// MARK: - Delivery

func (c *Client) ListDeliveryAddress(settlement int64, page, count int) (ListResponse[DeliveryPoint], error) {
	res := ListResponse[DeliveryPoint]{}
	err := c.get(
		"delivery-address", &res,
		SetQuery("settlement", fmt.Sprint(settlement)),
		SetPage(page),
		SetCount(count),
	)
	return res, err
}

// FindSettlement - Поиск поселения для получения списка доставки
func (c *Client) FindSettlement(name string) (ListResponse[Settlement], error) {
	res := ListResponse[Settlement]{}
	err := c.get("settlement", &res, SetQuery("name", name))
	return res, err
}

// CheckuotByProducts - создание заказа без использования корзины
func (c *Client) CheckoutByProducts(email, phone, name string, cart map[int64]int) (*Order, error) {
	res := &Order{}
	err := c.post("order/checkout-by-products/", map[string]any{
		"comment":                  "", //
		"contact_email":            email,
		"contact_phone":            phone,
		"contact_name":             name,
		"items_data":               newCartItemFromMap(cart),
		"phone":                    phone,
		"person_type":              1,
		"manager_action":           3,
		"contact_person":           name,
		"is_use_digital_signature": false,
		"delivery_type_id":         1,
		"pickup_type_id":           1,
		"deliveryTypeId":           1,
		"paymentTypeId":            1, // Идентификатор оплаты картой
	}, res)
	return res, err
}

// GetOrder - Получение заказа по его идентификатору
func (c *Client) GetOrder(id int64) (*Order, error) {
	res := &Order{}
	err := c.get(fmt.Sprintf("order/%v", id), res)
	return res, err
}
