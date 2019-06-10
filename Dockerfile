FROM golang:1.12-alpine as builder
WORKDIR /app
COPY . .
RUN apk add --no-cache --virtual git
RUN env GOOS=linux GOARCH=amd64 && go mod tidy && go build -o demogo .

FROM alpine
WORKDIR /app
COPY --from=builder /app/demogo /app/
CMD [ "./demogo"  ]
EXPOSE 1112