begin transaction;

create table pastes
(
    id         integer primary key autoincrement,
    title      text,
    body       text,
    created_at int -- timestamp
);

create table changelog
(
    id             integer primary key autoincrement,
    paste_id       int,
    creator_ip     text,
    paste_body_len int
);

commit;