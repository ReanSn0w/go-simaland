package simaland

import (
	"fmt"
	"net/http"
)

type Options func(*Client) (*Client, error)

// WithLogger - Добавляет логгер с клиенту
func WithLogger(log Logger) Options {
	return func(c *Client) (*Client, error) {
		c.log = log
		return c, nil
	}
}

// WithClient - Настраивает ваш http клиент для обработки запросов
func WithClient(cl *http.Client) Options {
	return func(c *Client) (*Client, error) {
		c.http = cl
		return c, nil
	}
}

// WithAPIKey - Добавляет ключ API к HTTP запросу
func WithApiKey(key string) Options {
	return func(c *Client) (*Client, error) {
		c.mods = append(c.mods, SetHeader("x-api-key", key))
		return c, nil
	}
}

// WithJWT - Добавляет к запросу заголовок для авторизации по JWT токену
func WithJWT(login, password string) Options {
	return func(c *Client) (*Client, error) {
		jwt_auth, err := newJWTAuth(c, login, password)
		if err != nil {
			return nil, err
		}

		c.mods = append(c.mods, jwt_auth.modificator())
		return c, nil
	}
}

type RequestModificator func(*http.Request)

func SetHeader(key, val string) RequestModificator {
	return func(r *http.Request) {
		r.Header.Add(key, val)
	}
}

func SetQuery(key, val string) RequestModificator {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set(key, val)
		r.URL.RawQuery = q.Encode()
	}
}

// SetPage Устанавливает номер выводимой страницы
func SetPage(page int) RequestModificator {
	return SetQuery("page", fmt.Sprint(page))
}

// SetCount устанавливает кол-во товаров на странице
// ENUM: 10, 20, 50, 100
// defaultValue: 50
func SetCount(count int) RequestModificator {
	switch count {
	case 10, 20, 50, 100:
		return SetQuery("per-page", fmt.Sprint(count))
	default:
		return SetQuery("per-page", fmt.Sprint(50))
	}
}
