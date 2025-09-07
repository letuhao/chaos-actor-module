// MongoDB Index Creation Script for RPG Stats Sub-System
// Run this script after creating the collections to optimize query performance

// Switch to the RPG Stats database
use rpg_stats;

print("Creating indexes for RPG Stats Sub-System...");

// 1. Player Progress Collection Indexes
print("Creating player_progress indexes...");

// Unique index on actor_id
db.player_progress.createIndex(
  { "actor_id": 1 },
  { 
    unique: true,
    name: "idx_player_progress_actor_id"
  }
);

// Index on level for level-based queries
db.player_progress.createIndex(
  { "level": 1 },
  { 
    name: "idx_player_progress_level"
  }
);

// Index on last_updated for cleanup queries
db.player_progress.createIndex(
  { "last_updated": 1 },
  { 
    name: "idx_player_progress_last_updated"
  }
);

// 2. Player Effects Active Collection Indexes
print("Creating player_effects_active indexes...");

// Compound index on actor_id and expires_at for active effects queries
db.player_effects_active.createIndex(
  { "actor_id": 1, "expires_at": 1 },
  { 
    name: "idx_player_effects_actor_expires"
  }
);

// Index on effect_id for effect lookups
db.player_effects_active.createIndex(
  { "effect_id": 1 },
  { 
    name: "idx_player_effects_effect_id"
  }
);

// Index on expires_at for cleanup of expired effects
db.player_effects_active.createIndex(
  { "expires_at": 1 },
  { 
    name: "idx_player_effects_expires_at"
  }
);

// Index on actor_id for player-specific queries
db.player_effects_active.createIndex(
  { "actor_id": 1 },
  { 
    name: "idx_player_effects_actor_id"
  }
);

// 3. Player Equipment Collection Indexes
print("Creating player_equipment indexes...");

// Compound unique index on actor_id and slot
db.player_equipment.createIndex(
  { "actor_id": 1, "slot": 1 },
  { 
    unique: true,
    name: "idx_player_equipment_actor_slot"
  }
);

// Index on actor_id for player equipment queries
db.player_equipment.createIndex(
  { "actor_id": 1 },
  { 
    name: "idx_player_equipment_actor_id"
  }
);

// Index on slot for slot-based queries
db.player_equipment.createIndex(
  { "slot": 1 },
  { 
    name: "idx_player_equipment_slot"
  }
);

// Index on equipped_at for equipment history
db.player_equipment.createIndex(
  { "equipped_at": 1 },
  { 
    name: "idx_player_equipment_equipped_at"
  }
);

// 4. Titles Owned Collection Indexes
print("Creating titles_owned indexes...");

// Compound unique index on actor_id and title_id
db.titles_owned.createIndex(
  { "actor_id": 1, "title_id": 1 },
  { 
    unique: true,
    name: "idx_titles_owned_actor_title"
  }
);

// Index on actor_id for player titles
db.titles_owned.createIndex(
  { "actor_id": 1 },
  { 
    name: "idx_titles_owned_actor_id"
  }
);

// Index on title_id for title lookups
db.titles_owned.createIndex(
  { "title_id": 1 },
  { 
    name: "idx_titles_owned_title_id"
  }
);

// Index on granted_at for title history
db.titles_owned.createIndex(
  { "granted_at": 1 },
  { 
    name: "idx_titles_owned_granted_at"
  }
);

// 5. Content Stat Registry Collection Indexes
print("Creating content_stat_registry indexes...");

// Unique index on stat_key
db.content_stat_registry.createIndex(
  { "stat_key": 1 },
  { 
    unique: true,
    name: "idx_content_registry_stat_key"
  }
);

// Index on category for category-based queries
db.content_stat_registry.createIndex(
  { "category": 1 },
  { 
    name: "idx_content_registry_category"
  }
);

// Index on is_primary for primary/derived stat queries
db.content_stat_registry.createIndex(
  { "is_primary": 1 },
  { 
    name: "idx_content_registry_is_primary"
  }
);

// 6. Additional Performance Indexes
print("Creating additional performance indexes...");

// Text index on display names for search
db.content_stat_registry.createIndex(
  { "display_name": "text", "description": "text" },
  { 
    name: "idx_content_registry_text_search"
  }
);

// Sparse index on modifier conditions for conditional effects
db.player_effects_active.createIndex(
  { "modifier.conditions.stack_id": 1 },
  { 
    sparse: true,
    name: "idx_player_effects_stack_id"
  }
);

// Index on modifier key for stat-specific queries
db.player_effects_active.createIndex(
  { "modifier.key": 1 },
  { 
    name: "idx_player_effects_modifier_key"
  }
);

db.player_equipment.createIndex(
  { "modifier.key": 1 },
  { 
    name: "idx_player_equipment_modifier_key"
  }
);

db.titles_owned.createIndex(
  { "modifier.key": 1 },
  { 
    name: "idx_titles_owned_modifier_key"
  }
);

print("All indexes created successfully!");
print("Database is now optimized for RPG Stats Sub-System queries.");

// Display index information
print("\nIndex Summary:");
print("==============");

print("\nPlayer Progress Indexes:");
db.player_progress.getIndexes().forEach(function(index) {
  print("  - " + index.name + ": " + JSON.stringify(index.key));
});

print("\nPlayer Effects Active Indexes:");
db.player_effects_active.getIndexes().forEach(function(index) {
  print("  - " + index.name + ": " + JSON.stringify(index.key));
});

print("\nPlayer Equipment Indexes:");
db.player_equipment.getIndexes().forEach(function(index) {
  print("  - " + index.name + ": " + JSON.stringify(index.key));
});

print("\nTitles Owned Indexes:");
db.titles_owned.getIndexes().forEach(function(index) {
  print("  - " + index.name + ": " + JSON.stringify(index.key));
});

print("\nContent Stat Registry Indexes:");
db.content_stat_registry.getIndexes().forEach(function(index) {
  print("  - " + index.name + ": " + JSON.stringify(index.key));
});
