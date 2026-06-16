# 🍱 FoodChain — Decentralized Food Supply Chain Tracker
## Project Plan & Technical Specification

---

## 📌 Deskripsi Sistem

**FoodChain** adalah sistem pelacakan rantai pasok makanan berbasis web yang mengintegrasikan dua domain akademik secara sinergis:

- **Strategi Algoritma (Stima):** Optimasi rute distribusi makanan menggunakan perbandingan algoritma **Greedy** dan **Dynamic Programming (DP)** pada model graf berarah multi-tahap.
- **Kriptografi:** Penjaminan **integritas dan immutability** data logistik pangan melalui struktur data **Blockchain lokal** berbasis SHA-256 dan ECDSA.

Sistem ini dibangun sebagai aplikasi web interaktif yang dapat mensimulasikan, memvisualisasikan, dan membuktikan keunggulan masing-masing pendekatan secara empiris — menjadi fondasi eksperimen untuk **dua makalah akademik** sekaligus.

---

## 🎯 Tujuan Sistem

| Tujuan | Keterangan |
|---|---|
| Akademik Stima | Membuktikan perbedaan performa Greedy vs DP dalam konteks distribusi pangan dengan constraint deadline (expiry date) |
| Akademik Kripto | Membuktikan sifat tamper-proof blockchain melalui simulasi serangan dan validasi hash chain |
| Praktis | Menyediakan sistem pelacakan rantai pasok makanan yang aman, transparan, dan dapat diaudit |

---

## 🏗️ Arsitektur Sistem

```
┌─────────────────────────────────────────────────────────┐
│                     FRONTEND (React + TypeScript)        │
│  ┌───────────────┐ ┌─────────────────┐ ┌─────────────┐  │
│  │ Dashboard     │ │ Algorithm Lab   │ │ Blockchain  │  │
│  │ (Input Data)  │ │ (Greedy vs DP)  │ │ Explorer    │  │
│  └───────────────┘ └─────────────────┘ └─────────────┘  │
└────────────────────────┬────────────────────────────────┘
                         │ REST API (HTTP/JSON)
┌────────────────────────▼────────────────────────────────┐
│                     BACKEND (Go)                         │
│  ┌──────────────────┐      ┌──────────────────────────┐  │
│  │ Algorithm Layer  │      │ Kripto Layer              │  │
│  │ ├── greedy.go    │      │ ├── blockchain.go         │  │
│  │ └── dp.go        │      │ ├── ecdsa.go              │  │
│  └──────────────────┘      │ └── hash.go               │  │
│                             └──────────────────────────┘  │
│  ┌──────────────────────────────────────────────────────┐  │
│  │ Service Layer (supply_chain.go) — Orchestrator       │  │
│  └──────────────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────────────┐  │
│  │ API Layer (routes.go) — HTTP Handler                 │  │
│  └──────────────────────────────────────────────────────┘  │
│  ┌──────────────────────────────────────────────────────┐  │
│  │ Models Layer (models.go) — Struct Definitions        │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────┘
```

---

## 📁 Struktur Direktori Final

```
SUPPLYCHAIN/
├── backend/
│   ├── algorithm/
│   │   ├── greedy.go          # Greedy EDF scheduling
│   │   └── dp.go              # DP cost/time optimization
│   ├── kripto/
│   │   ├── blockchain.go      # Block struct, chain management, validation
│   │   ├── ecdsa.go           # Key generation, signing, verification
│   │   └── hash.go            # SHA-256 hashing untuk product fingerprint
│   ├── service/
│   │   └── supply_chain.go    # Orchestrator: hubungkan algo + kripto
│   ├── api/
│   │   └── routes.go          # REST endpoint handler
│   ├── models/
│   │   └── models.go          # Struct: Food, Block, Order, Route, Node
│   ├── storage/
│   │   └── store.go           # In-memory chain & graph storage
│   └── main.go                # Entry point, server init
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Dashboard/
│   │   │   │   ├── FoodInputForm.tsx      # Form input data makanan
│   │   │   │   ├── FoodTable.tsx          # Tabel daftar makanan terdaftar
│   │   │   │   └── CourierCheckIn.tsx     # Form check-in kurir (suhu, lokasi)
│   │   │   ├── Graph/
│   │   │   │   ├── GraphCanvas.tsx        # Visualisasi graf interaktif
│   │   │   │   ├── NodeEditor.tsx         # Form tambah/edit node
│   │   │   │   └── EdgeWeightEditor.tsx   # Editor bobot cost & time
│   │   │   ├── AlgorithmLab/
│   │   │   │   ├── AlgoComparison.tsx     # Panel DP vs Greedy side-by-side
│   │   │   │   ├── RouteVisualizer.tsx    # Highlight rute di graf
│   │   │   │   └── PerformanceChart.tsx   # Line chart: Node count vs Time(ms)
│   │   │   └── BlockchainExplorer/
│   │   │       ├── BlockCard.tsx          # Card satu blok blockchain
│   │   │       ├── ChainViewer.tsx        # Tampilan chain secara keseluruhan
│   │   │       └── TamperSimulator.tsx    # Tombol & UI serangan siber
│   │   ├── api/
│   │   │   └── client.ts                  # Axios/fetch wrapper ke backend
│   │   ├── types/
│   │   │   └── index.ts                   # TypeScript type definitions
│   │   ├── App.tsx
│   │   └── main.tsx
│   ├── public/
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json
│   └── vite.config.ts
└── README.md
```

---

## 🖥️ Modul Frontend — Detail Spesifikasi

### Modul 1: Dashboard Utama (Manajemen & Input Data Makanan)

**Tujuan:** Antarmuka input data bagi Admin Logistik atau Kurir.

#### Sub-fitur: Input Data Makanan

| Field | Tipe | Keterangan |
|---|---|---|
| ID Barang | String (auto) | Generate otomatis: `FOOD-001`, `FOOD-002`, dst |
| Nama Makanan | String | Contoh: `Daging Sapi Segar`, `Susu Pasteur` |
| Expiry Date | Date Picker | Tanggal kedaluwarsa |
| Sisa Hari | Number (auto) | Dihitung: `expiry_date - today` |
| Lokasi Tujuan | Dropdown | Pilih node retailer dari graf |
| Berat (kg) | Number | Untuk constraint kapasitas DP |
| Urgensi | 1–10 | Nilai urgensi pengiriman |

**Logika Klasifikasi Status:**
```
Sisa Hari > 7   → 🟢 NORMAL   → DP mode: Cost Minimization
Sisa Hari 3–7   → 🟡 WARNING  → DP mode: Hybrid
Sisa Hari < 3   → 🔴 CRITICAL → DP mode: Time Minimization (rute tercepat)
```

#### Sub-fitur: Check-In Kurir

Form yang diisi kurir di setiap checkpoint:

| Field | Tipe | Keterangan |
|---|---|---|
| ID Barang | Dropdown | Pilih barang yang di-check-in |
| Lokasi Saat Ini | Dropdown | Node saat ini di graf |
| Suhu Aktual (°C) | Number | Suhu penyimpanan saat ini |
| Kelembapan (%) | Number | Kelembapan container |
| Timestamp | Auto | Diisi otomatis saat submit |
| Digital Signature | Auto | ECDSA sign oleh sistem atas nama kurir |

**Aksi submit** → data langsung dibungkus menjadi blok baru di blockchain.

---

### Modul 2: Graf Dinamis & Kustomisasi Rute

**Model Graf:** Multi-Stage Directed Graph (Graf Berarah Multi-Tahap)

```
Stage 0          Stage 1           Stage 2          Stage 3
(Sumber)         (Gudang Transit)  (Hub Regional)   (Retailer)

[Pabrik A] ──→ [Gudang Jakarta] ──→ [Hub Bandung] ──→ [Retailer 1]
               [Gudang Surabaya] ──→ [Hub Semarang] ──→ [Retailer 2]
                                                    ──→ [Retailer 3]
```

**Default Graph (bawaan sistem):**
- 1 node sumber (Pabrik)
- 2 node gudang transit
- 2 node hub regional
- 3 node retailer akhir

**Fitur Kustomisasi:**

| Fitur | Keterangan |
|---|---|
| Tambah Node | Form: pilih stage, nama node, tambahkan edge ke node lain |
| Edit Bobot Edge | Dua nilai per edge: **Cost** (Rp) dan **Time** (jam) |
| Generate Random Weights | Tombol randomize semua bobot untuk skenario eksperimen |
| Reset ke Default | Kembalikan graf ke konfigurasi awal |

**Visualisasi Graf:**
- Node berbentuk lingkaran, dibedakan warna per stage
- Edge ditampilkan dengan label `Cost / Time`
- Rute terpilih di-highlight berbeda per algoritma

---

### Modul 3: Algorithm Lab — Eksperimen DP vs Greedy

**Tujuan:** Membuktikan perbedaan performa dan output Greedy vs DP secara visual dan empiris.

#### Cara Kerja

**Input:** Daftar makanan yang sudah diinput di Dashboard + struktur graf saat ini.

**Greedy Algorithm:**
- Strategi: Pilih selalu edge dengan nilai terkecil secara lokal (greedy local optimal)
- Mode: selalu minimasi biaya OR waktu, tidak adaptif
- Kompleksitas: O(E log E) per produk

**Dynamic Programming:**
- Strategi: Hitung semua kemungkinan jalur optimal dari source ke destination per stage
- Mode **adaptif berdasarkan sisa hari:**
  - Sisa hari banyak → minimasi **Cost**
  - Sisa hari kritis → minimasi **Time**
- Kompleksitas: O(V × E) per produk

#### Tampilan Hasil

```
┌─────────────────────────┬─────────────────────────┐
│      GREEDY RESULT      │        DP RESULT         │
│  ─────────────────────  │  ─────────────────────   │
│  Rute: A→G1→H1→R2       │  Rute: A→G2→H2→R2        │
│  Total Cost: Rp 850.000  │  Total Cost: Rp 620.000  │
│  Total Time: 14 jam      │  Total Time: 11 jam      │
│  Exec Time: 0.8 ms       │  Exec Time: 2.3 ms       │
│  [Garis MERAH di graf]   │  [Garis HIJAU di graf]   │
└─────────────────────────┴─────────────────────────┘
```

#### Performance Chart

Line chart interaktif (Recharts):
- **Sumbu X:** Jumlah node dalam graf (5, 10, 15, 20, 25, 30)
- **Sumbu Y:** Waktu eksekusi (ms)
- **Dua garis:** Greedy (merah) vs DP (hijau)
- **Tujuan:** Membuktikan pertumbuhan kompleksitas secara empiris

---

### Modul 4: Blockchain Explorer & Tamper Simulation

**Tujuan:** Membuktikan integritas dan immutability data logistik pangan.

#### Struktur Blok

```go
type Block struct {
    Index        int
    Timestamp    string
    Data         TransactionData
    PrevHash     string    // SHA-256 dari blok sebelumnya
    Hash         string    // SHA-256(Index + Timestamp + Data + PrevHash)
    Signature    string    // ECDSA sign(Hash, private_key_kurir)
}

type TransactionData struct {
    FoodID       string
    FoodName     string
    Location     string    // Node saat ini
    Temperature  float64   // Suhu aktual (°C)
    Humidity     float64   // Kelembapan (%)
    ExpiryDate   string
    CourierID    string
    EventType    string    // "DEPARTURE" | "ARRIVAL" | "CHECK_IN"
}
```

#### Tampilan Block Card

Setiap blok ditampilkan sebagai kartu dengan:
- Header: `Block #3 — 2026-06-14 10:32:01`
- Body: semua field TransactionData
- Footer: `Prev Hash: a3f2...` | `Curr Hash: 9c1e...`
- Status badge: `✅ VALID` atau `❌ INVALID`

#### Simulasi Serangan Siber (Tamper Test)

**Skenario:** Kurir ingin memanipulasi data suhu dari 15°C → 3°C agar tidak disalahkan atas kerusakan daging.

**Alur simulasi:**
1. Klik tombol **"⚠️ Edit Data Secara Ilegal"** pada suatu blok
2. Modal muncul: edit field suhu (atau field lain)
3. Klik **"Konfirmasi Manipulasi"**
4. Backend rekalkulasi hash blok tersebut → hash berubah
5. Blok berikutnya `prev_hash` tidak cocok → chain patah
6. Frontend response:
   - Blok yang diretas + semua blok setelahnya → **berubah warna merah**
   - Banner besar muncul: **❌ BLOCKCHAIN CORRUPTED — CHAIN INTEGRITY VIOLATED**
   - Detail: `Block #3 hash mismatch: expected a3f2... got 7b91...`
7. Tombol **"🔄 Restore Chain"** untuk reset ke kondisi valid

---

## ⚙️ Backend — Detail Teknis

### models/models.go

```go
// Node dalam graf
type Node struct {
    ID    string
    Name  string
    Stage int    // 0=Sumber, 1=Gudang, 2=Hub, 3=Retailer
}

// Edge dalam graf
type Edge struct {
    From string
    To   string
    Cost float64  // Biaya (Rp)
    Time float64  // Waktu (jam)
}

// Data makanan
type Food struct {
    ID          string
    Name        string
    ExpiryDate  time.Time
    DaysLeft    int
    Destination string
    Weight      float64
    Urgency     int
    Status      string  // "NORMAL" | "WARNING" | "CRITICAL"
}

// Hasil rute algoritma
type RouteResult struct {
    Algorithm   string    // "GREEDY" | "DP"
    Path        []string  // Urutan node ID
    TotalCost   float64
    TotalTime   float64
    ExecTimeMs  float64
    Mode        string    // "COST" | "TIME"
}
```

### algorithm/greedy.go

```
Fungsi: GreedyRoute(graph, source, destination, mode)
- mode = "COST" → pilih edge dengan cost terkecil di setiap langkah
- mode = "TIME" → pilih edge dengan time terkecil di setiap langkah
- Return: RouteResult
```

### algorithm/dp.go

```
Fungsi: DPRoute(graph, source, destination, food)
- Jika food.DaysLeft > 7  → optimize Cost
- Jika food.DaysLeft <= 3 → optimize Time  
- Hitung DP table per stage: dp[node] = min cost/time untuk mencapai node tsb
- Return: RouteResult dengan path reconstruction
```

### kripto/hash.go

```
Fungsi: HashBlock(block) → string
- Concatenate: Index + Timestamp + JSON(Data) + PrevHash
- Return: SHA-256 hex string

Fungsi: HashFood(food) → string
- Fingerprint produk: SHA-256(ID + Name + ExpiryDate)
```

### kripto/ecdsa.go

```
Fungsi: GenerateKeyPair() → (privateKey, publicKey)
Fungsi: SignData(data, privateKey) → signature string
Fungsi: VerifySignature(data, signature, publicKey) → bool
- Menggunakan curve: P-256 (secp256r1)
```

### kripto/blockchain.go

```
Fungsi: AddBlock(chain, data, privateKey) → Block
- Hash blok baru: SHA-256(prev_hash + data + timestamp)
- Sign hash dengan private key kurir
- Append ke chain

Fungsi: ValidateChain(chain) → (bool, invalidIndex)
- Loop setiap blok: cek curr_hash == SHA-256(block_fields)
- Cek prev_hash blok[i] == hash blok[i-1]
- Return false + index pertama yang tidak valid

Fungsi: TamperBlock(chain, index, newData) → chain
- Ubah data blok[index] tanpa recalculate hash (simulasi serangan)
```

---

## 🔌 REST API Endpoints

### Food Management
| Method | Endpoint | Keterangan |
|---|---|---|
| `GET` | `/api/foods` | Ambil semua data makanan |
| `POST` | `/api/foods` | Tambah makanan baru |
| `DELETE` | `/api/foods/:id` | Hapus makanan |

### Graph Management
| Method | Endpoint | Keterangan |
|---|---|---|
| `GET` | `/api/graph` | Ambil struktur graf saat ini |
| `POST` | `/api/graph/node` | Tambah node baru |
| `POST` | `/api/graph/edge` | Tambah/edit edge |
| `POST` | `/api/graph/randomize` | Randomize semua bobot |
| `POST` | `/api/graph/reset` | Reset ke default |

### Algorithm
| Method | Endpoint | Keterangan |
|---|---|---|
| `POST` | `/api/algo/compare` | Jalankan Greedy + DP, return keduanya |
| `POST` | `/api/algo/benchmark` | Benchmark dengan N node berbeda, return data chart |

### Blockchain
| Method | Endpoint | Keterangan |
|---|---|---|
| `GET` | `/api/chain` | Ambil semua blok |
| `POST` | `/api/chain/checkin` | Tambah blok check-in kurir |
| `GET` | `/api/chain/validate` | Validasi seluruh chain |
| `POST` | `/api/chain/tamper` | Simulasi manipulasi blok |
| `POST` | `/api/chain/restore` | Restore chain ke kondisi valid |

---

## 🧪 Rencana Eksperimen

### Eksperimen 1 — Stima: Greedy vs DP (Perbandingan Kualitas Hasil)

**Setup:**
- Graf dengan 3 skenario: Normal (sisa hari > 7), Warning (3–7 hari), Critical (< 3 hari)
- 10 produk per skenario dengan variasi berat dan urgensi

**Metrik:**
- Total cost rute (Rp)
- Total time rute (jam)
- Apakah deadline terpenuhi

**Hipotesis:** DP selalu menghasilkan rute lebih optimal, Greedy lebih cepat namun suboptimal.

---

### Eksperimen 2 — Stima: Kompleksitas Waktu (Benchmark)

**Setup:**
- Variasi jumlah node: 5, 10, 15, 20, 25, 30
- Setiap konfigurasi dijalankan 10 kali, diambil rata-rata
- Graf di-generate secara acak (random weights)

**Metrik:**
- Waktu eksekusi (ms) untuk Greedy dan DP
- Pertumbuhan relatif setiap penambahan node

**Output:** Line chart Node Count vs Execution Time (ms)

---

### Eksperimen 3 — Kripto: Tamper Detection

**Setup:**
- Buat chain dengan 10 blok valid
- Manipulasi blok ke-3, ke-5, dan ke-8 secara terpisah
- Jalankan validasi

**Metrik:**
- Apakah manipulasi terdeteksi (true positive rate: harus 100%)
- Index blok pertama yang terdeteksi invalid
- Waktu deteksi (ms)

**Hipotesis:** Setiap manipulasi sekecil apapun pasti terdeteksi karena sifat SHA-256.

---

### Eksperimen 4 — Kripto: Overhead Kriptografi

**Setup:**
- Ukur waktu operasi: hashing, signing ECDSA, verification
- Bandingkan throughput: dengan kriptografi vs tanpa

**Metrik:**
- Waktu hash per blok (ms)
- Waktu signing per blok (ms)
- Waktu verifikasi per blok (ms)
- Transaksi per detik (TPS) dengan dan tanpa kripto

---

## 📋 Judul Makalah

### Makalah Strategi Algoritma (IF2211)

> **"Comparative Analysis of Greedy and Dynamic Programming Algorithms for Deadline-Aware Route Optimization in Food Supply Chain Distribution"**

atau versi Indonesia:

> **"Analisis Komparatif Algoritma Greedy dan Dynamic Programming untuk Optimasi Rute Berbasis Tenggat Waktu pada Distribusi Rantai Pasok Makanan"**

**Topik yang dicakup:** Dynamic Programming + Greedy (kategori i & b dari daftar topik)

---

### Makalah Kriptografi (II4021)

> **"Blockchain-Based Data Integrity and Tamper Detection for Food Supply Chain Traceability Using SHA-256 Hash Chain and ECDSA Digital Signature"**

atau versi Indonesia:

> **"Integritas Data dan Deteksi Manipulasi Berbasis Blockchain untuk Keterlacakan Rantai Pasok Makanan Menggunakan SHA-256 Hash Chain dan Tanda Tangan Digital ECDSA"**

**Topik yang dicakup:** Blockchain + Fungsi Hash + Tanda Tangan Digital (kategori 7, 8, 13)

---

## 🗓️ Timeline Pengerjaan

| Fase | Durasi | Output |
|---|---|---|
| **Fase 1:** Setup & Models | 1 hari | Struktur proyek, models.go, types/index.ts |
| **Fase 2:** Kripto Layer | 2 hari | hash.go, ecdsa.go, blockchain.go |
| **Fase 3:** Algorithm Layer | 2 hari | greedy.go, dp.go |
| **Fase 4:** Service & API | 1 hari | supply_chain.go, routes.go, main.go |
| **Fase 5:** Frontend — Dashboard | 1 hari | FoodInputForm, FoodTable, CourierCheckIn |
| **Fase 6:** Frontend — Graf | 1 hari | GraphCanvas, NodeEditor, EdgeWeightEditor |
| **Fase 7:** Frontend — Algo Lab | 1 hari | AlgoComparison, RouteVisualizer, PerformanceChart |
| **Fase 8:** Frontend — Blockchain | 1 hari | BlockCard, ChainViewer, TamperSimulator |
| **Fase 9:** Eksperimen & Data | 1 hari | Jalankan semua eksperimen, catat hasil |
| **Fase 10:** Penulisan Makalah | 2 hari | Draft makalah Stima + Kripto |

**Total estimasi: ~13 hari** (sebelum deadline 19 Juni 2026)

---

## ✅ Checklist Persyaratan Akademik

### Stima (IF2211)
- [ ] Algoritma yang dibahas termasuk dalam daftar topik (Greedy ✓, DP ✓)
- [ ] Ada kontribusi berupa implementasi + eksperimen
- [ ] Bukan studi literatur murni
- [ ] Makalah 6–10 halaman format IEEE
- [ ] Ditulis dalam Bahasa Indonesia atau Inggris
- [ ] Tidak ada Wikipedia sebagai referensi
- [ ] Ada pernyataan anti-plagiarisme + tanda tangan digital
- [ ] Email @std.stei.itb.ac.id dan @gmail.com tercantum
- [ ] Judul didaftarkan di spreadsheet sebelum 13 Juni 2026

### Kriptografi (II4021)
- [ ] Topik termasuk dalam daftar (Blockchain ✓, Hash ✓, Digital Signature ✓)
- [ ] Ada kontribusi: rancangan, implementasi, eksperimen
- [ ] Bukan studi literatur murni
- [ ] Makalah minimal 6 halaman format IEEE
- [ ] Tidak ada Wikipedia sebagai referensi
- [ ] Ada pernyataan anti-plagiarisme + tanda tangan digital
- [ ] Judul didaftarkan di spreadsheet sebelum 13 Juni 2026
- [ ] Dikumpulkan ke Google Drive yang benar sebelum 19 Juni 2026

---

*Dokumen ini adalah living document — update sesuai perkembangan implementasi.*
