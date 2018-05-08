FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./... \
    && go install -v ./...

EXPOSE 7000
CMD ["klusternaut"]