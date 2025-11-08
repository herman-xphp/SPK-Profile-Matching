# SPK Profile Matching

Sistem Pendukung Keputusan (SPK) untuk Profile Matching menggunakan metode Profile Matching untuk menentukan ranking tenaga kerja berdasarkan kesesuaian dengan profil jabatan.

**Repository**: [https://github.com/herman-xphp/spk-profile-matching](https://github.com/herman-xphp/spk-profile-matching)

## 📋 Daftar Isi

- [Fitur](#-fitur)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Persyaratan Sistem](#-persyaratan-sistem)
- [Cara Download dari GitHub](#-cara-download-dari-github)
- [Instalasi](#-instalasi)
- [Setup Database](#-setup-database)
- [Menjalankan Aplikasi](#-menjalankan-aplikasi)
- [Cara Masuk ke Aplikasi](#-cara-masuk-ke-aplikasi)
- [Struktur Project](#-struktur-project)
- [Troubleshooting](#-troubleshooting)

## ✨ Fitur

- **Manajemen Jabatan**: Kelola data jabatan yang tersedia
- **Manajemen Aspek & Kriteria**: Kelola aspek dan kriteria penilaian
- **Manajemen Tenaga Kerja**: Kelola data tenaga kerja
- **Target Profile**: Tentukan nilai target per kriteria untuk setiap jabatan
- **Nilai Tenaga Kerja**: Input nilai aktual tenaga kerja per kriteria
- **Profile Matching**: Perhitungan otomatis untuk menentukan ranking
- **Hasil Ranking**: Tampilkan hasil ranking dengan detail perhitungan
- **Export CSV**: Export hasil ranking ke file CSV

## 🛠 Teknologi yang Digunakan

### Backend
- **Go 1.25.3** - Bahasa pemrograman
- **Gin** - Web framework
- **GORM** - ORM untuk database
- **MySQL 8.0** - Database
- **JWT** - Authentication

### Frontend
- **React 18.2.0** - UI framework
- **React Router** - Routing
- **Axios** - HTTP client
- **Tailwind CSS** - Styling
- **Radix UI** - UI components
- **Sonner** - Toast notifications

## 📦 Persyaratan Sistem

Sebelum memulai, pastikan Anda telah menginstall:

1. **Go** (versi 1.25.3 atau lebih baru)
   - Download: https://golang.org/dl/
   - Verifikasi: `go version`

2. **Node.js** (versi 16 atau lebih baru)
   - Download: https://nodejs.org/
   - Verifikasi: `node --version`

3. **Yarn** (package manager untuk frontend)
   - Install: `npm install -g yarn`
   - Verifikasi: `yarn --version`

4. **MySQL** (versi 8.0 atau lebih baru)
   - Download: https://dev.mysql.com/downloads/mysql/
   - Atau gunakan Docker (disarankan)

5. **Docker & Docker Compose** (opsional, untuk database)
   - Download: https://www.docker.com/get-started
   - Verifikasi: `docker --version` dan `docker-compose --version`

6. **Git**
   - Download: https://git-scm.com/downloads
   - Verifikasi: `git --version`

## 📥 Cara Download dari GitHub

### Metode 1: Clone Repository (Disarankan)

```bash
# Clone repository
git clone https://github.com/herman-xphp/spk-profile-matching.git

# Masuk ke direktori project
cd spk-profile-matching
```

### Metode 2: Download ZIP

1. Buka repository di GitHub
2. Klik tombol **"Code"** → **"Download ZIP"**
3. Extract file ZIP ke direktori yang diinginkan
4. Buka terminal di direktori tersebut

## 🔧 Instalasi

### 1. Setup Backend

```bash
# Masuk ke direktori backend
cd backend

# Install dependencies Go
go mod download

# Copy file template environment
cp ENV_TEMPLATE.txt .env

# Edit file .env sesuai konfigurasi Anda
# Gunakan text editor favorit Anda (nano, vim, code, dll)
nano .env
```

**Isi file `.env` minimal seperti ini:**

```env
PORT=8000
GIN_MODE=debug
SECRET_KEY=your-secret-key-here-change-in-production
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=spk_profile_matching
TEST_DB_NAME=spk_profile_matching_test
```

**⚠️ PENTING**: Ganti `SECRET_KEY` dengan string random yang kuat (minimal 32 karakter)!

### 2. Setup Frontend

```bash
# Kembali ke root project
cd ..

# Masuk ke direktori frontend
cd frontend

# Install dependencies
yarn install

# Buat file .env untuk frontend
# File ini diperlukan untuk konfigurasi URL backend
echo "REACT_APP_BACKEND_URL=http://localhost:8000" > .env
```

**Catatan**: Pastikan `REACT_APP_BACKEND_URL` sesuai dengan URL backend Anda. Default adalah `http://localhost:8000`.

## 🗄️ Setup Database

### Opsi 1: Menggunakan Docker (Disarankan)

```bash
# Masuk ke direktori backend
cd backend

# Jalankan MySQL dengan Docker Compose
docker-compose up -d

# Verifikasi container berjalan
docker ps
```

Docker akan otomatis:
- Membuat database `spk_profile_matching`
- Membuat database `spk_profile_matching_test`
- Setup user dan password

### Opsi 2: Setup Manual MySQL

1. **Buat database secara manual:**

```bash
# Login ke MySQL
mysql -u root -p

# Buat database
CREATE DATABASE spk_profile_matching;
CREATE DATABASE spk_profile_matching_test;

# Keluar dari MySQL
exit;
```

2. **Atau gunakan script SQL:**

```bash
# Masuk ke direktori backend
cd backend

# Jalankan script SQL
mysql -u root -p < db-init/create-dbs.sql
```

### 3. Seed Database (Mengisi Data Awal)

```bash
# Masih di direktori backend
cd backend

# Jalankan seed untuk mengisi data awal
go run cmd/seed/main.go
```

Seed akan membuat:
- ✅ User admin: `admin@kpsggroup.com` / `admin123`
- ✅ User biasa: `user@kpsggroup.com` / `user123`
- ✅ Data jabatan, aspek, kriteria
- ✅ Data target profile
- ✅ Data tenaga kerja dan nilai

## 🚀 Menjalankan Aplikasi

### Menjalankan Backend

Buka terminal pertama:

```bash
# Masuk ke direktori backend
cd backend

# Jalankan server
go run cmd/api/main.go
```

Backend akan berjalan di: **http://localhost:8000**

Anda akan melihat output seperti:
```
[GIN-debug] Listening and serving HTTP on :8000
```

### Menjalankan Frontend

Buka terminal kedua (terminal baru):

```bash
# Masuk ke direktori frontend
cd frontend

# Jalankan development server
yarn start
```

Frontend akan berjalan di: **http://localhost:3000**

Browser akan otomatis terbuka. Jika tidak, buka manual di browser.

## 🔐 Cara Masuk ke Aplikasi

1. **Buka browser** dan akses: **http://localhost:3000**

2. **Halaman Login** akan muncul

3. **Gunakan kredensial berikut:**

   **Admin:**
   - Email: `admin@kpsggroup.com`
   - Password: `admin123`

   **User Biasa:**
   - Email: `user@kpsggroup.com`
   - Password: `user123`

4. **Klik tombol "Masuk"**

5. Setelah login berhasil, Anda akan diarahkan ke **Dashboard**

### Fitur berdasarkan Role

**Admin** memiliki akses penuh:
- ✅ Dashboard
- ✅ Manajemen Jabatan
- ✅ Manajemen Aspek
- ✅ Manajemen Kriteria
- ✅ Manajemen Tenaga Kerja
- ✅ Manajemen Target Profile
- ✅ Manajemen Nilai Tenaga Kerja
- ✅ Perhitungan Profile Matching
- ✅ Hasil Ranking

**User** memiliki akses terbatas:
- ✅ Dashboard
- ✅ Melihat Jabatan
- ✅ Melihat Aspek
- ✅ Melihat Kriteria
- ✅ Melihat Tenaga Kerja
- ✅ Melihat Target Profile
- ✅ Melihat Nilai Tenaga Kerja
- ✅ Perhitungan Profile Matching
- ✅ Hasil Ranking

## 📁 Struktur Project

```
SPK-Profile-Matching/
├── backend/                 # Backend Go application
│   ├── cmd/
│   │   ├── api/            # Main API server
│   │   └── seed/           # Database seeder
│   ├── internal/
│   │   ├── controllers/    # HTTP handlers
│   │   ├── services/       # Business logic
│   │   ├── repositories/   # Data access layer
│   │   ├── models/         # Database models
│   │   ├── dto/            # Data transfer objects
│   │   └── middleware/     # Middleware (auth, CORS)
│   ├── pkg/
│   │   └── database/       # Database connection
│   ├── db-init/            # Database initialization scripts
│   ├── tests/              # Integration tests
│   ├── go.mod              # Go dependencies
│   ├── docker-compose.yml  # Docker configuration
│   └── .env                # Environment variables (buat sendiri)
│
└── frontend/               # Frontend React application
    ├── public/             # Static files
    ├── src/
    │   ├── components/     # React components
    │   ├── pages/         # Page components
    │   ├── App.js          # Main app component
    │   └── index.js        # Entry point
    ├── package.json        # Node dependencies
    └── yarn.lock           # Lock file
```

## 🔍 Troubleshooting

### Backend tidak bisa connect ke database

**Error**: `Could not connect to database`

**Solusi**:
1. Pastikan MySQL berjalan:
   ```bash
   # Cek status MySQL
   sudo systemctl status mysql
   # atau untuk Docker
   docker ps
   ```

2. Pastikan konfigurasi di `.env` benar:
   - `DB_HOST=127.0.0.1` (atau `localhost`)
   - `DB_PORT=3306`
   - `DB_USER=root`
   - `DB_PASSWORD=` (kosong jika tidak ada password)

3. Test koneksi manual:
   ```bash
   mysql -u root -p -h 127.0.0.1 -P 3306
   ```

### Frontend tidak bisa connect ke backend

**Error**: `Network Error` atau `CORS Error`

**Solusi**:
1. Pastikan backend berjalan di `http://localhost:8000`
2. Buat file `.env` di folder `frontend` jika backend berjalan di port/URL berbeda:
   ```bash
   cd frontend
   echo "REACT_APP_BACKEND_URL=http://localhost:8000" > .env
   ```
3. Restart frontend setelah membuat/mengubah `.env`
4. Pastikan CORS sudah di-enable di backend (sudah ada di middleware)

### Port sudah digunakan

**Error**: `address already in use`

**Solusi**:
1. Cari process yang menggunakan port:
   ```bash
   # Untuk port 8000 (backend)
   lsof -i :8000
   # atau
   netstat -ano | grep 8000
   
   # Untuk port 3000 (frontend)
   lsof -i :3000
   ```

2. Kill process tersebut atau ubah port di konfigurasi

### Database migration error

**Error**: `Error 1452: Foreign key constraint fails`

**Solusi**:
1. Hapus semua tabel dan jalankan seed ulang:
   ```bash
   cd backend
   go run cmd/seed/main.go
   ```
   Seed akan otomatis drop dan recreate semua tabel

2. Atau hapus database dan buat ulang:
   ```bash
   mysql -u root -p
   DROP DATABASE spk_profile_matching;
   CREATE DATABASE spk_profile_matching;
   exit;
   ```

### Module not found (Go)

**Error**: `cannot find module`

**Solusi**:
```bash
cd backend
go mod tidy
go mod download
```

### Module not found (Node)

**Error**: `Cannot find module`

**Solusi**:
```bash
cd frontend
rm -rf node_modules yarn.lock
yarn install
```

### JWT Token Error

**Error**: `Invalid credentials` atau `Failed to generate token`

**Solusi**:
1. Pastikan `SECRET_KEY` sudah di-set di file `.env`
2. Pastikan `SECRET_KEY` tidak kosong
3. Restart backend setelah mengubah `.env`

## 📝 Catatan Penting

1. **Jangan commit file `.env`** ke Git - file ini berisi informasi sensitif
2. **Ganti `SECRET_KEY`** dengan nilai yang kuat di production
3. **Gunakan database terpisah** untuk development dan production
4. **Backup database** secara berkala
5. **Jalankan test** sebelum deploy:
   ```bash
   cd backend
   go test ./...
   ```

## 📞 Support

Jika mengalami masalah, silakan:
1. Cek bagian [Troubleshooting](#-troubleshooting)
2. Buka issue di [GitHub repository](https://github.com/herman-xphp/spk-profile-matching)
3. Hubungi tim development

## 📄 License

[Tambahkan informasi license di sini]

---

**Selamat menggunakan SPK Profile Matching! 🎉**

