create database turf;

CREATE TABLE jockey (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL,
   apellido VARCHAR (50) NOT NULL,
   ranking integer NOT NULL DEFAULT 0,
   fechaModificacion VARCHAR (20) NOT NULL
);
CREATE UNIQUE INDEX jockey_nombre_apellido_unique ON jockey (nombre, apellido);


CREATE TABLE cuidador (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL,
   apellido VARCHAR (50) NOT NULL,
   ranking integer NOT NULL DEFAULT 0,
   fechaModificacion VARCHAR (20) NOT NULL
);
CREATE UNIQUE INDEX cuidador_nombre_apellido_unique ON cuidador (nombre, apellido);


CREATE TABLE caballo (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) ,
   fecha_nac INTEGER ,
   sexo VARCHAR (1) ,
   pelaje VARCHAR (10) ,
   padre VARCHAR (10) ,
   madre VARCHAR (10) ,
   peso INTEGER NULL,
   fechaModificacion VARCHAR (20) NOT NULL
);
CREATE UNIQUE INDEX caballo_nombre_unique ON caballo (nombre);


CREATE TABLE carrera (
   id SERIAL PRIMARY KEY,
   fecha VARCHAR (20) NOT NULL DEFAULT '1900-01-01',
   nombre VARCHAR (50) NOT NULL UNIQUE,
   pista VARCHAR (20) NULL,
   distancia INTEGER NOT NULL
 ); 


CREATE TABLE carrera_detalle (
   id SERIAL PRIMARY KEY,
   nombre VARCHAR (50) NOT NULL REFERENCES carrera (nombre) ON DELETE CASCADE,
   competidor INTEGER NOT NULL REFERENCES caballo (id) ON DELETE CASCADE,
   jockey INTEGER NOT NULL REFERENCES jockey (id) ON DELETE CASCADE,
   cuidador INTEGER NOT NULL REFERENCES cuidador (id) ON DELETE CASCADE,
   handicap INTEGER NOT NULL,
   resultado INTEGER 

 ); 


 CREATE TABLE totales (
   id SERIAL PRIMARY KEY,
   aciertos_1pos INTEGER NOT NULL, -- * 5
   aciertos_2pos INTEGER NOT NULL, -- * 3
   aciertos_3pos INTEGER NOT NULL, -- * 2
   carreras INTEGER NOT NULL,
   eficiencia NUMERIC (10,2) 
 ); 

-- *************



CREATE OR REPLACE FUNCTION merge (v_nom TEXT, v_ape TEXT, v_table INTEGER)
RETURNS integer AS $$

DECLARE
    v_id integer;
    fecha TEXT;
BEGIN 

-- creo tabla temporal e inserto registro de jockey correspondiente.

DROP TABLE IF EXISTS mytemp;

CREATE TEMP TABLE mytemp (
    nombre   TEXT,
    apellido TEXT,
    fechaModificacion TEXT
);

fecha = (SELECT CURRENT_DATE);

INSERT INTO mytemp (nombre, apellido, fechaModificacion)
VALUES (v_nom, v_ape, fecha);

BEGIN
    CASE v_table
        WHEN 1 THEN
            INSERT INTO jockey (nombre, apellido, fechaModificacion)
            SELECT nombre, apellido, fechaModificacion
            FROM mytemp
            ON CONFLICT (nombre, apellido) DO NOTHING;
            SELECT id INTO v_id FROM jockey WHERE nombre = v_nom AND apellido = v_ape;
        WHEN 2 THEN
            INSERT INTO cuidador (nombre, apellido, fechaModificacion)
            SELECT nombre, apellido, fechaModificacion
            FROM mytemp
            ON CONFLICT (nombre, apellido) DO NOTHING;
            SELECT id INTO v_id FROM cuidador WHERE nombre = v_nom AND apellido = v_ape;
        WHEN 3 THEN
            INSERT INTO caballo (nombre, fechaModificacion)
            SELECT nombre, fechaModificacion
            FROM mytemp
            ON CONFLICT (nombre) DO NOTHING;   
            SELECT id INTO v_id FROM caballo WHERE nombre = v_nom;
        ELSE
            RAISE NOTICE 'Unknown parameter';
    END CASE;
END;    



RETURN v_id;

END;
    
$$ LANGUAGE 'plpgsql';
