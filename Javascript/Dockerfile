FROM golang:latest

WORKDIR /go/src/github.com/jinwo-o/PatientData/Javascript
COPY . .

# RUN go build main.go

EXPOSE 8080
ENTRYPOINT [ "/go/src/github.com/jinwo-o/PatientData/Javascript/bash_script" ]
