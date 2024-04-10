CREATE TABLE users
(
    id int primary key,
    email          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE twits
(
    id serial not null unique,
    name varchar(255) not null,
    user_id int,
    foreign key (user_id) references users (id)
);