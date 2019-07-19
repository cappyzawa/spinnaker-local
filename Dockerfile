FROM golang:1

ENV CGO_ENABLED 0

RUN mkdir /src
WORKDIR /src
COPY . .
RUN go build -o spinnaker-sample .

EXPOSE 8080
CMD ["./spinnaker-sample"]
