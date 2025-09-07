# 03 — Spirit Sea System (Hệ Thống Thức Hải)

**Generated:** 2025-01-27  
**Based on:** Tinh-Khí-Thần system with Spirit Sea as consciousness center

## Tổng quan

**Thức Hải (Spirit Sea)** là trung tâm ý thức và thần thức, ảnh hưởng đến các primary stats liên quan đến **Thần** và khả năng điều khiển trong hệ thống Tinh-Khí-Thần.

## 🎯 Mục tiêu

- **Consciousness Center**: Trung tâm ý thức
- **Divine Control**: Điều khiển thần thức
- **Mental Force**: Lực lượng tinh thần
- **Spell Enhancement**: Tăng cường pháp thuật

## 🧠 Thần Thức (Divine Consciousness)

**Thần Thức** là khả năng cốt lõi của Thức Hải, tương tự như chỉ số **Magic/Intelligence** trong các RPG:

- **Spell Power**: Sức mạnh pháp thuật
- **Spell Accuracy**: Độ chính xác pháp thuật  
- **Spell Range**: Phạm vi pháp thuật
- **Spell Control**: Điều khiển pháp thuật

## 🏗️ Cấu Trúc Dữ Liệu

### **SpiritSeaSystem - Hệ Thống Thức Hải**

```go
// Spirit Sea System (识海系统) - Hệ thống thức hải
type SpiritSeaSystem struct {
    // Spirit Sea core (识海核心) - Lõi thức hải
    SpiritSeaCore SpiritSeaCore `json:"spirit_sea_core"`
    
    // Consciousness (意识) - Ý thức
    Consciousness Consciousness `json:"consciousness"`
    
    // Mental Force (念力) - Lực lượng tinh thần
    MentalForce MentalForce `json:"mental_force"`
    
    // Wisdom (智慧) - Trí tuệ
    Wisdom Wisdom `json:"wisdom"`
    
    // Willpower (意志力) - Ý chí
    Willpower Willpower `json:"willpower"`
    
    // Divine Sense (神觉) - Thần giác
    DivineSense DivineSense `json:"divine_sense"`
    
    // Spiritual Control (精神控制) - Điều khiển tinh thần
    SpiritualControl SpiritualControl `json:"spiritual_control"`
    
    // Memory Palace (记忆宫殿) - Cung điện ký ức
    MemoryPalace MemoryPalace `json:"memory_palace"`
    
    // Cultivation Mind (修炼心境) - Tâm cảnh tu luyện
    CultivationMind CultivationMind `json:"cultivation_mind"`
}
```

### **SpiritSeaCore - Lõi Thức Hải**

```go
// Spirit Sea Core (识海核心) - Lõi thức hải
type SpiritSeaCore struct {
    // Basic capacity (基础容量) - Dung lượng cơ bản
    BaseCapacity    float64 `json:"base_capacity"`    // Base capacity (基础容量) - Dung lượng cơ bản
    CurrentCapacity float64 `json:"current_capacity"` // Current capacity (当前容量) - Dung lượng hiện tại
    MaxCapacity     float64 `json:"max_capacity"`     // Maximum capacity (最大容量) - Dung lượng tối đa
    
    // Shen clarity (神清) - Thần thanh
    ShenClarity     float64 `json:"shen_clarity"`     // Shen clarity (神清) - Thần thanh
    ShenControl     float64 `json:"shen_control"`     // Shen control (神控) - Thần khống
    ShenCapacity    float64 `json:"shen_capacity"`    // Shen capacity (神容) - Thần dung
    ShenStability   float64 `json:"shen_stability"`   // Shen stability (神稳) - Thần ổn
    ShenEfficiency  float64 `json:"shen_efficiency"`  // Shen efficiency (神效) - Thần hiệu
}
```

### **Consciousness - Ý Thức**

```go
// Consciousness (意识) - Ý thức
type Consciousness struct {
    // Basic consciousness (基础意识) - Ý thức cơ bản
    BasicConsciousness    float64 `json:"basic_consciousness"`    // Basic consciousness (基础意识) - Ý thức cơ bản
    AdvancedConsciousness float64 `json:"advanced_consciousness"` // Advanced consciousness (高级意识) - Ý thức nâng cao
    DivineConsciousness   float64 `json:"divine_consciousness"`   // Divine consciousness (神意识) - Thần ý thức
    TranscendentConsciousness float64 `json:"transcendent_consciousness"` // Transcendent consciousness (超脱意识) - Siêu thoát ý thức
    
    // Awareness levels (觉知层次) - Cấp độ nhận thức
    SelfAwareness    float64 `json:"self_awareness"`    // Self awareness (自我觉知) - Tự giác
    WorldAwareness   float64 `json:"world_awareness"`   // World awareness (世界觉知) - Thế giới giác
    UniversalAwareness float64 `json:"universal_awareness"` // Universal awareness (宇宙觉知) - Vũ trụ giác
}
```

### **MentalForce - Lực Lượng Tinh Thần**

```go
// Mental Force (念力) - Lực lượng tinh thần
type MentalForce struct {
    // Nian Li capacity (念力容量) - Dung lượng niệm lực
    NianLiCapacity  float64 `json:"nian_li_capacity"`  // Nian Li capacity (念力容量) - Dung lượng niệm lực
    NianLiCurrent   float64 `json:"nian_li_current"`   // Current Nian Li (当前念力) - Niệm lực hiện tại
    NianLiRegen     float64 `json:"nian_li_regen"`     // Nian Li regeneration (念力恢复) - Tốc độ hồi niệm lực
    NianLiEfficiency float64 `json:"nian_li_efficiency"` // Nian Li efficiency (念力效率) - Hiệu suất niệm lực
    
    // Mental strength (精神力量) - Sức mạnh tinh thần
    MentalStrength  float64 `json:"mental_strength"`   // Mental strength (精神力量) - Sức mạnh tinh thần
    MentalEndurance float64 `json:"mental_endurance"`  // Mental endurance (精神耐力) - Sức bền tinh thần
    MentalResilience float64 `json:"mental_resilience"` // Mental resilience (精神韧性) - Sự dẻo dai tinh thần
}
```

### **Wisdom - Trí Tuệ**

```go
// Wisdom (智慧) - Trí tuệ
type Wisdom struct {
    // Knowledge (知识) - Tri thức
    Knowledge       float64 `json:"knowledge"`         // Knowledge (知识) - Tri thức
    Understanding   float64 `json:"understanding"`     // Understanding (理解) - Sự hiểu biết
    Analysis        float64 `json:"analysis"`          // Analysis (分析) - Phân tích
    Synthesis       float64 `json:"synthesis"`         // Synthesis (综合) - Tổng hợp
    Intuition       float64 `json:"intuition"`         // Intuition (直觉) - Trực giác
    
    // Advanced wisdom (高级智慧) - Trí tuệ nâng cao
    Insight         float64 `json:"insight"`           // Insight (洞察) - Sự thấu hiểu
    Foresight       float64 `json:"foresight"`         // Foresight (预见) - Sự tiên tri
    Enlightenment   float64 `json:"enlightenment"`     // Enlightenment (开悟) - Sự khai ngộ
    Transcendence   float64 `json:"transcendence"`     // Transcendence (超脱) - Sự siêu thoát
}
```

### **Willpower - Ý Chí**

```go
// Willpower (意志力) - Ý chí
type Willpower struct {
    // Basic willpower (基础意志力) - Ý chí cơ bản
    Determination   float64 `json:"determination"`     // Determination (决心) - Sự quyết tâm
    Perseverance    float64 `json:"perseverance"`      // Perseverance (毅力) - Sự kiên trì
    Resilience      float64 `json:"resilience"`        // Resilience (韧性) - Sự dẻo dai
    Courage         float64 `json:"courage"`           // Courage (勇气) - Lòng dũng cảm
    Discipline      float64 `json:"discipline"`        // Discipline (纪律) - Kỷ luật
    
    // Advanced willpower (高级意志力) - Ý chí nâng cao
    UnbreakableWill float64 `json:"unbreakable_will"`  // Unbreakable will (不可破意志) - Ý chí không thể phá vỡ
    IronResolve     float64 `json:"iron_resolve"`      // Iron resolve (钢铁决心) - Quyết tâm sắt thép
    TranscendentWill float64 `json:"transcendent_will"` // Transcendent will (超脱意志) - Ý chí siêu thoát
}
```

### **DivineSense - Thần Giác**

```go
// Divine Sense (神觉) - Thần giác
type DivineSense struct {
    // Sense range (感知范围) - Phạm vi cảm giác
    SenseRange      float64 `json:"sense_range"`       // Sense range (感知范围) - Phạm vi cảm giác
    SensePrecision  float64 `json:"sense_precision"`   // Sense precision (感知精度) - Độ chính xác cảm giác
    SenseSpeed      float64 `json:"sense_speed"`       // Sense speed (感知速度) - Tốc độ cảm giác
    SensePenetration float64 `json:"sense_penetration"` // Sense penetration (感知穿透) - Khả năng xuyên thấu
    SenseStealth    float64 `json:"sense_stealth"`     // Sense stealth (感知隐蔽) - Khả năng ẩn nấp
    
    // Advanced senses (高级感知) - Cảm giác nâng cao
    DivineVision    float64 `json:"divine_vision"`     // Divine vision (神视) - Thần thị
    DivineHearing   float64 `json:"divine_hearing"`    // Divine hearing (神听) - Thần thính
    DivineTouch     float64 `json:"divine_touch"`      // Divine touch (神触) - Thần xúc
    DivineSmell     float64 `json:"divine_smell"`      // Divine smell (神嗅) - Thần khứu
    DivineTaste     float64 `json:"divine_taste"`      // Divine taste (神味) - Thần vị
}
```

### **SpiritualControl - Điều Khiển Tinh Thần**

```go
// Spiritual Control (精神控制) - Điều khiển tinh thần
type SpiritualControl struct {
    // Self control (自我控制) - Điều khiển bản thân
    SelfControl     float64 `json:"self_control"`      // Self control (自我控制) - Điều khiển bản thân
    EmotionControl  float64 `json:"emotion_control"`   // Emotion control (情绪控制) - Điều khiển cảm xúc
    ThoughtControl  float64 `json:"thought_control"`   // Thought control (思维控制) - Điều khiển tư duy
    MemoryControl   float64 `json:"memory_control"`    // Memory control (记忆控制) - Điều khiển ký ức
    DreamControl    float64 `json:"dream_control"`     // Dream control (梦境控制) - Điều khiển giấc mơ
    
    // External control (外部控制) - Điều khiển bên ngoài
    MindControl     float64 `json:"mind_control"`      // Mind control (心灵控制) - Điều khiển tâm trí
    IllusionControl float64 `json:"illusion_control"`  // Illusion control (幻术控制) - Điều khiển ảo thuật
    Telepathy       float64 `json:"telepathy"`         // Telepathy (心灵感应) - Tâm linh cảm ứng
    Telekinesis     float64 `json:"telekinesis"`       // Telekinesis (念力移物) - Niệm lực di vật
}
```

### **MemoryPalace - Cung Điện Ký ức**

```go
// Memory Palace (记忆宫殿) - Cung điện ký ức
type MemoryPalace struct {
    // Basic memory (基础记忆) - Ký ức cơ bản
    MemoryCapacity  float64 `json:"memory_capacity"`   // Memory capacity (记忆容量) - Dung lượng ký ức
    MemorySpeed     float64 `json:"memory_speed"`      // Memory speed (记忆速度) - Tốc độ ký ức
    MemoryAccuracy  float64 `json:"memory_accuracy"`   // Memory accuracy (记忆准确性) - Độ chính xác ký ức
    MemoryRetention float64 `json:"memory_retention"`  // Memory retention (记忆保持) - Duy trì ký ức
    MemoryRecall    float64 `json:"memory_recall"`     // Memory recall (记忆回忆) - Hồi tưởng ký ức
    
    // Advanced memory (高级记忆) - Ký ức nâng cao
    EideticMemory   float64 `json:"eidetic_memory"`    // Eidetic memory (图像记忆) - Hình ảnh ký ức
    PhotographicMemory float64 `json:"photographic_memory"` // Photographic memory (照相记忆) - Chụp ảnh ký ức
    PerfectMemory   float64 `json:"perfect_memory"`    // Perfect memory (完美记忆) - Hoàn mỹ ký ức
    TranscendentMemory float64 `json:"transcendent_memory"` // Transcendent memory (超脱记忆) - Siêu thoát ký ức
}
```

### **CultivationMind - Tâm Cảnh Tu Luyện**

```go
// Cultivation Mind (修炼心境) - Tâm cảnh tu luyện
type CultivationMind struct {
    // Basic cultivation mind (基础修炼心境) - Tâm cảnh tu luyện cơ bản
    Calmness        float64 `json:"calmness"`          // Calmness (平静) - Bình tĩnh
    Serenity        float64 `json:"serenity"`          // Serenity (宁静) - Tĩnh lặng
    Tranquility     float64 `json:"tranquility"`       // Tranquility (宁静) - An tĩnh
    Peace           float64 `json:"peace"`             // Peace (和平) - Hòa bình
    Harmony         float64 `json:"harmony"`           // Harmony (和谐) - Hài hòa
    
    // Advanced cultivation mind (高级修炼心境) - Tâm cảnh tu luyện nâng cao
    Enlightenment   float64 `json:"enlightenment"`     // Enlightenment (开悟) - Khai ngộ
    Transcendence   float64 `json:"transcendence"`     // Transcendence (超脱) - Siêu thoát
    Nirvana         float64 `json:"nirvana"`           // Nirvana (涅槃) - Niết bàn
    Buddhahood      float64 `json:"buddhahood"`        // Buddhahood (佛性) - Phật tính
    DaoHeart        float64 `json:"dao_heart"`         // Dao Heart (道心) - Đạo tâm
}
```

## 📊 Primary Stats từ Thức Hải

```go
// Spirit Sea primary stats (识海主要属性) - Các primary stats từ thức hải
var SpiritSeaPrimaryStats = []string{
    // Core stats (核心属性) - Thuộc tính cốt lõi
    "shen_clarity",          // Shen clarity (神清) - Thần thanh
    "shen_control",          // Shen control (神控) - Thần khống
    "shen_capacity",         // Shen capacity (神容) - Thần dung
    "shen_stability",        // Shen stability (神稳) - Thần ổn
    "shen_efficiency",       // Shen efficiency (神效) - Thần hiệu
    
    // Advanced stats (高级属性) - Thuộc tính nâng cao
    "divine_awareness",      // Divine awareness (神觉) - Thần giác
    "spiritual_range",       // Spiritual range (精神范围) - Phạm vi tinh thần
    "mental_resistance",     // Mental resistance (精神抗性) - Khả năng chống tinh thần
    "thought_speed",         // Thought speed (思维速度) - Tốc độ tư duy
    "memory_capacity",       // Memory capacity (记忆容量) - Dung lượng ký ức
    
    // Consciousness stats (意识属性) - Thuộc tính ý thức
    "basic_consciousness",   // Basic consciousness (基础意识) - Ý thức cơ bản
    "advanced_consciousness", // Advanced consciousness (高级意识) - Ý thức nâng cao
    "divine_consciousness",  // Divine consciousness (神意识) - Thần ý thức
    "transcendent_consciousness", // Transcendent consciousness (超脱意识) - Siêu thoát ý thức
    
    // Mental Force stats (念力属性) - Thuộc tính niệm lực
    "nian_li_capacity",      // Nian Li capacity (念力容量) - Dung lượng niệm lực
    "nian_li_current",       // Current Nian Li (当前念力) - Niệm lực hiện tại
    "nian_li_regen",         // Nian Li regeneration (念力恢复) - Tốc độ hồi niệm lực
    "nian_li_efficiency",    // Nian Li efficiency (念力效率) - Hiệu suất niệm lực
    
    // Wisdom stats (智慧属性) - Thuộc tính trí tuệ
    "knowledge",             // Knowledge (知识) - Tri thức
    "understanding",         // Understanding (理解) - Sự hiểu biết
    "analysis",              // Analysis (分析) - Phân tích
    "synthesis",             // Synthesis (综合) - Tổng hợp
    "intuition",             // Intuition (直觉) - Trực giác
    
    // Willpower stats (意志力属性) - Thuộc tính ý chí
    "determination",         // Determination (决心) - Sự quyết tâm
    "perseverance",          // Perseverance (毅力) - Sự kiên trì
    "resilience",            // Resilience (韧性) - Sự dẻo dai
    "courage",               // Courage (勇气) - Lòng dũng cảm
    "discipline",            // Discipline (纪律) - Kỷ luật
    
    // Divine Sense stats (神觉属性) - Thuộc tính thần giác
    "sense_range",           // Sense range (感知范围) - Phạm vi cảm giác
    "sense_precision",       // Sense precision (感知精度) - Độ chính xác cảm giác
    "sense_speed",           // Sense speed (感知速度) - Tốc độ cảm giác
    "sense_penetration",     // Sense penetration (感知穿透) - Khả năng xuyên thấu
    "sense_stealth",         // Sense stealth (感知隐蔽) - Khả năng ẩn nấp
    
    // Spiritual Control stats (精神控制属性) - Thuộc tính điều khiển tinh thần
    "self_control",          // Self control (自我控制) - Điều khiển bản thân
    "emotion_control",       // Emotion control (情绪控制) - Điều khiển cảm xúc
    "thought_control",       // Thought control (思维控制) - Điều khiển tư duy
    "memory_control",        // Memory control (记忆控制) - Điều khiển ký ức
    "dream_control",         // Dream control (梦境控制) - Điều khiển giấc mơ
    
    // Memory Palace stats (记忆宫殿属性) - Thuộc tính cung điện ký ức
    "memory_speed",          // Memory speed (记忆速度) - Tốc độ ký ức
    "memory_accuracy",       // Memory accuracy (记忆准确性) - Độ chính xác ký ức
    "memory_retention",      // Memory retention (记忆保持) - Duy trì ký ức
    "memory_recall",         // Memory recall (记忆回忆) - Hồi tưởng ký ức
    
    // Cultivation Mind stats (修炼心境属性) - Thuộc tính tâm cảnh tu luyện
    "calmness",              // Calmness (平静) - Bình tĩnh
    "serenity",              // Serenity (宁静) - Tĩnh lặng
    "tranquility",           // Tranquility (宁静) - An tĩnh
    "peace",                 // Peace (和平) - Hòa bình
    "harmony",               // Harmony (和谐) - Hài hòa
}
```

## 🧮 Công Thức Tính Toán

### **CalculateSpiritSeaPrimaryStats - Tính Primary Stats Thức Hải**

```go
// Calculate Spirit Sea primary stats (计算识海主要属性) - Tính toán các primary stats của thức hải
func CalculateSpiritSeaPrimaryStats(spiritSea SpiritSeaSystem) map[string]float64 {
    stats := make(map[string]float64)
    
    // Core stats (核心属性) - Thuộc tính cốt lõi
    stats["shen_clarity"] = spiritSea.SpiritSeaCore.ShenClarity
    stats["shen_control"] = spiritSea.SpiritSeaCore.ShenControl
    stats["shen_capacity"] = spiritSea.SpiritSeaCore.ShenCapacity
    stats["shen_stability"] = spiritSea.SpiritSeaCore.ShenStability
    stats["shen_efficiency"] = spiritSea.SpiritSeaCore.ShenEfficiency
    
    // Advanced stats (高级属性) - Thuộc tính nâng cao
    stats["divine_awareness"] = spiritSea.Consciousness.DivineConsciousness
    stats["spiritual_range"] = spiritSea.DivineSense.SenseRange
    stats["mental_resistance"] = spiritSea.MentalForce.MentalResilience
    stats["thought_speed"] = spiritSea.Wisdom.Analysis
    stats["memory_capacity"] = spiritSea.MemoryPalace.MemoryCapacity
    
    // Consciousness stats (意识属性) - Thuộc tính ý thức
    stats["basic_consciousness"] = spiritSea.Consciousness.BasicConsciousness
    stats["advanced_consciousness"] = spiritSea.Consciousness.AdvancedConsciousness
    stats["divine_consciousness"] = spiritSea.Consciousness.DivineConsciousness
    stats["transcendent_consciousness"] = spiritSea.Consciousness.TranscendentConsciousness
    
    // Mental Force stats (念力属性) - Thuộc tính niệm lực
    stats["nian_li_capacity"] = spiritSea.MentalForce.NianLiCapacity
    stats["nian_li_current"] = spiritSea.MentalForce.NianLiCurrent
    stats["nian_li_regen"] = spiritSea.MentalForce.NianLiRegen
    stats["nian_li_efficiency"] = spiritSea.MentalForce.NianLiEfficiency
    
    // Wisdom stats (智慧属性) - Thuộc tính trí tuệ
    stats["knowledge"] = spiritSea.Wisdom.Knowledge
    stats["understanding"] = spiritSea.Wisdom.Understanding
    stats["analysis"] = spiritSea.Wisdom.Analysis
    stats["synthesis"] = spiritSea.Wisdom.Synthesis
    stats["intuition"] = spiritSea.Wisdom.Intuition
    
    // Willpower stats (意志力属性) - Thuộc tính ý chí
    stats["determination"] = spiritSea.Willpower.Determination
    stats["perseverance"] = spiritSea.Willpower.Perseverance
    stats["resilience"] = spiritSea.Willpower.Resilience
    stats["courage"] = spiritSea.Willpower.Courage
    stats["discipline"] = spiritSea.Willpower.Discipline
    
    // Divine Sense stats (神觉属性) - Thuộc tính thần giác
    stats["sense_range"] = spiritSea.DivineSense.SenseRange
    stats["sense_precision"] = spiritSea.DivineSense.SensePrecision
    stats["sense_speed"] = spiritSea.DivineSense.SenseSpeed
    stats["sense_penetration"] = spiritSea.DivineSense.SensePenetration
    stats["sense_stealth"] = spiritSea.DivineSense.SenseStealth
    
    // Spiritual Control stats (精神控制属性) - Thuộc tính điều khiển tinh thần
    stats["self_control"] = spiritSea.SpiritualControl.SelfControl
    stats["emotion_control"] = spiritSea.SpiritualControl.EmotionControl
    stats["thought_control"] = spiritSea.SpiritualControl.ThoughtControl
    stats["memory_control"] = spiritSea.SpiritualControl.MemoryControl
    stats["dream_control"] = spiritSea.SpiritualControl.DreamControl
    
    // Memory Palace stats (记忆宫殿属性) - Thuộc tính cung điện ký ức
    stats["memory_speed"] = spiritSea.MemoryPalace.MemorySpeed
    stats["memory_accuracy"] = spiritSea.MemoryPalace.MemoryAccuracy
    stats["memory_retention"] = spiritSea.MemoryPalace.MemoryRetention
    stats["memory_recall"] = spiritSea.MemoryPalace.MemoryRecall
    
    // Cultivation Mind stats (修炼心境属性) - Thuộc tính tâm cảnh tu luyện
    stats["calmness"] = spiritSea.CultivationMind.Calmness
    stats["serenity"] = spiritSea.CultivationMind.Serenity
    stats["tranquility"] = spiritSea.CultivationMind.Tranquility
    stats["peace"] = spiritSea.CultivationMind.Peace
    stats["harmony"] = spiritSea.CultivationMind.Harmony
    
    return stats
}
```

### **CalculateNianLiFromSpiritSea - Tính Niệm Lực từ Thức Hải**

```go
// Calculate Nian Li from Spirit Sea (根据识海计算念力) - Tính niệm lực từ thức hải
func CalculateNianLiFromSpiritSea(spiritSea SpiritSeaSystem) float64 {
    baseNianLi := spiritSea.MentalForce.NianLiCapacity
    shenMultiplier := 1.0 + (spiritSea.SpiritSeaCore.ShenControl * 0.5)
    consciousnessMultiplier := 1.0 + (spiritSea.Consciousness.DivineConsciousness * 0.3)
    wisdomMultiplier := 1.0 + (spiritSea.Wisdom.Intuition * 0.2)
    
    return baseNianLi * shenMultiplier * consciousnessMultiplier * wisdomMultiplier
}
```

### **CalculateSpellDamageFromShen - Tính Sát Thương Pháp Thuật từ Thần**

```go
// Calculate spell damage from Shen (根据神计算法术伤害) - Tính sát thương pháp thuật từ thần
func CalculateSpellDamageFromShen(spiritSea SpiritSeaSystem) float64 {
    baseDamage := 100.0
    shenClarity := spiritSea.SpiritSeaCore.ShenClarity
    shenControl := spiritSea.SpiritSeaCore.ShenControl
    divineConsciousness := spiritSea.Consciousness.DivineConsciousness
    wisdom := spiritSea.Wisdom.Knowledge
    
    shenMultiplier := 1.0 + (shenClarity * 0.4) + (shenControl * 0.3)
    consciousnessMultiplier := 1.0 + (divineConsciousness * 0.2)
    wisdomMultiplier := 1.0 + (wisdom * 0.1)
    
    return baseDamage * shenMultiplier * consciousnessMultiplier * wisdomMultiplier
}
```

## 🔗 Tích Hợp Với Actor Core v3

### **ConvertToSubsystemOutput - Chuyển Đổi Thành SubsystemOutput**

```go
// Convert Spirit Sea to SubsystemOutput (转换识海到子系统输出) - Chuyển đổi thức hải thành SubsystemOutput
func ConvertSpiritSeaToSubsystemOutput(spiritSea SpiritSeaSystem) SubsystemOutput {
    // Calculate primary stats (计算主要属性) - Tính toán primary stats
    primaryStats := CalculateSpiritSeaPrimaryStats(spiritSea)
    
    // Calculate derived stats (计算派生属性) - Tính toán derived stats
    derivedStats := make(map[string]float64)
    
    // Nian Li calculations (念力计算) - Tính toán niệm lực
    derivedStats["nian_li_max"] = CalculateNianLiFromSpiritSea(spiritSea)
    derivedStats["nian_li_current"] = spiritSea.MentalForce.NianLiCurrent
    derivedStats["nian_li_regen"] = spiritSea.MentalForce.NianLiRegen
    
    // Spell calculations (法术计算) - Tính toán pháp thuật
    derivedStats["spell_damage"] = CalculateSpellDamageFromShen(spiritSea)
    derivedStats["spell_accuracy"] = CalculateSpellAccuracyFromShen(spiritSea)
    derivedStats["spell_range"] = CalculateSpellRangeFromShen(spiritSea)
    derivedStats["spell_cooldown"] = CalculateSpellCooldownFromShen(spiritSea)
    
    // Create contributions (创建贡献) - Tạo contributions
    var contributions []Contribution
    
    // Primary stats contributions (主要属性贡献) - Contributions primary stats
    for stat, value := range primaryStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Primary",
            Value:     value,
            System:    "SpiritSeaSystem",
            Priority:  1,
        })
    }
    
    // Derived stats contributions (派生属性贡献) - Contributions derived stats
    for stat, value := range derivedStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Derived",
            Value:     value,
            System:    "SpiritSeaSystem",
            Priority:  2,
        })
    }
    
    return SubsystemOutput{
        Primary:   primaryStats,
        Derived:   derivedStats,
        Caps:      make(map[string]float64),
        Context:   map[string]interface{}{"spirit_sea": spiritSea},
        Contributions: contributions,
    }
}
```

## 💡 Đặc Điểm Nổi Bật

### **Consciousness Center:**
- **Divine Consciousness**: Thần ý thức
- **Mental Force**: Lực lượng tinh thần
- **Spiritual Control**: Điều khiển tinh thần

### **Spell Enhancement:**
- **Spell Power**: Sức mạnh pháp thuật
- **Spell Accuracy**: Độ chính xác pháp thuật
- **Spell Range**: Phạm vi pháp thuật
- **Spell Control**: Điều khiển pháp thuật

### **Memory & Wisdom:**
- **Memory Palace**: Cung điện ký ức
- **Knowledge Management**: Quản lý tri thức
- **Intuition System**: Hệ thống trực giác
