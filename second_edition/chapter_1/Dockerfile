FROM golang:latest

RUN mkdir /files
COPY hw.go /files
WORKDIR /files

RUN go build -o /files/hw hw.go
RUN chmod +x /files/hw
ENTRYPOINT [ "/files/hw" ]