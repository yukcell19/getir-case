package main

import (
	"context" // Context işlemleri ve timeout yönetimi için kullanıyoruz.
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time" // Tarih ve saat işlemleri için kullanıyoruz.

	"go.mongodb.org/mongo-driver/bson"          // BSON veri tipleri için kullanıyoruz.
	"go.mongodb.org/mongo-driver/mongo"         // MongoDB istemcisi için kullanıyoruz.
	"go.mongodb.org/mongo-driver/mongo/options" // MongoDB bağlantı ayarları için kullanıyoruz.
)

// Tüm fonksiyonlarda kullanabilmek için global koleksiyon değişkeni oluşturuyoruz
var mongoCollection *mongo.Collection

// MongoDB bağlantısını başlatıyoruz
func initMongoDB() {
	// MongoDB bağlantı URI'sini tanımlıyoruz
	uri := os.Getenv("MONGO_URI")

	// Bağlantı işlemi için context ve timeout ayarlıyoruz.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Fonksiyon bitince context'i iptal ediyoruz.

	// Bağlantı seçeneklerini belirliyoruz ve URI'yı ekliyoruz.
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB istemcisi ile bağlantı kuruyoruz.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err) // Bağlantı başarısızsa uygulamayı sonlandırıyoruz ve hata mesajı iletiyoruz.
	}

	// İlgili veritabanı ve koleksiyona erişiyoruz.
	mongoCollection = client.Database("getir-case-study").Collection("records")
	fmt.Println("MongoDB connection successful!")
}

// MongoDB'den filtrelenmiş kayıtları döndüren endpoint'i oluşturuyoruz
func mongoHandler(w http.ResponseWriter, r *http.Request) {
	// Frontendle çalışırken karşılaşılan hatayı gidermek için böyle bir yol izliyoruz.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Yalnızca POST metoduna izin veriyoruz.
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Gelen JSON isteğini okuyoruz ve struct'a aktarıyoruz.
	var req struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int    `json:"minCount"`
		MaxCount  int    `json:"maxCount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Tarihleri string'den go'daki time.Time tipine çeviriyoruz.
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		http.Error(w, "Invalid format for startDate", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		http.Error(w, "Invalid format for endDate", http.StatusBadRequest)
		return
	}

	// MongoDB aggregation pipeline oluşturuyoruz.
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{
			"createdAt": bson.M{"$gte": startDate, "$lte": endDate}, // Kayıtları tarih aralığına göre filtreliyoruz.
		}}},
		bson.D{{Key: "$addFields", Value: bson.M{
			"totalCount": bson.M{"$sum": "$counts"}, // Her kayda counts içindeki değerlerin toplamını ekliyoruz.
		}}},
		bson.D{{Key: "$match", Value: bson.M{
			"totalCount": bson.M{"$gte": req.MinCount, "$lte": req.MaxCount}, // Kayıtları toplam değere göre tekrar filtreliyoruz.
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":        0, // _id alanını döndürmüyoruz.
			"key":        1, // key alanını döndürüyoruz.
			"createdAt":  1, // createdAt alanını döndürüyoruz.
			"totalCount": 1, // totalCount alanını döndürüyoruz.
		}}},
	}

	// Sorguyu çalıştırmak için context ve timeout oluşturuyoruz.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Aggregation pipeline ile MongoDB sorgusu yapıyoruz.
	cursor, err := mongoCollection.Aggregate(ctx, pipeline)
	if err != nil {
		http.Error(w, "Error occurred during MongoDB query", http.StatusInternalServerError)
		log.Println("MongoDB query failed:", err)
		return
	}
	defer cursor.Close(ctx) // İş bitince cursor'u kapatıyoruz.

	// Sonuçları slice içine aktarıyoruz.
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		http.Error(w, "Could not read results", http.StatusInternalServerError)
		log.Println("Result reading failed:", err)
		return
	}

	// Sonuç yoksa boş bir slice döndürüyoruz.
	if results == nil {
		results = []bson.M{}
	}

	// Başarı ve hata durumuna göre kod ve mesaj belirliyoruz.
	code := 0
	msg := "Success"
	if len(results) == 0 {
		code = 1
		msg = "No records found"
	}

	// Yanıt objesini oluşturuyoruz.
	response := map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"records": results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
