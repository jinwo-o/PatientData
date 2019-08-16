# **PatientData**

PatientData RESTAPI was written in Golang, with a mySQL database, and routed using chi. The frontend is created using jquery datatables and Ajax calls

**Docker-Compose**

Must have mySQL running, golang, and other go-dependencies

$docker-compose up --build

Then open up localhost:10001 to view the datatables

Use Postman to make different requests

Some info when running Postman:
- 100 patients are already loaded
- when patients are created they will be ID:101 and onwards
- Get: localhost:8080/patients
- GetByID: localhost:8080/patients/5 (for example)
- Create: use Post, localhost:8080/patients and put {"name":"Sample","gender":"M","disease":"meningitis"} in the Body
- Delete: use Delete, localhost:8080/patients/86

Note: Go to https://travis-ci.com/jinwo-o/PatientData to see the test-suite CI