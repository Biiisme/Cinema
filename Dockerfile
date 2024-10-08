# Giai đoạn 1: Build ứng dụng Go
FROM golang:1.23.0-alpine as builder

WORKDIR /Cinema-1

# Sao chép file Go mod và sum
COPY go.mod go.sum ./

# Tải về các phụ thuộc Go
RUN go mod download

# Sao chép các file khác của ứng dụng
COPY . .

# Build ứng dụng Go
RUN go build -o Cinema-1 main.go

# Giai đoạn 2: Tạo một image nhỏ để chạy ứng dụng Go
FROM alpine:latest

# Cài đặt client Postgres (tùy chọn, nếu ứng dụng Go của bạn cần)
RUN apk --no-cache add postgresql-client

WORKDIR /Cinema-1

# Sao chép binary Go đã build từ giai đoạn trước
COPY --from=builder /Cinema-1 .

# Mở cổng mà ứng dụng sẽ chạy
EXPOSE 8080

# Lệnh để chạy ứng dụng Go
CMD ["./Cinema-1"]
