create table users
(
    user_id       bigint auto_increment
        primary key,
    user_email    varchar(200) null,
    user_phone    varchar(100) null,
    user_password varchar(200) null,
    user_verify   int          null,
    created_at    datetime     null,
    updated_at    datetime     null,
    created_by    int          null,
    updated_by    int          null
);

