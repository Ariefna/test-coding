#### Soal 3
### Compare Folder

Pada problem kali ini saya memulai dengan membuat klon dari folder structure yang tertera di readme test. Setelah itu saya memulai coding saya dengan membuat fungsi untuk mencari ada folder apa saja di dalam folder source dan target. Setelah ditemukan folder foldernya, kemudian saya masukkan nama folder tersebut ke dalam sebuah array dengan tipe data string, yang mana array ini akan menjadi patokan ada file apa saja di dalam folder yang dideteksi. Setelah semua tersimpan dalam sebuah array, saya mengembalikan array tersebut.

Setelah pembuatan fungsi untuk melacak ada direktori apa saja pada folder source dan target. Saya melanjutkan dengan membuat fungsi untuk mengetahui ada file apa saja dalam direktori tersebut. Kemudian file itu saya masukkan ke dalam sebuah array, yang mana indexnya sesuai dengan tata folder dari fungsi folder yang sudah saya tentukan. Setelah pencarian selesai dilakukan, saya mengembalikan array file yang berupa array 2 dimensi, karena file dalam setiap folder saya masukkan ke dalam array lagi.

Setelah selesai melacak file file dalam folder dan subfolder. Saya melakukan komparasi antar index yang mengandung file, untuk mengetahui apakah file ada di target atau tidak dan sebaliknya. Saat melakukan komparasi, jika file tidak ada pada target atau sebaliknya, maka sistem akan mencetak pemberitahuan ke user.