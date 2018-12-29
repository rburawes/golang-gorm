-- Table: authors

drop table if exists authors cascade;

create table authors
(
  author_id serial,
  firstname character varying(255) not null,
  lastname character varying(255) not null,
  middlename character varying(255),
  about character varying(1000) not null,
  constraint authors_pkey primary key (author_id)
)
with (
  oids=false
);
alter table authors
  owner to postgres;