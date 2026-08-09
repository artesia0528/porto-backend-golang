# Panduan Deploy ke VPS

Panduan lengkap untuk deploy portfolio backend ke VPS (Ubuntu/Debian) menggunakan Docker.

## Prasyarat di VPS

- VPS dengan Ubuntu 22.04+ atau Debian 12+
- Minimal RAM 1GB
- Akses SSH ke VPS
- Domain yang sudah di-point ke IP VPS (opsional, tapi direkomendasikan)

---

## Step 1: Install Docker di VPS

SSH ke VPS kamu, lalu jalankan:

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com | sudo sh

# Tambahkan user ke group docker (supaya tidak perlu sudo)
sudo usermod -aG docker $USER

# Logout dan login lagi supaya group berlaku
exit
```

Setelah login kembali, verifikasi:

```bash
docker --version
docker compose version
```

---

## Step 2: Upload Project ke VPS

### Opsi A: Via Git (Rekomendasi)

```bash
# Di VPS
cd ~
git clone <URL_REPO_KAMU> portfolio-backend
cd portfolio-backend
```

### Opsi B: Via SCP (jika belum pakai Git)

```bash
# Di mesin lokal kamu
scp -r /home/aditya/exercise/portfolio/porto-backend-golang user@IP_VPS:~/portfolio-backend
```

---

## Step 3: Buat file `.env` di VPS

```bash
cd ~/portfolio-backend
cp .env.example .env
nano .env
```

Isi dengan value production:

```env
PORT=8080
JWT_SECRET=GANTI_DENGAN_STRING_RANDOM_YANG_PANJANG_MINIMAL_32_KARAKTER
APP_ENV=production

DATABASE_URL=host=db user=postgres password=GANTI_PASSWORD_KUAT dbname=portfolio port=5432 sslmode=disable

POSTGRES_USER=postgres
POSTGRES_PASSWORD=GANTI_PASSWORD_KUAT
POSTGRES_DB=portfolio
```

> [!CAUTION]
> - **GANTI** `JWT_SECRET` dengan string random yang panjang. Generate dengan: `openssl rand -hex 32`
> - **GANTI** `POSTGRES_PASSWORD` dengan password yang kuat. Generate dengan: `openssl rand -base64 24`
> - Pastikan `POSTGRES_PASSWORD` di baris `DATABASE_URL` dan `POSTGRES_PASSWORD` **SAMA**

---

## Step 4: Deploy!

```bash
# Build & start semua services
make prod-up

# Atau jika make belum terinstall:
docker compose -f docker-compose.prod.yml up -d --build
```

Cek apakah semua berjalan:

```bash
# Lihat status containers
docker compose -f docker-compose.prod.yml ps

# Lihat logs
docker compose -f docker-compose.prod.yml logs -f
```

Test API:

```bash
curl http://localhost:8080/api/experiences
# Harus return: {"success":true,"message":"Berhasil","data":[]}
```

---

## Step 5: Setup Nginx Reverse Proxy (Rekomendasi)

Nginx berfungsi sebagai "pintu depan" yang menerima traffic dari internet di port 80/443, lalu meneruskannya ke app di port 8080.

```bash
# Install Nginx
sudo apt install nginx -y

# Buat config
sudo nano /etc/nginx/sites-available/portfolio-api
```

Isi dengan:

```nginx
server {
    listen 80;
    server_name api.domainmu.com;  # Ganti dengan domain/IP kamu

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Aktifkan:

```bash
sudo ln -s /etc/nginx/sites-available/portfolio-api /etc/nginx/sites-enabled/
sudo nginx -t          # Test config
sudo systemctl reload nginx
```

---

## Step 6: SSL / HTTPS (Gratis via Let's Encrypt)

```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx -y

# Generate SSL (otomatis update config Nginx)
sudo certbot --nginx -d api.domainmu.com
```

Certbot akan otomatis renew sertifikat sebelum expired.

---

## Perintah Berguna

| Perintah | Fungsi |
|---|---|
| `docker compose -f docker-compose.prod.yml logs -f` | Lihat logs realtime |
| `docker compose -f docker-compose.prod.yml logs app` | Logs app saja |
| `docker compose -f docker-compose.prod.yml restart app` | Restart app tanpa restart DB |
| `docker compose -f docker-compose.prod.yml down` | Stop semua |
| `docker compose -f docker-compose.prod.yml up -d --build` | Rebuild & deploy ulang |

## Update/Deploy Ulang

Ketika ada perubahan code:

```bash
cd ~/portfolio-backend
git pull                    # Ambil perubahan terbaru
make prod-up                # Rebuild & restart
```

---

## Troubleshooting

### App tidak bisa konek ke database
```bash
# Cek apakah DB container sehat
docker compose -f docker-compose.prod.yml ps

# Lihat log DB
docker compose -f docker-compose.prod.yml logs db
```

### Port 8080 sudah dipakai
Ubah `PORT` di `.env`, dan update Nginx proxy_pass sesuai port baru.

### Mau reset database total
```bash
docker compose -f docker-compose.prod.yml down -v   # -v = hapus volume/data
docker compose -f docker-compose.prod.yml up -d --build
```
