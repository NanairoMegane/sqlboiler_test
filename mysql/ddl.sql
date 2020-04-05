create table user(
    user_id int         not null auto_increment,
    name    varchar(10) not null,
    age     int         not null,
    primary key(user_id)
);

create table division(
    division_id int auto_increment,
    name        varchar(10) not null,
    user_id     int         not null,
    primary key(division_id),
    constraint fk_user_div
        foreign key(user_id)
        references user (user_id)
        on delete no action
        on update no action       
);

create table branch(
    branch_id   int auto_increment,
    name        varchar(10) not null,
    user_id     int         null,
    primary key(branch_id)
);