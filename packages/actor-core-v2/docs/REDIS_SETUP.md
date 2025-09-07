# Redis (Memurai) Setup Guide

## 📋 Overview

This guide will help you set up Redis using Memurai on Windows for Actor Core v2.0 testing.

## 🚀 Quick Start

### **Option 1: Using Memurai (Recommended)**

1. **Download Memurai**
   - Go to [Memurai.com](https://www.memurai.com/)
   - Download the free version for Windows
   - Install Memurai

2. **Start Memurai**
   ```bash
   # Method 1: Command Line
   memurai-server.exe --port 6379 --maxmemory 100mb --save 60 1
   
   # Method 2: Windows Service (if installed as service)
   net start memurai
   
   # Method 3: GUI
   # Open Memurai GUI and start server
   ```

3. **Test Connection**
   ```bash
   # Run our test script
   go run ./scripts/simple_redis.go
   
   # Or run integration tests
   go test ./tests/integration/redis_integration_test.go -v
   ```

### **Option 2: Using Docker**

1. **Install Docker Desktop**
   - Download from [Docker.com](https://www.docker.com/products/docker-desktop)

2. **Run Redis Container**
   ```bash
   docker run -d --name redis-test -p 6379:6379 redis:latest
   ```

3. **Test Connection**
   ```bash
   go run ./scripts/simple_redis.go
   ```

### **Option 3: Using WSL2 with Redis**

1. **Install WSL2**
   ```bash
   wsl --install
   ```

2. **Install Redis in WSL2**
   ```bash
   sudo apt update
   sudo apt install redis-server
   sudo service redis-server start
   ```

3. **Test Connection**
   ```bash
   go run ./scripts/simple_redis.go
   ```

## 🔧 Configuration

### **Default Configuration**
- **Host**: localhost
- **Port**: 6379
- **Password**: (none)
- **Database**: 0
- **Max Memory**: 100MB
- **Save**: 60 seconds, 1 change

### **Custom Configuration**

You can modify the Redis configuration in:
- `tests/integration/redis_helper.go`
- `tests/integration/redis_integration_test.go`
- `scripts/simple_redis.go`

```go
config := cache.RedisConfig{
    Host:         "localhost",
    Port:         6379,
    Password:     "", // Set password if needed
    DB:           0,
    MaxRetries:   3,
    DialTimeout:  5 * time.Second,
    ReadTimeout:  3 * time.Second,
    WriteTimeout: 3 * time.Second,
    PoolSize:     10,
    MinIdleConns: 5,
    EnableTLS:    false,
}
```

## 🧪 Testing

### **Run All Redis Tests**
```bash
# Integration tests
go test ./tests/integration/redis_integration_test.go -v

# Simple test
go run ./scripts/simple_redis.go

# Using Makefile
make test-redis
make test-redis-integration
make test-redis-script
```

### **Test Categories**

1. **Connection Test**
   - Basic connectivity
   - Health check
   - Authentication

2. **Basic Operations**
   - Set/Get/Delete
   - Exists check
   - TTL operations

3. **Complex Data**
   - JSON serialization
   - Nested objects
   - Arrays and maps

4. **Performance Test**
   - Bulk operations
   - Throughput measurement
   - Memory usage

5. **Advanced Features**
   - Tag-based invalidation
   - Pub/Sub messaging
   - Statistics collection

## 🐛 Troubleshooting

### **Common Issues**

1. **Connection Refused**
   ```
   dial tcp [::1]:6379: connectex: No connection could be made because the target machine actively refused it.
   ```
   **Solution**: Start Redis server first

2. **Port Already in Use**
   ```
   bind: address already in use
   ```
   **Solution**: Use different port or stop existing Redis instance

3. **Authentication Failed**
   ```
   NOAUTH Authentication required
   ```
   **Solution**: Set password in configuration or disable authentication

4. **Memory Issues**
   ```
   OOM command not allowed when used memory > 'maxmemory'
   ```
   **Solution**: Increase maxmemory or clear existing data

### **Debug Commands**

```bash
# Check if Redis is running
netstat -an | findstr :6379

# Test Redis CLI
redis-cli ping

# Check Redis info
redis-cli info

# Clear all data
redis-cli flushall

# Monitor Redis commands
redis-cli monitor
```

## 📊 Performance Expectations

### **Expected Performance**
- **Connection Time**: < 100ms
- **Set Operations**: > 10,000 ops/sec
- **Get Operations**: > 15,000 ops/sec
- **Memory Usage**: < 100MB for test data
- **Latency**: < 1ms for local operations

### **Benchmark Results**
```
Operations: 2000 (set + get)
Duration: 150ms
Ops/sec: 13,333
Memory Usage: 2.5MB
```

## 🔒 Security Considerations

### **Production Setup**
1. **Set Password**
   ```bash
   memurai-server.exe --requirepass yourpassword
   ```

2. **Enable TLS**
   ```go
   config := cache.RedisConfig{
       EnableTLS: true,
       // ... other config
   }
   ```

3. **Network Security**
   - Use firewall rules
   - Bind to specific interfaces
   - Use VPN for remote access

## 📚 Additional Resources

- [Memurai Documentation](https://docs.memurai.com/)
- [Redis Documentation](https://redis.io/docs/)
- [Go Redis Client](https://github.com/go-redis/redis)
- [Actor Core v2.0 Cache Design](./CACHE_DESIGN_V2.md)

## 🆘 Support

If you encounter issues:

1. Check Redis server status
2. Verify network connectivity
3. Check configuration settings
4. Review error logs
5. Test with Redis CLI first

For Actor Core v2.0 specific issues, check the [Troubleshooting Guide](./TROUBLESHOOTING_GUIDE_V2.md).
