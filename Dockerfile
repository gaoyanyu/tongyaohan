FROM golang:1.24

# 创建目标目录并设置权限
RUN mkdir -p /var/log/containers /var/containers /data /app && \
    chmod 755 /var/log/containers /var/containers /data /app


