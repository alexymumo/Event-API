

create table events(
    eventid int not null auto_increment primary key, 
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    location bigint references location(location_id),
    amount bigint not null,
    date varchar(255) not null,
    creator varchar(255) not null,
    liked boolean not null
);


create table location(
    longitude bigint not null,
    latitude bigint not null
)

insert into events(`title`,`description`, `location`) values ('Test title', 'test description', 'test location');


create table if not exists users(userid int not null auto_increment primary key,name varchar(255) not null,email varchar(255) not null,password varchar(255) not null)