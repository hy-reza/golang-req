# Go Water and Wind Status

Aplikasi ini menggunakan bahasa pemrograman Go untuk mengirim data water dan wind secara acak dalam format JSON ke server melalui POST request. Data ini juga akan mencetak status water dan wind berdasarkan nilai yang dihasilkan.

## Struktur Data

Data yang dikirim dalam bentuk JSON memiliki dua atribut:

- `water` (dalam satuan meter) - Nilai acak antara 1-100.
- `wind` (dalam satuan meter per detik) - Nilai acak antara 1-100.

## Status Water dan Wind

Status water dan wind ditentukan berdasarkan nilai-nilai yang dihasilkan:

- Water:
  - Aman: nilai ≤ 5
  - Siaga: 6 ≤ nilai ≤ 8
  - Bahaya: nilai > 8
- Wind:
  - Aman: nilai ≤ 6
  - Siaga: 7 ≤ nilai ≤ 15
  - Bahaya: nilai > 15

