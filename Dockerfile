FROM golang:latest

WORKDIR /go/src/github.com/jinwo-o/PatientData
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build main.go

EXPOSE 8080
ENTRYPOINT [ "/go/src/github.com/jinwo-o/PatientData/main" ]

# RUN apt-get update
# RUN apt-get install vim -y
# RUN go get "github.com/go-sql-driver/mysql"


# WORKDIR /go/src/github.com/jinwo-o/PatientData/restAPI

# RUN make run

# CMD /go/src/github.com/jinwo-o/PatientData/bin/restAPI


# RUN go get "github.com/jinwo-o/PatientData"


# ADD ./src /go/src/PatientData
# WORKDIR /go/src/PatientData

# CMD ["go", "run", "main.go"]

# WORKDIR /go/src/github.com/jinwo-o/PatientData/restAPI
# COPY . ./
# RUN go get -d -v ./...
# RUN go install -v ./...

# RUN go install .
# CMD ["/go/bin/restAPI"]

# CMD ["go", "run", "main.go"]