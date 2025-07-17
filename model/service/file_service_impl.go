package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"sippm/model/domain"
	"strings"
)

type FileServImpl struct {
}

func NewFileServImpl() *FileServImpl {
	return &FileServImpl{}
}

func (serv *FileServImpl) UploadFile(files []*multipart.FileHeader) ([]domain.File, error) {
	uploadDir := "assets"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	var uploaded []domain.File

	// Valid extension
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".pdf":  true,
		".doc":  true,
	}

	for _, fileHeader := range files {
		// Validation Exts
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !allowedExts[ext] {
			return nil, fmt.Errorf("file type not allowed: %s", ext)
		}

		// Open file
		src, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}

		// --- Modifikasi Mulai Di Sini ---
		// Dapatkan nama file asli tanpa ekstensi
		baseName := strings.TrimSuffix(fileHeader.Filename, ext)
		fileName := fileHeader.Filename // Inisialisasi dengan nama file asli
		dstPath := filepath.Join(uploadDir, fileName)
		counter := 1

		// Periksa apakah file sudah ada, jika ya, tambahkan (counter)
		for {
			_, err := os.Stat(dstPath)
			if os.IsNotExist(err) {
				// File tidak ada, nama ini bisa dipakai
				break
			}
			// File sudah ada, buat nama baru
			fileName = fmt.Sprintf("%s(%d)%s", baseName, counter, ext)
			dstPath = filepath.Join(uploadDir, fileName)
			counter++
		}
		// --- Modifikasi Berakhir Di Sini ---

		// Save to disk
		dst, errDst := os.Create(dstPath)
		if errDst != nil {
			_ = src.Close()
			return nil, fmt.Errorf("failed to create file: %w", errDst)
		}

		if _, err := io.Copy(dst, src); err != nil {
			_ = dst.Close()
			_ = src.Close()
			return nil, fmt.Errorf("failed to save file: %w", err)
		}

		// Close File
		_ = dst.Close()
		_ = src.Close()

		// Create URL FILE
		fileURL := fmt.Sprintf("/assets/%s", fileName)

		uploadedFile := domain.File{
			FileName: fileName,
			FileURL:  fileURL,
		}

		uploaded = append(uploaded, uploadedFile)
	}

	return uploaded, nil
}

func (serv *FileServImpl) DeleteFile(fileURL string) error {
	// Parse URL
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return fmt.Errorf("invalid file URL: %w", err)
	}

	// Make sure initial path is "/assets/"
	if !strings.HasPrefix(parsedURL.Path, "/assets/") {
		return fmt.Errorf("file URL must start with /assets/")
	}

	// Get File Name
	fileName := strings.TrimPrefix(parsedURL.Path, "/assets/")
	if fileName == "" {
		return fmt.Errorf("file name is empty in URL")
	}

	// Full path file on disk
	fullPath := filepath.Join("assets", fileName)

	// Check file is exist
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}

	// Delete file
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
