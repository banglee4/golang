package service

import (
	"fmt"
	"sippm/helper"
	"sippm/model/repository"
	"sippm/model/web/pengumuman"

	"gorm.io/gorm"
)

type PengumumanServiceImpl struct {
	DB       *gorm.DB
	PengRepo repository.PengumumanRepository
	DokRepo  repository.PengDokumenRepository
}

func NewPengumumanServiceImpl(pengRepo repository.PengumumanRepository, dokRepo repository.PengDokumenRepository, db *gorm.DB) *PengumumanServiceImpl {
	return &PengumumanServiceImpl{PengRepo: pengRepo, DokRepo: dokRepo, DB: db}
}

func (serv *PengumumanServiceImpl) Create(request pengumuman.CreateRequest) error {
	// Model
	model := pengumuman.CreateRequestToDomain(request)

	// Call Repo
	err := serv.PengRepo.Create(model, serv.DB)
	if err != nil {
		return err
	}

	return err
}

func (serv *PengumumanServiceImpl) Update(request pengumuman.UpdateRequest) error {
	// Ubah request menjadi model domain
	model := pengumuman.UpdateRequestToDomain(request)

	// Mulai transaksi GORM
	tx := serv.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	// Handle panic/rollback otomatis
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update Pengumuman utama
	if err := serv.PengRepo.Update(model, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update pengumuman: %w", err)
	}

	// Update Dokumen terkait
	if err := serv.DokRepo.DeleteByPengumumanId(model.Id, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete dokumen: %w", err)
	}

	for _, dokReq := range model.Dokumen {
		if err := serv.DokRepo.Update(model.Id, dokReq, tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update dokumen: %w", err)
		}
	}

	// Commit jika semua update berhasil
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (serv *PengumumanServiceImpl) GetAll() ([]pengumuman.Response, error) {
	// Call Repo
	results, err := serv.PengRepo.GetAll(serv.DB)
	if err != nil {
		return nil, err
	}

	return pengumuman.ToResponses(results), nil
}

func (serv *PengumumanServiceImpl) Delete(id int) (err error) {
	// Mulai Transaksi GORM
	tx := serv.DB.Begin()
	errTransaction := helper.TransactionErrorHandler(tx)
	if errTransaction != nil {
		return errTransaction
	}

	// Call Repo
	errDok := serv.DokRepo.DeleteByPengumumanId(id, tx)
	if errDok != nil {
		err = errDok
		return err
	}

	errPeng := serv.PengRepo.Delete(id, tx)
	if errPeng != nil {
		err = errPeng
		return err
	}

	// Commit Transaksi jika semua operasi di atas berhasil
	errCommit := helper.TransactionCommitHandler(tx)
	if errCommit != nil {
		err = errCommit
	}

	return nil
}
