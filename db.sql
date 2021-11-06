	create table users (
		user_id serial primary key,
		username varchar (50) unique not null,
		password varchar (100) not null,
		created_on timestamp default current_timestamp,
		last_login timestamp,
		session varchar (100) 
	);

	create table dogs (
		dog_id serial primary key,
		name varchar (150) not null
	);

	
	create table cats (
		dog_id serial primary key,
		name varchar (150) not null
	);

	-- Select all tables names
	-- https://www.postgresql.org/docs/current/infoschema-tables.html
	select table_name
  	from information_schema.tables
 	where table_schema='public'
   	and table_type='BASE TABLE';