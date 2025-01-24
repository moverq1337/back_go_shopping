# back_go_shopping

# Product API

Product API — это REST API для работы с пользователями, товарами и корзиной. API использует **PostgreSQL** для хранения данных и **pgAdmin** для удобного управления базой. Проект поддерживает авторизацию пользователей через JWT.

## **Функциональность**

- Регистрация и авторизация пользователей.
- Управление товарами (CRUD, с правами администратора).
- Работа с корзиной: добавление, обновление и удаление товаров.
- JWT-аутентификация.
- Связь между пользователями, товарами и корзиной.

---

## **Как запустить проект**

### **1. Установите зависимости**

Убедитесь, что у вас установлены:

- [Docker](https://www.docker.com/)
- [Go (1.19+)](https://golang.org/dl/)

### **2. Клонируйте репозиторий**

```bash
git clone https://github.com/moverq1337/back_go_shopping.git
cd product-api
```

### **3. Настройте Docker**

Проект уже содержит файл `docker-compose.yml` для запуска PostgreSQL и pgAdmin.

#### Запуск контейнеров

```bash
docker-compose up -d
```

- PostgreSQL будет доступен на порту 5432.
- pgAdmin будет доступен на [http://localhost:5050](http://localhost:5050).
  - Логин: `admin@admin.com`
  - Пароль: `admin`

### **4. Запустите сервер**

#### 1. Настройте переменные окружения

Создайте файл `.env` в корне проекта и добавьте в него:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=product_db
SERVER_PORT=8080
```

#### 2. Установите зависимости

```bash
go mod tidy
```

#### 3. Запустите сервер

```bash
go run cmd/server/main.go
```

Сервер запустится на [http://localhost:8080](http://localhost:8080).

---

## **API Эндпоинты**

### **Авторизация**

#### Регистрация пользователя

**POST** `/api/auth/register`

**Пример запроса:**

```json
{
	"email": "user@example.com",
	"username": "user",
	"password": "password123"
}
```

#### Авторизация пользователя

**POST** `/api/auth/login`

**Пример запроса:**

```json
{
	"email": "user@example.com",
	"password": "password123"
}
```

- **Ответ:** JWT-токен, который используется для аутентификации.

---

### **Управление товарами**

#### Получить все товары

**GET** `/api/products`

#### Создать товар (**только для администратора**)

**POST** `/api/admin/products`

**Пример запроса:**

```json
{
	"name": "New Product",
	"description": "Description",
	"imageUrl": "http://example.com/image.jpg",
	"sex": true,
	"isNew": true,
	"price": 99.99
}
```

- **Требуется:** Заголовок `Authorization: Bearer <admin_token>`

#### Обновить товар (**только для администратора**)

**PUT** `/api/admin/products/:id`

#### Удалить товар (**только для администратора**)

**DELETE** `/api/admin/products/:id`

---

### **Работа с корзиной**

#### Добавить товар в корзину

**POST** `/api/cart`

**Пример запроса:**

```json
{
	"productId": 1,
	"quantity": 2
}
```

- **Требуется:** Заголовок `Authorization: Bearer <user_token>`

#### Получить корзину пользователя

**GET** `/api/cart`

- **Требуется:** Заголовок `Authorization: Bearer <user_token>`

#### Обновить количество товара в корзине

**PUT** `/api/cart/:id`

- **Пример запроса:**

```json
{
	"quantity": 3
}
```

- **Требуется:** Заголовок `Authorization: Bearer <user_token>`

#### Удалить товар из корзины

**DELETE** `/api/cart/:id`

- **Требуется:** Заголовок `Authorization: Bearer <user_token>`

---

## **Тестирование API**

Вы можете использовать Insomnia для тестирования API. Импортируйте предоставленный JSON-файл (`insomnia_export.json`) в Insomnia, чтобы быстро настроить запросы.

---

## **Структура проекта**

```
product-api/
│
├── cmd/
│   └── server/
│       └── main.go          # Точка входа для запуска сервера
│
├── internal/
│   ├── handlers/            # Обработчики для эндпоинтов
│   ├── models/              # Определения моделей (User, Product, Cart)
│   ├── repository/          # Репозитории для работы с базой данных
│   ├── router/              # Настройка маршрутов API
│   ├── middleware/          # Middleware (JWT, Role-based access)
│   └── config/              # Настройка и загрузка конфигураций
│
├── pkg/
│   ├── utils/               # Утилиты (JWT, хэширование паролей)
│   └── database/            # Подключение к базе данных
│
├── docker-compose.yml       # Настройка PostgreSQL и pgAdmin
├── init.sql                 # Скрипт для создания таблиц
├── go.mod                   # Зависимости Go
├── go.sum                   # Контрольная сумма зависимостей
└── README.md                # Этот файл
```

---

## **Полезные команды**

- Остановить контейнеры Docker:

```bash
docker-compose down
```

- Перезапустить сервер:

```bash
go run cmd/server/main.go
```

- Проверить подключение к базе:

```bash
docker exec -it postgres psql -U admin -d product_db
```

---

## **TODO**

- Расширить API для работы с заказами.
