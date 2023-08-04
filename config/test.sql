create table `events`
(
    eventid bigint auto_increment,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    location varchar(255) NOT NULL,
    primary key(`eventid`)
);

insert into events(`title`,`description`, `location`)
values
('Test title', 'test description', 'test location');