// MongoDB Schema Definitions for RPG Stats Sub-System
// Run this script to set up the database collections and indexes

// Switch to the RPG Stats database
use rpg_stats;

// 1. Player Progress Collection
db.createCollection("player_progress", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["actor_id", "level", "xp", "allocations", "last_updated"],
      properties: {
        actor_id: {
          bsonType: "string",
          description: "Unique actor identifier"
        },
        level: {
          bsonType: "int",
          minimum: 1,
          maximum: 100,
          description: "Character level"
        },
        xp: {
          bsonType: "long",
          minimum: 0,
          description: "Total experience points"
        },
        allocations: {
          bsonType: "object",
          description: "Stat point allocations",
          patternProperties: {
            "^[A-Z_]+$": {
              bsonType: "int",
              minimum: 0
            }
          }
        },
        last_updated: {
          bsonType: "long",
          description: "Last update timestamp"
        }
      }
    }
  }
});

// 2. Player Effects Active Collection
db.createCollection("player_effects_active", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["actor_id", "effect_id", "modifier", "expires_at", "created_at"],
      properties: {
        actor_id: {
          bsonType: "string",
          description: "Actor identifier"
        },
        effect_id: {
          bsonType: "string",
          description: "Unique effect identifier"
        },
        modifier: {
          bsonType: "object",
          required: ["key", "op", "value", "source"],
          properties: {
            key: {
              bsonType: "string",
              description: "Stat key being modified"
            },
            op: {
              bsonType: "string",
              enum: ["ADD_FLAT", "ADD_PCT", "MULTIPLY", "OVERRIDE"],
              description: "Modifier operation"
            },
            value: {
              bsonType: "double",
              description: "Modifier value"
            },
            source: {
              bsonType: "object",
              required: ["kind", "id", "label"],
              properties: {
                kind: {
                  bsonType: "string",
                  description: "Source type"
                },
                id: {
                  bsonType: "string",
                  description: "Source identifier"
                },
                label: {
                  bsonType: "string",
                  description: "Human-readable label"
                }
              }
            },
            conditions: {
              bsonType: "object",
              properties: {
                requires_tags_all: {
                  bsonType: "array",
                  items: { bsonType: "string" }
                },
                requires_tags_any: {
                  bsonType: "array",
                  items: { bsonType: "string" }
                },
                forbid_tags: {
                  bsonType: "array",
                  items: { bsonType: "string" }
                },
                duration_ms: {
                  bsonType: "long",
                  minimum: 0
                },
                stack_id: {
                  bsonType: "string"
                },
                max_stacks: {
                  bsonType: "int",
                  minimum: 1
                }
              }
            },
            priority: {
              bsonType: "int",
              description: "Modifier priority"
            }
          }
        },
        expires_at: {
          bsonType: "long",
          description: "Effect expiration timestamp"
        },
        created_at: {
          bsonType: "long",
          description: "Effect creation timestamp"
        }
      }
    }
  }
});

// 3. Player Equipment Collection
db.createCollection("player_equipment", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["actor_id", "slot", "modifier", "equipped_at"],
      properties: {
        actor_id: {
          bsonType: "string",
          description: "Actor identifier"
        },
        slot: {
          bsonType: "string",
          description: "Equipment slot"
        },
        modifier: {
          bsonType: "object",
          required: ["key", "op", "value", "source"],
          properties: {
            key: {
              bsonType: "string",
              description: "Stat key being modified"
            },
            op: {
              bsonType: "string",
              enum: ["ADD_FLAT", "ADD_PCT", "MULTIPLY", "OVERRIDE"],
              description: "Modifier operation"
            },
            value: {
              bsonType: "double",
              description: "Modifier value"
            },
            source: {
              bsonType: "object",
              required: ["kind", "id", "label"],
              properties: {
                kind: {
                  bsonType: "string",
                  description: "Source type"
                },
                id: {
                  bsonType: "string",
                  description: "Source identifier"
                },
                label: {
                  bsonType: "string",
                  description: "Human-readable label"
                }
              }
            },
            priority: {
              bsonType: "int",
              description: "Modifier priority"
            }
          }
        },
        equipped_at: {
          bsonType: "long",
          description: "Equipment timestamp"
        }
      }
    }
  }
});

// 4. Titles Owned Collection
db.createCollection("titles_owned", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["actor_id", "title_id", "modifier", "granted_at"],
      properties: {
        actor_id: {
          bsonType: "string",
          description: "Actor identifier"
        },
        title_id: {
          bsonType: "string",
          description: "Title identifier"
        },
        modifier: {
          bsonType: "object",
          required: ["key", "op", "value", "source"],
          properties: {
            key: {
              bsonType: "string",
              description: "Stat key being modified"
            },
            op: {
              bsonType: "string",
              enum: ["ADD_FLAT", "ADD_PCT", "MULTIPLY", "OVERRIDE"],
              description: "Modifier operation"
            },
            value: {
              bsonType: "double",
              description: "Modifier value"
            },
            source: {
              bsonType: "object",
              required: ["kind", "id", "label"],
              properties: {
                kind: {
                  bsonType: "string",
                  description: "Source type"
                },
                id: {
                  bsonType: "string",
                  description: "Source identifier"
                },
                label: {
                  bsonType: "string",
                  description: "Human-readable label"
                }
              }
            },
            priority: {
              bsonType: "int",
              description: "Modifier priority"
            }
          }
        },
        granted_at: {
          bsonType: "long",
          description: "Title grant timestamp"
        }
      }
    }
  }
});

// 5. Content Stat Registry Collection
db.createCollection("content_stat_registry", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["stat_key", "category", "display_name", "description", "is_primary"],
      properties: {
        stat_key: {
          bsonType: "string",
          description: "Unique stat key"
        },
        category: {
          bsonType: "string",
          description: "Stat category"
        },
        display_name: {
          bsonType: "string",
          description: "Human-readable name"
        },
        description: {
          bsonType: "string",
          description: "Stat description"
        },
        rounding: {
          bsonType: "string",
          description: "Rounding rule"
        },
        min_value: {
          bsonType: "double",
          description: "Minimum allowed value"
        },
        max_value: {
          bsonType: "double",
          description: "Maximum allowed value"
        },
        default_value: {
          bsonType: "double",
          description: "Default value"
        },
        is_primary: {
          bsonType: "bool",
          description: "Whether this is a primary stat"
        }
      }
    }
  }
});

print("Database collections created successfully!");
print("Run the indexes.js script to create indexes.");
