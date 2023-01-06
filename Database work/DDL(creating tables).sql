CREATE TABLE IF NOT EXISTS events(
    events_id serial PRIMARY KEY not null,
    event_name varchar(250) not null,
    event_time time not null,
    event_result varchar(50) not null
);
create table if not exists subscriptions(
    subs_id serial primary key not null,
    subs_name varchar(250) not null,
    subs_code varchar(100) not null,
    category1 varchar(100) not null,
    category2 varchar(100) not null,
    parents_sub_id int not null
);
CREATE TYPE status AS ENUM ('0', '1');
create table if not exists users(
    user_id serial primary key not null,
    username varchar(250) not null,
    user_email varchar(100) not null,
    user_password varchar(300) not null,
    user_activation_status varchar(300) not null,
    status status not null
);
create table if not exists user_subscription(
    user_id integer references users(user_id),
    subs_id integer references subscriptions(subs_id)
);
create table if not exists fin_stats(
    id serial primary key  not null,
    date_id integer not null,
    fin_value float not null,
    fin_type varchar(50) not null,
    load_date time not null
);
insert into users(username, user_email, user_password, user_activation_status, status) values('Tokesh', 't_niyazbek@kbtu.kz', 'tokeshtokesh', 'OFDFDOFODFODSFWQPRQr32rR2', '1');
