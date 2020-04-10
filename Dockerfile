FROM golang:1.14.0 AS build_stage

WORKDIR /app

RUN go get github.com/oxequa/realize

COPY ./ /app

ADD .realize.yaml .

EXPOSE 1323

CMD ["realize", "start", "--run"]