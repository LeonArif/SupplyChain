# SupplyChain

Repository ini adalah implementasi prototipe supply-chain yang memadukan komponen algoritma rute/optimasi dan modul kriptografi (blockchain & signature) untuk eksperimen dan visualisasi.

Daftar isi
- Deskripsi singkat
- Fitur utama
- Struktur repo
- Persiapan & cara menjalankan
- Backend (Go)
- Frontend (React + Vite)
- Kriptografi (Kripto)
- Algoritma (Algorithm)
- Pengembangan & kontribusi

Deskripsi singkat
This project menyatukan dua domain utama:
- Algoritma optimasi dan visualisasi graf untuk perbandingan metode (dynamic programming, greedy).
- Modul kriptografi/blockchain untuk menyimpan blok, tanda tangan ECDSA, dan demonstrasi tampering.

Fitur utama
- Backend API ditulis dengan Go, menyediakan layanan supply-chain, storage, dan endpoint untuk eksperimen.
- Modul `kripto` berisi implementasi hashing, ECDSA, dan struktur blockchain sederhana.
- Modul `algorithm` berisi contoh implementasi algoritma DP dan greedy.
- Frontend menggunakan React + Vite (TypeScript) untuk visualisasi graf, perbandingan algoritma, dan eksplorasi blockchain.

Struktur repo (inti)
- `backend/` — kode Go untuk API, layanan, dan implementasi algoritma.
	- `backend/algorithm/` — `dp.go`, `greedy.go` (implementasi algoritma yang dapat diuji dan dibandingkan)
	- `backend/kripto/` — `blockchain.go`, `ecdsa.go`, `hash.go` (komponen kriptografi)
	- `backend/api/` — routing dan handler API
	- `backend/service/` — logika supply_chain dan service
- `frontend/` — aplikasi React + Vite (TS)
	- `frontend/src/components/AlgorithmLab` — komponen perbandingan algoritma dan visualisasi
	- `frontend/src/api/client.ts` — client yang berinteraksi dengan backend

Persiapan & ketergantungan
- Prasyarat:
	- Go (1.20+ direkomendasikan)
	- Node.js (16+) + npm/yarn/pnpm

Menjalankan backend (lokal)
1. Masuk ke direktori `backend`:

```powershell
cd backend
```

2. (Opsional) Set `GOCACHE` lokal agar cache Go tidak bercampur global:

PowerShell:
```powershell
$env:GOCACHE = (Join-Path (Get-Location) '.gocache')
```

Bash (WSL/git-bash):
```bash
export GOCACHE=$(pwd)/.gocache
```

3. Jalankan server:

```powershell
go run .
# atau dari root repo: go run ./backend
```

Catatan: repository sudah menambahkan `.gocache` ke `.gitignore` sehingga folder cache lokal tidak akan ter-commit.

Menjalankan frontend
1. Masuk ke `frontend` dan install dependensi:

```bash
cd frontend
npm install
```

2. Jalankan dev server:

```bash
npm run dev
```

Endpoint & interaksi
- Periksa `backend/api/routes.go` untuk daftar route API (handler utama dan endpoint yang tersedia).
- Client frontend menggunakan `frontend/src/api/client.ts` untuk memanggil API.

Kriptografi (Kripto)
--------------------
Folder: `backend/kripto`

Tujuan: menyediakan utilitas kriptografi untuk eksperimen blockchain, tanda tangan, dan hashing.

- `hash.go` — utilitas hashing (mis. SHA-256) yang digunakan untuk membuat fingerprint blok dan data.
- `ecdsa.go` — implementasi pembuatan pasangan kunci, penandatanganan, dan verifikasi menggunakan ECDSA (elliptic curve). Berguna untuk menandatangani transaksi atau metadata blok.
- `blockchain.go` — struktur blok sederhana, fungsi untuk menambah blok, verifikasi rantai, dan demonstrasi tamper detection.

Rekomendasi keamanan:
- Kode ini bertujuan edukasi dan eksperimen — jangan gunakan implementasi ini untuk produksi tanpa review keamanan pihak ketiga.
- Gunakan library kriptografi standar bila menyiapkan sistem nyata.

Algoritma (Algorithm)
----------------------
Folder: `backend/algorithm`

Tujuan: menampung contoh algoritma yang berkaitan dengan supply-chain, rute, dan optimasi.

- `dp.go` — contoh teknik dynamic programming untuk masalah yang relevan (mis. optimasi biaya, path with constraints). Gunakan untuk membandingkan solusi optimal vs heuristik.
- `greedy.go` — heuristik greedy yang cepat namun tidak selalu optimal; cocok untuk perbandingan performa dan kompleksitas waktu.

Frontend visualization
- Komponen di `frontend/src/components/AlgorithmLab` memvisualisasikan perbandingan performa, route reconstruction, dan chart performa (`PerformanceChart.tsx`).

Pengembangan & kontribusi
- Workflow singkat:
	1. Buat branch fitur: `git checkout -b feat/your-feature`
	2. Lakukan perubahan pada `backend` atau `frontend`
	3. Jalankan lokal: `go run ./backend` dan `npm run dev` di `frontend`
	4. Buka PR dan sertakan deskripsi perubahan serta contoh penggunaan

- Menambahkan tests: tambahkan file `_test.go` di `backend` untuk unit test Go; untuk frontend gunakan testing library sesuai stack (React Testing Library / Vitest).

Penutup
Jika Anda ingin, saya bisa:
- Menambahkan dokumentasi endpoint API yang diambil langsung dari `backend/api/routes.go`.
- Menulis contoh request cURL untuk endpoint penting.
- Menambahkan file CONTRIBUTING.md atau contoh skrip untuk menjalankan seluruh stack.

Terima kasih — beri tahu saya bagian mana yang mau diperluas lagi.
