# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20
ENV WORKDIR="/app"
# create a working directory inside the image
WORKDIR ${WORKDIR}

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ${WORKDIR}
CMD /bin/bash