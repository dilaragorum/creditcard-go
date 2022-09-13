# Pay

---

## Amac

Kullanicilar kredi karti bilgilerinin ilk 6 hanesini girdiklerinde kredi kartina ait banka bilgilerini ve bankanin sundugu taksit seceneklerini gorebilmeliler. Bu seceneklere gore kredi karti bilgilerini girdiginde ise sectigi seceneklere gore bankanin sagladigi odeme yonetemi ile odeme yapabilmeli.

## Kapsam

Ornek bir kredi karti numarasi asagidaki gibidir:

[x][xxx xx][xx xxxx xxx][x]

Bu alanlarin temsil ettigi degerler ise asagidaki gibidir:

[x]             : Kart tipi \
[xxx xx]        : Banka kodu \
[xx xxxx xxx]   : Musteri numarasi \
[x]             : Checksum 

Bankalar kart tipine ve hizmet turune gore farkli banka koduna sahip olabilir. Bu banka koduna gore taksit ve vade secenekleri degisiklik gosterebilir.

## Kabul Kriterleri
### Faz 0.5

Bu faz ile birlikte 2 farkli banka desteklenmeli. Bu bankalar ve destekledigi odeme yonemleri ise asagidaki sekilde olmali: 
#### ABank

Kart tipi 4 (VISA) ve 5 (MasterCard) olan kartlari destekler.

3 farkli hizmet sunar. Bunlar ve banka kodlari:\
**[242 10] Free:** Bu hizmet turu musterilere sadece 100TL altindaki harcamalari icin toplam fiyatin uzerine %5 faiz ile 3 taksit sunar.

**[242 20] Gold:** Bu hizmet turu musterilere 100TL altindaki harcamalari icin faizsiz 3 taksit, 100-500TL araligindaki harcamalari icin toplam fiyatin uzerine %5 faiz ile 6 taksit sunar.

**[242 30] Platinum:**  Bu hizmet turu musterilerine 500TL ve altindaki harcamalari icin faizsiz 3 ve 6 taksit, 500TL uzerindeki harcamalari icin toplam fiyatin uzerine %5 faiz ile 9 taksit sunar.


#### BBank

Kart tipi sadece 5 (MasterCard) olan kartlari destekler.

3 farkli hizmet sunar. Bunlar ve banka kodlari:\
**[567 10] Free:** Bu hizmet turu musterilere 1000TL altindaki tum harcamalari icin 3, 6 ve 9 taksit sunar. 3 taksit icin toplam fiyat uzerine %5, 6 taksit icin toplam fiyat uzerine %8, 9 taksit icin toplam fiyat uzerine %11 faiz uygular.

**[567 20] Platinum:** Bu hizmet turu musterilere 1500TL altindaki tum harcamalari icin 3, 6 ve 9 taksit sunar. 3 taksit icin %0, 6 taksit icin toplam fiyat uzerine %5, 9 taksit icin toplam fiyat uzerine %8 faiz uygular.

#### Beklenen Ciktilar

**-> Kart Numarasi: 1242 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 2242 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 3242 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 4242 1000 0000 0000 Fiyat: 100TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 105, Taksit Tutari: 35

**-> Kart Numarasi: 4242 1000 0000 0000 Fiyat: 500TL :**
- Taksit Sayisi: 1, Toplam Geri Odeme: 500, Taksit Tutari: 500 

**-> Kart Numarasi: 4242 2000 0000 0000 Fiyat: 100TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 100, Taksit Tutari: 33.33

**-> Kart Numarasi: 4242 2000 0000 0000 Fiyat: 300TL :**
- Taksit Sayisi: 6, Toplam Geri Odeme: 315, Taksit Tutari: 52.5

**-> Kart Numarasi: 4242 2000 0000 0000 Fiyat: 1000TL :**
- Taksit Sayisi: 1, Toplam Geri Odeme: 1000, Taksit Tutari: 1000

**-> Kart Numarasi: 4242 3000 0000 0000 Fiyat: 100TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 100, Taksit Tutari: 33.33
- Taksit Sayisi: 6, Toplam Geri Odeme: 100, Taksit Tutari: 16.66

**-> Kart Numarasi: 4242 3000 0000 0000 Fiyat: 300TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 300, Taksit Tutari: 100
- Taksit Sayisi: 6, Toplam Geri Odeme: 300, Taksit Tutari: 50

**-> Kart Numarasi: 4242 3000 0000 0000 Fiyat: 600TL :**
- Taksit Sayisi: 9, Toplam Geri Odeme: 630, Taksit Tutari: 70

**-> Kart Numarasi: 1576 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 2576 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 3576 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 4576 1000 0000 0000 Fiyat: 100TL :** desteklenmeyen kart

**-> Kart Numarasi: 4242 1000 0000 0000 Fiyat: 1100TL :**
- Taksit Sayisi: 1, Toplam Geri Odeme: 1100, Taksit Tutari: 1100

**-> Kart Numarasi: 4242 1000 0000 0000 Fiyat: 300TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 315, Taksit Tutari: 105
- Taksit Sayisi: 6, Toplam Geri Odeme: 324, Taksit Tutari: 54
- Taksit Sayisi: 9, Toplam Geri Odeme: 333, Taksit Tutari: 37

**-> Kart Numarasi: 4242 2000 0000 0000 Fiyat: 300TL :**
- Taksit Sayisi: 3, Toplam Geri Odeme: 300, Taksit Tutari: 100
- Taksit Sayisi: 6, Toplam Geri Odeme: 315, Taksit Tutari: 52,5
- Taksit Sayisi: 9, Toplam Geri Odeme: 324, Taksit Tutari: 36

**-> Kart Numarasi: 4242 1000 0000 0000 Fiyat: 2100TL :**
- Taksit Sayisi: 1, Toplam Geri Odeme: 2100, Taksit Tutari: 2100


### Faz 1

Bu faz ile bilrikte bankalar 2 farkli odeme tipi sunabilir. Odeme tipi kart tipine gore degisiklik gosterir. Odeme turleri asagidaki gibidir: \

**Direkt**: Odeme istegi servis saglayici odeme sisteminin sagladigi API'a direkt olarak gonderilir. Buradan donen cevaba gore odeme sonucu listelenir.

```shell
curl --location --request POST 'https://trendyol-case.mocklab.io/v1/pay' \
--header 'Content-Type: application/json' \
--data-raw '{
  "price": 123,
  "cardNumber": "xxxx-xxxx-xxxx-xxxx"
}'
```

**Callback**: Odeme araci bir kurulus tarafindan gerceklestirilir. Araci kurulusun odeme sayfasinin baglantisi kullaniciya donulur.

```http request
https://typay.com/v1/pay/{bankCode}/{uuid}
```

#### ABank

Kart tipi 4 (VISA) olan kartlar icin direkt, 5 (MasterCard) olan kartlar icin callback odeme yontemini kullanir.


#### BBank

Kart tipi sadece 5 (MasterCard) olan kartlari destekler ve callback odeme yontemini kullanir.
