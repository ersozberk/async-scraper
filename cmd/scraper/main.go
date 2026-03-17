package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// checkWebsite, hedefe HTTP isteği atıp durumunu kanala gönderir
func checkWebsite(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Eğer bir site çöktüyse programın sonsuza kadar beklemesini önlemek için
	// 5 saniyelik bir zaman aşımı (timeout) süresi belirliyoruz.
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Hedef siteye GET isteği atıyoruz
	resp, err := client.Get("https://" + url)
	if err != nil {
		ch <- fmt.Sprintf("❌ [%s] Ulaşılamıyor veya Çökmüş: %v", url, err)
		return
	}
	defer resp.Body.Close() // İşlem bitince bağlantıyı kapat

	// HTTP durum koduna göre (200 OK) başarılı olup olmadığını kontrol ediyoruz
	if resp.StatusCode == http.StatusOK {
		ch <- fmt.Sprintf("✅ [%s] Sorunsuz Çalışıyor (Status: %d)", url, resp.StatusCode)
	} else {
		ch <- fmt.Sprintf("⚠️ [%s] Hata Veriyor (Status: %d)", url, resp.StatusCode)
	}
}

func main() {
	baslangic := time.Now()
	fmt.Println("Sistem Sağlık Taraması Başlatılıyor...\n")

	// Kontrol edilecek hedefler
	siteler := []string{
		"google.com",
		"apple.com",
		"bing.com",
		"amazon.com",
	}

	sonucKanali := make(chan string, len(siteler))
	var wg sync.WaitGroup

	// Sitelerin hepsine AYNI ANDA istek atıyoruz
	for _, site := range siteler {
		wg.Add(1)
		go checkWebsite(site, sonucKanali, &wg)
	}

	// Bekleme ve Kanal Kapatma Goroutine'i
	go func() {
		wg.Wait()
		close(sonucKanali)
	}()

	// Gelen sonuçları anında ekrana basıyoruz
	for mesaj := range sonucKanali {
		fmt.Println(mesaj)
	}

	gecenSure := time.Since(baslangic)
	fmt.Printf("\n🚀 Tarama tamamlandı! Toplam süre: %v\n", gecenSure)
}
