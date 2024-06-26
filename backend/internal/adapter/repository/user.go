package repository

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/shironxn/blanknotes/internal/core/domain"
	"github.com/shironxn/blanknotes/internal/core/port"
	"github.com/shironxn/blanknotes/internal/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	db         *gorm.DB
	pagination util.Pagination
}

func NewUserRepository(db *gorm.DB, pagination util.Pagination) port.UserRepository {
	return &UserRepository{
		db:         db,
		pagination: pagination,
	}
}

func (r *UserRepository) GetAll(req domain.UserQuery, metadata *domain.Metadata) ([]domain.User, error) {
	var entity []domain.User
	query := r.db.Model(&domain.User{})

	if !req.Details {
		query = r.db.Model(&domain.User{}).Select("id", "name", "created_at", "updated_at")
	}

	if err := query.
		Where(&domain.UserQuery{Name: req.Name}).
		Count(&metadata.TotalRecords).
		Scopes(r.pagination.Paginate(metadata)).
		Find(&entity).
		Error; err != nil {
		return nil, err
	}

	if reflect.DeepEqual(entity, []domain.User{}) {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return entity, nil
}

func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
	var entity domain.User

	if err := r.db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return nil, err
	}

	return &entity, nil
}

func (r *UserRepository) Update(req domain.UserRequest, user *domain.User) (*domain.User, error) {
	entity := user

	if req.Bio == "" {
		r.db.Model(entity).Update("bio", "")
	}

	if req.AvatarURL == "" {
		r.db.Model(entity).Update("avatar_url", "")
	}

	if err := r.db.Model(entity).Updates(req).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			if strings.Contains(mysqlErr.Message, "name") {
				return nil, fiber.NewError(fiber.StatusBadRequest, "user with the same name already exists")
			}
			if strings.Contains(mysqlErr.Message, "email") {
				return nil, fiber.NewError(fiber.StatusBadRequest, "user with the same email already exists")
			}
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return nil, err
	}

	return entity, nil
}

func (r *UserRepository) Delete(user *domain.User) error {
	entity := user

	if err := r.db.Delete(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		return err
	}

	return nil
}
