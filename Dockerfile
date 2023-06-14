FROM golang

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o restaurant-reservation-system ./cmd/main

CMD ["./restaurant-reservation-system"]

# docker-compose up --build restaurant-reservation-system