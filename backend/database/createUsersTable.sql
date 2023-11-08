create table if not exists users (
    id integer unsigned auto_increment primary key,
    email varchar(100) unique, not null,
    password_hash varchar(100) not null,
    uuid varchar(100),
    activate_flag boolean not null,
    created_at datetime
);
