# API данного сервера
## `api` - корневой эндпоинт для API-обращений
---
### `product` - подкорень продуктового эндпоинта
#### GET `product/{id}`
Получение информации о продукте по параметру `{id}`.

**Возвращаемые данные**
- `id` **int64** - ID продукта в системе
- `supplier_id` **int64** - ID продавца данного продукта в системе
- `name` **string** - наименование продукта
- `description` **string** - описание продукта
- `price` **int32** - цена в рублях
- `quantity` **int32** - доступное для заказа количество

**Примеры ответов**

**`200` - OK | `product/4`**
```json
{
    "id": 4,
    "supplier_id": 1,
    "name": "Творог",
    "description": "Он вкусный, вы обязаны его попробовать!",
    "price": 100,
    "quantity": 20
}
```

**`404` - NOT FOUND | `product/14510959`**
```json
{
    "code": 404,
    "message": "Product by ID \"14510959\" not exists"
}
```

#### GET `product/supplier{id}`
Получение списка продуктов по ID продавцу

**Возвращаемые данные**
- `id` **int64** - ID продавца
- `name` **string** - имя продавца 
- `products` **array[product]** - список продуктов в объекте product:
    - `id` **int64** - ID продукта в системе
    - `name` **string** - наименование продукта
    - `description` **string** - описание продукта
    - `price` **int32** - цена в рублях
    - `quantity` **int32** - доступное для заказа количество

**Примеры ответов**

**`200` - OK | `product/supplier1`**
```json
{
    "id": 1,
    "name": "Ланс",
    "products": [
            {
                "id": 3,
                "name": "Топлёное сгущенное молоко в банке, 500г.",
                "description": "Использован только натуральный сахар и натуральное топлёное молоко, ничего лишнего!",
                "price": 250,
                "quantity": 100
            },
            {
                "id": 4,
                "name": "Творог",
                "description": "Он вкусный, вы обязаны его попробовать!",
                "price": 250,
                "quantity": 100
            },
            {
                "id": 5,
                "name": "Молока отборное 3,2%, 1,5 л.",
                "description": "Только натуральное коровье молоко из села Двориных без каких-либо разбавлений!",
                "price": 130,
                "quantity": 50
            }
    ]
}
```

**`404` - NOT FOUND | `product/supplier4`**
```json
{
    "code": 404,
    "message": "Supplier by ID \"4\" not exists"
}
```

#### GET `products/suppliers`
Получение продавцов по limit и offset Query-параметрам

**Query-параметры**
- `limit` **int32** - макс. список продавцов
- `offset` **int32** - поинтер начального индекса списка продавцов

**Возвращаемые данные**
- `suppliers` **array[suppliers]** - список продавцов
    - `id` **int64** - ID продавца
    - `name` **string** - имя продавца 
    - `products` **array[product]** - список продуктов в объекте product:
        - `id` **int64** - ID продукта в системе
        - `name` **string** - наименование продукта
        - `description` **string** - описание продукта
        - `price` **int32** - цена в рублях
        - `quantity` **int32** - доступное для заказа количество

**Примеры ответов**

**`200` - OK | `product/suppliers?limit=5&offset=0`**
```json
[
    {
        "id": 1,
        "name": "Ланс",
        "products": [
                {
                    "id": 3,
                    "name": "Топлёное сгущенное молоко в банке, 500г.",
                    "description": "Использован только натуральный сахар и натуральное топлёное молоко, ничего лишнего!",
                    "price": 250,
                    "quantity": 100
                },
                {
                    "id": 4,
                    "name": "Творог",
                    "description": "Он вкусный, вы обязаны его попробовать!",
                    "price": 250,
                    "quantity": 100
                },
                {
                    "id": 5,
                    "name": "Молока отборное 3,2%, 1,5 л.",
                    "description": "Только натуральное коровье молоко из села Двориных без каких-либо разбавлений!",
                    "price": 130,
                    "quantity": 50
                }
        ]
    },
    {
        "id": 2,
        "name": "Вегадор",
        "products": [
                {
                    "id": 6,
                    "name": "Виноград Кишмиш розовый без косточки",
                    "description": "Выращен в садах герцога",
                    "price": 2500,
                    "quantity": 10
                },
                {
                    "id": 7,
                    "name": "Авокадо Хасс \"ВВ отборное\", шт",
                    "description": "",
                    "price": 210,
                    "quantity": 100
                },
                {
                    "id": 8,
                    "name": "Виноград Мускат премиум",
                    "description": "Made in Sweden",
                    "price": 600,
                    "quantity": 50
                },
                {
                    "id": 9,
                    "name": "Сметана 15%, 250 г",
                    "description": "Пискарёвское, сделано в г. Санкт-Петербург",
                    "price": 89,
                    "quantity": 100
                }
        ]
    }
]
```

**`400` - Bad request | `product/suppliers?limit=5&offset=4`**
```json
{
    "code": 400,
    "message": "Request offset is greater than suppliers count"
}
```

### `supplier` - подкорень эндпоинта среды продавцов
#### GET `supplier/products`
Получение списка своих продуктов

***- Должны быть куки файлы с supplierId параметром!***

**Возвращаемые данные**
- `products` **array[product]** - список продуктов в объекте product:
    - `id` **int64** - ID продукта в системе
    - `name` **string** - наименование продукта
    - `description` **string** - описание продукта
    - `price` **int32** - цена в рублях
    - `quantity` **int32** - доступное для заказа количество

**Примеры ответов**

**`200` - OK | `supplier/products`**
```json
[
    {
        "id": 3,
        "name": "Топлёное сгущенное молоко в банке, 500г.",
        "description": "Использован только натуральный сахар и натуральное топлёное молоко, ничего лишнего!",
        "price": 250,
        "quantity": 100
    },
    {
        "id": 4,
        "name": "Творог",
        "description": "Он вкусный, вы обязаны его попробовать!",
        "price": 250,
        "quantity": 100
    },
    {
        "id": 5,
        "name": "Молока отборное 3,2%, 1,5 л.",
        "description": "Только натуральное коровье молоко из села Двориных без каких-либо разбавлений!",
        "price": 130,
        "quantity": 50
    }
]
```

#### POST `supplier/products/{id}`
Обновление информации о данном продукте по query-параметру `{id}`

***- Должны быть куки файлы с supplierId параметром!***

**Данные запроса**
- `price` **int32** ***optional*** - новая цена продукта
- `description` **int32** ***optional*** - новое описание продукта

**Возвращаемые данные**
- `id` **int64** - ID продукта в системе
- `name` **string** - наименование продукта
- `description` **string** ***nullable*** - *Если менялся данный параметр.* Описание продукта
- `price` **int32** ***nullable*** - *Если менялся данный параметр.* Цена в рублях

**Примеры ответов**

**`200` - OK | `supplier/products/4`**

**Request body**
```json
{
    "price": 130
}
```

**Response**
```json
{
    "id": 4,
    "name": "Творог",
    "price": 100
}
```

**`200` - OK | `supplier/products/3`**

**Request body**
```json
{
    "price": 200,
    "description": "Использован только натуральный сахар и натуральное топлёное молоко. Сделано в Краснодаре"
}
```

**Response**
```json
{
    "id": 3,
    "name": "Топлёное сгущенное молоко в банке, 500г.",
    "description": "Использован только натуральный сахар и натуральное топлёное молоко. Сделано в Краснодаре",
    "price": 200
}
```

**`400` - Bad request | `supplier/products/3`**

**Response**
```json
{
    "code": 400,
    "message": "Missed request body"
}
```

**`401` - Unauthorized request | `supplier/products/3`**

**Request body**
```json
{
    "price": 200,
    "description": "Использован только натуральный сахар и натуральное топлёное молоко. Сделано в Краснодаре"
}
```

**Response**
```json
{
    "code": 401,
    "message": "Missed cookie session, request declined"
}
```

#### GET `supplier/shipments`
Получение информации о незавершенных поставках

***- Должны быть куки файлы с supplierId параметром!***

#### POST `supplier/shipments/ship`
Оформление новой поставки продуктов
