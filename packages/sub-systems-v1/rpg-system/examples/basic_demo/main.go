package main

import (
"fmt"
"log"

"rpg-system/internal/integration"
"rpg-system/internal/model"
)

func main() {
fmt.Println("=== RPG Stats Sub-System Demo ===")
fmt.Println()

provider := integration.NewSnapshotProvider()

snapshot, err := provider.BuildForActor("player1", &model.SnapshotOptions{
WithBreakdown: false,
})
if err != nil {
log.Fatalf("Failed to build snapshot: %v", err)
}

fmt.Printf("Actor ID: %s\n", snapshot.ActorID)
fmt.Printf("Level: 5\n")
fmt.Println()

// Show primary stats
fmt.Println("Primary Stats:")
primaryStats := model.PrimaryStats()
for _, stat := range primaryStats {
if value, exists := snapshot.Stats[stat]; exists {
fmt.Printf("  %s: %.0f\n", stat, value)
}
}

fmt.Println()

// Show derived stats
fmt.Println("Derived Stats:")
derivedStats := []string{"HP_MAX", "MANA_MAX", "ATK", "MATK", "DEF", "EVASION", "MOVE_SPEED", "CRIT_CHANCE", "CRIT_DAMAGE"}
for _, statName := range derivedStats {
if value, exists := snapshot.Stats[model.StatKey(statName)]; exists {
fmt.Printf("  %s: %.2f\n", statName, value)
}
}

fmt.Println()
fmt.Println("=== Demo Complete ===")
}
