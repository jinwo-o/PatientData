# SQL Script that creates fake patient data
# Create the database parameters; ID, Name, Gender, Disease
CREATE DATABASE IF NOT EXISTS patients;
USE patients;

CREATE TABLE IF NOT EXISTS patient_table ( ID smallint unsigned not null AUTO_INCREMENT, name varchar(20) not null, gender varchar(20) not null, \
disease varchar(50) not null, primary key (ID) );


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

WHILE id_int  <= 100 DO

SET name_str =  (select elt(0.5 + RAND() * 5, 'Klaw', 'FVV', 'PSkills', 'KLow', 'Green Ranger', 'Big Spain'));
set gender_str = (select elt(0.5 + RAND() * 2, 'M', 'F'));
SET disease_str = (select elt(0.5 + RAND() * 6, 'cold', 'flu', 'mengitis', 'hepititis B', 'Laker Fan', 'Kobe'));

INSERT INTO patient_table ( name, gender, disease ) VALUES ( name_str, gender_str, disease_str );
SET  id_int = id_int + 1; 
END WHILE;

SELECT * FROM patient_table;
END$$

DELIMITER ;

# Call the loop/generate patients function
USE patients;

CALL test_mysql_while_loop();
