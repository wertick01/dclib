
CREATE TABLE authors (
    author_id INTEGER NOT NULL AUTO_INCREMENT, 
    author_name VARCHAR(100) NOT NULL, 
    author_surname VARCHAR(100) NOT NULL, 
    author_patrynomic VARCHAR(100), 
    author_photo VARCHAR(100) NOT NULL, 
    author_stars INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY(author_id)
);

CREATE TABLE books ( 
    book_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    book_name VARCHAR(100) NOT NULL, 
    book_count INTEGER, 
    book_photo VARCHAR(100) NOT NULL,
    book_stars INTEGER NOT NULL DEFAULT 0);

CREATE TABLE books_authors (
    book_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    FOREIGN KEY(book_id) REFERENCES books(book_id),
    FOREIGN KEY(author_id) REFERENCES authors(author_id)
    );

CREATE TABLE roles (
    role_id INTEGER NOT NULL PRIMARY KEY,
    user_role VARCHAR(100) NOT NULL
    );

CREATE TABLE users (
    userid INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    username VARCHAR(100) NOT NULL,
    usersurname VARCHAR(100) NOT NULL,
    userpatrynomic VARCHAR(100),
    userphone VARCHAR(100) NOT NULL,
    useremail VARCHAR(100) NOT NULL,
    userhash VARCHAR(100) NOT NULL,
    userrole INTEGER NOT NULL,
    FOREIGN KEY(userrole) REFERENCES roles(role_id)
    );

CREATE TABLE favoriete_books (
    userid INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    FOREIGN KEY(book_id) REFERENCES books(book_id),
    FOREIGN KEY(userid) REFERENCES users(userid)
    );

CREATE TABLE favoriete_authors (
    userid INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    FOREIGN KEY(userid) REFERENCES users(userid),
    FOREIGN KEY(author_id) REFERENCES authors(author_id)
    );

CREATE TABLE booking (
     id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
     book_id INTEGER NOT NULL, 
     userid INTEGER NOT NULL, 
     date_of_issue TIMESTAMP NOT NULL, 
     date_of_delivery TIMESTAMP, 
     is_confirm BOOLEAN,
     FOREIGN KEY(book_id) REFERENCES books(book_id),
     FOREIGN KEY(userid) REFERENCES users(userid)
    );

INSERT INTO authors (author_name, author_surname, author_patrynomic, author_photo) VALUES
    ('Lev', 'Tolstoy', 'Nickolaevich', ''),
    ('Steven', 'King', 'Edwid', ''),
    ('Adolph', 'Hitler', 'Alaizovich', ''),    
    ('Mikhail', 'Lermontov', 'Jurievich', ''),    
    ('Nikolay', 'Hohol', 'Vasilevich', ''),    
    ('Fedor', 'Dostoevsky', 'Mikhailovich', ''),    
    ('Ray', 'Bradbury', 'Douglas', ''),    
    ('Vladimir', 'Lenin', 'Ilich', '');

INSERT INTO books (book_name, book_count, book_photo) VALUES
    ('War and peace', 1, ''),
    ('It', 2, ''),
    ('Something', 1, ''),    
    ('Mein Kamph', 1488, ''),    
    ('Hero of our time', 1, ''),    
    ('Taras Bulba', 2, ''),    
    ('Died souls', 3, ''),    
    ('Crime and punishment', 1, ''),    
    ('Dandelion wine', 2, ''),    
    ('The State and the Revolution', 5, '');

INSERT INTO books_authors (book_id, author_id) VALUES
    (1, 1),
    (2, 2),
    (3, 2),    
    (4, 3),    
    (5, 4),    
    (6, 5),    
    (7, 5),    
    (8, 6),    
    (9, 7),    
    (10, 8);

INSERT INTO roles (role_id, user_role) VALUES 
    (1, 'admin'), 
    (2, 'user');