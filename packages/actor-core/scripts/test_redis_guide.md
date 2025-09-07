# Redis Testing Guide for Actor Core v2.0

## 🚀 Quick Start Options

### **Option 1: Memurai (Easiest for Windows)**

1. **Download & Install Memurai**
   - Go to [https://www.memurai.com/](https://www.memurai.com/)
   - Download the free version
   - Install with default settings

2. **Start Memurai**
   ```bash
   # Open Command Prompt as Administrator
   memurai-server.exe --port 6379 --maxmemory 100mb --save 60 1
   ```

3. **Test Connection**
   ```bash
   # In another terminal
   go run ./scripts/simple_redis.go
   ```

### **Option 2: Docker (If Available)**

1. **Install Docker Desktop**
   - Download from [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
   - Install and start Docker Desktop

2. **Start Redis with Docker**
   ```bash
   # Start Redis container
   docker-compose up -d redis
   
   # Test connection
   go run ./scripts/simple_redis.go
   
   # Stop when done
   docker-compose down
   ```

### **Option 3: WSL2 with Redis**

1. **Install WSL2**
   ```bash
   wsl --install
   ```

2. **Install Redis in WSL2**
   ```bash
   # In WSL2 terminal
   sudo apt update
   sudo apt install redis-server
   sudo service redis-server start
   ```

3. **Test Connection**
   ```bash
   # In Windows terminal
   go run ./scripts/simple_redis.go
   ```

## 🧪 Testing Commands

### **Basic Tests**
```bash
# Test Redis connection
go run ./scripts/simple_redis.go

# Run integration tests
go test ./tests/integration/redis_integration_test.go -v

# Run specific test
go test ./tests/integration/redis_integration_test.go -v -run TestRedisConnection
```

### **Using Makefile**
```bash
# Show available commands
make help

# Test Redis connection
make test-redis

# Run all Redis tests
make test-redis-integration

# Start Redis with Docker
make docker-redis

# Stop Redis Docker container
make docker-redis-stop
```

## 🔧 Troubleshooting

### **Common Issues & Solutions**

1. **"Connection Refused" Error**
   ```
   dial tcp [::1]:6379: connectex: No connection could be made because the target machine actively refused it.
   ```
   **Solution**: Start Redis server first

2. **"Docker not found" Error**
   ```
   docker: The term 'docker' is not recognized
   ```
   **Solution**: Install Docker Desktop or use Memurai instead

3. **"Port already in use" Error**
   ```
   bind: address already in use
   ```
   **Solution**: 
   ```bash
   # Find process using port 6379
   netstat -ano | findstr :6379
   
   # Kill the process (replace PID with actual process ID)
   taskkill /PID <PID> /F
   ```

4. **"Module not found" Error**
   ```
   go: cannot find module "actor-core-v2"
   ```
   **Solution**: 
   ```bash
   # Make sure you're in the correct directory
   cd chaos-actor-module/packages/actor-core
   
   # Initialize go module
   go mod init actor-core-v2
   go mod tidy
   ```

### **Debug Commands**

```bash
# Check if Redis is running
netstat -an | findstr :6379

# Check Redis process
tasklist | findstr redis

# Test Redis CLI (if available)
redis-cli ping

# Check Go modules
go mod verify

# Clean and rebuild
go clean -cache
go mod tidy
```

## 📊 Expected Results

### **Successful Test Output**
```
🚀 Testing Redis (Memurai) Connection...
=====================================
1. Testing connection...
✅ Redis connection successful! Pong: PONG

2. Testing basic operations...
✅ Set key successful
✅ Get key successful: test_value
✅ Key exists: 1
✅ TTL: 4m59.999s
✅ Delete key successful

3. Testing complex data...
✅ Set complex data successful
✅ Get complex data successful: map[active:true age:30 name:John Doe scores:[85 92 78]]

4. Testing performance...
✅ Performance test completed:
   Operations: 2000 (set + get)
   Duration: 150ms
   Ops/sec: 13333.33

5. Testing Redis info...
✅ Redis info retrieved (length: 1234 characters)
✅ Database size: 0 keys
✅ Redis connection closed

🎉 All Redis tests completed successfully!
=====================================
Redis (Memurai) is ready for use with Actor Core v2.0!
```

## 🎯 Next Steps

After successful Redis testing:

1. **Run Full Test Suite**
   ```bash
   make test-cache
   ```

2. **Test Actor Core Integration**
   ```bash
   go test ./... -v
   ```

3. **Start Development**
   - Redis is ready for Actor Core v2.0
   - Cache system is fully functional
   - Performance monitoring is active

## 📚 Additional Resources

- [Redis Setup Guide](./docs/REDIS_SETUP.md)
- [Cache Design Document](./docs/CACHE_DESIGN_V2.md)
- [API Reference](./docs/API_REFERENCE_V2.0.md)
- [Memurai Documentation](https://docs.memurai.com/)
- [Docker Documentation](https://docs.docker.com/)

## 🆘 Need Help?

If you're still having issues:

1. Check the [Troubleshooting Guide](./docs/TROUBLESHOOTING_GUIDE_V2.md)
2. Verify your Redis installation
3. Check network connectivity
4. Review error logs
5. Try different Redis setup options

**Remember**: Redis is required for Actor Core v2.0 cache system to function properly!
