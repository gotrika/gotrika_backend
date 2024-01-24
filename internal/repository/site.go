package repository

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"os"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateCounterCode(url string, siteID string) string {
	file, err := os.ReadFile("./internal/templates/tracker.tpl")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	t := template.Must(template.New("t").Parse(string(file)))
	_ = t.Execute(&tpl, map[string]string{"code": siteID, "url": url})
	script := tpl.String()
	return script
}

type SiteRepo struct {
	collection *mongo.Collection
}

func NewSiteRepo(db *mongo.Database) *SiteRepo {
	return &SiteRepo{
		collection: db.Collection(core.SiteCollectioName),
	}
}

func (r *SiteRepo) CreateSite(ctx context.Context, userID primitive.ObjectID, siteDTO dto.CreateSiteDTO, scriptUrl string) (*core.Site, error) {

	s := core.Site{
		Name:        siteDTO.Name,
		URL:         siteDTO.URL,
		OwnerId:     userID,
		CounterCode: siteDTO.CounterCode,
	}
	res, err := r.collection.InsertOne(ctx, &s)
	if err != nil {
		return nil, err
	}
	s.ID = res.InsertedID.(primitive.ObjectID)
	s.CounterCode = generateCounterCode(scriptUrl, s.ID.Hex())
	_, _ = r.collection.UpdateByID(ctx, s.ID, bson.D{bson.E{Key: "$set", Value: bson.M{"counter_code": s.CounterCode}}})
	return &s, nil
}

func (r *SiteRepo) GetSiteByID(ctx context.Context, siteID primitive.ObjectID) (*core.Site, error) {
	var site core.Site
	err := r.collection.FindOne(ctx, bson.M{"_id": siteID}).Decode(&site)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, core.ErrSiteNotFound
		}
		return nil, err
	}
	return &site, nil
}

func (r *SiteRepo) DeleteSite(ctx context.Context, siteID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": siteID})
	return err
}

func (r *SiteRepo) ListSites(ctx context.Context, listDTO *dto.ListSitesDTO) ([]core.Site, int, error) {
	var sites []core.Site
	opts := getPaginationOpts(listDTO.Limit, listDTO.Offset)
	searchQuery := bson.D{}
	if listDTO.Search != "" {
		expr := primitive.Regex{Pattern: listDTO.Search}
		searchQuery = bson.D{
			bson.E{Key: "$or", Value: bson.A{
				bson.D{
					bson.E{Key: "name", Value: expr},
				},
				bson.D{
					bson.E{Key: "url", Value: expr},
				},
			}},
		}
	}
	if listDTO.IsAdmin {
		filter := bson.M{}
		if len(searchQuery) != 0 {
			filter = bson.M{"$and": bson.A{searchQuery}}
		}
		cur, err := r.collection.Find(ctx, filter, opts)
		if err != nil {
			return nil, 0, err
		}
		if err := cur.All(ctx, &sites); err != nil {
			return nil, 0, err
		}

		count, err := r.collection.CountDocuments(ctx, filter)
		return sites, int(count), err
	}
	query := bson.D{
		bson.E{Key: "$or", Value: bson.A{
			bson.D{
				bson.E{Key: "owner_id", Value: listDTO.UserID},
			},
			bson.D{
				bson.E{Key: "access_users", Value: bson.A{listDTO.UserID}},
			},
		}},
	}
	filter := query
	if len(searchQuery) != 0 {
		filter = bson.D{
			bson.E{Key: "$and", Value: bson.A{searchQuery, query}},
		}
	}
	cur, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	if err := cur.All(ctx, &sites); err != nil {
		return nil, 0, err
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	return sites, int(count), err
}

func (r *SiteRepo) UpdateSite(ctx context.Context, siteID primitive.ObjectID, siteDTO *dto.UpdateSiteDTO) (*core.Site, error) {
	update := bson.M{}
	if siteDTO.Name != "" {
		update["name"] = siteDTO.Name
	}
	if siteDTO.URL != "" {
		update["url"] = siteDTO.URL
	}
	accessUsers := []primitive.ObjectID{}
	for _, userID := range siteDTO.AccessUsers {
		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return nil, err
		}
		accessUsers = append(accessUsers, objID)
	}
	update["access_users"] = accessUsers
	_, err := r.collection.UpdateByID(ctx, siteID, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	site, err := r.GetSiteByID(ctx, siteID)
	if err != nil {
		return nil, err
	}
	return site, nil
}
