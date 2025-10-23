FROM golang:1.21

WORKDIR /workspace

# Install basic tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy entrypoint
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Run the app
CMD ["/entrypoint.sh"]
