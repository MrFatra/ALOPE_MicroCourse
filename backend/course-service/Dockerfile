# Kita pakai 1.24 sebagai base karena 1.25 belum ada imagenya
FROM golang:1.24-alpine

# Wajib install git dan build-base agar Go bisa download toolchain
RUN apk add --no-cache git build-base

WORKDIR /app

# INI KUNCINYA:
# Mengizinkan Go mendownload compiler 1.25 sendiri saat build
ENV GOTOOLCHAIN=go1.25.0+auto
# Mengarahkan cache agar tidak hilang
ENV GOMODCACHE=/go/pkg/mod

COPY go.mod go.sum ./

# Saat RUN ini dipanggil, Go akan melihat 'go 1.25' di go.mod, 
# lalu dia akan OTOMATIS mendownload binary Go 1.25 ke dalam container.
RUN go mod download

COPY . .

# Build menggunakan compiler 1.25 yang baru didownload tadi
RUN go build -o main ./cmd/main.go

EXPOSE 8080
CMD ["./main"]