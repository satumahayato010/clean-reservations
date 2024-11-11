FROM golang:1.23

WORKDIR /app

COPY . .

COPY .air.toml .

RUN apt-get update && apt-get install -y curl \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
    && go mod tidy \
    && go install go.uber.org/mock/mockgen@v0.3.0 \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/mysql

CMD ["air"]
