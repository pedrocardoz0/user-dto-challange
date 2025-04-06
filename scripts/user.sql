create database if not exists user;

create schema if not exists public;

create table if not exists public.users (
    id uuid primary key default uuid_generate_v4(),
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    date_of_birth date not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

insert into public.users (first_name, last_name, email, password, date_of_birth) values ('John', 'Doe', 'john.doe@example.com', 'password123', '1990-01-01');
