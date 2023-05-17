FROM golang:1.19-alpine3.16

RUN mkdir /e-commerce

COPY . /e-commerce

WORKDIR /e-commerce

RUN go build -o e-commerce

LABEL Name="e-commerce" Version=1.0

EXPOSE  8080

CMD [ "./e-commerce", "--migrate=true" ]