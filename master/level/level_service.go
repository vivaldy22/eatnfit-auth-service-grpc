package level

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/queries"
	"strconv"
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
	var level = new(authservice.Level)
	row := s.db.QueryRow(queries.GET_BY_ID_LEVEL, id.Id)

	err := row.Scan(&level.LevelId, &level.LevelName, &level.LevelStatus)
	if err != nil {
		return nil, err
	}
	return level, nil
}

func (s *Service) Create(ctx context.Context, level *authservice.Level) (*empty.Empty, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.CREATE_LEVEL)

	if err != nil {
		return new(empty.Empty), err
	}

	res, err := stmt.Exec(level.LevelName)

	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	level.LevelId = int64(lastInsertID)
	stmt.Close()
	return new(empty.Empty), tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *authservice.LevelUpdateRequest) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.UPDATE_LEVEL)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(request.Level.LevelName, request.Id.Id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	convIdToNum, _ := strconv.Atoi(request.Id.Id)
	request.Level.LevelId = int64(convIdToNum)
	return new(empty.Empty), tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *authservice.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_LEVEL)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	return new(empty.Empty), tx.Commit()
}

