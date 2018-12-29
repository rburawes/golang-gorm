-- Table: authors

drop table if exists users cascade;

create table users
(
  id serial,
  firstname character varying(255) not null,
  lastname character varying(255) not null,
  middlename character varying(255),
  email character varying(255),
  username character varying(255) not null,
  password character varying(255) not null,
  constraint users_pkey primary key (id)
)
with (
  oids=false
);
alter table users owner to postgres;
drop index if exists username_index;
drop index if exists email_index;
create unique index username_index on users using btree (username);
create unique index email_index on users using btree (email);