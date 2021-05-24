FROM golang:latest as builder

RUN apt-get update && apt-get install -y ca-certificates openssl python3 python3-pip

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect storage.googleapis.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/storage.googleapis.crt
# Update certificates
RUN update-ca-certificates

ENV GO111MODULE=on

RUN mkdir -p /cakebox/cakebox

WORKDIR /cakebox

COPY ./cakebox/*.go /cakebox/cakebox
COPY ./birthday_server/server.go /cakebox

COPY go.mod .
COPY go.sum .
COPY gendb.py .
COPY birthdays.csv .
COPY requirements.txt .

RUN python3 -m pip install --trusted-host pypi.python.org --trusted-host files.pythonhosted.org --trusted-host pypi.org -r requirements.txt
RUN python3 gendb.py

RUN go build -o server .

EXPOSE 50051

CMD ./server
