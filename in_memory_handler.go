package main

import (
	"encoding/json" // JSON verilerini işlemek için kullanıyoruz.
	"net/http"
)

func inMemoryHandler(w http.ResponseWriter, r *http.Request) {
	// Frontendle çalışırken karşılaşılan hatayı gidermek için böyle bir yol izliyoruz.
	w.Header().Set("Access-Control-Allow-Origin", "*") // Farklı kaynaklardan gelen isteklere izin veriyoruz.
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS") // Hangi methodlara ve
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // headerlara izin oldugunu belirtiyoruz.
		w.WriteHeader(http.StatusOK)                                         // İşlemin olumlu olduğunu belirtiyoruz.
		return
	}

	if r.Method == http.MethodPost {
		handlePost(w, r) // Gelen istek POST ise handlePost methodu ile veri kaydetme işlemi yapılacağını
	} else if r.Method == http.MethodGet {
		handleGet(w, r) // GET ise handleGet methodu ile değer çekme işlemi yapılacağını belirtiyoruz.
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // Diğer methodlar için hata dönüyoruz.
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Key   string `json:"key"`   // JSON'daki "key" alanını bu değişkene yazıyoruz.
		Value string `json:"value"` // "value" alanını bu değişkene yazıyoruz.
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body) // İstek bodysini alıp go struct'ına çeviriyoruz.
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest) // Eğer JSON formatı yanlışsa hata dönüyoruz.
		return
	}

	store.Set(body.Key, body.Value) // Girilen anahtar ve değeri store adlı değişkene kaydediyoruz.

	w.Header().Set("Content-Type", "application/json") // Yanıtın JSON olduğunu belirtiyoruz.
	json.NewEncoder(w).Encode(body)                    // İstekle gelen veriyi tekrar yanıt olarak dönüyoruz.
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key") // URL'den "key" parametresini alıyoruz.
	if key == "" {
		http.Error(w, "Missing 'key' parameter", http.StatusBadRequest) // Key yoksa hata dönüyoruz.
		return
	}

	value, exists := store.Get(key) // Post işleminde oluşturmuş olduğumuz store adlı değişkenden belirtilen key içindeki değeri çekiyoruz.
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound) // Belirtilen anahtar store adlı değişkende bulunamazsa hata dönüyoruz.
		return
	}

	response := map[string]string{ // Anahtar bulunursa key-value şeklinde dönüyoruz.
		"key":   key,
		"value": value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
