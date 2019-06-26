USE patients

CREATE TABLE patient_table ( 
    patient_id smallint unsigned not null, 
    name varchar(20) not null, 
    gender varchar(20) not null, 
    disease varchar(50) not null, 
    PRIMARY KEY (patient_id)
);
