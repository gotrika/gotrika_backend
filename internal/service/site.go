package service

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

type SitesR interface {
	CreateSite(ctx context.Context, userID primitive.ObjectID, siteDTO dto.CreateSiteDTO) (*core.Site, error)
	UpdateSite(ctx context.Context, siteID primitive.ObjectID, siteDTO *dto.UpdateSiteDTO) (*core.Site, error)
	DeleteSite(ctx context.Context, siteID primitive.ObjectID) error
	GetSiteByID(ctx context.Context, siteID primitive.ObjectID) (*core.Site, error)
	ListSites(ctx context.Context, listDTO *dto.ListSitesDTO) ([]core.Site, int, error)
}

type SiteService struct {
	repo SitesR
}

func NewSiteService(repo SitesR) *SiteService {
	return &SiteService{repo: repo}
}

func (s *SiteService) convertCoreSiteToSiteRetrive(site *core.Site) *dto.SiteRetrieveDTO {
	accessUsers := make([]string, len(site.AccessUsers))
	for index, objID := range site.AccessUsers {
		accessUsers[index] = objID.Hex()
	}
	siteRetrieve := dto.SiteRetrieveDTO{
		ID:          site.ID.Hex(),
		Name:        site.Name,
		URL:         site.URL,
		OwnerID:     site.OwnerId.Hex(),
		AccessUsers: accessUsers,
	}
	return &siteRetrieve
}

func (s *SiteService) CreateSite(ctx context.Context, userID primitive.ObjectID, siteDTO dto.CreateSiteDTO) (*dto.SiteRetrieveDTO, error) {
	site, err := s.repo.CreateSite(ctx, userID, siteDTO)
	if err != nil {
		return nil, err
	}
	siteRetrieve := s.convertCoreSiteToSiteRetrive(site)
	return siteRetrieve, nil
}

func (s *SiteService) GetSiteByID(ctx context.Context, isAdmin bool, userID, siteID primitive.ObjectID) (*dto.SiteRetrieveDTO, error) {
	site, err := s.repo.GetSiteByID(ctx, siteID)
	if err != nil {
		return nil, err
	}
	if site.OwnerId != userID && slices.Contains(site.AccessUsers, userID) && !isAdmin {
		return nil, core.ErrSiteAccessDenied
	}
	siteRetrieve := s.convertCoreSiteToSiteRetrive(site)
	return siteRetrieve, nil
}

func (s *SiteService) DeleteSite(ctx context.Context, isAdmin bool, userID, siteID primitive.ObjectID) error {
	site, err := s.repo.GetSiteByID(ctx, siteID)
	if err != nil {
		return err
	}
	if site.OwnerId != userID && slices.Contains(site.AccessUsers, userID) && !isAdmin {
		return core.ErrSiteAccessDenied
	}
	return s.repo.DeleteSite(ctx, siteID)
}

func (s *SiteService) ListSites(ctx context.Context, search string, isAdmin bool, userID primitive.ObjectID, limit int, offset int) ([]*dto.SiteRetrieveDTO, int, error) {
	listSitesDTO := dto.ListSitesDTO{
		Limit:   limit,
		Offset:  offset,
		IsAdmin: isAdmin,
		UserID:  userID,
		Search:  search,
	}
	coreSites, count, err := s.repo.ListSites(ctx, &listSitesDTO)
	if err != nil {
		return nil, 0, err
	}
	siteRetriveDTOS := make([]*dto.SiteRetrieveDTO, len(coreSites))
	for index, site := range coreSites {
		siteRetriveDTOS[index] = s.convertCoreSiteToSiteRetrive(&site)
	}
	return siteRetriveDTOS, count, nil
}

func (s *SiteService) UpdateSite(ctx context.Context, isAdmin bool, userID primitive.ObjectID, siteID primitive.ObjectID, siteDTO *dto.UpdateSiteDTO) (*dto.SiteRetrieveDTO, error) {
	site, err := s.repo.GetSiteByID(ctx, siteID)
	if err != nil {
		return nil, err
	}
	if !isAdmin && userID != site.OwnerId {
		return nil, core.ErrSiteAccessDenied
	}
	newSite, err := s.repo.UpdateSite(ctx, siteID, siteDTO)
	if err != nil {
		return nil, err
	}
	return s.convertCoreSiteToSiteRetrive(newSite), nil
}
