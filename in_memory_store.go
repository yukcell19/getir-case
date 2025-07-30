package main

import "sync" // Eşzamanlı veri yönetimi için bu yapıyı kullanıyoruz.

// Basit bir anahtar-değer saklama yapısı oluşturuyoruz.
type InMemoryStore struct {
	data map[string]string // Veriler burada anahtar-değer olarak tutulur.
	mu   sync.RWMutex      // Veri okuma/yazma sırasında karışıklık olmaması ve eşzamanlılık için mutex adlı yapıyı kullanıyoruz.
}

// Yeni bir InMemoryStore nesnesi oluşturuyoruz.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]string),
	}
}

// Belirtilen anahtar için verilen değeri saklayan yapıyı oluşturuyoruz.
func (s *InMemoryStore) Set(key, value string) {
	s.mu.Lock()         // Yazma işlemi için kilitliyoruz.
	defer s.mu.Unlock() // Fonksiyon bitince kilidi açıyoruz.
	s.data[key] = value // Değeri map'e ekliyoruz.
}

// Belirtilen anahtar eğer varsa içerdiği değeri dönen, eğer anahtar yoksa false dönen yapıyı oluşturuyoruz.
func (s *InMemoryStore) Get(key string) (string, bool) {
	s.mu.RLock()         // Okuma işlemi için kilitliyoruz (yazmaya engel olmaz).
	defer s.mu.RUnlock() // Fonksiyon bitince kilidi açıyoruz.
	value, exists := s.data[key]
	return value, exists
}
