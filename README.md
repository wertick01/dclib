# Digital Clouds Library
![DCLIB](https://github.com/wertick01/dclib/blob/main/dclibrary_image.png?raw=true)

## REQUESTS

| METHOD | PATH | RETURN |
| ------ | ------ | ------ |
| POST | /api/login | Login (and JWT tocker) |
| GET | /api/refresh | Refresh the JWT tocken |
| POST | /api/Registration | User registration |
|  |  |  |
| GET | /api/users | List of users |
| GET | /api/users/{id:[0-9]+} | Find user |
| PUT | /api/users/{id:[0-9]+} | Change user |
| DELETE| /api/users/{id:[0-9]+} | Delete user |
|  |  |  |
| POST | /api/book | Creating the book |
| GET | /api/book | List of books |
| PUT | /api/book | Change book |
| GET | /api/book/{id:[0-9]+} | Find book by ID |
| POST | /api/book/{id:[0-9]+} | Put star to the book by ID |
| DELETE | /api/book/{id:[0-9]+} | Delete book by ID |
|  |  |  |
| GET | /api/authors | List of authors |
| POST | /api/authors | Create author |
| GET | /api/authors/{id:[0-9]+} | Find author by ID |
| PUT | /api/authors/{id:[0-9]+} | Change author |
| POST | /api/authors/{id:[0-9]+} | Put star to the author by ID |
| DELETE | /api/authors/{id:[0-9]+} | Delete author by ID |
| GET | /api/authors/books/{id:[0-9]+} | List of the authors books by author ID |
|  |  |  |
| POST | /api/favorietes/books/list | List of favoriete books |
| POST | /api/favorietes/authors/list | List of favoriete authors |
| POST | /api/favorietes/books/add | Add book to favorietes |
| POST | /api/favorietes/authors/add | Add author to favorietes |
| POST | /api/favorietes/books/delete | Delete book from favorietes |
| POST | /api/favorietes/authors/delete | Delete author from favorietes |
|  |  |  |
| GET | /api/reserved/list | List reserved books |
| POST | /api/reserved/reserve | Reserve book |
| PUT | /api/reserved/return | Return reserved book |
| PUT | /api/reserved/confirm | Confirm the return of the book |


## Description

- For what?
This api is a small online library within **Digital clouds** and is designed to share books between employees, drop them off and keep a record of them.
