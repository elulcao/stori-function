FROM alpine

# Install Python
RUN apk add --update python3

# Copy data files
COPY logo.svg /data/logo.svg
COPY txns.csv /data/txns.csv

# Expose port 80 for HTTP traffic
EXPOSE 80

# Start Python HTTP server
CMD ["python3", "-m", "http.server", "80", "--directory", "/data"]
