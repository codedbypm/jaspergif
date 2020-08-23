FROM gcr.io/jaspergif/extract:latest

RUN mkdir /app

ADD ../extract.go ../go.* /app/

WORKDIR /app

RUN go build -o extract .

CMD ["/app/extract"]