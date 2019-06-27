# SQL Script that creates fake patient data
# Create the database parameters; ID, Name, Gender, Disease
CREATE DATABASE IF NOT EXISTS patients;
USE patients;

CREATE TABLE IF NOT EXISTS patient_table ( patient_id smallint unsigned not null, name varchar(20) not null, gender varchar(20) not null, \
disease varchar(50) not null, constraint pk_example primary key (patient_id) );


# Create a loop function; Generates 100 Patients
DELIMITER $$

USE patients

CREATE PROCEDURE test_mysql_while_loop()
BEGIN
DECLARE id_int  INT;
DECLARE name_str  VARCHAR(255);
DECLARE gender_str  VARCHAR(255);
DECLARE disease_str  VARCHAR(255);

SET id_int = 1;
SET name_str =  'jin';
SET gender_str = 'M';
SET disease_str = 'flu';

WHILE id_int  <= 100 DO
INSERT INTO patient_table ( patient_id, name, gender, disease ) VALUES ( id_int, name_str, gender_str, disease_str );
SET  id_int = id_int + 1; 
END WHILE;

SELECT * FROM patient_table;
END$$

DELIMITER ;

# Call the loop/generate patients function
USE patients;

CALL test_mysql_while_loop();
