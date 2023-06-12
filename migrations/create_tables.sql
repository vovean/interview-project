begin transaction;

create table pastes
(
    id         integer,
    title      text,
    body       text,
    created_at int -- timestamp
);

create table changelog
(
    id             integer,
    paste_id       int,
    creator_ip     text,
    paste_body_len int
);

commit;