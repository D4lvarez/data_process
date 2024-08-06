package queries

// Master Tables
const create_table_roles string = `create table if not exists roles (
	id int auto_increment primary key,
	name varchar(12) unique not null,
	created_at datetime not null default current_timestamp
);`

const create_table_document_types string = `create table if not exists document_types (
	id int auto_increment primary key,
	name varchar(24) unique not null,
	created_at datetime not null default current_timestamp
);`

const create_table_ports string = `create table if not exists ports (
	id int auto_increment primary key,
	name varchar(24) unique not null,
	created_at datetime not null default current_timestamp
);`

const create_table_currencies string = `create table if not exists currencies (
	id int auto_increment primary key,
	name varchar(24) unique not null,
	slug varchar(12) unique not null,
	symbol char(4) not null,
	created_at datetime not null default current_timestamp
);`

const create_table_rates string = `create table if not exists rates(
	id int auto_increment primary key,
	currency_id int not null,
	amount decimal(10,6) not null,
	created_at datetime not null default current_timestamp,
	foreign key (currency_id)
		references currencies (id)
		on update restrict
		on delete restrict
);`

const create_table_tax_payer_types string = `create table if not exists tax_payer_type(
	id int auto_increment primary key,
	name varchar(100) not null unique,
	description varchar(255) not null,
	iva boolean default false,
	islr boolean default false,
	created_at datetime not null default current_timestamp
);`

const create_table_banks string = `create table if not exists banks(
	id int auto_increment primary key,
	name varchar(32) unique not null,
	code char(4) unique not null,
	created_at datetime not null default current_timestamp
);`

const create_table_bank_accounts string = `create table if not exists bank_accounts(
	id int auto_increment primary key,
	bank_id int not null,
	account_number varchar(20),
	is_available boolean default true,
	created_at datetime not null default current_timestamp,
	updated_at datetime default null,
	foreign key (bank_id)
		references banks (id)
		on update restrict
		on delete cascade
);`

const create_table_exchange_rates string = `create table if not exists exchange_rates(
	id int auto_increment primary key,
	currency_id int not null,
	amount decimal(10,6) not null,
	created_at datetime not null default current_timestamp,
	foreign key (currency_id)
		references currencies (id)
		on update restrict
		on delete restrict
);`

func GetCreateTablesQueries() []string {
	return []string{
		create_table_roles,
		create_table_document_types,
		create_table_ports,
		create_table_currencies,
		create_table_rates,
		create_table_tax_payer_types,
		create_table_banks,
		create_table_bank_accounts,
		create_table_exchange_rates,
	}
}
