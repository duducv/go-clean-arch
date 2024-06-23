drop schema if exists public;
create schema public;

create table public.event (
    event_id text primary key not null,
    description text,
    price numeric,
    capacity integer
);

create table public.ticket (
    ticket_id text,
    event_id text,
    email text,
    status text
);

create table public.transaction (
    transaction_id text,
    ticket_id text,
    event_id text,
    tid text,
    price numeric,
    status text
);

insert into public.event (event_id, description, price, capacity) values ('clxnut5z2000008lg9aombhm6', 'Foo Fighters 10/10/2024 22:00', 300, 100000);