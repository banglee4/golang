CREATE TABLE Users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role ENUM('Admin', 'Dosen', 'UPPM', 'LPPM') NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Dosen (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       user_id INT NOT NULL,
                       nidn VARCHAR(20),
                       nama VARCHAR(50),
                       photo TEXT,
                       email VARCHAR(50),
                       phone VARCHAR(50),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       CONSTRAINT fk_dosen_user FOREIGN KEY (user_id) REFERENCES Users(id),
                       UNIQUE INDEX idx_nidn (nidn)
);

CREATE TABLE Uppm (
                      id INT AUTO_INCREMENT PRIMARY KEY,
                      user_id INT NOT NULL,
                      fakultas ENUM('FTIK', 'FEB', 'FSIP') NOT NULL,
                      nama VARCHAR(50),
                      photo TEXT,
                      email VARCHAR(50),
                      phone VARCHAR(50),
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      CONSTRAINT fk_uppm_user FOREIGN KEY (user_id) REFERENCES Users(id),
                      INDEX idx_nama (nama)
);

CREATE TABLE Lppm (
                      id INT AUTO_INCREMENT PRIMARY KEY,
                      user_id INT NOT NULL,
                      nama VARCHAR(50),
                      photo TEXT,
                      email VARCHAR(50),
                      phone VARCHAR(50),
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      CONSTRAINT fk_lppm_user FOREIGN KEY (user_id) REFERENCES Users(id),
                      INDEX idx_nama (nama)
);

CREATE TABLE Usulan (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        pengusul_id INT NOT NULL,
                        fakultas ENUM('FTIK', 'FEB', 'FSIP') NOT NULL,
                        prodi ENUM(
                            'Informatika',
                            'Sistem Informasi',
                            'Teknik Komputer',
                            'Teknologi Informasi',
                            'Ilmu Komputer',
                            'Teknik Sipil',
                            'Teknik Elektro',
                            'Manajemen',
                            'Akutansi',
                            'Sastra Inggris',
                            'Bahasa Inggris',
                            'Matematika'
                            ) NOT NULL,
                        judul VARCHAR(250) NOT NULL UNIQUE,
                        jenis_usulan ENUM('PKM', 'Penelitian') NOT NULL,
                        tahun_usulan VARCHAR(10),
                        anggaran TEXT,
                        status ENUM('Draft', 'Usulan Pending', 'Usulan Diterima', 'Usulan Ditolak', 'Laporan Kemajuan Pending', 'Laporan Kemajuan Diterima', 'Laporan Kemajuan Ditolak', 'Laporan Akhir Pending', 'Laporan Akhir Diterima', 'Laporan Akhir Ditolak', 'Selesai') NOT NULL,
                        nama_anggota_1 VARCHAR(100),
                        nama_anggota_2 VARCHAR(100),
                        nama_anggota_3 VARCHAR(100),
                        nama_mahasiswa_1 VARCHAR(100),
                        nama_mahasiswa_2 VARCHAR(100),
                        nama_mahasiswa_3 VARCHAR(100),
                        nama_mahasiswa_4 VARCHAR(100),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                        CONSTRAINT fk_usulan_dosen FOREIGN KEY (pengusul_id) REFERENCES Users(id),

                        INDEX idx_judul (judul),
                        INDEX idx_fakultas (fakultas),
                        INDEX idx_prodi (prodi),
                        INDEX idx_jenis_usulan (jenis_usulan),
                        INDEX idx_tahun_usulan (tahun_usulan),
                        INDEX idx_status_usulan (status)
);

CREATE TABLE FileLaporan (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             usulan_id INT NOT NULL,
                             file_usulan TEXT,
                             status_usulan ENUM('Pending', 'Diterima', 'Ditolak'),
                             catatan_usulan TEXT,
                             file_laporan_kemajuan TEXT,
                             status_laporan_kemajuan ENUM('Pending', 'Diterima', 'Ditolak'),
                             catatan_laporan_kemajuan TEXT,
                             file_laporan_akhir TEXT,
                             status_laporan_akhir ENUM('Pending', 'Diterima', 'Ditolak'),
                             catatan_laporan_akhir TEXT,
                             CONSTRAINT fk_filelaporan_usulan FOREIGN KEY (usulan_id) REFERENCES Usulan(id)
);

CREATE TABLE Pengumuman (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             judul TEXT,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE DokumenPengumuman (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            pengumuman_id INT,
                            nama_dokument TEXT,
                            url TEXT,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            CONSTRAINT FOREIGN KEY (pengumuman_id) REFERENCES Pengumuman(id)
);

CREATE TABLE JabatanKetua (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            nama_dosen TEXT,
                            jabatan ENUM('Ketua UPPM FTIK', 'Ketua UPPM FEB', 'Ketua UPPM FSIP'),
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);