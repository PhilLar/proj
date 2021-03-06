-- CREATE USER philipp WITH PASSWORD 'qwedf12';
-- GRANT ALL PRIVILEGES ON DATABASE "TestDB" to philipp;


CREATE TABLE IF NOT EXISTS regions (
   id SERIAL NOT NULL UNIQUE,
   title TEXT NOT NULL,
   longtitude FLOAT NOT NULL,   
   latitude FLOAT NOT NULL,
   approxSquare FLOAT NOT NULL
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
   OF_type VARCHAR(30) NOt NULL,
   SAL FLOAT NOT NULL,
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

-- TESTING

-- CREATE TABLE IF NOT EXISTS farms (
--    id       SERIAL   NOT NULL UNIQUE,
--    title    TEXT NOT NULL
--    specialization VARCHAR(30) NOT NULL,
--    heads_of_animals     BIGINT NOT NULL,
--    heads_of_cows     BIGINT NOT NULL,
--    longitude FLOAT NOT NULL,     
--    latitude FLOAT NOT NULL,
--    address TEXT NOT NULL,
--    SAL BIGINT NOT NULL,
--    region_id BIGINT REFERENCES regions(id) NOT NULL,
--    farm_unions_id BIGINT REFERENCES regions(id)
-- );