# go-simaland
Библиотека для работы с API магазина simaland

#### Реализованы только необходимые мне методы, в случае необходимости, добавление других методов возможно с помощью метода DO структуры Client

## Как использовать?

```
cl, _ := simaland.NewClient(simaland.WithApiKey("<API_KEY>"))
list, _ := cl.ListCategory(1, 10)
```

## Лицензия MIT