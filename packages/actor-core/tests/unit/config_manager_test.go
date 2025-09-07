package unit

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"actor-core/services/configuration"
)

func TestNewConfigManager(t *testing.T) {
	cm := configuration.NewConfigManager()

	if cm == nil {
		t.Error("Expected ConfigManager to be created")
	}

	if cm.GetConfigCount() != 0 {
		t.Errorf("Expected config count to be 0, got %d", cm.GetConfigCount())
	}

	if cm.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", cm.GetVersion())
	}

	if cm.GetCreatedAt() == 0 {
		t.Error("Expected CreatedAt to be set")
	}

	if cm.GetUpdatedAt() == 0 {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestNewConfigManagerWithFile(t *testing.T) {
	filePath := "test_config.json"
	cm := configuration.NewConfigManagerWithFile(filePath)

	if cm == nil {
		t.Error("Expected ConfigManager to be created")
	}

	if cm.GetFilePath() != filePath {
		t.Errorf("Expected file path to be %s, got %s", filePath, cm.GetFilePath())
	}
}

func TestSetConfig(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test setting a string config
	err := cm.SetConfig("test_string", "hello world")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test setting an int config
	err = cm.SetConfig("test_int", 42)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test setting a float config
	err = cm.SetConfig("test_float", 3.14)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test setting a bool config
	err = cm.SetConfig("test_bool", true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test setting an empty key
	err = cm.SetConfig("", "value")
	if err == nil {
		t.Error("Expected error for empty key")
	}

	// Test that version was incremented
	if cm.GetVersion() != 5 {
		t.Errorf("Expected version to be 5, got %d", cm.GetVersion())
	}
}

func TestGetConfig(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfig("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing config
	cm.SetConfig("test_string", "hello world")
	value, err := cm.GetConfig("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "hello world" {
		t.Errorf("Expected value to be 'hello world', got %v", value)
	}
}

func TestGetConfigString(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigString("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing string config
	cm.SetConfig("test_string", "hello world")
	value, err := cm.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != "hello world" {
		t.Errorf("Expected value to be 'hello world', got %s", value)
	}

	// Test getting non-string config
	cm.SetConfig("test_int", 42)
	_, err = cm.GetConfigString("test_int")
	if err == nil {
		t.Error("Expected error for non-string config")
	}
}

func TestGetConfigInt(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigInt("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing int config
	cm.SetConfig("test_int", 42)
	value, err := cm.GetConfigInt("test_int")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 42 {
		t.Errorf("Expected value to be 42, got %d", value)
	}

	// Test getting int64 config
	cm.SetConfig("test_int64", int64(100))
	value, err = cm.GetConfigInt("test_int64")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 100 {
		t.Errorf("Expected value to be 100, got %d", value)
	}

	// Test getting float64 config
	cm.SetConfig("test_float", 3.14)
	value, err = cm.GetConfigInt("test_float")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 3 {
		t.Errorf("Expected value to be 3, got %d", value)
	}

	// Test getting non-number config
	cm.SetConfig("test_string", "hello")
	_, err = cm.GetConfigInt("test_string")
	if err == nil {
		t.Error("Expected error for non-number config")
	}
}

func TestGetConfigInt64(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigInt64("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing int64 config
	cm.SetConfig("test_int64", int64(100))
	value, err := cm.GetConfigInt64("test_int64")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 100 {
		t.Errorf("Expected value to be 100, got %d", value)
	}

	// Test getting int config
	cm.SetConfig("test_int", 42)
	value, err = cm.GetConfigInt64("test_int")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 42 {
		t.Errorf("Expected value to be 42, got %d", value)
	}

	// Test getting float64 config
	cm.SetConfig("test_float", 3.14)
	value, err = cm.GetConfigInt64("test_float")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 3 {
		t.Errorf("Expected value to be 3, got %d", value)
	}
}

func TestGetConfigFloat64(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigFloat64("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing float64 config
	cm.SetConfig("test_float", 3.14)
	value, err := cm.GetConfigFloat64("test_float")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 3.14 {
		t.Errorf("Expected value to be 3.14, got %f", value)
	}

	// Test getting int config
	cm.SetConfig("test_int", 42)
	value, err = cm.GetConfigFloat64("test_int")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 42.0 {
		t.Errorf("Expected value to be 42.0, got %f", value)
	}

	// Test getting int64 config
	cm.SetConfig("test_int64", int64(100))
	value, err = cm.GetConfigFloat64("test_int64")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != 100.0 {
		t.Errorf("Expected value to be 100.0, got %f", value)
	}
}

func TestGetConfigBool(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigBool("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing bool config
	cm.SetConfig("test_bool", true)
	value, err := cm.GetConfigBool("test_bool")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value != true {
		t.Errorf("Expected value to be true, got %t", value)
	}

	// Test getting non-bool config
	cm.SetConfig("test_string", "hello")
	_, err = cm.GetConfigBool("test_string")
	if err == nil {
		t.Error("Expected error for non-bool config")
	}
}

func TestGetConfigMap(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigMap("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing map config
	mapValue := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}
	cm.SetConfig("test_map", mapValue)
	value, err := cm.GetConfigMap("test_map")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(value) != 2 {
		t.Errorf("Expected map length to be 2, got %d", len(value))
	}

	if value["key1"] != "value1" {
		t.Errorf("Expected key1 to be 'value1', got %v", value["key1"])
	}

	if value["key2"] != 42 {
		t.Errorf("Expected key2 to be 42, got %v", value["key2"])
	}

	// Test getting non-map config
	cm.SetConfig("test_string", "hello")
	_, err = cm.GetConfigMap("test_string")
	if err == nil {
		t.Error("Expected error for non-map config")
	}
}

func TestGetConfigSlice(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test getting non-existent config
	_, err := cm.GetConfigSlice("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test getting existing slice config
	sliceValue := []interface{}{"item1", "item2", 42}
	cm.SetConfig("test_slice", sliceValue)
	value, err := cm.GetConfigSlice("test_slice")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(value) != 3 {
		t.Errorf("Expected slice length to be 3, got %d", len(value))
	}

	if value[0] != "item1" {
		t.Errorf("Expected item1 to be 'item1', got %v", value[0])
	}

	if value[1] != "item2" {
		t.Errorf("Expected item2 to be 'item2', got %v", value[1])
	}

	if value[2] != 42 {
		t.Errorf("Expected item3 to be 42, got %v", value[2])
	}

	// Test getting non-slice config
	cm.SetConfig("test_string", "hello")
	_, err = cm.GetConfigSlice("test_string")
	if err == nil {
		t.Error("Expected error for non-slice config")
	}
}

func TestHasConfig(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test non-existent config
	if cm.HasConfig("non_existent") {
		t.Error("Expected config to not exist")
	}

	// Test existing config
	cm.SetConfig("test_string", "hello world")
	if !cm.HasConfig("test_string") {
		t.Error("Expected config to exist")
	}
}

func TestRemoveConfig(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test removing non-existent config
	err := cm.RemoveConfig("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent config")
	}

	// Test removing existing config
	cm.SetConfig("test_string", "hello world")
	err = cm.RemoveConfig("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if cm.HasConfig("test_string") {
		t.Error("Expected config to be removed")
	}

	// Test that version was incremented
	if cm.GetVersion() != 3 {
		t.Errorf("Expected version to be 3, got %d", cm.GetVersion())
	}
}

func TestGetAllConfigs(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test empty configs
	configs := cm.GetAllConfigs()
	if len(configs) != 0 {
		t.Errorf("Expected 0 configs, got %d", len(configs))
	}

	// Test with configs
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	configs = cm.GetAllConfigs()
	if len(configs) != 2 {
		t.Errorf("Expected 2 configs, got %d", len(configs))
	}

	if configs["test_string"] != "hello world" {
		t.Errorf("Expected test_string to be 'hello world', got %v", configs["test_string"])
	}

	if configs["test_int"] != 42 {
		t.Errorf("Expected test_int to be 42, got %v", configs["test_int"])
	}
}

func TestGetConfigKeys(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test empty configs
	keys := cm.GetConfigKeys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}

	// Test with configs
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	keys = cm.GetConfigKeys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}

	// Check that both keys are present
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	if !keyMap["test_string"] {
		t.Error("Expected test_string key to be present")
	}

	if !keyMap["test_int"] {
		t.Error("Expected test_int key to be present")
	}
}

func TestGetConfigCount(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test empty configs
	count := cm.GetConfigCount()
	if count != 0 {
		t.Errorf("Expected 0 configs, got %d", count)
	}

	// Test with configs
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	count = cm.GetConfigCount()
	if count != 2 {
		t.Errorf("Expected 2 configs, got %d", count)
	}
}

func TestClearConfigs(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test clearing empty configs
	cm.ClearConfigs()

	// Test clearing with configs
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)
	cm.ClearConfigs()

	if cm.GetConfigCount() != 0 {
		t.Error("Expected configs to be cleared")
	}

	// Test that version was incremented (2 calls to SetConfig + 1 call to ClearConfigs = 3 increments)
	if cm.GetVersion() != 4 {
		t.Errorf("Expected version to be 4, got %d", cm.GetVersion())
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create a temporary config file
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "test_config.json")

	// Create test config data
	configData := map[string]interface{}{
		"test_string": "hello world",
		"test_int":    42,
		"test_float":  3.14,
		"test_bool":   true,
	}

	// Write config file
	configFileData := configuration.ConfigFile{
		Version:   1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Configs:   make(map[string]configuration.ConfigData),
	}

	for key, value := range configData {
		configFileData.Configs[key] = configuration.ConfigData{
			Key:       key,
			Value:     value,
			Type:      "string",
			Category:  "general",
			UpdatedAt: time.Now().Unix(),
		}
	}

	data, err := json.MarshalIndent(configFileData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test loading from file
	cm := configuration.NewConfigManager()
	err = cm.LoadFromFile(configFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify loaded configs
	if cm.GetConfigCount() != 4 {
		t.Errorf("Expected 4 configs, got %d", cm.GetConfigCount())
	}

	value, err := cm.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "hello world" {
		t.Errorf("Expected 'hello world', got %s", value)
	}

	// Test loading non-existent file
	err = cm.LoadFromFile("non_existent.json")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestSaveToFile(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()
	configFile := filepath.Join(tempDir, "test_config.json")

	// Test saving to file
	cm := configuration.NewConfigManager()
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	err := cm.SaveToFile(configFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Error("Expected config file to be created")
	}

	// Test loading the saved file
	cm2 := configuration.NewConfigManager()
	err = cm2.LoadFromFile(configFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify loaded configs
	if cm2.GetConfigCount() != 2 {
		t.Errorf("Expected 2 configs, got %d", cm2.GetConfigCount())
	}

	value, err := cm2.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "hello world" {
		t.Errorf("Expected 'hello world', got %s", value)
	}
}

func TestSetFilePath(t *testing.T) {
	cm := configuration.NewConfigManager()

	filePath := "test_config.json"
	cm.SetFilePath(filePath)

	if cm.GetFilePath() != filePath {
		t.Errorf("Expected file path to be %s, got %s", filePath, cm.GetFilePath())
	}
}

func TestGetFilePath(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test default file path
	if cm.GetFilePath() != "" {
		t.Errorf("Expected empty file path, got %s", cm.GetFilePath())
	}

	// Test setting file path
	filePath := "test_config.json"
	cm.SetFilePath(filePath)

	if cm.GetFilePath() != filePath {
		t.Errorf("Expected file path to be %s, got %s", filePath, cm.GetFilePath())
	}
}

func TestConfigManagerGetVersion(t *testing.T) {
	cm := configuration.NewConfigManager()

	if cm.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", cm.GetVersion())
	}

	cm.SetConfig("test", "value")

	if cm.GetVersion() != 2 {
		t.Errorf("Expected version to be 2, got %d", cm.GetVersion())
	}
}

func TestConfigManagerGetUpdatedAt(t *testing.T) {
	cm := configuration.NewConfigManager()
	originalUpdatedAt := cm.GetUpdatedAt()

	// Wait a bit to ensure timestamp changes
	time.Sleep(1 * time.Second)

	cm.SetConfig("test", "value")

	if cm.GetUpdatedAt() <= originalUpdatedAt {
		t.Errorf("Expected UpdatedAt to be updated. Original: %d, New: %d", originalUpdatedAt, cm.GetUpdatedAt())
	}
}

func TestConfigManagerGetCreatedAt(t *testing.T) {
	cm := configuration.NewConfigManager()

	if cm.GetCreatedAt() == 0 {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestValidateConfig(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test valid configs
	err := cm.ValidateConfig("max_connections", 100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = cm.ValidateConfig("timeout", 5.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = cm.ValidateConfig("debug", true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test invalid configs
	err = cm.ValidateConfig("", "value")
	if err == nil {
		t.Error("Expected error for empty key")
	}

	err = cm.ValidateConfig("max_connections", -1)
	if err == nil {
		t.Error("Expected error for negative max_connections")
	}

	err = cm.ValidateConfig("timeout", -1.0)
	if err == nil {
		t.Error("Expected error for negative timeout")
	}

	err = cm.ValidateConfig("debug", "true")
	if err == nil {
		t.Error("Expected error for non-boolean debug")
	}
}

func TestSetConfigWithValidation(t *testing.T) {
	cm := configuration.NewConfigManager()

	// Test valid config
	err := cm.SetConfigWithValidation("max_connections", 100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test invalid config
	err = cm.SetConfigWithValidation("max_connections", -1)
	if err == nil {
		t.Error("Expected error for invalid config")
	}
}

func TestConfigManagerClone(t *testing.T) {
	cm := configuration.NewConfigManager()
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	clone := cm.Clone()

	// Test that values are copied
	value, err := clone.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "hello world" {
		t.Errorf("Expected 'hello world', got %s", value)
	}

	valueInt, err := clone.GetConfigInt("test_int")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if valueInt != 42 {
		t.Errorf("Expected 42, got %d", valueInt)
	}

	// Test that modifying clone doesn't affect original
	clone.SetConfig("test_string", "modified")
	originalValue, err := cm.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if originalValue != "hello world" {
		t.Error("Modifying clone should not affect original")
	}
}

func TestMerge(t *testing.T) {
	cm1 := configuration.NewConfigManager()
	cm1.SetConfig("test_string", "hello world")
	cm1.SetConfig("test_int", 42)

	cm2 := configuration.NewConfigManager()
	cm2.SetConfig("test_float", 3.14)
	cm2.SetConfig("test_bool", true)

	cm1.Merge(cm2)

	// Test that all configs are merged
	if cm1.GetConfigCount() != 4 {
		t.Errorf("Expected 4 configs, got %d", cm1.GetConfigCount())
	}

	value, err := cm1.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "hello world" {
		t.Errorf("Expected 'hello world', got %s", value)
	}

	valueFloat, err := cm1.GetConfigFloat64("test_float")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if valueFloat != 3.14 {
		t.Errorf("Expected 3.14, got %f", valueFloat)
	}

	// Test that version was incremented (2 calls to SetConfig + 1 call to Merge = 3 increments)
	if cm1.GetVersion() != 4 {
		t.Errorf("Expected version to be 4, got %d", cm1.GetVersion())
	}
}

func TestToJSON(t *testing.T) {
	cm := configuration.NewConfigManager()
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	jsonData, err := cm.ToJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("Expected JSON data to be non-empty")
	}
}

func TestFromJSON(t *testing.T) {
	cm := configuration.NewConfigManager()
	cm.SetConfig("test_string", "hello world")
	cm.SetConfig("test_int", 42)

	jsonData, err := cm.ToJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	newCm := configuration.NewConfigManager()
	err = newCm.FromJSON(jsonData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test that values are restored
	value, err := newCm.GetConfigString("test_string")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "hello world" {
		t.Errorf("Expected 'hello world', got %s", value)
	}

	valueInt, err := newCm.GetConfigInt("test_int")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if valueInt != 42 {
		t.Errorf("Expected 42, got %d", valueInt)
	}
}
