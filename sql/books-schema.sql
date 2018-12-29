-- Table: books
drop table if exists books;
create table books
(
  isbn character(14) not null,
  title character varying(255) not null,
  price numeric(5,2) not null,
  constraint books_pkey primary key (isbn)
)
with (
  oids=false
);
alter table books
  owner to postgres;