// db/indexes.js
// Run with: `mongosh <connection> db.getSiblingDB("<DB>").load("db/indexes.js")`

db.player_progress.createIndex({_id: 1});
db.player_effects_active.createIndex({actorId: 1});
db.player_effects_active.createIndex({actorId: 1, expiresAt: 1});
db.player_equipment.createIndex({actorId: 1});
db.player_equipment.createIndex({actorId: 1, slot: 1}, {unique: true});
db.titles_owned.createIndex({actorId: 1});
db.titles_owned.createIndex({actorId: 1, isActive: 1});
db.content_stat_registry.createIndex({_id: 1});
