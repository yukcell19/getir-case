package main

// Gerekli kütüphaneleri import ediyoruz.
import (
	"fmt"      // Konsola bilgi yazdırmak için kullanılır.
	"log"      // Hata ve bilgi loglarını yazmak için kullanılır.
	"net/http" // HTTP sunucusu ve istekleri yönetmek için kullanılır.
	"os"

	"github.com/joho/godotenv"
)

// Global bir değişken olarak InMemoryStore nesnemizi oluşturuyoruz.
var store = NewInMemoryStore() // in_memory_store.go dosyasındaki yapıyı burada kullanıyoruz.

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// MongoDB bağlantısını başlatıyoruz.
	initMongoDB()

	// API uç noktalarını tanımlıyoruz.
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/in-memory", inMemoryHandler)
	http.HandleFunc("/mongo-records", mongoHandler)

	// Heroku'dan port alıyoruz, yoksa 8080 kullanıyoruz.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Sunucunun başladığını belirtiyoruz ve gelen istekleri dinliyoruz.
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil)) // Hata olursa programı sonlandırır.
}

// Sağlık kontrolü endpoint'i. Sunucunun çalıştığını basit bir mesajla dönüyoruz.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running!")
}
