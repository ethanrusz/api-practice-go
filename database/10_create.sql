-- Script to create test DB.
create table album
(
    id     varchar(2) primary key not null,
    title  varchar(50),
    artist varchar(50),
    price  decimal
);

insert into album
(id, title, artist, price)
values ('1', 'Blue Train', 'John Coltraine', 56.99),
       ('2', 'Jeru', 'Gerry Mulligan', 17.99),
       ('3', 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);
