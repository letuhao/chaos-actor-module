//go:build mongodb
// +build mongodb

package integration

import (
	"context"
	"errors"
	"time"

	"rpg-system/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoAdapter handles database operations for the RPG Stats Sub-System
type MongoAdapter struct {
	client              *mongo.Client
	database            *mongo.Database
	playerProgressColl  *mongo.Collection
	playerEffectsColl   *mongo.Collection
	playerEquipmentColl *mongo.Collection
	titlesOwnedColl     *mongo.Collection
	contentRegistryColl *mongo.Collection
}

// NewMongoAdapter creates a new MongoDB adapter
func NewMongoAdapter(uri, databaseName string) (*MongoAdapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseName)

	adapter := &MongoAdapter{
		client:              client,
		database:            db,
		playerProgressColl:  db.Collection("player_progress"),
		playerEffectsColl:   db.Collection("player_effects_active"),
		playerEquipmentColl: db.Collection("player_equipment"),
		titlesOwnedColl:     db.Collection("titles_owned"),
		contentRegistryColl: db.Collection("content_stat_registry"),
	}

	// Create indexes
	err = adapter.createIndexes()
	if err != nil {
		return nil, err
	}

	return adapter, nil
}

// NewMongoAdapterLocal creates a new MongoDB adapter for localhost:27017
func NewMongoAdapterLocal(databaseName string) (*MongoAdapter, error) {
	return NewMongoAdapter("mongodb://localhost:27017", databaseName)
}

// Close closes the MongoDB connection
func (ma *MongoAdapter) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return ma.client.Disconnect(ctx)
}

// createIndexes creates the necessary database indexes
func (ma *MongoAdapter) createIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Player Progress indexes
	_, err := ma.playerProgressColl.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "actor_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	// Player Effects indexes
	_, err = ma.playerEffectsColl.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "actor_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "expires_at", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "effect_id", Value: 1}},
		},
	})
	if err != nil {
		return err
	}

	// Player Equipment indexes
	_, err = ma.playerEquipmentColl.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "actor_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "slot", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "actor_id", Value: 1}, {Key: "slot", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}

	// Titles Owned indexes
	_, err = ma.titlesOwnedColl.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "actor_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "title_id", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "actor_id", Value: 1}, {Key: "title_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}

	// Content Registry indexes
	_, err = ma.contentRegistryColl.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "stat_key", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	return nil
}

// GetPlayerProgress retrieves player progression data
func (ma *MongoAdapter) GetPlayerProgress(ctx context.Context, actorID string) (*model.PlayerProgress, error) {
	var progress model.PlayerProgress
	err := ma.playerProgressColl.FindOne(ctx, bson.M{"actor_id": actorID}).Decode(&progress)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return &progress, nil
}

// SavePlayerProgress saves player progression data
func (ma *MongoAdapter) SavePlayerProgress(ctx context.Context, progress *model.PlayerProgress) error {
	progress.LastUpdated = time.Now().Unix()

	opts := options.Replace().SetUpsert(true)
	_, err := ma.playerProgressColl.ReplaceOne(
		ctx,
		bson.M{"actor_id": progress.ActorID},
		progress,
		opts,
	)
	return err
}

// GetActiveEffects retrieves active effects for a player
func (ma *MongoAdapter) GetActiveEffects(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	cursor, err := ma.playerEffectsColl.Find(ctx, bson.M{
		"actor_id":   actorID,
		"expires_at": bson.M{"$gt": time.Now().Unix()},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var effects []model.StatModifier
	for cursor.Next(ctx) {
		var effectDoc struct {
			Modifier model.StatModifier `bson:"modifier"`
		}
		if err := cursor.Decode(&effectDoc); err != nil {
			return nil, err
		}
		effects = append(effects, effectDoc.Modifier)
	}

	return effects, cursor.Err()
}

// AddEffect adds a new effect to a player
func (ma *MongoAdapter) AddEffect(ctx context.Context, actorID string, effect model.StatModifier, duration time.Duration) error {
	effectDoc := bson.M{
		"actor_id":   actorID,
		"effect_id":  effect.Source.ID,
		"modifier":   effect,
		"expires_at": time.Now().Add(duration).Unix(),
		"created_at": time.Now().Unix(),
	}

	_, err := ma.playerEffectsColl.InsertOne(ctx, effectDoc)
	return err
}

// RemoveEffect removes an effect from a player
func (ma *MongoAdapter) RemoveEffect(ctx context.Context, actorID, effectID string) error {
	_, err := ma.playerEffectsColl.DeleteOne(ctx, bson.M{
		"actor_id":  actorID,
		"effect_id": effectID,
	})
	return err
}

// GetEquippedItems retrieves equipped items for a player
func (ma *MongoAdapter) GetEquippedItems(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	cursor, err := ma.playerEquipmentColl.Find(ctx, bson.M{"actor_id": actorID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []model.StatModifier
	for cursor.Next(ctx) {
		var itemDoc struct {
			Modifier model.StatModifier `bson:"modifier"`
		}
		if err := cursor.Decode(&itemDoc); err != nil {
			return nil, err
		}
		items = append(items, itemDoc.Modifier)
	}

	return items, cursor.Err()
}

// EquipItem equips an item to a player
func (ma *MongoAdapter) EquipItem(ctx context.Context, actorID, slot string, item model.StatModifier) error {
	itemDoc := bson.M{
		"actor_id":    actorID,
		"slot":        slot,
		"modifier":    item,
		"equipped_at": time.Now().Unix(),
	}

	opts := options.Replace().SetUpsert(true)
	_, err := ma.playerEquipmentColl.ReplaceOne(
		ctx,
		bson.M{"actor_id": actorID, "slot": slot},
		itemDoc,
		opts,
	)
	return err
}

// UnequipItem unequips an item from a player
func (ma *MongoAdapter) UnequipItem(ctx context.Context, actorID, slot string) error {
	_, err := ma.playerEquipmentColl.DeleteOne(ctx, bson.M{
		"actor_id": actorID,
		"slot":     slot,
	})
	return err
}

// GetOwnedTitles retrieves owned titles for a player
func (ma *MongoAdapter) GetOwnedTitles(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	cursor, err := ma.titlesOwnedColl.Find(ctx, bson.M{"actor_id": actorID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var titles []model.StatModifier
	for cursor.Next(ctx) {
		var titleDoc struct {
			Modifier model.StatModifier `bson:"modifier"`
		}
		if err := cursor.Decode(&titleDoc); err != nil {
			return nil, err
		}
		titles = append(titles, titleDoc.Modifier)
	}

	return titles, cursor.Err()
}

// GrantTitle grants a title to a player
func (ma *MongoAdapter) GrantTitle(ctx context.Context, actorID, titleID string, title model.StatModifier) error {
	titleDoc := bson.M{
		"actor_id":   actorID,
		"title_id":   titleID,
		"modifier":   title,
		"granted_at": time.Now().Unix(),
	}

	opts := options.Replace().SetUpsert(true)
	_, err := ma.titlesOwnedColl.ReplaceOne(
		ctx,
		bson.M{"actor_id": actorID, "title_id": titleID},
		titleDoc,
		opts,
	)
	return err
}

// GetStatRegistry retrieves the stat registry from the database
func (ma *MongoAdapter) GetStatRegistry(ctx context.Context) ([]model.StatDef, error) {
	cursor, err := ma.contentRegistryColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var registry []model.StatDef
	for cursor.Next(ctx) {
		var statDef model.StatDef
		if err := cursor.Decode(&statDef); err != nil {
			return nil, err
		}
		registry = append(registry, statDef)
	}

	return registry, cursor.Err()
}

// SaveStatRegistry saves stat definitions to the database
func (ma *MongoAdapter) SaveStatRegistry(ctx context.Context, registry []model.StatDef) error {
	// Clear existing registry
	_, err := ma.contentRegistryColl.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	// Insert new registry
	var docs []interface{}
	for _, statDef := range registry {
		docs = append(docs, statDef)
	}

	if len(docs) > 0 {
		_, err = ma.contentRegistryColl.InsertMany(ctx, docs)
	}
	return err
}

// CleanupExpiredEffects removes expired effects from the database
func (ma *MongoAdapter) CleanupExpiredEffects(ctx context.Context) error {
	_, err := ma.playerEffectsColl.DeleteMany(ctx, bson.M{
		"expires_at": bson.M{"$lt": time.Now().Unix()},
	})
	return err
}

// GetPlayerStatsSummary returns a summary of all player stats from the database
func (ma *MongoAdapter) GetPlayerStatsSummary(ctx context.Context, actorID string) (*PlayerStatsSummary, error) {
	// Get player progress
	progress, err := ma.GetPlayerProgress(ctx, actorID)
	if err != nil {
		return nil, err
	}

	// Get all modifiers
	effects, err := ma.GetActiveEffects(ctx, actorID)
	if err != nil {
		return nil, err
	}

	equipment, err := ma.GetEquippedItems(ctx, actorID)
	if err != nil {
		return nil, err
	}

	titles, err := ma.GetOwnedTitles(ctx, actorID)
	if err != nil {
		return nil, err
	}

	return &PlayerStatsSummary{
		Progress:  progress,
		Effects:   effects,
		Equipment: equipment,
		Titles:    titles,
	}, nil
}
