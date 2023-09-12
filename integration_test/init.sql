create table accounts
(
    id    BIGSERIAL,
    name  varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    roles varchar(255) NOT NULL,
    PRIMARY KEY (id)
)
