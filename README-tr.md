# ⚡ Go Eşzamanlı Sağlık Denetleyicisi

[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/ersozberk/async-scraper/edit/main/README.md)
[![pt-br](https://img.shields.io/badge/lang-tr-green.svg)](https://github.com/ersozberk/async-scraper/edit/main/README-tr.md)

![Go [Sürüm](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Eşzamanlılık](https://img.shields.io/badge/Concurrency-Goroutines_%26_Channels-blue)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)

Go ile oluşturulmuş, yıldırım hızında, eşzamanlı bir HTTP sağlık kontrol aracı. Bu proje, birden fazla web servisine aynı anda ping göndererek ve gerçek zamanlı durumlarını raporlayarak Go'nun eşzamanlılık modelinin gücünü gösteriyor.

## ✨ Özellikler

* **Goroutine'ler:** Birden fazla HTTP GET isteğini eşzamanlı olarak yürütür ve toplam yürütme süresini en yavaş isteğin süresine indirir.

* **Kanallar:** İş parçacığı güvenli veri toplama ve senkronizasyonu için tamponlanmış kanallar kullanır.

* **Zaman Aşımları ve Kaynak Yönetimi:** Sunucular yanıt vermediğinde goroutine sızıntılarını önlemek için `http.Client` zaman aşımı uygular.

* **Standart Go Proje Düzeni:** `cmd/` dizin yapısına uyar.

## 📂 Proje Yapısı

```metin
.

├── cmd/
│ └── scraper/
│ └── main.go # Temel mantık ve eşzamanlılık uygulaması
├── go.mod # Go modül bildirimleri
├── .gitignore # Git yok sayma kuralları
└── README.md # Proje dokümantasyonu
```

## 🛠 Kurulum ve Kullanım
Depoyu klonlayın:

```Bash
git clone [https://github.com/ersozberk/async-scraper.git](https://github.com/ersozberk/async-scraper.git)
cd async-scraper

```
Eşzamanlılık aracını doğrudan çalıştırın:
```Bash
go run cmd/scraper/main.go
```
Bağımsız bir uygulama oluşturun Çalıştırılabilir dosya:

```Bash
go build -o health-checker cmd/scraper/main.go
./health-checker
```
## 🧠 Öğrenme Çıktıları
1. Arka plan görev yürütme için goroutine'leri ustaca kullanma.

2. Kanallar ve sync.WaitGroup aracılığıyla durum ve senkronizasyonu yönetme.

3. Özel HTTP istemcileri ve zaman aşımı süreleriyle ağ güvenilirliği sorunlarını ele alma.

EOF
