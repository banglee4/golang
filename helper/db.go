package helper

import "gorm.io/gorm"

func TransactionErrorHandler(tx *gorm.DB) (err error) {
	if tx.Error != nil {
		// Jika gagal memulai transaksi, langsung kembalikan error
		return tx.Error
	}

	// Defer function untuk menangani rollback atau panic recovery
	// Fungsi ini akan dieksekusi saat fungsi Update selesai (baik sukses atau gagal)
	defer func() {
		// Jika ada panic, rollback transaksi dan re-panic
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		// Jika ada error yang di-return dari fungsi Update (nilai 'err' tidak nil),
		// lakukan rollback transaksi.
		// Variabel 'err' di sini adalah named return value dari fungsi Update.
		if err != nil {
			tx.Rollback()
		}
	}()

	return nil
}

func TransactionCommitHandler(tx *gorm.DB) (err error) {
	err = tx.Commit().Error // GORM 2.x mengembalikan error melalui .Error()
	if err != nil {
		// Jika commit itu sendiri gagal (sangat jarang, tapi mungkin), kembalikan error
		return err
	}

	return nil
}
