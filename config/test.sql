create table event
(
    eventid int not null auto_increment, 
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    location varchar(255) NOT NULL,
    primary key(`eventid`)
);

insert into events(`name`,`description`, `location`) values ('Test title', 'test description', 'test location');