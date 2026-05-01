FROM alpine:latest

# Install Time Zone package
RUN apk add --no-cache tzdata

# Set working directory di dalam container
WORKDIR /app

# Copy file lainnya
COPY . .

# Set Production Environment
RUN mv .env.prod .env

# Set permit binnary execution
RUN chmod +x /app/bin/api

# Run service
CMD ["sh", "-c", "./bin/api"]
