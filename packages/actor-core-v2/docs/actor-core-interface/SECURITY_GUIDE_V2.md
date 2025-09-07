# Security Guide - Actor Core v2.0

## Tổng Quan

Security Guide cho Actor Core v2.0 cung cấp hướng dẫn chi tiết về bảo mật, bao gồm input validation, access control, data protection, và security best practices.

## 1. Security Threats & Mitigations

### 1.1 Input Validation Threats

#### 1.1.1 Stat Value Injection
**Threat**: Malicious stat values that could cause buffer overflow, integer overflow, or other security issues.

**Mitigation**:
```go
// ✅ GOOD: Input validation with bounds checking
func (ac *ActorCore) SetStat(statName string, value int64) error {
    // Validate stat name
    if !ac.isValidStatName(statName) {
        return ErrInvalidStatName
    }
    
    // Validate value bounds
    if err := ac.validateStatValue(statName, value); err != nil {
        return err
    }
    
    // Set stat
    ac.PrimaryStats[statName] = value
    return nil
}

func (ac *ActorCore) validateStatValue(statName string, value int64) error {
    statDef, exists := ac.statDefinitions[statName]
    if !exists {
        return ErrStatNotFound
    }
    
    // Check bounds
    if value < statDef.MinValue || value > statDef.MaxValue {
        return ErrValueOutOfBounds
    }
    
    // Check for suspicious values
    if value < 0 && statDef.MinValue >= 0 {
        return ErrNegativeValueNotAllowed
    }
    
    return nil
}
```

#### 1.1.2 Formula Injection
**Threat**: Malicious formulas that could execute arbitrary code or cause system crashes.

**Mitigation**:
```go
// ✅ GOOD: Formula validation and sandboxing
func (engine *FormulaEngine) ValidateFormula(formula string) error {
    // Check for dangerous functions
    dangerousFunctions := []string{
        "exec", "system", "eval", "import", "require",
        "os.", "sys.", "subprocess", "file", "open",
    }
    
    for _, dangerous := range dangerousFunctions {
        if strings.Contains(formula, dangerous) {
            return ErrDangerousFunction
        }
    }
    
    // Check for suspicious patterns
    suspiciousPatterns := []string{
        "`", "\\", "..", "//", "/*", "*/",
        "javascript:", "data:", "vbscript:",
    }
    
    for _, pattern := range suspiciousPatterns {
        if strings.Contains(formula, pattern) {
            return ErrSuspiciousPattern
        }
    }
    
    // Validate syntax
    if err := engine.validateFormulaSyntax(formula); err != nil {
        return err
    }
    
    return nil
}
```

### 1.2 Access Control Threats

#### 1.2.1 Unauthorized Stat Access
**Threat**: Unauthorized access to sensitive stats or configuration.

**Mitigation**:
```go
// ✅ GOOD: Role-based access control
type AccessControl struct {
    roles        map[string]*Role
    permissions  map[string][]string
    userRoles    map[string][]string
}

type Role struct {
    Name        string
    Permissions []string
    Level       int
}

func (ac *AccessControl) CanAccessStat(userID, statName string) bool {
    userRoles := ac.userRoles[userID]
    for _, roleName := range userRoles {
        role := ac.roles[roleName]
        for _, permission := range role.Permissions {
            if permission == "stat:"+statName || permission == "stat:*" {
                return true
            }
        }
    }
    return false
}

func (ac *ActorCore) GetStat(userID, statName string) (interface{}, error) {
    if !ac.accessControl.CanAccessStat(userID, statName) {
        return nil, ErrAccessDenied
    }
    
    return ac.getStatValue(statName)
}
```

#### 1.2.2 Configuration Tampering
**Threat**: Unauthorized modification of configuration files or settings.

**Mitigation**:
```go
// ✅ GOOD: Configuration integrity checking
type SecureConfigurationManager struct {
    configs        map[string]*ConfigDefinition
    checksums      map[string]string
    signatures     map[string]string
    publicKey      *rsa.PublicKey
    privateKey     *rsa.PrivateKey
}

func (scm *SecureConfigurationManager) LoadConfig(filename string) error {
    // Read configuration file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    
    // Verify checksum
    checksum := sha256.Sum256(data)
    if scm.checksums[filename] != hex.EncodeToString(checksum[:]) {
        return ErrChecksumMismatch
    }
    
    // Verify signature
    if err := scm.verifySignature(filename, data); err != nil {
        return err
    }
    
    // Parse configuration
    config := &ConfigDefinition{}
    if err := json.Unmarshal(data, config); err != nil {
        return err
    }
    
    scm.configs[filename] = config
    return nil
}

func (scm *SecureConfigurationManager) SaveConfig(filename string, config *ConfigDefinition) error {
    // Serialize configuration
    data, err := json.Marshal(config)
    if err != nil {
        return err
    }
    
    // Generate checksum
    checksum := sha256.Sum256(data)
    scm.checksums[filename] = hex.EncodeToString(checksum[:])
    
    // Generate signature
    signature, err := scm.signData(data)
    if err != nil {
        return err
    }
    scm.signatures[filename] = signature
    
    // Write file
    return ioutil.WriteFile(filename, data, 0644)
}
```

### 1.3 Data Protection Threats

#### 1.3.1 Sensitive Data Exposure
**Threat**: Exposure of sensitive data in logs, memory dumps, or network traffic.

**Mitigation**:
```go
// ✅ GOOD: Data encryption and masking
type SecureDataManager struct {
    encryptionKey []byte
    sensitiveFields map[string]bool
}

func (sdm *SecureDataManager) EncryptSensitiveData(data map[string]interface{}) map[string]interface{} {
    encrypted := make(map[string]interface{})
    
    for key, value := range data {
        if sdm.sensitiveFields[key] {
            // Encrypt sensitive fields
            encryptedValue, err := sdm.encrypt(value)
            if err != nil {
                // Log error but continue
                continue
            }
            encrypted[key] = encryptedValue
        } else {
            encrypted[key] = value
        }
    }
    
    return encrypted
}

func (sdm *SecureDataManager) MaskSensitiveData(data map[string]interface{}) map[string]interface{} {
    masked := make(map[string]interface{})
    
    for key, value := range data {
        if sdm.sensitiveFields[key] {
            // Mask sensitive fields
            masked[key] = sdm.mask(value)
        } else {
            masked[key] = value
        }
    }
    
    return masked
}

func (sdm *SecureDataManager) mask(value interface{}) string {
    switch v := value.(type) {
    case string:
        if len(v) <= 4 {
            return "****"
        }
        return v[:2] + "****" + v[len(v)-2:]
    case int64:
        return "****"
    default:
        return "****"
    }
}
```

#### 1.3.2 Memory Dump Protection
**Threat**: Sensitive data exposed in memory dumps or core files.

**Mitigation**:
```go
// ✅ GOOD: Secure memory management
type SecureMemoryManager struct {
    sensitiveData map[string][]byte
    mutex         sync.RWMutex
}

func (smm *SecureMemoryManager) StoreSensitiveData(key string, data []byte) {
    smm.mutex.Lock()
    defer smm.mutex.Unlock()
    
    // Encrypt data before storing
    encrypted, err := smm.encrypt(data)
    if err != nil {
        return
    }
    
    smm.sensitiveData[key] = encrypted
}

func (smm *SecureMemoryManager) GetSensitiveData(key string) ([]byte, error) {
    smm.mutex.RLock()
    defer smm.mutex.RUnlock()
    
    encrypted, exists := smm.sensitiveData[key]
    if !exists {
        return nil, ErrDataNotFound
    }
    
    // Decrypt data
    return smm.decrypt(encrypted)
}

func (smm *SecureMemoryManager) ClearSensitiveData(key string) {
    smm.mutex.Lock()
    defer smm.mutex.Unlock()
    
    if data, exists := smm.sensitiveData[key]; exists {
        // Overwrite with random data
        rand.Read(data)
        delete(smm.sensitiveData, key)
    }
}
```

## 2. Security Best Practices

### 2.1 Input Validation

#### 2.1.1 Stat Name Validation
```go
// ✅ GOOD: Comprehensive stat name validation
func (ac *ActorCore) isValidStatName(statName string) bool {
    // Check length
    if len(statName) == 0 || len(statName) > 100 {
        return false
    }
    
    // Check characters (only alphanumeric and underscore)
    for _, char := range statName {
        if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
            return false
        }
    }
    
    // Check for reserved names
    reservedNames := []string{
        "system", "admin", "root", "config", "secret",
        "password", "token", "key", "private",
    }
    
    for _, reserved := range reservedNames {
        if strings.EqualFold(statName, reserved) {
            return false
        }
    }
    
    return true
}
```

#### 2.1.2 Value Range Validation
```go
// ✅ GOOD: Strict value range validation
func (ac *ActorCore) validateValueRange(statName string, value int64) error {
    statDef, exists := ac.statDefinitions[statName]
    if !exists {
        return ErrStatNotFound
    }
    
    // Check absolute bounds
    if value < statDef.MinValue || value > statDef.MaxValue {
        return ErrValueOutOfBounds
    }
    
    // Check for suspicious values
    if value < 0 && statDef.MinValue >= 0 {
        return ErrNegativeValueNotAllowed
    }
    
    // Check for overflow/underflow
    if value > math.MaxInt32 || value < math.MinInt32 {
        return ErrValueTooLarge
    }
    
    return nil
}
```

### 2.2 Authentication & Authorization

#### 2.2.1 User Authentication
```go
// ✅ GOOD: Secure user authentication
type UserAuthenticator struct {
    users        map[string]*User
    sessions     map[string]*Session
    passwordHasher *bcrypt.Hasher
    tokenManager *TokenManager
}

type User struct {
    ID           string
    Username     string
    PasswordHash string
    Roles        []string
    CreatedAt    time.Time
    LastLogin    time.Time
    IsActive     bool
}

func (ua *UserAuthenticator) Authenticate(username, password string) (*Session, error) {
    user, exists := ua.users[username]
    if !exists {
        return nil, ErrUserNotFound
    }
    
    if !user.IsActive {
        return nil, ErrUserInactive
    }
    
    // Verify password
    if err := ua.passwordHasher.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, ErrInvalidPassword
    }
    
    // Create session
    session := &Session{
        ID:        generateSessionID(),
        UserID:    user.ID,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }
    
    ua.sessions[session.ID] = session
    user.LastLogin = time.Now()
    
    return session, nil
}
```

#### 2.2.2 Role-Based Access Control
```go
// ✅ GOOD: Comprehensive RBAC system
type RBACManager struct {
    roles        map[string]*Role
    permissions  map[string]*Permission
    userRoles    map[string][]string
    rolePermissions map[string][]string
}

type Role struct {
    ID          string
    Name        string
    Description string
    Level       int
    Permissions []string
}

type Permission struct {
    ID          string
    Name        string
    Resource    string
    Action      string
    Conditions  map[string]interface{}
}

func (rbac *RBACManager) CanAccess(userID, resource, action string) bool {
    userRoles := rbac.userRoles[userID]
    for _, roleName := range userRoles {
        role := rbac.roles[roleName]
        for _, permissionID := range role.Permissions {
            permission := rbac.permissions[permissionID]
            if permission.Resource == resource && permission.Action == action {
                return true
            }
        }
    }
    return false
}
```

### 2.3 Data Encryption

#### 2.3.1 At-Rest Encryption
```go
// ✅ GOOD: Data encryption at rest
type DataEncryption struct {
    encryptionKey []byte
    cipher        cipher.AEAD
}

func (de *DataEncryption) EncryptData(data []byte) ([]byte, error) {
    // Generate random nonce
    nonce := make([]byte, 12)
    if _, err := rand.Read(nonce); err != nil {
        return nil, err
    }
    
    // Encrypt data
    encrypted := de.cipher.Seal(nil, nonce, data, nil)
    
    // Prepend nonce to encrypted data
    result := make([]byte, len(nonce)+len(encrypted))
    copy(result, nonce)
    copy(result[len(nonce):], encrypted)
    
    return result, nil
}

func (de *DataEncryption) DecryptData(encryptedData []byte) ([]byte, error) {
    if len(encryptedData) < 12 {
        return nil, ErrInvalidEncryptedData
    }
    
    // Extract nonce
    nonce := encryptedData[:12]
    encrypted := encryptedData[12:]
    
    // Decrypt data
    decrypted, err := de.cipher.Open(nil, nonce, encrypted, nil)
    if err != nil {
        return nil, err
    }
    
    return decrypted, nil
}
```

#### 2.3.2 In-Transit Encryption
```go
// ✅ GOOD: Data encryption in transit
type TransportEncryption struct {
    tlsConfig *tls.Config
    cert      *x509.Certificate
    key       *rsa.PrivateKey
}

func (te *TransportEncryption) EncryptTransport(conn net.Conn) (net.Conn, error) {
    return tls.Client(conn, te.tlsConfig), nil
}

func (te *TransportEncryption) DecryptTransport(conn net.Conn) (net.Conn, error) {
    return tls.Server(conn, te.tlsConfig), nil
}
```

### 2.4 Audit Logging

#### 2.4.1 Security Event Logging
```go
// ✅ GOOD: Comprehensive security event logging
type SecurityLogger struct {
    logger    *log.Logger
    mutex     sync.Mutex
    logLevel  LogLevel
}

type SecurityEvent struct {
    Timestamp   time.Time
    EventType   string
    UserID      string
    Resource    string
    Action      string
    Result      string
    IPAddress   string
    UserAgent   string
    Details     map[string]interface{}
}

func (sl *SecurityLogger) LogSecurityEvent(event *SecurityEvent) {
    sl.mutex.Lock()
    defer sl.mutex.Unlock()
    
    // Log security event
    sl.logger.Printf("[SECURITY] %s %s %s %s %s %s %s %v",
        event.Timestamp.Format(time.RFC3339),
        event.EventType,
        event.UserID,
        event.Resource,
        event.Action,
        event.Result,
        event.IPAddress,
        event.Details,
    )
    
    // Check for suspicious patterns
    sl.checkSuspiciousActivity(event)
}

func (sl *SecurityLogger) checkSuspiciousActivity(event *SecurityEvent) {
    // Check for brute force attacks
    if event.EventType == "AUTHENTICATION_FAILED" {
        sl.checkBruteForceAttack(event.UserID, event.IPAddress)
    }
    
    // Check for privilege escalation attempts
    if event.Action == "ELEVATE_PRIVILEGES" && event.Result == "FAILED" {
        sl.logger.Printf("[SECURITY_ALERT] Privilege escalation attempt by user %s", event.UserID)
    }
    
    // Check for data access patterns
    if event.Action == "ACCESS_DATA" {
        sl.checkDataAccessPattern(event)
    }
}
```

### 2.5 Rate Limiting

#### 2.5.1 API Rate Limiting
```go
// ✅ GOOD: Rate limiting for API calls
type RateLimiter struct {
    limits      map[string]*RateLimit
    mutex       sync.RWMutex
    cleanupTicker *time.Ticker
}

type RateLimit struct {
    MaxRequests int
    Window      time.Duration
    Requests    []time.Time
}

func (rl *RateLimiter) IsAllowed(userID string, limitName string) bool {
    rl.mutex.Lock()
    defer rl.mutex.Unlock()
    
    limit, exists := rl.limits[limitName]
    if !exists {
        return true
    }
    
    now := time.Now()
    cutoff := now.Add(-limit.Window)
    
    // Remove old requests
    var validRequests []time.Time
    for _, reqTime := range limit.Requests {
        if reqTime.After(cutoff) {
            validRequests = append(validRequests, reqTime)
        }
    }
    limit.Requests = validRequests
    
    // Check if limit exceeded
    if len(limit.Requests) >= limit.MaxRequests {
        return false
    }
    
    // Add current request
    limit.Requests = append(limit.Requests, now)
    
    return true
}
```

## 3. Security Testing

### 3.1 Input Validation Tests
```go
// ✅ GOOD: Comprehensive input validation tests
func TestInputValidation(t *testing.T) {
    ac := NewActorCore()
    
    // Test valid inputs
    validInputs := []struct {
        statName string
        value    int64
    }{
        {"vitality", 100},
        {"strength", 50},
        {"intelligence", 75},
    }
    
    for _, input := range validInputs {
        err := ac.SetStat(input.statName, input.value)
        if err != nil {
            t.Errorf("Valid input rejected: %v", err)
        }
    }
    
    // Test invalid inputs
    invalidInputs := []struct {
        statName string
        value    int64
        expectedError error
    }{
        {"", 100, ErrInvalidStatName},
        {"vitality", -1, ErrNegativeValueNotAllowed},
        {"vitality", 999999999, ErrValueOutOfBounds},
        {"system", 100, ErrInvalidStatName},
        {"admin", 100, ErrInvalidStatName},
    }
    
    for _, input := range invalidInputs {
        err := ac.SetStat(input.statName, input.value)
        if err != input.expectedError {
            t.Errorf("Invalid input not properly rejected: %v", err)
        }
    }
}
```

### 3.2 Security Penetration Tests
```go
// ✅ GOOD: Security penetration tests
func TestSecurityPenetration(t *testing.T) {
    ac := NewActorCore()
    
    // Test SQL injection attempts
    sqlInjectionAttempts := []string{
        "'; DROP TABLE stats; --",
        "1' OR '1'='1",
        "admin'--",
        "1' UNION SELECT * FROM users--",
    }
    
    for _, attempt := range sqlInjectionAttempts {
        err := ac.SetStat(attempt, 100)
        if err == nil {
            t.Errorf("SQL injection attempt succeeded: %s", attempt)
        }
    }
    
    // Test XSS attempts
    xssAttempts := []string{
        "<script>alert('xss')</script>",
        "javascript:alert('xss')",
        "<img src=x onerror=alert('xss')>",
    }
    
    for _, attempt := range xssAttempts {
        err := ac.SetStat(attempt, 100)
        if err == nil {
            t.Errorf("XSS attempt succeeded: %s", attempt)
        }
    }
    
    // Test command injection attempts
    commandInjectionAttempts := []string{
        "stat; rm -rf /",
        "stat | cat /etc/passwd",
        "stat && whoami",
    }
    
    for _, attempt := range commandInjectionAttempts {
        err := ac.SetStat(attempt, 100)
        if err == nil {
            t.Errorf("Command injection attempt succeeded: %s", attempt)
        }
    }
}
```

## 4. Security Monitoring

### 4.1 Real-time Security Monitoring
```go
// ✅ GOOD: Real-time security monitoring
type SecurityMonitor struct {
    events      chan *SecurityEvent
    alerts      chan *SecurityAlert
    rules       []SecurityRule
    mutex       sync.RWMutex
}

type SecurityRule struct {
    Name        string
    Condition   func(*SecurityEvent) bool
    Action      func(*SecurityEvent)
    Severity    string
}

func (sm *SecurityMonitor) Start() {
    go sm.processEvents()
    go sm.processAlerts()
}

func (sm *SecurityMonitor) processEvents() {
    for event := range sm.events {
        sm.mutex.RLock()
        for _, rule := range sm.rules {
            if rule.Condition(event) {
                rule.Action(event)
                sm.alerts <- &SecurityAlert{
                    Rule:      rule,
                    Event:     event,
                    Timestamp: time.Now(),
                }
            }
        }
        sm.mutex.RUnlock()
    }
}
```

### 4.2 Security Metrics
```go
// ✅ GOOD: Security metrics collection
type SecurityMetrics struct {
    FailedLogins      int64
    SuccessfulLogins  int64
    PrivilegeEscalations int64
    DataAccessAttempts int64
    SuspiciousActivity int64
    LastReset         time.Time
}

func (sm *SecurityMetrics) RecordFailedLogin() {
    atomic.AddInt64(&sm.FailedLogins, 1)
}

func (sm *SecurityMetrics) RecordSuccessfulLogin() {
    atomic.AddInt64(&sm.SuccessfulLogins, 1)
}

func (sm *SecurityMetrics) GetFailureRate() float64 {
    total := sm.FailedLogins + sm.SuccessfulLogins
    if total == 0 {
        return 0
    }
    return float64(sm.FailedLogins) / float64(total)
}
```

## 5. Security Configuration

### 5.1 Security Settings
```go
// ✅ GOOD: Security configuration
type SecurityConfig struct {
    // Authentication settings
    PasswordMinLength    int
    PasswordMaxLength    int
    PasswordRequireSpecial bool
    PasswordRequireNumber bool
    PasswordRequireUpper bool
    PasswordRequireLower bool
    MaxLoginAttempts     int
    LockoutDuration      time.Duration
    
    // Session settings
    SessionTimeout       time.Duration
    SessionMaxAge        time.Duration
    SessionSecure        bool
    SessionHttpOnly      bool
    
    // Encryption settings
    EncryptionKey        []byte
    EncryptionAlgorithm  string
    KeyRotationInterval  time.Duration
    
    // Rate limiting
    RateLimitEnabled     bool
    RateLimitRequests    int
    RateLimitWindow      time.Duration
    
    // Logging
    LogLevel             string
    LogSecurityEvents    bool
    LogRetentionDays     int
}
```

### 5.2 Security Policies
```go
// ✅ GOOD: Security policy enforcement
type SecurityPolicy struct {
    Name        string
    Description string
    Rules       []PolicyRule
    Enforcement EnforcementLevel
}

type PolicyRule struct {
    Name        string
    Condition   func(*SecurityEvent) bool
    Action      PolicyAction
    Severity    string
}

type PolicyAction int
const (
    ACTION_ALLOW PolicyAction = iota
    ACTION_DENY
    ACTION_ALERT
    ACTION_BLOCK
    ACTION_QUARANTINE
)

func (sp *SecurityPolicy) Evaluate(event *SecurityEvent) PolicyAction {
    for _, rule := range sp.Rules {
        if rule.Condition(event) {
            return rule.Action
        }
    }
    return ACTION_ALLOW
}
```

---

*Tài liệu này cung cấp hướng dẫn chi tiết về bảo mật cho Actor Core v2.0, đảm bảo hệ thống được bảo vệ khỏi các mối đe dọa bảo mật phổ biến.*
