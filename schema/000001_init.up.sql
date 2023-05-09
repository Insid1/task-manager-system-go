CREATE TABLE users
(
    id            serial       not null unique primary key,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial       not null unique primary key,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique primary key,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique primary key,
    title       varchar(255) not null,
    description varchar(255),
    is_active   bool         not null default true
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique primary key,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);