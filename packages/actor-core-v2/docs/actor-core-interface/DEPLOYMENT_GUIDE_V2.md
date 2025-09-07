# Deployment Guide - Actor Core v2.0

## Tổng Quan

Deployment Guide cho Actor Core v2.0 cung cấp hướng dẫn chi tiết về triển khai, bao gồm environment setup, configuration, monitoring, và maintenance.

## 1. Environment Setup

### 1.1 System Requirements

#### 1.1.1 Minimum Requirements
| Component | Minimum | Recommended | Production |
|-----------|---------|-------------|------------|
| CPU | 2 cores | 4 cores | 8+ cores |
| RAM | 4GB | 8GB | 16GB+ |
| Storage | 20GB | 50GB | 100GB+ |
| Network | 100Mbps | 1Gbps | 10Gbps+ |
| OS | Linux/Windows/macOS | Linux | Linux |

#### 1.1.2 Software Dependencies
```bash
# Go runtime
go version 1.25.1+

# Database (MongoDB)
mongodb 6.0+

# Monitoring tools
prometheus 2.40+
grafana 9.0+
node_exporter 1.5+

# Logging tools
fluentd 1.15+
elasticsearch 8.0+
kibana 8.0+

# Security tools
vault 1.12+
consul 1.15+
```

### 1.2 Environment Configuration

#### 1.2.1 Development Environment
```yaml
# docker-compose.dev.yml
version: '3.8'
services:
  actor-core:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ENV=development
      - LOG_LEVEL=debug
      - MONGODB_URI=mongodb://localhost:27017
    volumes:
      - ./config:/app/config
      - ./logs:/app/logs
    depends_on:
      - mongodb
      - redis

  mongodb:
    image: mongo:6.0
    ports:
      - "27017:27017"
    environment:
      - MONGO_INITDB_ROOT_USERNAME=admin
      - MONGO_INITDB_ROOT_PASSWORD=password
    volumes:
      - mongodb_data:/data/db

  redis:
    image: redis:7.0
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./monitoring/prometheus.yml:/etc/prometheus/prometheus.yml

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana_data:/var/lib/grafana

volumes:
  mongodb_data:
  redis_data:
  grafana_data:
```

#### 1.2.2 Production Environment
```yaml
# kubernetes/production.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: actor-core

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: actor-core
  namespace: actor-core
spec:
  replicas: 3
  selector:
    matchLabels:
      app: actor-core
  template:
    metadata:
      labels:
        app: actor-core
    spec:
      containers:
      - name: actor-core
        image: actor-core:latest
        ports:
        - containerPort: 8080
        env:
        - name: ENV
          value: "production"
        - name: LOG_LEVEL
          value: "info"
        - name: MONGODB_URI
          valueFrom:
            secretKeyRef:
              name: mongodb-secret
              key: uri
        resources:
          requests:
            memory: "512Mi"
            cpu: "250m"
          limits:
            memory: "1Gi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: actor-core-service
  namespace: actor-core
spec:
  selector:
    app: actor-core
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

## 2. Configuration Management

### 2.1 Configuration Files

#### 2.1.1 Application Configuration
```yaml
# config/application.yml
app:
  name: "Actor Core v2.0"
  version: "2.0.0"
  environment: "production"
  
server:
  port: 8080
  host: "0.0.0.0"
  timeout: 30s
  
database:
  mongodb:
    uri: "mongodb://localhost:27017"
    database: "actor_core"
    timeout: 10s
    max_pool_size: 100
    min_pool_size: 10
    
cache:
  redis:
    host: "localhost"
    port: 6379
    password: ""
    db: 0
    timeout: 5s
    
logging:
  level: "info"
  format: "json"
  output: "stdout"
  file:
    path: "/var/log/actor-core"
    max_size: "100MB"
    max_backups: 5
    max_age: 30
    
monitoring:
  prometheus:
    enabled: true
    port: 9090
    path: "/metrics"
  health:
    enabled: true
    port: 8081
    path: "/health"
    
security:
  jwt:
    secret: "your-secret-key"
    expiration: "24h"
  rate_limit:
    enabled: true
    requests_per_minute: 1000
  encryption:
    enabled: true
    algorithm: "AES-256-GCM"
    key: "your-encryption-key"
```

#### 2.1.2 Actor Core Configuration
```yaml
# config/actor-core.yml
actor_core:
  stats:
    primary:
      vitality:
        min_value: 1
        max_value: 999999
        default_value: 10
      endurance:
        min_value: 1
        max_value: 999999
        default_value: 10
      # ... other primary stats
      
    derived:
      hp_max:
        formula: "(vitality * 15 + endurance * 10 + constitution * 5) * (1 + level / 100)"
        dependencies: ["vitality", "endurance", "constitution", "level"]
      stamina:
        formula: "(endurance * 20 + vitality * 5) * (1 + level / 100)"
        dependencies: ["endurance", "vitality", "level"]
      # ... other derived stats
      
  flexible_systems:
    speed_system:
      enabled: true
      categories:
        - movement
        - casting
        - crafting
        - learning
        - combat
        - social
        - administrative
        
    karma_system:
      enabled: true
      levels:
        - global
        - world
        - region
        - sect
        - nation
        
    administrative_system:
      enabled: true
      division_types:
        - world
        - continent
        - nation
        - race
        - sect
        - realm
        - zone
        - city
        - village
        
  proficiency_system:
    enabled: true
    max_skills: 1000
    categories:
      - combat
      - magic
      - crafting
      - social
      - movement
      - survival
      
  skill_system:
    enabled: true
    max_skills: 500
    categories:
      - combat
      - magic
      - profession
      - social
      - movement
      - survival
```

### 2.2 Environment Variables

#### 2.2.1 Required Environment Variables
```bash
# Database
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=actor_core

# Cache
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Security
JWT_SECRET=your-jwt-secret-key
ENCRYPTION_KEY=your-encryption-key

# Monitoring
PROMETHEUS_ENABLED=true
PROMETHEUS_PORT=9090

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

#### 2.2.2 Optional Environment Variables
```bash
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
SERVER_TIMEOUT=30s

# Cache
CACHE_TTL=3600s
CACHE_MAX_SIZE=1000

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS_PER_MINUTE=1000

# Security
PASSWORD_MIN_LENGTH=8
PASSWORD_REQUIRE_SPECIAL=true
MAX_LOGIN_ATTEMPTS=5
LOCKOUT_DURATION=15m
```

## 3. Deployment Strategies

### 3.1 Blue-Green Deployment

#### 3.1.1 Blue-Green Setup
```yaml
# kubernetes/blue-green.yaml
apiVersion: v1
kind: Service
metadata:
  name: actor-core-service
spec:
  selector:
    app: actor-core
    version: blue
  ports:
  - port: 80
    targetPort: 8080

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: actor-core-blue
spec:
  replicas: 3
  selector:
    matchLabels:
      app: actor-core
      version: blue
  template:
    metadata:
      labels:
        app: actor-core
        version: blue
    spec:
      containers:
      - name: actor-core
        image: actor-core:blue
        ports:
        - containerPort: 8080

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: actor-core-green
spec:
  replicas: 3
  selector:
    matchLabels:
      app: actor-core
      version: green
  template:
    metadata:
      labels:
        app: actor-core
        version: green
    spec:
      containers:
      - name: actor-core
        image: actor-core:green
        ports:
        - containerPort: 8080
```

#### 3.1.2 Blue-Green Switch Script
```bash
#!/bin/bash
# blue-green-switch.sh

CURRENT_VERSION=$(kubectl get service actor-core-service -o jsonpath='{.spec.selector.version}')
NEW_VERSION=""

if [ "$CURRENT_VERSION" = "blue" ]; then
    NEW_VERSION="green"
else
    NEW_VERSION="blue"
fi

echo "Switching from $CURRENT_VERSION to $NEW_VERSION"

# Update service selector
kubectl patch service actor-core-service -p '{"spec":{"selector":{"version":"'$NEW_VERSION'"}}}'

# Wait for rollout to complete
kubectl rollout status deployment/actor-core-$NEW_VERSION

# Verify switch
kubectl get pods -l app=actor-core,version=$NEW_VERSION

echo "Switch completed successfully"
```

### 3.2 Rolling Deployment

#### 3.2.1 Rolling Update Configuration
```yaml
# kubernetes/rolling-update.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: actor-core
spec:
  replicas: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 2
  selector:
    matchLabels:
      app: actor-core
  template:
    metadata:
      labels:
        app: actor-core
    spec:
      containers:
      - name: actor-core
        image: actor-core:latest
        ports:
        - containerPort: 8080
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
```

#### 3.2.2 Rolling Update Script
```bash
#!/bin/bash
# rolling-update.sh

NEW_IMAGE=$1
if [ -z "$NEW_IMAGE" ]; then
    echo "Usage: $0 <new-image>"
    exit 1
fi

echo "Starting rolling update to $NEW_IMAGE"

# Update deployment
kubectl set image deployment/actor-core actor-core=$NEW_IMAGE

# Wait for rollout to complete
kubectl rollout status deployment/actor-core

# Verify deployment
kubectl get pods -l app=actor-core

echo "Rolling update completed successfully"
```

### 3.3 Canary Deployment

#### 3.3.1 Canary Configuration
```yaml
# kubernetes/canary.yaml
apiVersion: argoproj.io/v1alpha1
kind: Rollout
metadata:
  name: actor-core
spec:
  replicas: 5
  strategy:
    canary:
      steps:
      - setWeight: 20
      - pause: {duration: 10m}
      - setWeight: 40
      - pause: {duration: 10m}
      - setWeight: 60
      - pause: {duration: 10m}
      - setWeight: 80
      - pause: {duration: 10m}
  selector:
    matchLabels:
      app: actor-core
  template:
    metadata:
      labels:
        app: actor-core
    spec:
      containers:
      - name: actor-core
        image: actor-core:latest
        ports:
        - containerPort: 8080
```

## 4. Monitoring & Observability

### 4.1 Metrics Collection

#### 4.1.1 Prometheus Configuration
```yaml
# monitoring/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  - "actor-core-rules.yml"

scrape_configs:
  - job_name: 'actor-core'
    static_configs:
      - targets: ['actor-core:8080']
    metrics_path: '/metrics'
    scrape_interval: 5s

  - job_name: 'node-exporter'
    static_configs:
      - targets: ['node-exporter:9100']

  - job_name: 'mongodb-exporter'
    static_configs:
      - targets: ['mongodb-exporter:9216']

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093
```

#### 4.1.2 Custom Metrics
```go
// metrics/custom_metrics.go
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // Actor Core Metrics
    ActorCount = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "actor_core_actor_count",
            Help: "Number of actors in the system",
        },
        []string{"type", "status"},
    )
    
    StatCalculationDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "actor_core_stat_calculation_duration_seconds",
            Help: "Duration of stat calculations",
            Buckets: prometheus.DefBuckets,
        },
        []string{"stat_type", "calculation_type"},
    )
    
    StatCalculationErrors = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "actor_core_stat_calculation_errors_total",
            Help: "Total number of stat calculation errors",
        },
        []string{"error_type", "stat_name"},
    )
    
    // Flexible Systems Metrics
    FlexibleSystemUsage = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "actor_core_flexible_system_usage",
            Help: "Usage of flexible systems",
        },
        []string{"system_name", "component"},
    )
    
    // Performance Metrics
    MemoryUsage = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "actor_core_memory_usage_bytes",
            Help: "Memory usage in bytes",
        },
        []string{"component"},
    )
    
    CacheHitRate = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "actor_core_cache_hit_rate",
            Help: "Cache hit rate",
        },
        []string{"cache_type"},
    )
)
```

### 4.2 Logging Configuration

#### 4.2.1 Structured Logging
```go
// logging/structured_logger.go
package logging

import (
    "github.com/sirupsen/logrus"
    "github.com/fluent/fluent-logger-golang/fluent"
)

type StructuredLogger struct {
    logger  *logrus.Logger
    fluent  *fluent.Fluent
}

func NewStructuredLogger() *StructuredLogger {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
    })
    
    fluent, err := fluent.New(fluent.Config{
        FluentHost: "localhost",
        FluentPort: 24224,
    })
    if err != nil {
        logger.WithError(err).Error("Failed to create fluent logger")
    }
    
    return &StructuredLogger{
        logger: logger,
        fluent: fluent,
    }
}

func (sl *StructuredLogger) LogActorEvent(actorID, eventType string, data map[string]interface{}) {
    entry := sl.logger.WithFields(logrus.Fields{
        "actor_id":    actorID,
        "event_type":  eventType,
        "timestamp":   time.Now().Unix(),
    })
    
    for key, value := range data {
        entry = entry.WithField(key, value)
    }
    
    entry.Info("Actor event")
    
    // Send to Fluentd
    if sl.fluent != nil {
        sl.fluent.Post("actor.events", map[string]interface{}{
            "actor_id":   actorID,
            "event_type": eventType,
            "data":       data,
            "timestamp":  time.Now().Unix(),
        })
    }
}
```

#### 4.2.2 Log Aggregation
```yaml
# monitoring/fluentd.conf
<source>
  @type tail
  path /var/log/actor-core/*.log
  pos_file /var/log/fluentd/actor-core.log.pos
  tag actor-core.*
  format json
  time_key timestamp
  time_format %Y-%m-%dT%H:%M:%S.%L%z
</source>

<filter actor-core.**>
  @type record_transformer
  <record>
    service_name "actor-core"
    environment "#{ENV['ENVIRONMENT'] || 'development'}"
  </record>
</filter>

<match actor-core.**>
  @type elasticsearch
  host elasticsearch
  port 9200
  index_name actor-core
  type_name _doc
  include_timestamp true
</match>
```

### 4.3 Health Checks

#### 4.3.1 Health Check Endpoints
```go
// health/health_checker.go
package health

import (
    "context"
    "net/http"
    "time"
)

type HealthChecker struct {
    checks map[string]HealthCheck
}

type HealthCheck interface {
    Check(ctx context.Context) error
}

type HealthStatus struct {
    Status    string            `json:"status"`
    Timestamp time.Time         `json:"timestamp"`
    Checks    map[string]string `json:"checks"`
}

func (hc *HealthChecker) CheckHealth(ctx context.Context) *HealthStatus {
    status := &HealthStatus{
        Status:    "healthy",
        Timestamp: time.Now(),
        Checks:    make(map[string]string),
    }
    
    for name, check := range hc.checks {
        if err := check.Check(ctx); err != nil {
            status.Status = "unhealthy"
            status.Checks[name] = err.Error()
        } else {
            status.Checks[name] = "ok"
        }
    }
    
    return status
}

// Database health check
type DatabaseHealthCheck struct {
    db Database
}

func (dhc *DatabaseHealthCheck) Check(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    return dhc.db.Ping(ctx)
}

// Cache health check
type CacheHealthCheck struct {
    cache Cache
}

func (chc *CacheHealthCheck) Check(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    return chc.cache.Ping(ctx)
}
```

## 5. Backup & Recovery

### 5.1 Database Backup

#### 5.1.1 MongoDB Backup Script
```bash
#!/bin/bash
# backup-mongodb.sh

BACKUP_DIR="/backups/mongodb"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_NAME="actor_core_backup_$DATE"

# Create backup directory
mkdir -p $BACKUP_DIR

# Create backup
mongodump --host localhost:27017 --db actor_core --out $BACKUP_DIR/$BACKUP_NAME

# Compress backup
tar -czf $BACKUP_DIR/$BACKUP_NAME.tar.gz -C $BACKUP_DIR $BACKUP_NAME

# Remove uncompressed backup
rm -rf $BACKUP_DIR/$BACKUP_NAME

# Upload to cloud storage
aws s3 cp $BACKUP_DIR/$BACKUP_NAME.tar.gz s3://actor-core-backups/

# Clean up old backups (keep last 30 days)
find $BACKUP_DIR -name "*.tar.gz" -mtime +30 -delete

echo "Backup completed: $BACKUP_NAME.tar.gz"
```

#### 5.1.2 Automated Backup Schedule
```yaml
# kubernetes/backup-cronjob.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: mongodb-backup
spec:
  schedule: "0 2 * * *"  # Daily at 2 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: mongodb-backup
            image: mongo:6.0
            command:
            - /bin/bash
            - -c
            - |
              mongodump --host mongodb:27017 --db actor_core --out /backup
              tar -czf /backup/actor_core_backup_$(date +%Y%m%d_%H%M%S).tar.gz -C /backup actor_core
              aws s3 cp /backup/*.tar.gz s3://actor-core-backups/
            volumeMounts:
            - name: backup-storage
              mountPath: /backup
          volumes:
          - name: backup-storage
            persistentVolumeClaim:
              claimName: backup-pvc
          restartPolicy: OnFailure
```

### 5.2 Configuration Backup

#### 5.2.1 Configuration Backup Script
```bash
#!/bin/bash
# backup-config.sh

CONFIG_DIR="/app/config"
BACKUP_DIR="/backups/config"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_NAME="config_backup_$DATE"

# Create backup directory
mkdir -p $BACKUP_DIR

# Create backup
tar -czf $BACKUP_DIR/$BACKUP_NAME.tar.gz -C $CONFIG_DIR .

# Upload to cloud storage
aws s3 cp $BACKUP_DIR/$BACKUP_NAME.tar.gz s3://actor-core-backups/config/

# Clean up old backups (keep last 7 days)
find $BACKUP_DIR -name "*.tar.gz" -mtime +7 -delete

echo "Config backup completed: $BACKUP_NAME.tar.gz"
```

## 6. Maintenance & Updates

### 6.1 Regular Maintenance

#### 6.1.1 Maintenance Checklist
```bash
#!/bin/bash
# maintenance-checklist.sh

echo "=== Actor Core v2.0 Maintenance Checklist ==="

# 1. Check system resources
echo "1. Checking system resources..."
df -h
free -h
top -bn1 | head -20

# 2. Check application logs
echo "2. Checking application logs..."
tail -n 100 /var/log/actor-core/application.log | grep -i error

# 3. Check database status
echo "3. Checking database status..."
mongosh --eval "db.stats()"

# 4. Check cache status
echo "4. Checking cache status..."
redis-cli ping

# 5. Check monitoring metrics
echo "5. Checking monitoring metrics..."
curl -s http://localhost:9090/metrics | grep actor_core

# 6. Check health endpoints
echo "6. Checking health endpoints..."
curl -s http://localhost:8080/health | jq .

# 7. Check disk space
echo "7. Checking disk space..."
du -sh /var/log/actor-core
du -sh /backups

# 8. Check security updates
echo "8. Checking security updates..."
apt list --upgradable | grep security

echo "=== Maintenance checklist completed ==="
```

#### 6.1.2 Automated Maintenance
```yaml
# kubernetes/maintenance-cronjob.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: maintenance
spec:
  schedule: "0 3 * * 0"  # Weekly on Sunday at 3 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: maintenance
            image: actor-core:latest
            command:
            - /bin/bash
            - -c
            - |
              # Log rotation
              find /var/log/actor-core -name "*.log" -mtime +7 -delete
              
              # Database optimization
              mongosh --eval "db.runCommand({compact: 'actors'})"
              
              # Cache cleanup
              redis-cli FLUSHDB
              
              # Metrics cleanup
              find /var/lib/prometheus -name "*.db" -mtime +30 -delete
          restartPolicy: OnFailure
```

### 6.2 Update Procedures

#### 6.2.1 Zero-Downtime Update
```bash
#!/bin/bash
# zero-downtime-update.sh

NEW_VERSION=$1
if [ -z "$NEW_VERSION" ]; then
    echo "Usage: $0 <new-version>"
    exit 1
fi

echo "Starting zero-downtime update to version $NEW_VERSION"

# 1. Build new image
echo "1. Building new image..."
docker build -t actor-core:$NEW_VERSION .

# 2. Push to registry
echo "2. Pushing to registry..."
docker push actor-core:$NEW_VERSION

# 3. Update deployment
echo "3. Updating deployment..."
kubectl set image deployment/actor-core actor-core=actor-core:$NEW_VERSION

# 4. Wait for rollout
echo "4. Waiting for rollout..."
kubectl rollout status deployment/actor-core

# 5. Verify update
echo "5. Verifying update..."
kubectl get pods -l app=actor-core
kubectl get service actor-core-service

# 6. Run health checks
echo "6. Running health checks..."
curl -f http://localhost:8080/health || exit 1

echo "Zero-downtime update completed successfully"
```

#### 6.2.2 Rollback Procedure
```bash
#!/bin/bash
# rollback.sh

PREVIOUS_VERSION=$1
if [ -z "$PREVIOUS_VERSION" ]; then
    echo "Usage: $0 <previous-version>"
    exit 1
fi

echo "Starting rollback to version $PREVIOUS_VERSION"

# 1. Rollback deployment
echo "1. Rolling back deployment..."
kubectl rollout undo deployment/actor-core

# 2. Wait for rollback
echo "2. Waiting for rollback..."
kubectl rollout status deployment/actor-core

# 3. Verify rollback
echo "3. Verifying rollback..."
kubectl get pods -l app=actor-core
kubectl get service actor-core-service

# 4. Run health checks
echo "4. Running health checks..."
curl -f http://localhost:8080/health || exit 1

echo "Rollback completed successfully"
```

## 7. Troubleshooting

### 7.1 Common Issues

#### 7.1.1 Performance Issues
```bash
#!/bin/bash
# troubleshoot-performance.sh

echo "=== Performance Troubleshooting ==="

# Check CPU usage
echo "CPU usage:"
top -bn1 | grep "Cpu(s)"

# Check memory usage
echo "Memory usage:"
free -h

# Check disk I/O
echo "Disk I/O:"
iostat -x 1 5

# Check network usage
echo "Network usage:"
iftop -t -s 10

# Check application metrics
echo "Application metrics:"
curl -s http://localhost:9090/metrics | grep actor_core

# Check database performance
echo "Database performance:"
mongosh --eval "db.currentOp()"

# Check cache performance
echo "Cache performance:"
redis-cli info stats
```

#### 7.1.2 Database Issues
```bash
#!/bin/bash
# troubleshoot-database.sh

echo "=== Database Troubleshooting ==="

# Check MongoDB status
echo "MongoDB status:"
systemctl status mongod

# Check MongoDB logs
echo "MongoDB logs:"
tail -n 50 /var/log/mongodb/mongod.log

# Check database connections
echo "Database connections:"
mongosh --eval "db.serverStatus().connections"

# Check database locks
echo "Database locks:"
mongosh --eval "db.currentOp()"

# Check database size
echo "Database size:"
mongosh --eval "db.stats()"

# Check slow queries
echo "Slow queries:"
mongosh --eval "db.setProfilingLevel(2, {slowms: 100})"
```

### 7.2 Emergency Procedures

#### 7.2.1 Emergency Shutdown
```bash
#!/bin/bash
# emergency-shutdown.sh

echo "=== Emergency Shutdown ==="

# 1. Stop application
echo "1. Stopping application..."
kubectl scale deployment actor-core --replicas=0

# 2. Stop database
echo "2. Stopping database..."
kubectl scale deployment mongodb --replicas=0

# 3. Stop cache
echo "3. Stopping cache..."
kubectl scale deployment redis --replicas=0

# 4. Stop monitoring
echo "4. Stopping monitoring..."
kubectl scale deployment prometheus --replicas=0
kubectl scale deployment grafana --replicas=0

echo "Emergency shutdown completed"
```

#### 7.2.2 Emergency Recovery
```bash
#!/bin/bash
# emergency-recovery.sh

echo "=== Emergency Recovery ==="

# 1. Start database
echo "1. Starting database..."
kubectl scale deployment mongodb --replicas=1
kubectl wait --for=condition=ready pod -l app=mongodb --timeout=300s

# 2. Start cache
echo "2. Starting cache..."
kubectl scale deployment redis --replicas=1
kubectl wait --for=condition=ready pod -l app=redis --timeout=300s

# 3. Start application
echo "3. Starting application..."
kubectl scale deployment actor-core --replicas=3
kubectl wait --for=condition=ready pod -l app=actor-core --timeout=300s

# 4. Start monitoring
echo "4. Starting monitoring..."
kubectl scale deployment prometheus --replicas=1
kubectl scale deployment grafana --replicas=1

# 5. Verify recovery
echo "5. Verifying recovery..."
curl -f http://localhost:8080/health || exit 1

echo "Emergency recovery completed"
```

---

*Tài liệu này cung cấp hướng dẫn chi tiết về triển khai Actor Core v2.0, đảm bảo hệ thống được triển khai một cách an toàn và hiệu quả.*
