package auction

import (
	"context"
	"time"

	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/config"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {

	auctionMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	go ar.scheduleAuctionClose(auctionEntity.Id)

	return nil
}

func (ar *AuctionRepository) scheduleAuctionClose(auctionID string) {
	duration := config.GetAuctionDuration()
	time.Sleep(duration)

	ctx := context.Background()

	auction, err := ar.GetByID(ctx, auctionID)
	if err != nil {
		return
	}

	if auction.Status == auction_entity.Closed {
		return
	}

	err = ar.CloseAuction(ctx, auctionID)
	if err != nil {
		logger.Error("Error closing auction", err)
	}
}

func (ar *AuctionRepository) GetByID(ctx context.Context, id string) (*AuctionEntityMongo, error) {
	var auction AuctionEntityMongo

	err := ar.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&auction)
	if err != nil {
		return nil, err
	}

	return &auction, nil
}

func (ar *AuctionRepository) CloseAuction(ctx context.Context, id string) error {
	_, err := ar.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": auction_entity.Closed}},
	)

	return err
}
