package user

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/queries"
)

type Service struct{
	db *sql.DB
}

func NewService(db *sql.DB) authservice.UserCRUDServer {
	return &Service{db}
}

func (s *Service) GetAll(ctx context.Context, e *empty.Empty) (*authservice.UserListResponse, error) {
	var users = new(authservice.UserListResponse)
	rows, err := s.db.Query(queries.GET_ALL_USER)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(authservice.UserResponse)
		if err := rows.Scan(&each.UserId, &each.UserEmail, &each.UserPassword, &each.UserFName, &each.UserLName,
			&each.UserGender, &each.UserPhoto, &each.UserBalance, &each.UserLevel, &each.UserStatus); err != nil {
			return nil, err
		}
		users.List = append(users.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetByID(ctx context.Context, id *authservice.ID) (*authservice.User, error) {
	var user = new(authservice.User)
	row := s.db.QueryRow(queries.GET_BY_ID_USER, id.Id)

	err := row.Scan(&user.UserId, &user.UserEmail, &user.UserPassword, &user.UserFName, &user.UserLName,
		&user.UserGender, &user.UserPhoto, &user.UserBalance, &user.UserLevel, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByEmail(ctx context.Context, email *authservice.Email) (*authservice.User, error) {
	var user = new(authservice.User)
	row := s.db.QueryRow(queries.GET_BY_EMAIL_USER, email.Email)

	err := row.Scan(&user.UserId, &user.UserEmail, &user.UserPassword, &user.UserFName, &user.UserLName,
		&user.UserGender, &user.UserPhoto, &user.UserBalance, &user.UserLevel, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, user *authservice.User) (*empty.Empty, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.CREATE_USER)

	if err != nil {
		return new(empty.Empty), err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, user.UserEmail, user.UserPassword, user.UserFName, user.UserLName,
		user.UserGender, user.UserPhoto, user.UserBalance, user.UserLevel)

	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	user.UserId = id
	stmt.Close()
	return new(empty.Empty), tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *authservice.UserUpdateRequest) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.UPDATE_USER)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(request.User.UserEmail, request.User.UserPassword, request.User.UserFName, request.User.UserLName,
		request.User.UserGender, request.User.UserPhoto, request.User.UserBalance, request.User.UserLevel, request.Id.Id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	request.User.UserId = request.Id.Id
	return new(empty.Empty), tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *authservice.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_USER)
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
