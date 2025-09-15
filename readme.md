# 📚 Bookstore RESTful API (Golang + Gin + PostgreSQL)

RESTful API untuk sistem bookstore, dikembangkan dengan **Golang**, **Gin**, **GORM**, dan **PostgreSQL**.  

### DB Seeder
```bash
go run .\cmd\main.go seed:db

### Default User
| Role  | Email                                         | Password |
| ----- | --------------------------------------------- | -------- |
| Admin | [admin@example.com](mailto:admin@example.com) | admin123 |
| User  | [user@example.com](mailto:user@example.com)   | user123  |

### Swagger URL
http://localhost:8080/swagger/index.html#/Books/post_books

