package service

import (
	"fmt"

	"github.com/presidendjakarta/swantara-gate/internal/model"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
)

// APIConsumerService menangani business logic untuk API Consumer
type APIConsumerService struct {
	ConsumerRepo *repository.APIConsumerRepository
}

// NewAPIConsumerService membuat instance baru APIConsumerService
func NewAPIConsumerService(consumerRepo *repository.APIConsumerRepository) *APIConsumerService {
	return &APIConsumerService{ConsumerRepo: consumerRepo}
}

// CreateConsumer membuat consumer baru
func (s *APIConsumerService) CreateConsumer(req *model.CreateAPIConsumerRequest) (*model.APIConsumer, error) {
	consumer, err := s.ConsumerRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat consumer: %w", err)
	}

	return consumer, nil
}

// GetConsumerByID mengambil consumer berdasarkan ID
func (s *APIConsumerService) GetConsumerByID(id int64) (*model.APIConsumer, error) {
	consumer, err := s.ConsumerRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil consumer: %w", err)
	}

	return consumer, nil
}

// GetAllConsumers mengambil semua consumer dengan pagination
func (s *APIConsumerService) GetAllConsumers(page, limit int) ([]model.APIConsumer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	total, err := s.ConsumerRepo.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("gagal menghitung consumer: %w", err)
	}

	consumers, err := s.ConsumerRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mengambil daftar consumer: %w", err)
	}

	return consumers, total, nil
}

// UpdateConsumer memperbarui data consumer
func (s *APIConsumerService) UpdateConsumer(id int64, req *model.UpdateAPIConsumerRequest) error {
	_, err := s.ConsumerRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("consumer tidak ditemukan: %w", err)
	}

	err = s.ConsumerRepo.Update(id, req)
	if err != nil {
		return fmt.Errorf("gagal mengupdate consumer: %w", err)
	}

	return nil
}

// DeleteConsumer menghapus consumer
func (s *APIConsumerService) DeleteConsumer(id int64) error {
	_, err := s.ConsumerRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("consumer tidak ditemukan: %w", err)
	}

	err = s.ConsumerRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus consumer: %w", err)
	}

	return nil
}
