# Go Mini Projeler

Bu repo, Go dilini öğrenmek amacıyla hazırlanmış 5 farklı mini projeyi içermektedir. Her proje, farklı Go konseptlerini pratik ederek öğrenmeyi amaçlar. Aşağıda her projenin kısa açıklamasını ve kapsadığı konuları bulabilirsiniz.

---

## Mini Proje 1: Kişisel Görev Takip Konsol Uygulaması

### Konular
- Variables, Data Types  
- Conditionals, Loops  
- Arrays/Slices, Maps  
- Functions, Structs  

### Açıklama
Konsol tabanlı bir görev takip sistemi:

- Kullanıcıdan görev ismi, açıklama ve durum (`planned`, `ongoing`, `done`) bilgileri alınır.  
- Görevler slice içinde saklanır.  
- Görev ekleme, listeleme ve tamamlama işlemleri yapılır.  
- Duruma göre filtreleme özelliği içerir.

---

## Mini Proje 2: Basit Banka Hesap Simülatörü

### Konular
- Structs, Methods  
- Pointers  
- Error Handling  

### Açıklama
Banka hesabı işlemlerini yöneten basit bir simülatör:

- `Account` yapısı: isim ve bakiye içerir.  
- Para yatırma, para çekme ve bakiye görüntüleme methodları tanımlanmıştır.  
- Pointer kullanımıyla bakiye güncellenir.  
- Negatif bakiye işlemlerine karşı hata kontrolü yapılır.

---

## Mini Proje 3: JSON Tabanlı Kullanıcı Kaydı

### Konular
- File Operations  
- JSON İşleme (marshal/unmarshal)  
- Structs  
- Error Handling  

### Açıklama
JSON dosyası üzerinde kullanıcı kayıt işlemi:

- Kullanıcıdan isim ve yaş gibi bilgiler alınır.  
- Veriler JSON formatında dosyaya kaydedilir.  
- Dosyadan veriler okunarak kullanıcı listesi gösterilir.

---

## Mini Proje 4: Web Üzerinden Basit Not Ekleme Servisi

### Konular
- HTTP Sunucuları  
- REST API Geliştirme  
- JSON İşleme  
- Structs, Functions  
- Error Handling  

### Açıklama
Basit bir not servisi sağlayan REST API:

- `/addNote` ve `/getNotes` endpoint'leri mevcuttur.  
- Gelen JSON istekleri parse edilir, yanıtlar da JSON olarak döner.  
- Notlar slice içinde geçici olarak saklanır.

---

## Mini Proje 5: Görev Dağıtım Sistemi (Go Routines + Channels)

### Konular
- Go Routines  
- Channels  
- Structs, Maps 

### Açıklama
Çoklu çalışanlar arasında görev dağıtımı:

- Görevler çalışanlara sırayla atanır.  
- Her görev ayrı bir goroutine içinde çalışır.  
- Görev tamamlandığında `channel` ile sonuç bildirilir.  
- Hangi görev, hangi çalışan tarafından tamamlandığı ekrana yazılır.

---

## Nasıl Çalıştırılır?

1. Bu repoyu klonlayın:
   ```bash
   git clone https://github.com/melihkl/go-mini-projects.git
   cd go-mini-projeler

2. Her proje ayrı klasördedir. İlgili klasöre girin ve çalıştırın:
    ```bash
   cd mini-proje-1
   go run main.go

