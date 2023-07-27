
CREATE TABLE `events`
(
    EventID int auto_increment,
    Title varchar(255) NOT NULL,
    Description varchar(255) NOT NULL,
    Location varchar(255) NOT NULL,
    CreatedAt varchar(255) NOT NULL,
    UpdatedAt varchar(255) NOT NULL,
    PRIMARY KEY(`EventID`)
);

CREATE TABLE `users`
(
    ID int auto_increment,
    NAME varchar(255) NOT NULL,
    EMAIL varchar(255) NOT NULL,
    PASSWORD varchar(255) NOT NULL
    PRIMARY KEY(`ID`)   
);
