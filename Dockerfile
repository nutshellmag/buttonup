FROM golang:1.18.3-alpine3.16

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o buttonup
CMD [ "/app/buttonup" ]
