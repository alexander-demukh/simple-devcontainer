FROM golang:1.21

WORKDIR /workspace

# Install basic tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Keep container running
CMD ["sleep", "infinity"]
