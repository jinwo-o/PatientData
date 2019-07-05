# **PatientData**

Outputs an array of patient_id, name, gender, and disease. PatientData RESTAPI was written in Golang, with a mySQL database, and routed using chi

**Prerequisites**

Must have mySQL running, golang, and other go-dependencies

**1. Setup Database**

Go to the SQL directory and run database_rebuild.sh, $source database_rebuild.sh

You may need to edit database_rebuild.sh if your root password is not root1234

(Run $source destroy_database.sh to drop the database)

**2. Start Server**

Go to the restAPI directory and run the makefile, $make run

Open up localhost:8080/Posts
