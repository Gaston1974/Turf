create database turf;

CREATE TABLE jockey (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL,
   apellido VARCHAR (50),
   ranking integer NOT NULL DEFAULT 0
);

CREATE TABLE cuidador (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL,
   apellido VARCHAR (50),
   ranking integer NOT NULL DEFAULT 0
);


CREATE TABLE caballo (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL,
   fecha_nac INTEGER NOT NULL,
   sexo VARCHAR (1) NOT NULL,
   pelaje VARCHAR (10) NOT NULL,
   padre VARCHAR (10) NOT NULL,
   madre VARCHAR (10) NOT NULL,
   peso INTEGER NULL
);


CREATE TABLE carrera (
   id SERIAL PRIMARY KEY,
   fecha date NOT NULL DEFAULT '1900-01-01',
   nombre VARCHAR (20) NOT NULL,
   pista VARCHAR (20) NULL,
   distancia INTEGER NOT NULL
 ); 


CREATE TABLE carrera_detalle (
   id INT NOT NULL,
   nombre VARCHAR (20) NOT NULL,
   competidor INTEGER NOT NULL REFERENCES caballo (id) ON DELETE CASCADE,
   jockey INTEGER NOT NULL REFERENCES jockey (id) ON DELETE CASCADE,
   cuidador INTEGER NOT NULL REFERENCES cuidador (id) ON DELETE CASCADE,
   handicap INTEGER NOT NULL,
   resultado INTEGER NOT NULL

 ); 


 CREATE TABLE totales (
   id SERIAL PRIMARY KEY,
   aciertos_1pos INTEGER NOT NULL, -- * 5
   aciertos_2pos INTEGER NOT NULL, -- * 3
   aciertos_3pos INTEGER NOT NULL, -- * 2
   carreras INTEGER NOT NULL,
   eficiencia NUMERIC (10,2) 
 ); 
