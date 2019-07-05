***# PatientData***
Outputs an array of patient_id, name, gender, and disease 

**Prerequisites**

Need mySQL, golang

**1. Setup Database**

Go to the SQL directory and run database_rebuild.sh, $source database_rebuild.sh

You may need to edit database_rebuild.sh if your root password is not root1234

**2. Start Server**

Go to the restAPI directory and run the makefile, $make run

Open up localhost:8080/Patients

