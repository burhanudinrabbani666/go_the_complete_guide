# maps vs structs

Di Go (Golang), **struct** digunakan untuk menyimpan data yang bentuknya sudah tetap dan field-nya sudah jelas sejak awal. Misalnya data User yang selalu memiliki Name, Age, dan Email. Karena strukturnya pasti, struct lebih aman, lebih cepat, lebih hemat memori, dan cocok untuk merepresentasikan object seperti user, produk, transaksi, atau response API. Selain itu, struct juga bisa memiliki method sehingga sering dipakai dalam pengembangan aplikasi yang lebih besar dan terstruktur.

Sedangkan **map** digunakan untuk menyimpan data dalam bentuk pasangan key-value yang lebih fleksibel dan dinamis. Key pada map bisa bertambah atau berubah sesuai kebutuhan, misalnya untuk menyimpan nilai pelajaran, konfigurasi, cache, atau hasil parsing JSON yang tidak pasti strukturnya. map sangat cocok untuk kebutuhan lookup cepat atau data yang tidak memiliki field tetap, tetapi kurang type-safe dibanding struct. Singkatnya, struct dipakai untuk data yang jelas bentuknya, sedangkan map untuk data yang lebih bebas dan fleksibel.

Next: [Using make function](./13-using-make-function.md)