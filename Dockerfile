FROM golang:1.22-alpine as build

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o ./music-backend ./cmd/web-app/


FROM golang:1.22-alpine

WORKDIR /app

COPY --from=build /app/music-backend ./

EXPOSE 2020
CMD ["./music-backend"]