# Bookstore RESTful API (Golang + Gin + PostgreSQL)

RESTful API untuk sistem bookstore, dikembangkan dengan **Golang**, **Gin**, **GORM**, dan **PostgreSQL**.  

### Running with Seeder + Faker
```bash
go run .\cmd\main.go seed:db
```

### Default User
| Role  | Email                                         | Password |
| ----- | --------------------------------------------- | -------- |
| Admin | [admin.bookstore@mail.com] | admin123 |
| User  | [john_doe@mail.com]   | password  |

### Swagger URL
http://localhost:8080/swagger/index.html#/Books/post_books

