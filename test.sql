-- CREATE USER philipp WITH PASSWORD 'qwedf12';
-- GRANT ALL PRIVILEGES ON DATABASE "TestDB" to philipp;

-- CREATE TABLE films (
--     code        char(5),
--     title       varchar(40),
--     did         integer,
--     date_prod   date,
--     kind        varchar(10),
--     len         interval hour to minute,
--     CONSTRAINT production UNIQUE(date_prod)
-- );

CREATE TABLE IF NOT EXISTS regions (
   id SERIAL NOT NULL UNIQUE,
   title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS farms (
   id       SERIAL   NOT NULL UNIQUE,
   title    TEXT NOT NULL,
   specialization VARCHAR(30) NOT NULL,
   heads_of_animals     BIGINT NOT NULL,
   heads_of_cows     BIGINT NOT NULL,
   longitude FLOAT NOT NULL,     
   latitude FLOAT NOT NULL,
   address TEXT NOT NULL,
   SAL BIGINT NOT NULL,
   region_id BIGINT REFERENCES regions(id) NOT NULL,
   farm_unions_id BIGINT REFERENCES regions(id)
);

CREATE TABLE IF NOT EXISTS brigades (
   id SERIAL NOT NULL,
   title TEXT NOT NULL,
   farm_id BIGINT REFERENCES farms(id) NOT NULL
);



CREATE TABLE IF NOT EXISTS farm_unions (
   id SERIAL NOT NULL,
   title TEXT NOT NULL,
   adress TEXT NOT NULL,
   phone INT NOT NULL,
   email TEXT NOT NULL,
   region_id BIGINT REFERENCES regions(id) UNIQUE NOT NULL
);

