create table if not exists "GenTable"
(
    id serial not null,
    url text not null,
    short_url text not null primary key
    );