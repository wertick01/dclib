# Digital Clouds Library
![DCLIB](https://github.com/wertick01/dclib/blob/main/dclibrary_image.png?raw=true)

## REQUESTS

| NUMBER | METHOD | PATH | RETURN |
| ------ | ------ | ------ | ------ |
| **1** |  | **/api/[login/refresh/registration]** |  |
| 1.1 | POST | /api/login | Login (and JWT tocker) |
| 1.2 | GET | /api/refresh | Refresh the JWT tocken |
| 1.3 | POST | /api/Registration | User registration |
|  |  |  |  |
| **2** |  | **/api/users[/]** |  |
| 2.1 | GET | /api/users | List of users |
| 2.2 | GET | /api/users/{id:[0-9]+} | Find user |
| 2.3 | PUT | /api/users/{id:[0-9]+} | Change user |
| 2.4 | DELETE| /api/users/{id:[0-9]+} | Delete user |
|  |  |  |  |
| **3** |  | **/api/book[/]** |  |
| 3.1 | POST | /api/book | Creating the book |
| 3.2 | GET | /api/book | List of books |
| 3.3 | PUT | /api/book | Change book |
| 3.4 | GET | /api/book/{id:[0-9]+} | Find book by ID |
| 3.5 | POST | /api/book/{id:[0-9]+} | Put star to the book by ID |
| 3.6 | DELETE | /api/book/{id:[0-9]+} | Delete book by ID |
|  |  |  |  |
| **4** |  | **/api/authors/** |  |
| 4.1 | GET | /api/authors | List of authors |
| 4.2 | POST | /api/authors | Create author |
| 4.3 | GET | /api/authors/{id:[0-9]+} | Find author by ID |
| 4.4 | PUT | /api/authors/{id:[0-9]+} | Change author |
| 4.5 | POST | /api/authors/{id:[0-9]+} | Put star to the author by ID |
| 4.6 | DELETE | /api/authors/{id:[0-9]+} | Delete author by ID |
| 4.7 | GET | /api/authors/books/{id:[0-9]+} | List of the authors books by author ID |
|  |  |  |  |
| **5** |  | **/api/favorietes/** |  |
| 5.1 | POST | /api/favorietes/books/list | List of favoriete books |
| 5.2 | POST | /api/favorietes/authors/list | List of favoriete authors |
| 5.3 | POST | /api/favorietes/books/add | Add book to favorietes |
| 5.4 | POST | /api/favorietes/authors/add | Add author to favorietes |
| 5.5 | POST | /api/favorietes/books/delete | Delete book from favorietes |
| 5.6 | POST | /api/favorietes/authors/delete | Delete author from favorietes |
|  |  |  |  |
| **6** |  | **/api/reserved/** |  |
| 6.1 | GET | /api/reserved/list | List reserved books |
| 6.2 | POST | /api/reserved/reserve | Reserve book |
| 6.3 | PUT | /api/reserved/return | Return reserved book |
| 6.4 | PUT | /api/reserved/confirm | Confirm the return of the book |


## Description

- For what?
This api is a small online library within **Digital clouds** and is designed to share books between employees, drop them off and keep a record of them.

## Examples of requests


### - &ensp; 1.1 &ensp;|&ensp; /api/login &ensp;|&ensp; Method: POST
```sh
curl -v --cookie ... http://localhost/api/login
```
### - &ensp; 1.2 &ensp;|&ensp; /api/refresh &ensp;|&ensp; Method: GET
```sh
curl -v --cookie ... http://localhost/api/refresh
```
### - &ensp; 1.3 &ensp;|&ensp; /api/registration &ensp;|&ensp; Method: POST
```sh
curl -v --cookie ... http://localhost/api/registration
```
