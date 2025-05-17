FROM alpine:latest

# Set working directory di dalam container
WORKDIR /app

# Copy file lainnya
COPY . .

# Set izin eksekusi untuk binary
RUN chmod +x /app/bin/api

# Jalankan aplikasi
CMD ["sh", "-c", "./bin/api"]
