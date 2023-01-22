CREATE TABLE users
(
    user_id           serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);


CREATE TABLE todo_lists
(
    todo_lists_id          serial primary key,
    title       varchar(255) not null,
    description varchar(255)
);


CREATE TABLE todo_items
(
    todo_items_id          serial primary key,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);

ALTER TABLE todo_lists 
ADD COLUMN user_id INTEGER,
ADD FOREIGN KEY (user_id) 
REFERENCES users (user_id);

ALTER TABLE todo_items
ADD COLUMN todo_lists_id INTEGER,
ADD FOREIGN KEY (todo_lists_id)
REFERENCES todo_lists (todo_lists_id);
