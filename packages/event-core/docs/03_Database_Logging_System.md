# 03 — Database & Logging System (Hệ Thống Database & Logging)

**Generated:** 2025-01-27  
**Status:** Database Design  
**Based on:** Event sourcing & audit trail patterns

## Tổng quan

Database & Logging System lưu trữ toàn bộ event log vào database, hỗ trợ hệ thống nhân quả, thiên cơ thuật, và monitoring trong game online. Mọi hoạt động trong game đều được lưu vào các event tables để theo dõi.

## 🗄️ Database Schema (Cấu Trúc Database)

### **Core Event Tables (Bảng Event Cốt Lõi)**

#### **1. Event_Log - Bảng Log Event Chính**

```sql
-- Event_Log - Main event logging table (Bảng log event chính)
CREATE TABLE event_log (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Event identification (Nhận dạng event)
    event_id VARCHAR(64) NOT NULL UNIQUE,
    event_type VARCHAR(32) NOT NULL,
    event_category VARCHAR(32) NOT NULL,
    event_subcategory VARCHAR(32),
    
    -- Event participants (Người tham gia event)
    actor_id VARCHAR(64),
    target_id VARCHAR(64),
    world_id VARCHAR(64),
    server_id VARCHAR(32),
    
    -- Event timing (Thời gian event)
    timestamp DATETIME(6) NOT NULL,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    -- Event data (Dữ liệu event)
    data JSON,
    metadata JSON,
    
    -- Event relationships (Mối quan hệ event)
    parent_event_id VARCHAR(64),
    root_event_id VARCHAR(64),
    chain_id VARCHAR(64),
    
    -- Event status (Trạng thái event)
    status ENUM('pending', 'processing', 'completed', 'failed', 'cancelled', 'rolled_back') DEFAULT 'pending',
    priority INT DEFAULT 0,
    weight FLOAT DEFAULT 1.0,
    
    -- Event processing (Xử lý event)
    processing_time_ms INT,
    retry_count INT DEFAULT 0,
    error_message TEXT,
    
    -- Indexes (Chỉ mục)
    INDEX idx_event_type (event_type),
    INDEX idx_event_category (event_category),
    INDEX idx_actor_id (actor_id),
    INDEX idx_target_id (target_id),
    INDEX idx_world_id (world_id),
    INDEX idx_timestamp (timestamp),
    INDEX idx_parent_event_id (parent_event_id),
    INDEX idx_chain_id (chain_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### **2. Event_Chain - Bảng Chuỗi Event**

```sql
-- Event_Chain - Event chain table (Bảng chuỗi event)
CREATE TABLE event_chain (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Chain identification (Nhận dạng chuỗi)
    chain_id VARCHAR(64) NOT NULL UNIQUE,
    chain_type VARCHAR(32) NOT NULL,
    chain_name VARCHAR(128),
    chain_description TEXT,
    
    -- Chain participants (Người tham gia chuỗi)
    creator_id VARCHAR(64),
    world_id VARCHAR(64),
    server_id VARCHAR(32),
    
    -- Chain timing (Thời gian chuỗi)
    start_time DATETIME(6),
    end_time DATETIME(6),
    duration_ms BIGINT,
    
    -- Chain status (Trạng thái chuỗi)
    status ENUM('pending', 'executing', 'paused', 'completed', 'failed', 'cancelled') DEFAULT 'pending',
    current_event_index INT DEFAULT 0,
    total_events INT DEFAULT 0,
    
    -- Chain configuration (Cấu hình chuỗi)
    config JSON,
    metadata JSON,
    
    -- Chain metrics (Chỉ số chuỗi)
    success_rate FLOAT DEFAULT 0.0,
    average_processing_time_ms INT DEFAULT 0,
    total_processing_time_ms BIGINT DEFAULT 0,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_chain_type (chain_type),
    INDEX idx_creator_id (creator_id),
    INDEX idx_world_id (world_id),
    INDEX idx_status (status),
    INDEX idx_start_time (start_time),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### **3. Event_Causality - Bảng Nhân Quả Event**

```sql
-- Event_Causality - Event causality table (Bảng nhân quả event)
CREATE TABLE event_causality (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Causality identification (Nhận dạng nhân quả)
    causality_id VARCHAR(64) NOT NULL UNIQUE,
    cause_event_id VARCHAR(64) NOT NULL,
    effect_event_id VARCHAR(64) NOT NULL,
    causality_type VARCHAR(32) NOT NULL,
    
    -- Causality strength (Sức mạnh nhân quả)
    strength FLOAT NOT NULL DEFAULT 0.0,
    confidence FLOAT NOT NULL DEFAULT 0.0,
    latency_ms INT DEFAULT 0,
    
    -- Causality analysis (Phân tích nhân quả)
    impact_score FLOAT DEFAULT 0.0,
    influence_score FLOAT DEFAULT 0.0,
    correlation_score FLOAT DEFAULT 0.0,
    
    -- Causality metadata (Siêu dữ liệu nhân quả)
    analysis_method VARCHAR(32),
    analysis_algorithm VARCHAR(32),
    analysis_parameters JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_cause_event_id (cause_event_id),
    INDEX idx_effect_event_id (effect_event_id),
    INDEX idx_causality_type (causality_type),
    INDEX idx_strength (strength),
    INDEX idx_confidence (confidence),
    INDEX idx_created_at (created_at),
    
    -- Foreign keys (Khóa ngoại)
    FOREIGN KEY (cause_event_id) REFERENCES event_log(event_id) ON DELETE CASCADE,
    FOREIGN KEY (effect_event_id) REFERENCES event_log(event_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### **4. Event_Butterfly_Effect - Bảng Hiệu Ứng Cánh Bướm**

```sql
-- Event_Butterfly_Effect - Butterfly effect table (Bảng hiệu ứng cánh bướm)
CREATE TABLE event_butterfly_effect (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Effect identification (Nhận dạng hiệu ứng)
    effect_id VARCHAR(64) NOT NULL UNIQUE,
    trigger_event_id VARCHAR(64) NOT NULL,
    effect_chain_id VARCHAR(64) NOT NULL,
    effect_type VARCHAR(32) NOT NULL,
    
    -- Effect propagation (Lan truyền hiệu ứng)
    propagation_path JSON,
    propagation_time_ms INT DEFAULT 0,
    amplification_factor FLOAT DEFAULT 1.0,
    
    -- Effect analysis (Phân tích hiệu ứng)
    impact_score FLOAT DEFAULT 0.0,
    reach_count INT DEFAULT 0,
    duration_ms BIGINT DEFAULT 0,
    
    -- Effect metadata (Siêu dữ liệu hiệu ứng)
    analysis_method VARCHAR(32),
    analysis_algorithm VARCHAR(32),
    analysis_parameters JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_trigger_event_id (trigger_event_id),
    INDEX idx_effect_chain_id (effect_chain_id),
    INDEX idx_effect_type (effect_type),
    INDEX idx_impact_score (impact_score),
    INDEX idx_created_at (created_at),
    
    -- Foreign keys (Khóa ngoại)
    FOREIGN KEY (trigger_event_id) REFERENCES event_log(event_id) ON DELETE CASCADE,
    FOREIGN KEY (effect_chain_id) REFERENCES event_chain(chain_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### **Specialized Event Tables (Bảng Event Chuyên Biệt)**

#### **5. Combat_Event_Log - Bảng Log Event Chiến Đấu**

```sql
-- Combat_Event_Log - Combat event logging table (Bảng log event chiến đấu)
CREATE TABLE combat_event_log (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Event reference (Tham chiếu event)
    event_id VARCHAR(64) NOT NULL UNIQUE,
    
    -- Combat participants (Người tham gia chiến đấu)
    attacker_id VARCHAR(64) NOT NULL,
    defender_id VARCHAR(64) NOT NULL,
    combat_type VARCHAR(32) NOT NULL,
    
    -- Combat details (Chi tiết chiến đấu)
    skill_used VARCHAR(64),
    weapon_used VARCHAR(64),
    damage_dealt FLOAT DEFAULT 0.0,
    healing_done FLOAT DEFAULT 0.0,
    
    -- Combat state (Trạng thái chiến đấu)
    combat_state VARCHAR(32) NOT NULL,
    combat_phase VARCHAR(32) NOT NULL,
    combat_duration_ms INT DEFAULT 0,
    
    -- Combat results (Kết quả chiến đấu)
    victory BOOLEAN DEFAULT FALSE,
    defeat BOOLEAN DEFAULT FALSE,
    experience_gained FLOAT DEFAULT 0.0,
    loot_dropped JSON,
    
    -- Combat metadata (Siêu dữ liệu chiến đấu)
    combat_location VARCHAR(128),
    combat_environment VARCHAR(32),
    combat_conditions JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_attacker_id (attacker_id),
    INDEX idx_defender_id (defender_id),
    INDEX idx_combat_type (combat_type),
    INDEX idx_combat_state (combat_state),
    INDEX idx_victory (victory),
    INDEX idx_created_at (created_at),
    
    -- Foreign key (Khóa ngoại)
    FOREIGN KEY (event_id) REFERENCES event_log(event_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### **6. Item_Event_Log - Bảng Log Event Vật Phẩm**

```sql
-- Item_Event_Log - Item event logging table (Bảng log event vật phẩm)
CREATE TABLE item_event_log (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Event reference (Tham chiếu event)
    event_id VARCHAR(64) NOT NULL UNIQUE,
    
    -- Item details (Chi tiết vật phẩm)
    item_id VARCHAR(64) NOT NULL,
    item_type VARCHAR(32) NOT NULL,
    item_rarity VARCHAR(16) NOT NULL,
    item_quantity INT DEFAULT 1,
    item_value FLOAT DEFAULT 0.0,
    
    -- Item operation (Thao tác vật phẩm)
    operation VARCHAR(32) NOT NULL,
    source_location VARCHAR(128),
    destination_location VARCHAR(128),
    
    -- Item enhancement (Cường hóa vật phẩm)
    enhancement_level INT DEFAULT 0,
    enhancement_success BOOLEAN DEFAULT FALSE,
    enhancement_cost FLOAT DEFAULT 0.0,
    
    -- Item metadata (Siêu dữ liệu vật phẩm)
    item_attributes JSON,
    item_effects JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_item_id (item_id),
    INDEX idx_item_type (item_type),
    INDEX idx_item_rarity (item_rarity),
    INDEX idx_operation (operation),
    INDEX idx_enhancement_level (enhancement_level),
    INDEX idx_created_at (created_at),
    
    -- Foreign key (Khóa ngoại)
    FOREIGN KEY (event_id) REFERENCES event_log(event_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

#### **7. Cultivation_Event_Log - Bảng Log Event Tu Luyện**

```sql
-- Cultivation_Event_Log - Cultivation event logging table (Bảng log event tu luyện)
CREATE TABLE cultivation_event_log (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Event reference (Tham chiếu event)
    event_id VARCHAR(64) NOT NULL UNIQUE,
    
    -- Cultivation details (Chi tiết tu luyện)
    cultivation_type VARCHAR(32) NOT NULL,
    realm VARCHAR(32) NOT NULL,
    substage VARCHAR(32) NOT NULL,
    cultivation_method VARCHAR(32) NOT NULL,
    
    -- Cultivation progress (Tiến độ tu luyện)
    progress_gained FLOAT DEFAULT 0.0,
    experience_gained FLOAT DEFAULT 0.0,
    energy_consumed FLOAT DEFAULT 0.0,
    time_spent_ms BIGINT DEFAULT 0,
    
    -- Cultivation results (Kết quả tu luyện)
    breakthrough BOOLEAN DEFAULT FALSE,
    new_realm VARCHAR(32),
    new_substage VARCHAR(32),
    skills_learned JSON,
    
    -- Cultivation metadata (Siêu dữ liệu tu luyện)
    cultivation_location VARCHAR(128),
    cultivation_environment VARCHAR(32),
    cultivation_conditions JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_cultivation_type (cultivation_type),
    INDEX idx_realm (realm),
    INDEX idx_substage (substage),
    INDEX idx_breakthrough (breakthrough),
    INDEX idx_created_at (created_at),
    
    -- Foreign key (Khóa ngoại)
    FOREIGN KEY (event_id) REFERENCES event_log(event_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## 📊 Analytics & Monitoring Tables (Bảng Phân Tích & Giám Sát)

### **8. Event_Analytics - Bảng Phân Tích Event**

```sql
-- Event_Analytics - Event analytics table (Bảng phân tích event)
CREATE TABLE event_analytics (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Analytics identification (Nhận dạng phân tích)
    analysis_id VARCHAR(64) NOT NULL UNIQUE,
    analysis_type VARCHAR(32) NOT NULL,
    analysis_period VARCHAR(16) NOT NULL,
    
    -- Analytics scope (Phạm vi phân tích)
    world_id VARCHAR(64),
    server_id VARCHAR(32),
    event_type VARCHAR(32),
    event_category VARCHAR(32),
    
    -- Analytics metrics (Chỉ số phân tích)
    total_events BIGINT DEFAULT 0,
    successful_events BIGINT DEFAULT 0,
    failed_events BIGINT DEFAULT 0,
    average_processing_time_ms INT DEFAULT 0,
    total_processing_time_ms BIGINT DEFAULT 0,
    
    -- Analytics results (Kết quả phân tích)
    trends JSON,
    patterns JSON,
    anomalies JSON,
    predictions JSON,
    
    -- Analytics metadata (Siêu dữ liệu phân tích)
    analysis_algorithm VARCHAR(32),
    analysis_parameters JSON,
    confidence_score FLOAT DEFAULT 0.0,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    analysis_start_time DATETIME(6) NOT NULL,
    analysis_end_time DATETIME(6) NOT NULL,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_analysis_type (analysis_type),
    INDEX idx_analysis_period (analysis_period),
    INDEX idx_world_id (world_id),
    INDEX idx_event_type (event_type),
    INDEX idx_analysis_start_time (analysis_start_time),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### **9. Event_Monitoring - Bảng Giám Sát Event**

```sql
-- Event_Monitoring - Event monitoring table (Bảng giám sát event)
CREATE TABLE event_monitoring (
    -- Primary key (Khóa chính)
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    
    -- Monitoring identification (Nhận dạng giám sát)
    monitoring_id VARCHAR(64) NOT NULL UNIQUE,
    monitoring_type VARCHAR(32) NOT NULL,
    monitoring_level VARCHAR(16) NOT NULL,
    
    -- Monitoring scope (Phạm vi giám sát)
    world_id VARCHAR(64),
    server_id VARCHAR(32),
    event_type VARCHAR(32),
    event_category VARCHAR(32),
    
    -- Monitoring metrics (Chỉ số giám sát)
    event_count BIGINT DEFAULT 0,
    error_count BIGINT DEFAULT 0,
    warning_count BIGINT DEFAULT 0,
    success_rate FLOAT DEFAULT 0.0,
    average_response_time_ms INT DEFAULT 0,
    
    -- Monitoring thresholds (Ngưỡng giám sát)
    error_threshold INT DEFAULT 100,
    warning_threshold INT DEFAULT 50,
    response_time_threshold_ms INT DEFAULT 1000,
    
    -- Monitoring status (Trạng thái giám sát)
    status ENUM('normal', 'warning', 'error', 'critical') DEFAULT 'normal',
    alert_sent BOOLEAN DEFAULT FALSE,
    alert_message TEXT,
    
    -- Monitoring metadata (Siêu dữ liệu giám sát)
    monitoring_config JSON,
    metadata JSON,
    
    -- Timestamps (Thời gian)
    monitoring_start_time DATETIME(6) NOT NULL,
    monitoring_end_time DATETIME(6) NOT NULL,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    
    -- Indexes (Chỉ mục)
    INDEX idx_monitoring_type (monitoring_type),
    INDEX idx_monitoring_level (monitoring_level),
    INDEX idx_world_id (world_id),
    INDEX idx_status (status),
    INDEX idx_monitoring_start_time (monitoring_start_time),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

## 🔧 Database Optimization (Tối Ưu Database)

### **Partitioning Strategy (Chiến Lược Phân Vùng)**

```sql
-- Partition event_log by month (Phân vùng event_log theo tháng)
ALTER TABLE event_log PARTITION BY RANGE (YEAR(created_at) * 100 + MONTH(created_at)) (
    PARTITION p202501 VALUES LESS THAN (202502),
    PARTITION p202502 VALUES LESS THAN (202503),
    PARTITION p202503 VALUES LESS THAN (202504),
    -- ... more partitions
    PARTITION p_future VALUES LESS THAN MAXVALUE
);

-- Partition event_causality by month (Phân vùng event_causality theo tháng)
ALTER TABLE event_causality PARTITION BY RANGE (YEAR(created_at) * 100 + MONTH(created_at)) (
    PARTITION p202501 VALUES LESS THAN (202502),
    PARTITION p202502 VALUES LESS THAN (202503),
    PARTITION p202503 VALUES LESS THAN (202504),
    -- ... more partitions
    PARTITION p_future VALUES LESS THAN MAXVALUE
);
```

### **Indexing Strategy (Chiến Lược Lập Chỉ Mục)**

```sql
-- Composite indexes for common queries (Chỉ mục tổ hợp cho truy vấn thường dùng)
CREATE INDEX idx_event_type_timestamp ON event_log(event_type, timestamp);
CREATE INDEX idx_actor_id_timestamp ON event_log(actor_id, timestamp);
CREATE INDEX idx_world_id_timestamp ON event_log(world_id, timestamp);
CREATE INDEX idx_status_timestamp ON event_log(status, timestamp);

-- Full-text search indexes (Chỉ mục tìm kiếm toàn văn)
CREATE FULLTEXT INDEX idx_event_data ON event_log(data);
CREATE FULLTEXT INDEX idx_event_metadata ON event_log(metadata);
```

### **Archiving Strategy (Chiến Lược Lưu Trữ)**

```sql
-- Archive old events (Lưu trữ event cũ)
CREATE TABLE event_log_archive LIKE event_log;
ALTER TABLE event_log_archive ENGINE=ARCHIVE;

-- Move old events to archive (Chuyển event cũ sang archive)
INSERT INTO event_log_archive 
SELECT * FROM event_log 
WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);

-- Delete archived events from main table (Xóa event đã archive khỏi bảng chính)
DELETE FROM event_log 
WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);
```

## 📈 Performance Monitoring (Giám Sát Hiệu Suất)

### **Query Performance (Hiệu Suất Truy Vấn)**

```sql
-- Monitor slow queries (Giám sát truy vấn chậm)
SELECT 
    query_time,
    lock_time,
    rows_sent,
    rows_examined,
    sql_text
FROM mysql.slow_log 
WHERE start_time > DATE_SUB(NOW(), INTERVAL 1 HOUR)
ORDER BY query_time DESC;

-- Monitor table sizes (Giám sát kích thước bảng)
SELECT 
    table_name,
    table_rows,
    data_length,
    index_length,
    (data_length + index_length) as total_size
FROM information_schema.tables 
WHERE table_schema = 'event_core'
ORDER BY total_size DESC;
```

### **Index Usage (Sử Dụng Chỉ Mục)**

```sql
-- Monitor index usage (Giám sát sử dụng chỉ mục)
SELECT 
    object_schema,
    object_name,
    index_name,
    count_read,
    count_write,
    count_read / (count_read + count_write) as read_ratio
FROM performance_schema.table_io_waits_summary_by_index_usage
WHERE object_schema = 'event_core'
ORDER BY count_read DESC;
```

## 🔄 Data Retention & Cleanup (Lưu Trữ & Dọn Dẹp Dữ Liệu)

### **Retention Policies (Chính Sách Lưu Trữ)**

```sql
-- Create retention policy (Tạo chính sách lưu trữ)
CREATE EVENT cleanup_old_events
ON SCHEDULE EVERY 1 DAY
STARTS CURRENT_TIMESTAMP
DO
BEGIN
    -- Archive events older than 1 year (Lưu trữ event cũ hơn 1 năm)
    INSERT INTO event_log_archive 
    SELECT * FROM event_log 
    WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);
    
    -- Delete archived events (Xóa event đã archive)
    DELETE FROM event_log 
    WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);
    
    -- Clean up old analytics (Dọn dẹp phân tích cũ)
    DELETE FROM event_analytics 
    WHERE created_at < DATE_SUB(NOW(), INTERVAL 6 MONTH);
    
    -- Clean up old monitoring data (Dọn dẹp dữ liệu giám sát cũ)
    DELETE FROM event_monitoring 
    WHERE created_at < DATE_SUB(NOW(), INTERVAL 3 MONTH);
END;
```

### **Data Compression (Nén Dữ Liệu)**

```sql
-- Compress old partitions (Nén phân vùng cũ)
ALTER TABLE event_log PARTITION p202401 COMPRESSION='zlib';
ALTER TABLE event_log PARTITION p202402 COMPRESSION='zlib';
ALTER TABLE event_log PARTITION p202403 COMPRESSION='zlib';
```

## 🚀 Implementation Examples (Ví Dụ Triển Khai)

### **Event Logger Implementation**

```go
// EventLogger - Event logging implementation (Triển khai ghi log event)
type EventLogger struct {
    db          *sql.DB
    config      *LoggerConfig
    metrics     *LoggerMetrics
    buffer      *EventBuffer
    compressor  *EventCompressor
}

// LoggerConfig - Logger configuration (Cấu hình logger)
type LoggerConfig struct {
    // Database settings (Cài đặt database)
    DatabaseURL     string        `json:"database_url"`
    MaxConnections  int           `json:"max_connections"`
    ConnectionTimeout time.Duration `json:"connection_timeout"`
    
    // Buffer settings (Cài đặt buffer)
    BufferSize      int           `json:"buffer_size"`
    FlushInterval   time.Duration `json:"flush_interval"`
    FlushThreshold  int           `json:"flush_threshold"`
    
    // Compression settings (Cài đặt nén)
    EnableCompression bool        `json:"enable_compression"`
    CompressionLevel  int         `json:"compression_level"`
    
    // Retention settings (Cài đặt lưu trữ)
    RetentionDays    int         `json:"retention_days"`
    ArchiveAfterDays int         `json:"archive_after_days"`
}

// LogEvent - Log event to database (Ghi log event vào database)
func (el *EventLogger) LogEvent(event IEvent) error {
    // Validate event (Xác thực event)
    if err := event.Validate(); err != nil {
        return fmt.Errorf("event validation failed: %w", err)
    }
    
    // Prepare event data (Chuẩn bị dữ liệu event)
    eventData := &EventLogData{
        EventID:       event.GetEventID(),
        EventType:     event.GetEventType(),
        EventCategory: event.GetEventCategory(),
        ActorID:       event.GetActorID(),
        TargetID:      event.GetTargetID(),
        WorldID:       event.GetWorldID(),
        Timestamp:     event.GetTimestamp(),
        Data:          event.GetData(),
        Metadata:      event.GetMetadata(),
        Status:        EventStatusPending,
        Priority:      event.GetPriority(),
        Weight:        event.GetWeight(),
    }
    
    // Add to buffer (Thêm vào buffer)
    return el.buffer.AddEvent(eventData)
}

// FlushBuffer - Flush buffer to database (Đẩy buffer vào database)
func (el *EventLogger) FlushBuffer() error {
    events := el.buffer.GetEvents()
    if len(events) == 0 {
        return nil
    }
    
    // Prepare batch insert (Chuẩn bị insert theo lô)
    query := `
        INSERT INTO event_log (
            event_id, event_type, event_category, actor_id, target_id, world_id,
            timestamp, data, metadata, status, priority, weight
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
    
    stmt, err := el.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("prepare statement failed: %w", err)
    }
    defer stmt.Close()
    
    // Execute batch insert (Thực thi insert theo lô)
    for _, event := range events {
        _, err := stmt.Exec(
            event.EventID,
            event.EventType,
            event.EventCategory,
            event.ActorID,
            event.TargetID,
            event.WorldID,
            event.Timestamp,
            event.Data,
            event.Metadata,
            event.Status,
            event.Priority,
            event.Weight,
        )
        if err != nil {
            return fmt.Errorf("insert event failed: %w", err)
        }
    }
    
    // Clear buffer (Xóa buffer)
    el.buffer.Clear()
    
    // Update metrics (Cập nhật chỉ số)
    el.metrics.IncrementEventsLogged(len(events))
    
    return nil
}
```

### **Event Analytics Implementation**

```go
// EventAnalytics - Event analytics implementation (Triển khai phân tích event)
type EventAnalytics struct {
    db      *sql.DB
    config  *AnalyticsConfig
    metrics *AnalyticsMetrics
}

// AnalyticsConfig - Analytics configuration (Cấu hình phân tích)
type AnalyticsConfig struct {
    // Analysis settings (Cài đặt phân tích)
    AnalysisInterval time.Duration `json:"analysis_interval"`
    AnalysisDepth    int           `json:"analysis_depth"`
    
    // Algorithm settings (Cài đặt thuật toán)
    TrendAlgorithm   string        `json:"trend_algorithm"`
    PatternAlgorithm string        `json:"pattern_algorithm"`
    AnomalyAlgorithm string        `json:"anomaly_algorithm"`
    
    // Threshold settings (Cài đặt ngưỡng)
    TrendThreshold   float64       `json:"trend_threshold"`
    PatternThreshold float64       `json:"pattern_threshold"`
    AnomalyThreshold float64       `json:"anomaly_threshold"`
}

// AnalyzeEvents - Analyze events for trends and patterns (Phân tích event để tìm xu hướng và mẫu)
func (ea *EventAnalytics) AnalyzeEvents(analysisType string, startTime, endTime time.Time) (*AnalysisResult, error) {
    // Query events in time range (Truy vấn event trong khoảng thời gian)
    events, err := ea.queryEventsInRange(startTime, endTime)
    if err != nil {
        return nil, fmt.Errorf("query events failed: %w", err)
    }
    
    // Perform analysis (Thực hiện phân tích)
    result := &AnalysisResult{
        AnalysisType: analysisType,
        StartTime:    startTime,
        EndTime:      endTime,
        TotalEvents:  len(events),
    }
    
    // Analyze trends (Phân tích xu hướng)
    if trends, err := ea.analyzeTrends(events); err == nil {
        result.Trends = trends
    }
    
    // Analyze patterns (Phân tích mẫu)
    if patterns, err := ea.analyzePatterns(events); err == nil {
        result.Patterns = patterns
    }
    
    // Analyze anomalies (Phân tích bất thường)
    if anomalies, err := ea.analyzeAnomalies(events); err == nil {
        result.Anomalies = anomalies
    }
    
    // Save analysis result (Lưu kết quả phân tích)
    if err := ea.saveAnalysisResult(result); err != nil {
        return nil, fmt.Errorf("save analysis result failed: %w", err)
    }
    
    return result, nil
}
```

## 💡 Best Practices (Thực Hành Tốt Nhất)

### **Database Design (Thiết Kế Database)**
1. **Normalize Data** - Chuẩn hóa dữ liệu
2. **Use Appropriate Indexes** - Sử dụng chỉ mục phù hợp
3. **Partition Large Tables** - Phân vùng bảng lớn
4. **Archive Old Data** - Lưu trữ dữ liệu cũ

### **Performance Optimization (Tối Ưu Hiệu Suất)**
1. **Batch Operations** - Thao tác theo lô
2. **Connection Pooling** - Pool kết nối
3. **Query Optimization** - Tối ưu truy vấn
4. **Caching** - Cache dữ liệu

### **Data Retention (Lưu Trữ Dữ Liệu)**
1. **Define Retention Policies** - Định nghĩa chính sách lưu trữ
2. **Automate Cleanup** - Tự động dọn dẹp
3. **Compress Old Data** - Nén dữ liệu cũ
4. **Monitor Storage Usage** - Giám sát sử dụng lưu trữ
