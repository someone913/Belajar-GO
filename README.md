# Go Gin Todo API

REST API sederhana untuk manajemen Todo (CRUD) menggunakan **Go**, **Gin**, dan **PostgreSQL**.

## 📋 Fitur

- `GET /api/v1/todos` — Mengambil semua todo (dapat difilter dengan query `?done=true` atau `?done=false`)
- `GET /api/v1/todos/:id` — Mengambil satu todo berdasarkan ID
- `POST /api/v1/todos` — Membuat todo baru
- `PUT /api/v1/todos/:id` — Memperbarui todo yang sudah ada
- `DELETE /api/v1/todos/:id` — Menghapus todo

## 🛠️ Tech Stack

| Komponen | Versi / Paket |
|----------|---------------|
| Go | 1.27.1 |
| Web Framework | Gin v1.12.0 |
| Database Driver | pgx v5.11.0 |
| Validation | go-playground/validator v10 |
| Env Loader | joho/godotenv v1.5.1 |
| Database | PostgreSQL |

## 📦 Struktur Project

```
Go_Hari_7/
├── main.go                 # Entry point aplikasi
├── go.mod                  # Modul dan dependensi Go
├── go.sum
├── .env                    # Konfigurasi environment (DATABASE_URL)
├── README.md               # Dokumentasi ini
├── database/
│   └── database.go         # Koneksi PostgreSQL
├── router/
│   └── router.go           # Setup router dan grup route
└── todo/
    ├── model.go            # Struct model & request
    ├── routes.go           # Pendaftaran route /todos
    └── handler.go          # Handler CRUD (List, GetByID, Create, Update, Delete)
```

## ⚙️ Instalasi & Menjalankan

### 1. Clone / Buka project
Pastikan kamu berada di folder `Go_Hari_7`.

### 2. Install dependensi
```bash
go mod tidy
```

### 3. Konfigurasi `.env`
Buat file `.env` di root project dengan isi:

```env
DATABASE_URL=postgresql://user:password@host:port/dbname?sslmode=require
```

> Contoh untuk Neon (sudah tersedia di project ini):
> ```env
> DATABASE_URL=postgresql://neondb_owner:npg_xxx@ep-xxx.c-xxx.aws.neon.tech/neondb?sslmode=require&channel_binding=require
> ```

Pastikan database dan tabel `todos` sudah dibuat sebelum menjalankan aplikasi.

### 4. Membuat tabel di PostgreSQL
Jalankan query berikut di database kamu:

```sql
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Jalankan aplikasi
```bash
go run main.go
```

Jika berhasil, akan muncul pesan:
```
✅ Terhubung ke database PostgreSQL
🚀 Server Gin berjalan di http://localhost:8080
```

## 🌐 API Endpoints

Base URL: `http://localhost:8080/api/v1/todos`

### GET /todos
Mengambil semua todo.  
Query opsional: `?done=true` atau `?done=false`.

**Contoh request:**
```bash
curl http://localhost:8080/api/v1/todos
```

**Contoh response:**
```json
[
  {"id": 1, "title": "Belajar Go", "is_done": false},
  {"id": 2, "title": "Buat REST API", "is_done": true}
]
```

### GET /todos/:id
Mengambil satu todo berdasarkan ID.

**Contoh request:**
```bash
curl http://localhost:8080/api/v1/todos/1
```

**Contoh response (200):**
```json
{"id": 1, "title": "Belajar Go", "is_done": false}
```

**Contoh response (404):**
```json
{"error": "Todo tidak ditemukan"}
```

### POST /todos
Membuat todo baru.

**Contoh request:**
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go", "is_done": false}'
```

**Contoh response (201):**
```json
{"id": 3, "title": "Belajar Go", "is_done": false}
```

**Validasi:** `title` wajib diisi minimal 3 karakter.

### PUT /todos/:id
Memperbarui todo yang sudah ada.

**Contoh request:**
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go Advanced", "is_done": true}'
```

**Contoh response (200):**
```json
{"id": 1, "title": "Belajar Go Advanced", "is_done": true}
```

### DELETE /todos/:id
Menghapus todo berdasarkan ID.

**Contoh request:**
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

**Contoh response (200):**
```json
{"message": "Todo berhasil dihapus"}
```

## 🧪 Testing Manual (cURL)

```bash
# Lihat semua todo
curl http://localhost:8080/api/v1/todos

# Lihat todo yang selesai saja
curl "http://localhost:8080/api/v1/todos?done=true"

# Lihat satu todo
curl http://localhost:8080/api/v1/todos/1

# Tambah todo
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go", "is_done": false}'

# Update todo
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go", "is_done": true}'

# Hapus todo
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## 📝 Catatan Pengembangan

- Kode menggunakan `database/sql` dengan driver `pgx` untuk koneksi PostgreSQL.
- Semua response error menggunakan format seragam: `{"error": "pesan error"}`.
- Validasi input menggunakan `go-playground/validator/v10`.
- File `.env` **tidak** boleh di-commit ke version control (tambahkan ke `.gitignore`).

## 🚀 Roadmap (Opsional)

- [ ] Health check endpoint (`GET /health`)
- [ ] Pagination untuk daftar todo (`?page=1&limit=10`)
- [ ] Unit test untuk handler
- [ ] Dockerfile untuk deployment
- [ ] Middleware logging / CORS

---

Dibuat dengan ❤️ menggunakan Go & Gin.