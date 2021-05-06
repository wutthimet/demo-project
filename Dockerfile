FROM golang:latest

ENV PACKAGE_PATH=/demo-project
RUN mkdir -p /app/go/src/

WORKDIR /app/go/src/$PACKAGE_PATH

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8082

RUN go build

CMD ["./demo-project", "dev"]