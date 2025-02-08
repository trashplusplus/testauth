# testauth

testauth - базовий сервіс авторизації написаний за допомогою фреймворка Gin. при реєстраціє валідуються дані та хешується пароль за допомогою bcrypt. при авторизації генерується JWT-токен з клеймами id, email та exp(час валідності токену).

проєкт розбито на пакети:
cmd - головний пакет, який імпортує internal.
internal - все, що стосується поточного проєкту, сутності, сервер та сполучення з бд.
pkg - код, який можна повторно використати з мінімальними змінами або без змін в коді в незалежності від проєкту. 

В init.sql описана схема бази даних для СУБД PostgreSQL та додана вставка корисних даних в таблицю Products.
У сервіса є 3 єндпоінта:

🟡 [POST] /signup
```
{
  "username": "user123",
  "email": "email@email.com",
  "password": "testtesttest"
}
```
🟡 [POST] /login
```
{
  "email": "email@email.com",
  "password": "testtesttest"
}
```

🟢 [GET] /protected
Якщо є валідний Bearer JWT-токен в хедері Authorization, то сервіс відповість такою json-відповіддю.
```
[
    {
        "Id": 6,
        "Title": "AP",
        "Price": 2000000
    },
    {
        "Id": 8,
        "Title": "IWC",
        "Price": 500000
    },
    {
        "Id": 5,
        "Title": "Panerai",
        "Price": 100000
    },
    {
        "Id": 9,
        "Title": "Certina",
        "Price": 80000
    },
    {
        "Id": 7,
        "Title": "Brew",
        "Price": 15000
    },
    {
        "Id": 10,
        "Title": "Spinnaker",
        "Price": 10000
    },
    {
        "Id": 4,
        "Title": "Seiko",
        "Price": 8000
    },
    {
        "Id": 1,
        "Title": "Citizen",
        "Price": 5000
    },
    {
        "Id": 3,
        "Title": "Timex",
        "Price": 3000
    },
    {
        "Id": 2,
        "Title": "Casio",
        "Price": 3000
    }
]
```
Приклад запиту без наявності токену в хедері.

![image](https://github.com/user-attachments/assets/7f2fcce9-1ae5-45d6-95be-09da9da2ee79)


З можливих покращень я б додав .env файл, в який виніс би всю конфігурацію - ключ підпису jwt-токенів, ip, port та дані БД.

