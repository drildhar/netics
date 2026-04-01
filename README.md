# Laporan Penugasan Modul 1 - Open Recruitment NETICS 2026

**Nama:** Adriel Mahira Dharma  
**NRP:** 5025241097

---

## 1. Deskripsi Tugas
Implementasi sistem CI/CD sederhana yang mencakup pembuatan API publik, containerization menggunakan Docker, konfigurasi web server (Nginx) menggunakan Ansible sebagai reverse proxy, dan otomasi deployment menggunakan GitHub Actions.

## 2. Implementasi API
API dibangun menggunakan bahasa pemrograman **Golang** dengan spesifikasi sebagai berikut:
- **Endpoint:** `/health`
- **Port Internal:** `3000`
- **Fitur:** Menampilkan Nama, NRP, Status (UP), Timestamp saat ini, dan durasi Uptime server.

## 3. Infrastruktur & Containerization
### Docker
Aplikasi dibungkus ke dalam Docker image menggunakan *multi-stage build* untuk efisiensi ukuran image.
- **Docker Image:** `docker.io/redacted/netics-api:latest`

### Nginx & Ansible
Ansible digunakan untuk mengotomasi konfigurasi VPS agar:
1. Menginstal Nginx.
2. Mengonfigurasi Nginx sebagai **Reverse Proxy** yang meneruskan request dari port 80 ke port 3000 (API container).
3. Memastikan konfigurasi aktif tanpa intervensi manual.

## 4. Alur CI/CD (GitHub Actions)
Workflow CI/CD diatur untuk berjalan otomatis setiap kali ada `push` ke branch `main`:
1. **Build & Push:** Membuat image Docker baru dan mengunggahnya ke Docker Hub.
2. **Deploy:** Melakukan SSH ke VPS, menarik image terbaru, dan menjalankan ulang container.
3. **Configuration:** Menjalankan Ansible Playbook untuk memastikan Nginx selalu terkonfigurasi dengan benar.

## 5. Informasi Akses
- **URL API:** `http://redacted/health`
- **Port API:** `3000`
- **Status Deployment:** Tidak berhasil melalui GitHub Actions.
