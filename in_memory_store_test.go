package main

import "testing"

func TestSetAndGet(t *testing.T) {
	store := NewInMemoryStore()
	// Yeni, boş bir in-memory store oluşturuyoruz.

	store.Set("testkey", "testvalue")
	// "testkey" anahtarına "testvalue" değerini ekliyoruz.

	value, exists := store.Get("testkey")
	// Şimdi "testkey" anahtarını okuyup, hem değeri hem de var mı diye kontrol ediyoruz.

	if !exists {
		t.Error("Anahtar bulunamadı!")
		// Eğer exists false ise, yani "testkey" eklenmemiş gibi davranıyorsa test başarısız.
	}

	if value != "testvalue" {
		t.Errorf("Beklenen değer 'testvalue', ama '%s' bulundu", value)
		// "testkey" anahtarının değeri "testvalue" değilse, hata mesajı veriyoruz.
	}

	// Olmayan bir anahtar deniyoruz.
	_, exists = store.Get("not-exist")
	if exists {
		t.Error("Var olmayan anahtar için true dönüyor!")
		// Hiç eklenmemiş bir anahtar sorduğumuzda exists true ise hata var demektir.
	}
}
