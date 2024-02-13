# MyGram - Aplikasi Media Sosial untuk Berbagi Foto dan Komentar

MyGram adalah sebuah aplikasi media sosial yang memungkinkan pengguna untuk menyimpan foto serta membuat komentar untuk foto orang lain. Aplikasi ini dilengkapi dengan proses CRUD (Create, Read, Update, Delete) yang memungkinkan pengguna untuk berinteraksi dengan data secara penuh. Dalam pengembangan aplikasi ini, digunakan framework Gin Gonic dan ORM Gorm untuk mempermudah proses pengembangan.

## Daftar Table
Dalam proyek MyGram, terdapat empat table utama yang diperlukan. Berikut merupakan detailnya:

#### 1. User
```
- Field-field:
  - Email: Validasi format email yang valid, unique index, tidak boleh kosong.
  - Username: Unique index, tidak boleh kosong.
  - Password: Minimal panjang 6 karakter, tidak boleh kosong.
  - Age: Minimal nilai di atas 8, tidak boleh kosong.
```

#### 2. Photo
```
= Field-field:
  - Title: Tidak boleh kosong.
  - Photo_url: Tidak boleh kosong.
```
