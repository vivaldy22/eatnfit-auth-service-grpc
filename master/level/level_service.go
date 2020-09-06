package level

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/queries"
)

type Service struct{
	db *sql.DB
}

func NewService(db *sql.DB) authservice.LevelCRUDServer {
	return &Service{db}
}

func (s *Service) GetAll(ctx context.Context, empty *empty.Empty) (*authservice.LevelList, error) {
	var levels = new(authservice.LevelList)
	rows, err := s.db.Query(queries.GET_ALL_LEVEL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(authservice.Level)
		if err := rows.Scan(&each.LevelId, &each.LevelName, &each.LevelStatus); err != nil {
			return nil, err
		}
		levels.List = append(levels.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return levels, nil
}

func (s *Service) GetByID(ctx context.Context, id *authservice.ID) (*authservice.Level, error) {
	panic("implement me")
}

func (s *Service) Create(ctx context.Context, level *authservice.Level) (*empty.Empty, error) {
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, request *authservice.LevelUpdateRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, id *authservice.ID) (*empty.Empty, error) {
	panic("implement me")
}

