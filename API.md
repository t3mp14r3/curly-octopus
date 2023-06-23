# 🐙 API



#### Авторизация
Авторизация необходма на всех запросах, кроме регистрации и логина. Для этого нужно добавить header `Authorization` со значением `Bearer <token>`, где `<token>` это полученный jwt token. На всякий случай, запросы в которых **нужна** авторизация помечены 🔐, в которых нет - 🔓.

#### 🔓 Регистрация
`POST` `http://localhost:5000/register`
```
{
  "login": "user",
  "name": "Mr. User",
  "email": "user@website.com",
  "password": "user123"
}
```
Ответы
```
{"token": "mylongsecrettoken"}
```
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные или уже использованные данные|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔓 Логин
`POST` `http://localhost:5000/login`
```
{
  "login": "user",
  "password": "user123"
}
```
Ответы
```
{"token": "mylongsecrettoken"}
```
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные данные|
|403 - Forbidden|Неверный пароль|
|404 - NotFound|Пользователь не найден|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔐 Профиль
`GET` `http://localhost:5000/me`

Ответы
```
{
  "login": "user",
  "name": "Mr. User",
  "email": "user@website.com",
}
```
|Статус|Значение|
|:--|:--|
|404 - NotFound|Пользователь не найден|
|200 - OK|Все хорошо|

#### 🔐 Создать продукт
`POST` `http://localhost:5000/products`
```
{
  "name": "my product",
  "desc": "product description is optional",
  "cost": 4500,
  "barcode": "1234567890"
}
```
*поле `desc` опциональное, поле `cost` должно быть целым числом*

Ответы
```
{
  "id": "product-uuid",
  "name": "my product",
  "desc": "product description is optional",
  "cost": 4500,
  "barcode": "1234567890"
}
```
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные данные|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔐 Получить продукт
`GET` `http://localhost:5000/products/<id>`
*параметр `id` это идентификатор продукта, он обязателен*

Ответы
```
{
  "id": "product-uuid",
  "name": "my product",
  "desc": "product description is optional",
  "cost": 4500,
  "barcode": "1234567890"
}
```
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные данные|
|404 - NotFound|Продукт не найден|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔐 Получить продукты
`GET` `http://localhost:5000/products`

Ответы
```
[
  {
    "id": "product-uuid",
    "name": "my product",
    "desc": "product description is optional",
    "cost": 4500,
    "barcode": "1234567890"
  },
  ...
]
```
|Статус|Значение|
|:--|:--|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔐 Удалить продукт
`DELETE` `http://localhost:5000/products/<id>`

*параметр `id` это идентификатор продукта, он обязателен*

Ответы
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные данные|
|404 - NotFound|Продукт не найден|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|

#### 🔐 Сгенерировать чек
`GET` `http://localhost:5000/check/<id>`
*параметр `id` это идентификатор продукта, по которому необходимо сгенерировать чек, он обязателен*

В ответ сервер присылает `.pdf` файл с чеком.
|Статус|Значение|
|:--|:--|
|400 - BadRequest|Неправильные данные|
|404 - NotFound|Продукт не найден|
|500 - InternalServerError|Внутренняя ошибка сервера|
|200 - OK|Все хорошо|
