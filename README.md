# Getir-Case

Bu proje, Getir Araç tarafından hazırlanan teknik değerlendirme case'i doğrultusunda yapılmıştır. Backend, Go ile framework’süz (vanilla net/http), frontend ise vanilla JS/HTML/CSS ile yazıldı. Canlıda test edilebilen, **production-ready**, basit ve şık bir full-stack demo sunar.

---

## 🟣 Canlı Demo

- **Backend (Heroku):**
- [Health Endpoint'i (Canlılık Testi)](https://getir-case-enes-674776c8c0ea.herokuapp.com/health)
  
  <img width="1916" height="1033" alt="healthEndpoint" src="https://github.com/user-attachments/assets/20d6c19b-17d2-42e6-97b4-69692b784f93" />

- [In-Memory POST Endpoint'i](https://getir-case-enes-674776c8c0ea.herokuapp.com/in-memory)
  
  <img width="1918" height="1032" alt="in-memoryPostEndpoint" src="https://github.com/user-attachments/assets/2824656a-7f67-4961-8261-39e3f90bcc8f" />

- [In-Memory GET Endpoint'i](https://getir-case-enes-674776c8c0ea.herokuapp.com/in-memory?key=deneme)
  *(Deneme yerine post ile gönderdiğiniz herhangi bir keyi yazabilirsiniz.)*

  <img width="1918" height="1031" alt="in-memoryGetEndpoint" src="https://github.com/user-attachments/assets/b15ac066-1099-4ad0-9bb2-10d921be0611" />

- [MongoDB Endpoint'i](https://getir-case-enes-674776c8c0ea.herokuapp.com/mongo-records)
  *(Bu endpointi postman, swagger gibi uygulamalar aracılığıyla veya canlı frontend linki üzerinden test edebilirsiniz.)*

  <img width="1917" height="1033" alt="mongoEndpoint" src="https://github.com/user-attachments/assets/02101ad2-9903-40e1-8346-ca7498372a3d" />

- **[Frontend (Vercel)](https://getir-case.vercel.app)**  
  
  <img width="1916" height="1034" alt="frontend" src="https://github.com/user-attachments/assets/4abc31a7-6a75-4961-95b4-710a5f985323" />

---

## 🚀 Nasıl Çalıştırılır?

### 1. Backend (Go API)

**Gereksinimler:**  
- Herhangi bir kod editörü, IDE (VS Code vs.) 
- Go 1.20+  
- MongoDB bağlantı URI’si
- Bilgisayarınızda git kurulu ise aşşağıdaki komutlar ile projeyi indirebilirsiniz. Fakat isterseniz githubdaki dosyaları indirerek çalıştırmakta mümkün.
- .env dosyasına kendi MongoDB URI'nizi eklemeyi unutmayın.
- Dosya adının tam olarak .env olduğundan (ek uzantı veya boşluk olmadığından) emin olun.
- Eğer "go run." yazdığınızda "Error loading .env file" gibi bir hata alırsanız .env dosyasını komutla oluşturmak yerine manuel oluşturmayı deneyebilirsiniz.
```
git clone https://github.com/yukcell19/getir-case.git
cd getir-case
echo "MONGO_URI=buraya kendi MongoDB bağlantınızı yazın" > .env
go run .
```
---

## 🟪 API Endpoint'leri

### MongoDB Kayıt Filtreleme (POST)

- **URL:** `/mongo-records`
- **Method:** `POST`
- **Body:**
    ```json
    {
      "startDate": "2016-01-26",
      "endDate": "2018-02-02",
      "minCount": 2700,
      "maxCount": 3000
    }
    ```
    <img width="405" height="448" alt="mongopost2" src="https://github.com/user-attachments/assets/f4dd1523-a07b-4171-af8c-feecce7ee4b3" />
    
- **Response:**
    ```json
      {
        "code": 0,
        "msg": "Success",
        "records": [
          {
            "createdAt": "2017-01-28T01:22:14.398Z",
            "key": "TAKwGc6Jr4i8Z487",
            "totalCount": 2800
          },
          {
            "createdAt": "2017-01-27T08:19:14.135Z",
            "key": "NAeQ8eX7e5TEg7oH",
            "totalCount": 2900
          }
        ]
      }
    ```
    <img width="366" height="383" alt="mongopost3" src="https://github.com/user-attachments/assets/ab5c9b14-c311-48a7-b141-cdfa32c850a5" />
---

### In-Memory Key-Value Servisi

- **URL:** `/in-memory`

- **POST**
    - **Body:**  
      ```json
      { "key": "active-tabs", "value": "getir" }
      ```
    - **Response:**  
      ```json
      { "key": "active-tabs", "value": "getir" }
      ```
  <img width="371" height="444" alt="inmemorypost2" src="https://github.com/user-attachments/assets/355a9f62-64ce-4732-93e1-215d209870e8" />
  
- **GET**
    - **Param:**  
      `?key=active-tabs`
    - **Response:**  
      ```json
      { "key": "active-tabs", "value": "getir" }
      ```
  <img width="356" height="332" alt="inmemoryget2" src="https://github.com/user-attachments/assets/fa26d3e4-a3f2-4af3-8061-0663560950ed" />
---

## 🟪 Frontend Hakkında

- Frontend, vanilla JS ve HTML/CSS ile yazıldı.
- [Canlı frontend linki](https://getir-case.vercel.app) üzerinden tüm endpointleri rahatça test edebilirsiniz.
- Frontend, doğrudan canlı backend’e entegre edilmiştir.

---

## 🟪 Diğer Detaylar

- **CORS Desteği:**  
  - Sunucu tüm origin’lerden gelen isteklere açıktır (CORS ayarları backend’de aktiftir).  
  - Frontend ile herhangi bir tarayıcıdan test yapılabilir.

- **Deploy:**  
  - Backend Heroku üzerinde, frontend ise Vercel’de yayınlanmıştır.  
  - API uç noktaları ve frontend adresi tamamen herkese açıktır.

- **.env ve Güvenlik:**  
  - MongoDB bağlantı URI’sı `.env` dosyası ile yönetilir.  
  - Bu dosya `.gitignore` ile gizlenmiştir.
  - Kendi localinizde çalıştırmak için MongoDB urinizi .env dosyanıza ekleyin. 

- **Kod Kalitesi:**  
  - Tüm backend ve frontend kodlarında açıklamalar (yorum satırları) mevcuttur.
  - Kodlar mümkün olduğunca sade, okunabilir ve production-ready standartlarda yazılmıştır.
  - Hata yönetimi ve validasyonlar detaylı şekilde yapılmaya çalışılmıştır.

- **Test:**  
  - In-memory endpoint'i için basit bir unit test eklenmiştir.
  -  ```bash
      go test # Bu komut sayesinde test işlemini gerçekleştirebilirsiniz.
  - MongoDB endpoint’i, frontend arayüzü üzerinden canlı test edilebilir.
    
---

## 🟪 Hakkımda
- **İsim:** Enes
- **Mail:** yukselenesemre@gmail.com
- **Not:** Bu case'i sıfırdan Golang ve MongoDB öğrenerek tamamladım. Kodda yalınlık, açıklamalar ve hata yönetimine önem verdim.
- Özellikle RWmutex gibi thread safe ve günümüzde yaygın olan yöntemleri, teknolojileri kullanmaya çalıştım. 
  

