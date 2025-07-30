document.addEventListener("DOMContentLoaded", function () {
    // İhtiyacımız olan tüm HTML elemanlarını seçiyoruz.
    const mongoForm = document.getElementById("mongo-form");
    const mongoOutput = document.getElementById("mongo-output");

    const inmemoryPostForm = document.getElementById("inmemory-post-form");
    const inmemoryPostOutput = document.getElementById("inmemory-post-output");

    const inmemoryGetForm = document.getElementById("inmemory-get-form");
    const inmemoryGetOutput = document.getElementById("inmemory-get-output");

    mongoForm.addEventListener("submit", async function (e) {
        e.preventDefault(); // Submit butonunun varsayılan davranışını değiştirerek sayfanın yenilenmesini engelliyoruz.

        // Formdan girilen tarih ve sayısal değerleri çekiyoruz.
        const startDate = document.getElementById("startDate").value;
        const endDate = document.getElementById("endDate").value;
        const minCount = document.getElementById("minCount").value;
        const maxCount = document.getElementById("maxCount").value;

        // Backend'e göndermek için bir JSON nesnesi oluşturuyoruz.
        const data = {
            startDate: startDate,
            endDate: endDate,
            minCount: parseInt(minCount),
            maxCount: parseInt(maxCount)
        };

        // Girilen veriyi backende gönderiyoruz.
        try {
            mongoOutput.textContent = "Yükleniyor..."; // Kullanıcıya bilgi veriyoruz.

            // Backend'e POST isteği atıyoruz.
            const response = await fetch("http://localhost:8080/mongo-records", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            });

            // Dönen cevabı JSON olarak okuyoruz.
            const result = await response.json();

            // Sonucu ekranda okunabilir şekilde gösteriyoruz.
            mongoOutput.textContent = JSON.stringify(result, null, 2);

        } catch (err) {
            // Hata olursa kullanıcıya bildiriyoruz.
            mongoOutput.textContent = "Bir hata oluştu: " + err;
        }
    });

    // In-Memory key-value POST formunu dinliyoruz
    inmemoryPostForm.addEventListener("submit", async function (e) {
        e.preventDefault(); // Yine sayfa yenilenmesin diye engelliyoruz.

        // Kullanıcının girmiş olduğu key ve value değerlerini alıyoruz.
        const key = document.getElementById("inmemory-key").value;
        const value = document.getElementById("inmemory-value").value;

        // Göndereceğimiz veri nesnesini oluşturuyoruz.
        const data = { key: key, value: value };

        try {
            inmemoryPostOutput.textContent = "Ekleniyor..."; // Kullanıcıya durum bildiriyoruz.

            // POST isteği ile backend'e ekleme yapıyoruz.
            const response = await fetch("http://localhost:8080/in-memory", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            });

            // Sonucu bekliyoruz ve JSON olarak alıyoruz.
            const result = await response.json();

            // Sonucu ekranda okunabilir bir şekilde gösteriyoruz.
            inmemoryPostOutput.textContent = JSON.stringify(result, null, 2);

        } catch (err) {
            // Hata oluşursa ekrana yazıyoruz.
            inmemoryPostOutput.textContent = "Bir hata oluştu: " + err;
        }
    });

    // In-Memory key-value GET formunu dinliyoruz.
    inmemoryGetForm.addEventListener("submit", async function (e) {
        e.preventDefault(); // Sayfa yenilenmesini engelliyoruz.

        // Kullanıcının girmiş olduğu anahtarı alıyoruz.
        const key = document.getElementById("inmemory-get-key").value;

        try {
            inmemoryGetOutput.textContent = "Sorgulanıyor..."; // Durum bilgisini gösteriyoruz.

            // GET isteği ile girilen anahtarı backend'e gönderiyoruz.
            // encodeURIComponent ile güvenli bir şekilde URL'ye ekliyoruz.
            const response = await fetch(`http://localhost:8080/in-memory?key=${encodeURIComponent(key)}`);

            // Eğer cevap başarılı değilse, hatayı ekrana basıyoruz.
            if (!response.ok) {
                const text = await response.text();
                inmemoryGetOutput.textContent = "Hata: " + text;
                return;
            }

            // Cevap başarılıysa JSON olarak alıyoruz ve ekrana yazıyoruz.
            const result = await response.json();
            inmemoryGetOutput.textContent = JSON.stringify(result, null, 2);

        } catch (err) {
            // Hata oluşursa kullanıcıya gösteriyoruz.
            inmemoryGetOutput.textContent = "Bir hata oluştu: " + err;
        }
    });

});
