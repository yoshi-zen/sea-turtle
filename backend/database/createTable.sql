create table if not exists problems (
    id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    problem_statement text not null,
    answer text not null,
    created_at datetime
);

insert into problems (title, problem_statement, answer, created_at) values 
('test1', '赤くて甘い果物は?', 'りんご', now());
insert into problems (title, problem_statement, answer, created_at) values 
('test2', '黄色で酸っぱい果物は?', 'レモン', now());
insert into problems (title, problem_statement, answer, created_at) values 
('test3', '紫で甘い果物は?', 'ぶどう', now());

create table if not exists users (
    id integer unsigned auto_increment primary key,
    email varchar(100) unique, not null,
    password_hash varchar(100) not null,
    uuid varchar(100),
    activate_flag boolean not null,
    created_at datetime
);
