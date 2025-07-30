package main

// Gerekli kütüphaneleri import ediyoruz.
import (
	"fmt"      // Konsola bilgi yazdırmak için kullanılır.
	"log"      // Hata ve bilgi loglarını yazmak için kullanılır.
	"net/http" // HTTP sunucusu ve istekleri yönetmek için kullanılır.
)

// Global bir değişken olarak InMemoryStore nesnemizi oluşturuyoruz.
var store = NewInMemoryStore() // in_memory_store.go dosyasındaki yapıyı burada kullanıyoruz.

func main() {
	// MongoDB bağlantısını başlatıyoruz.
	initMongoDB()

	// API uç noktalarını tanımlıyoruz.
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/in-memory", inMemoryHandler)
	http.HandleFunc("/mongo-records", mongoHandler)

	// Sunucunun başladığını belirtiyoruz ve gelen istekleri dinliyoruz.
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Hata olursa programı sonlandırır.
}

// Sağlık kontrolü endpoint'i. Sunucunun çalıştığını basit bir mesajla dönüyoruz.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running!")
}
