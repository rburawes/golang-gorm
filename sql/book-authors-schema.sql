-- Table: book_authors
drop table if exists book_authors;
create table book_authors
(
  id serial not null,
  book_isbn character varying(14) not null,
  author_id integer not null,
  constraint book_authors_pkey primary key (id)
)
with (
  oids=false
);
alter table book_authors
  owner to postgres;
