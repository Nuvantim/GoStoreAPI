FROM alpine:latest

# Set working directory di dalam container
WORKDIR /app

# Copy binary file terlebih dahulu karena jarang berubah
COPY bin/main /app

# Copy environment file
COPY .env /app

# Copy file lainnya
COPY . .

# Set izin eksekusi untuk binary
RUN chmod +x /app/main

# Jalankan aplikasi
CMD ["./main"]
