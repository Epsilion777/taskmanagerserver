CREATE TABLE todos 
(
    id serial primary key,
    title varchar(100),
    progress int,
    date date, 
    userid varchar(100)

)
