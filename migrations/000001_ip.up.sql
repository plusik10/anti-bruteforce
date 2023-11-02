CREATE TABLE ip_list(
    id SERIAL PRIMARY KEY,
    ip varchar(64),
    block_ip integer
);
CREATE TABLE Bucket(
    id SERIAL PRIMARY KEY,
    param varchar(60),
    len integer,
    cap integer,
    updateAt date
)