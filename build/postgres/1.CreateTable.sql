CREATE TABLE tasks(
    id serial NOT NULL,
    name varchar(100) NOT NULL,
    status integer default 0,
    due_date_time TIMESTAMP  NOT NULL,
    deleted_at TIMESTAMP ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    primary key (id)
);
