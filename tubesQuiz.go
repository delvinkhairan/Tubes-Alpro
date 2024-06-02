package main

import (
  "fmt"
  "math/rand"
  "time"
)

type data struct {
  nama, pass string
  score      int
}

type dataSoal struct {
  soal    string
  jawaban string
  pilihan [4]string
}

const NMAX int = 100

type arrayPeserta [NMAX]data
type arrayAdmin [NMAX]data
type questionsArr [NMAX]dataSoal

var questions = questionsArr{
  {"Siapakah penemu bola lampu?", "Thomas Edison", [4]string{"Nikola Tesla", "Thomas Edison", "Alexander Graham Bell", "James Watt"}},
  {"Apa ibu kota dari Perancis?", "Paris", [4]string{"Berlin", "Madrid", "Roma", "Paris"}},
  {"Planet mana yang dikenal sebagai 'Planet Merah'?", "Mars", [4]string{"Venus", "Mars", "Jupiter", "Saturnus"}},
  {"Apa nama ilmiah dari air?", "H2O", [4]string{"CO2", "H2O", "O2", "NaCl"}},
  {"Siapakah penulis buku 'Harry Potter'?", "J.K. Rowling", [4]string{"J.R.R. Tolkien", "George R.R. Martin", "J.K. Rowling", "Suzanne Collins"}},
  {"Gunung apa yang tertinggi di dunia?", "Gunung Everest", [4]string{"Gunung Kilimanjaro", "Gunung K2", "Gunung Everest", "Gunung Fuji"}},
  {"Apa nama benua terkecil di dunia?", "Australia", [4]string{"Afrika", "Antartika", "Australia", "Amerika Selatan"}},
  {"Siapa presiden pertama Amerika Serikat?", "George Washington", [4]string{"Thomas Jefferson", "Abraham Lincoln", "George Washington", "John Adams"}},
  {"Berapakah jumlah warna dalam pelangi?", "7", [4]string{"5", "6", "7", "8"}},
  {"Apa nama kimia dari garam dapur?", "NaCl", [4]string{"H2SO4", "NaCl", "KCl", "H2O"}},
  {"Apa nama mata uang Jepang?", "Yen", [4]string{"Yen", "Won", "Baht", "Peso"}},
  {"Siapa penulis novel 'To Kill a Mockingbird'?", "Harper Lee", [4]string{"Harper Lee", "Mark Twain", "J.D. Salinger", "F. Scott Fitzgerald"}},
  {"Berapa banyak planet dalam tata surya kita?", "8", [4]string{"7", "8", "9", "10"}},
  {"Siapa yang menciptakan karakter Sherlock Holmes?", "Arthur Conan Doyle", [4]string{"Agatha Christie", "Arthur Conan Doyle", "Edgar Allan Poe", "Ian Fleming"}},
  {"Apa nama zat yang memberikan warna pada daun hijau?", "Chlorophyll", [4]string{"Hemoglobin", "Chlorophyll", "Carotene", "Anthocyanin"}},
  {"Di negara manakah Menara Eiffel berada?", "Perancis", [4]string{"Spanyol", "Italia", "Inggris", "Perancis"}},
  {"Siapa ilmuwan yang menemukan teori relativitas?", "Albert Einstein", [4]string{"Isaac Newton", "Galileo Galilei", "Albert Einstein", "Nikola Tesla"}},
  {"Apa nama ibu kota Australia?", "Canberra", [4]string{"Sydney", "Melbourne", "Canberra", "Brisbane"}},
  {"Apa simbol kimia dari emas?", "Au", [4]string{"Ag", "Au", "Pb", "Fe"}},
  {"Siapa yang melukis Mona Lisa?", "Leonardo da Vinci", [4]string{"Vincent van Gogh", "Michelangelo", "Leonardo da Vinci", "Pablo Picasso"}},
  {"Siapakah yang menemukan benua Amerika?", "Christopher Columbus", [4]string{"Marco Polo", "Christopher Columbus", "William Shakespeare", "Leonardo da Vinci"}},
  {"Apa nama planet terbesar di tata surya kita?", "Jupiter", [4]string{"Saturnus", "Jupiter", "Neptunus", "Uranus"}},
  {"Siapa yang menemukan telepon?", "Alexander Graham Bell", [4]string{"Thomas Edison", "Alexander Graham Bell", "Nikola Tesla", "Guglielmo Marconi"}},
  {"Apa bahasa resmi di Brasil?", "Portugis", [4]string{"Spanyol", "Portugis", "Prancis", "Inggris"}},
  {"Siapa yang menulis 'Hamlet' dan 'Romeo and Juliet'?", "William Shakespeare", [4]string{"Christopher Marlowe", "William Shakespeare", "Geoffrey Chaucer", "John Milton"}},
  {"Apa nama piramida terkenal di Mesir?", "Piramida Giza", [4]string{"Piramida Cheops", "Piramida Khafre", "Piramida Menkaure", "Piramida Giza"}},
  {"Apa nama sungai terpanjang di dunia?", "Sungai Nil", [4]string{"Sungai Amazon", "Sungai Yangtze", "Sungai Nil", "Sungai Mississippi"}},
  {"Apa nama ilmuwan yang menemukan hukum gravitasi?", "Isaac Newton", [4]string{"Albert Einstein", "Isaac Newton", "Galileo Galilei", "Nikola Tesla"}},
  {"Siapa yang menulis 'The Odyssey'?", "Homer", [4]string{"Virgil", "Homer", "Sophocles", "Euripides"}},
  {"Apa nama ibu kota Kanada?", "Ottawa", [4]string{"Toronto", "Vancouver", "Ottawa", "Montreal"}},
  {"Siapa yang menciptakan karakter James Bond?", "Ian Fleming", [4]string{"John Le Carré", "Ian Fleming", "Agatha Christie", "Arthur Conan Doyle"}},
  {"Apa nama gas yang paling melimpah di atmosfer bumi?", "Nitrogen", [4]string{"Oksigen", "Karbon dioksida", "Nitrogen", "Hidrogen"}},
  {"Di negara manakah terdapat Tembok Besar?", "Tiongkok", [4]string{"Jepang", "India", "Tiongkok", "Korea Selatan"}},
  {"Lontong dapat dengan mudah ditemukan di negara?", "Semua jawaban benar", [4]string{"Malaysia", "Indonesia", "Singapura", "Semua jawaban benar"}},
  {"Apa warna dari teh bunga telang?", "Biru", [4]string{"Merah", "Coklat", "Biru", "Kuning"}},
  {"Apa warna dari Nasi Kerabu?", "Biru", [4]string{"Putih", "Kuning", "Biru", "Coklat"}},
  {"Warna mata manusia yang tidak mungkin adalah?", "Hitam", [4]string{"Hitam", "Ungu", "Hijau", "Pink"}},
  {"Letusan vulkanik terbesar dalam catatan sejarah adalah?", "Tambora", [4]string{"Krakatoa", "Vesuvius", "Tambora", "St. Helens"}},
  {"Apakah bulan lebih besar dari planet lain?", "ya", [4]string{"Tidak", "Ya", "Semua jawaban benar", "Rahasia"}},
  {"Berapa jumlah bulan yang dimiliki oleh bumi?", "1", [4]string{"0", "1", "2", "3"}},
  {"Apa yang lebih berat, satu kilogram kapas atau satu kilogram besi?", "Sama berat", [4]string{"Kapas", "Besi", "Sama berat", "Tidak bisa dibandingkan"}},
  {"Manakah yang lebih dekat ke bumi, bulan atau matahari?", "Bulan", [4]string{"Bulan", "Matahari", "Sama saja", "Tergantung waktu"}},
  {"Di negara mana sebagian besar hutan Amazon berada?", "Brasil", [4]string{"Kolombia", "Venezuela", "Brasil", "Peru"}},
  {"Apa warna buah jeruk sebelum matang?", "Hijau", [4]string{"Oranye", "Kuning", "Hijau", "Merah"}},
  {"Siapa yang menulis 'Pride and Prejudice'?", "Jane Austen", [4]string{"Charlotte Bronte", "Emily Bronte", "Jane Austen", "Mary Shelley"}},
  {"Manakah yang lebih besar, atom atau molekul?", "Molekul", [4]string{"Atom", "Molekul", "Sama besar", "Tidak bisa dibandingkan"}},
  {"Berapa jumlah tulang dalam tubuh manusia dewasa?", "206", [4]string{"Tidak tahu", "206", "207", "208"}},
  {"Apa yang terbentuk lebih dulu, telur atau ayam?", "Telur", [4]string{"Telur", "Ayam", "Tidak bisa dijawab", "Keduanya bersamaan"}},
  {"Mana yang lebih dingin, Kutub Utara atau Kutub Selatan?", "Kutub Selatan", [4]string{"Kutub Utara", "Kutub Selatan", "Sama saja", "Bergantung musim"}},
  {"Hewan apa yang dikenal bisa berubah warna untuk beradaptasi dengan lingkungannya?", "Bunglon", [4]string{"Bunglon", "Manusia", "Kupu-kupu", "Ikan badut"}},
}

func menu(Ap *arrayPeserta, Ad *arrayAdmin, npeserta, nadmin *int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (npeserta), dan jumlah admin (nadmin)
     F.S. Menampilkan menu utama untuk login atau daftar, kemudian memanggil fungsi login atau daftar sesuai pilihan pengguna
  */
  var nmenu int
  fmt.Println("1. LOGIN")
  fmt.Println("2. DAFTAR")
  fmt.Scan(&nmenu)
  if nmenu == 1 {
    caripemain(Ap, Ad, npeserta, nadmin)
  } else if nmenu == 2 {
    daftar(Ap, Ad, npeserta, nadmin)
  }
}

func caripemain(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), dan jumlah admin (na)
     F.S. Menampilkan menu login untuk peserta atau admin, dan memanggil fungsi login yang sesuai
  */
  var nmenu int
  fmt.Println("1. LOGIN SEBAGAI PESERTA")
  fmt.Println("2. LOGIN SEBAGAI ADMIN")
  fmt.Scan(&nmenu)
  if nmenu == 1 {
    caripeserta(Ap, np, na, Ad)
  } else if nmenu == 2 {
    cariadmin(Ad, na, np, Ap)
  }
}

func caripeserta(A *arrayPeserta, np, na *int, Ad *arrayAdmin) {
  /*
     I.S. Terdefinisi array peserta (A), jumlah peserta (np), jumlah admin (na), dan array admin (Ad)
     F.S. Melakukan proses login untuk peserta dan memanggil fungsi homePeserta jika login berhasil, atau menampilkan pesan gagal login
  */
  var nama, pass string
  var passfound = -1
  fmt.Println("MASUKKAN NAMA ANDA")
  fmt.Scan(&nama)
  fmt.Println("MASUKKAN PASSWORD ANDA")
  fmt.Scan(&pass)
  for i := 0; i < *np && passfound == -1; i++ {
    if A[i].nama == nama && A[i].pass == pass {
      passfound = i
      fmt.Println("LOGIN BERHASIL")
      fmt.Println("Selamat datang peserta", A[i].nama, "!")
      homePeserta(A, Ad, np, na, i)
    }
  }
  if passfound == -1 {
    fmt.Println("LOGIN GAGAL")
    menu(A, Ad, np, na)
  }
}

func cariadmin(A *arrayAdmin, na, np *int, Ap *arrayPeserta) {
  /*
     I.S. Terdefinisi array admin (A), jumlah admin (na), jumlah peserta (np), dan array peserta (Ap)
     F.S. Melakukan proses login untuk admin dan memanggil fungsi homeAdmin jika login berhasil, atau menampilkan pesan gagal login
  */
  var nama, pass string
  var passfound = -1
  fmt.Println("MASUKKAN NAMA ANDA")
  fmt.Scan(&nama)
  fmt.Println("MASUKKAN PASSWORD ANDA")
  fmt.Scan(&pass)
  for i := 0; i < *na && passfound == -1; i++ {
    if A[i].nama == nama && A[i].pass == pass {
      passfound = i
      fmt.Println("LOGIN BERHASIL")
      fmt.Println("Selamat datang admin", A[i].nama, "!")
      homeAdmin(Ap, A, np, na, i)
    }
  }
  if passfound == -1 {
    fmt.Println("LOGIN GAGAL")
    menu(Ap, A, np, na)
  }
}

func daftar(Ap *arrayPeserta, Ad *arrayAdmin, npeserta, nadmin *int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (npeserta), dan jumlah admin (nadmin)
     F.S. Menampilkan menu pendaftaran untuk peserta atau admin, dan memanggil fungsi pendaftaran yang sesuai
  */
  var nmenu int
  fmt.Println("1. DAFTAR SEBAGAI PESERTA")
  fmt.Println("2. DAFTAR SEBAGAI ADMIN")
  fmt.Scan(&nmenu)
  if nmenu == 1 {
    daftarpeserta(Ap, npeserta, Ad, nadmin)
  } else if nmenu == 2 {
    daftaradmin(Ad, nadmin, Ap, npeserta)
  }
}

func daftarpeserta(A *arrayPeserta, np *int, Ad *arrayAdmin, na *int) {
  /*
     I.S. Terdefinisi array peserta (A), jumlah peserta (np), array admin (Ad), dan jumlah admin (na)
     F.S. Mendaftarkan peserta baru dengan memasukkan nama dan password, kemudian kembali ke menu utama
  */
  var nama, pass string
  fmt.Println("MASUKKAN NAMA BARU")
  fmt.Scan(&nama)
  fmt.Println("MASUKKAN PASSWORD BARU")
  fmt.Scan(&pass)
  A[*np].nama = nama
  A[*np].pass = pass
  (*np)++
  fmt.Println("PENDAFTARAN BERHASIL")
  menu(A, Ad, np, na)
}

func daftaradmin(A *arrayAdmin, na *int, Ap *arrayPeserta, np *int) {
  /*
     I.S. Terdefinisi array admin (A), jumlah admin (na), array peserta (Ap), dan jumlah peserta (np)
     F.S. Mendaftarkan admin baru dengan memasukkan nama dan password, kemudian kembali ke menu utama
  */
  var nama, pass string
  fmt.Println("MASUKKAN NAMA BARU")
  fmt.Scan(&nama)
  fmt.Println("MASUKKAN PASSWORD BARU")
  fmt.Scan(&pass)
  A[*na].nama = nama
  A[*na].pass = pass
  (*na)++
  fmt.Println("PENDAFTARAN BERHASIL")
  menu(Ap, A, np, na)
}

func homePeserta(Ap *arrayPeserta, Ad *arrayAdmin, np *int, na *int, index int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index peserta yang sedang login
     F.S. Menampilkan menu utama peserta dan memanggil fungsi yang sesuai dengan pilihan peserta
  */
  var n int
  fmt.Println("▕▔▔╲╱▋▔▋▔▋╲╱▔▔▏")
  fmt.Println("▕┈▔╲▍┈▋┈▍┈▋╱▔┈▏")
  fmt.Println("▔╲╱┳▅╮┊┊┊╭▅┳╲╱")
  fmt.Println("┈▕▋╰━┫┊┊┊┣━╯▋▏")
  fmt.Println("▋┈╲▍╱┈▂▂▂┈╲▍╱")
  fmt.Println("┈▊┈╲▏┈╲▂╱┈▕╱")
  fmt.Println("▋┈┈▋╲▂╱▔╲▂╱")
  fmt.Println("1. START")
  fmt.Println("2. LEADERBOARD")
  fmt.Println("3. LOGOUT")
  fmt.Scan(&n)
  if n == 1 {
    start(Ap, Ad, np, na, index)
  } else if n == 2 {
    leaderboard(Ap, Ad, *np, *na, index)
  } else if n == 3 {
    menu(Ap, Ad, np, na)
  }
}

func homeAdmin(Ap *arrayPeserta, Ad *arrayAdmin, np *int, na *int, index int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index admin yang sedang login
     F.S. Menampilkan menu utama admin dan memanggil fungsi yang sesuai dengan pilihan admin
  */
  var n int
  fmt.Println("1. TAMPILKAN SOAL")
  fmt.Println("2. UBAH SOAL")
  fmt.Println("3. TAMBAH SOAL")
  fmt.Println("4. HAPUS SOAL")
  fmt.Println("5. LOGOUT")
  fmt.Scan(&n)
  if n == 1 {
    tampilkanPertanyaan(Ap, Ad, np, na, index)
  } else if n == 2 {
    ubahsoal(Ap, Ad, np, na, index)
  } else if n == 3 {
    tambahsoal(Ap, Ad, np, na, index)
  } else if n == 4 {
    hapussoal(Ap, Ad, np, na, index)
  } else if n == 5 {
    menu(Ap, Ad, np, na)
  }
}

func tampilkanPertanyaan(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int, index int) {
  /*
       I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index admin yang sedang login
       F.S. Menampilkan semua soal yang ada, lengkap dengan jawaban dan pilihan jawaban, kemudian menunggu input untuk kembali ke menu admin
    */
  var questionCount, n, nsoal int
  nsoal = 0
  for i := 0; questions[i].soal != ""; i++ {
      questionCount++
  }
  for i := 0; i <= questionCount; i++ {
    if questions[i].soal != "" {
    fmt.Printf("Soal %d: %s\n", nsoal+1, questions[i].soal)
    fmt.Printf("Jawaban: %s\n", questions[i].jawaban)
    fmt.Printf("Pilihan:\n")
    for j := 0; j < 4; j++ {
      fmt.Printf("%d. %s\n", j+1, questions[i].pilihan[j])
    }
    fmt.Println()
    nsoal++
    }
  }
  fmt.Println("1. BACK")
  fmt.Scan(&n)
  if n == 1 {
    homeAdmin(Ap, Ad, np, na, index)
  }
}

func ubahsoal(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int, index int) {
  /*
       I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index admin yang sedang login
       F.S. Mengubah soal pada nomor yang dipilih oleh admin dengan data baru yang diinputkan, kemudian kembali ke menu admin
    */
  var idx int
  var soal, jawaban, pilihan string
  fmt.Println("Masukkan nomor soal yang ingin diubah:")
  fmt.Scan(&idx)
  if questions[idx-1].soal == "" {
    fmt.Println("Nomor soal tidak valid.")
    homeAdmin(Ap, Ad, np, na, index)
    return
  }
  idx--
  fmt.Println("Note: Ganti spasi dengan underscore (_)")
  fmt.Println("Masukkan soal baru:")
  
  fmt.Scan(&soal)
  questions[idx].soal = soal

  fmt.Println("Masukkan jawaban benar:")
  
  fmt.Scan(&jawaban)
  questions[idx].jawaban = jawaban

  fmt.Println("Masukkan 4 pilihan jawaban:")
  for i := 0; i < 4; i++ {
    fmt.Printf("Pilihan %d: ", i+1)
    
    fmt.Scanln(&pilihan)
    questions[idx].pilihan[i] = pilihan
  }
  fmt.Println("Soal berhasil diubah.")
  homeAdmin(Ap, Ad, np, na, index)
}

func tambahsoal(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int, index int) {
  /*
       I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index admin yang sedang login
       F.S. Menambahkan soal baru yang diinputkan oleh admin ke dalam daftar soal, kemudian kembali ke menu admin
    */
  var soal, jawaban, pilihan string
  var questionCount int
  for i := 0; questions[i].soal != ""; i++ {
      questionCount++
  }
  for i := 0; i <= questionCount; i++ {
    if questions[i].soal == "" {
      fmt.Println("Note: Ganti spasi dengan underscore (_)")
      fmt.Println("Masukkan soal baru:")
      fmt.Scan(&soal)
      questions[i].soal = soal

      fmt.Println("Masukkan jawaban benar:")
      
      fmt.Scan(&jawaban)
      questions[i].jawaban = jawaban

      fmt.Println("Masukkan 4 pilihan jawaban:")
      for j := 0; j < 4; j++ {
        fmt.Printf("Pilihan %d: ", j+1)
        
        fmt.Scan(&pilihan)
        questions[i].pilihan[j] = pilihan
      }
      fmt.Println("Soal berhasil ditambahkan.")
      homeAdmin(Ap, Ad, np, na, index)
      return
    }
  }
  homeAdmin(Ap, Ad, np, na, index)
}

func hapussoal(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int, index int) {
  /*
       I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index admin yang sedang login
       F.S. Menghapus soal pada nomor yang dipilih oleh admin, kemudian menggeser soal-soal berikutnya ke atas, lalu kembali ke menu admin
    */
  var idx int
  fmt.Println("Masukkan nomor soal yang ingin dihapus:")
  fmt.Scan(&idx)
  if questions[idx-1].soal == "" {
    fmt.Println("Nomor soal tidak valid.")
    homeAdmin(Ap, Ad, np, na, index)
    return
  }
  for questions[idx-1].soal != "" {
    questions[idx-1] = questions[idx]
    idx++
    }
  fmt.Println("Soal berhasil dihapus.")
  homeAdmin(Ap, Ad, np, na, index)
}

func start(Ap *arrayPeserta, Ad *arrayAdmin, np, na *int, index int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index peserta yang sedang login
     F.S. Memulai kuis untuk peserta yang sedang login dan menghitung skornya berdasarkan jawaban yang benar
  */
  var prize, questionCount int
  var selesai bool
  rand.Seed(time.Now().UnixNano())  
  for i := 0; questions[i].soal != ""; i++ {
      questionCount++
  }
  for !selesai && prize < 10000000 {
    randomIdx := rand.Intn(questionCount)
    question := questions[randomIdx]
    fmt.Println("𓃚 ===== Rp. ", prize+1000000, " ===== 𓃚")
    fmt.Println("Question:", question.soal)

    choices := question.pilihan
    rand.Shuffle(len(choices), func(i, j int) {
      choices[i], choices[j] = choices[j], choices[i]
    })

    for i := 0; i < 4; i++ {
      fmt.Printf("%d. %s\n", i+1, choices[i])
    }

    var answer int
    fmt.Print("Pilih jawaban (1-4): ")
    fmt.Scan(&answer)

    if answer >= 1 && answer <= 4 {
      if choices[answer-1] == question.jawaban {
        fmt.Println("Jawaban benar!")
        prize += 1000000
      } else {
        fmt.Println("██╗░░██╗")
        fmt.Println("╚██╗██╔╝")
        fmt.Println("░╚███╔╝░")
        fmt.Println("░██╔██╗░")
        fmt.Println("██╔╝╚██╗")
        fmt.Println("╚═╝░░╚═╝")
        fmt.Println("Jawaban salah!")
        selesai = true
        if prize < 5000000 {
          prize = 0
        } else if prize >= 5000000 && prize < 10000000 {
          prize = 5000000
        } else if prize == 10000000 {
          prize = 10000000
        }
        fmt.Println("ANDA MEMBAWA PULANG HADIAH UANG TUNAI SEBESAR: Rp. ", prize)
        Ap[index].score += prize
        leaderboard(Ap, Ad, *np, *na, index)
      }
    } else {
      fmt.Println("Pilihan tidak valid.")
    }
  }
  if prize == 10000000 {
    fmt.Println("CONGRATULATION ! ! !")
    fmt.Println("▕▔╲┈╱▔╲┈┈╱╲╱▔▏┈")
    fmt.Println("▕▏┈▏╱▉╲┈┈╱▉╲▕▏┈")
    fmt.Println("┈╲▃▏▔▔▔╲▂▂▂▕╱┈┈")
    fmt.Println("┈┈┈▏┊┊┳┊╲▂╱┳▏┈┈")
    fmt.Println("┈┈▕╲▂┊╰━━┻━╱┈┈┈")
    fmt.Println("┈┈╱┈┈▔▔╲▂▂╱╲┈┈┈")
    fmt.Println("ANDA MEMBAWA PULANG HADIAH UANG TUNAI SEBESAR: Rp. ", prize)
    Ap[index].score += prize
    leaderboard(Ap, Ad, *np, *na, index)
  }
}

func leaderboard(Ap *arrayPeserta, Ad *arrayAdmin, np int, na int, index int) {
  /*
     I.S. Terdefinisi array peserta (Ap), array admin (Ad), jumlah peserta (np), jumlah admin (na), dan index peserta yang sedang login
     F.S. Menampilkan hasil kuis (skor) dari semua peserta
  */
  var n int
  if np == 0 {
    fmt.Println("Belum ada peserta yang terdaftar.")
    return
  }

  sortedParticipants := make([]data, np)
  copy(sortedParticipants, Ap[:np])

  for i := 1; i < np; i++ {
    idx := i - 1
    for j := i ; j < np; j++ {
      if sortedParticipants[idx].score > sortedParticipants[j].score {
        idx = j
      }
    }
    if idx != i {
      sortedParticipants[i], sortedParticipants[idx] = sortedParticipants[idx], sortedParticipants[i]
    }
  }

  fmt.Println("===== LEADERBOARD =====")
  fmt.Println("Rank | Name              | Score")
  fmt.Println("------------------------------")
  for i := 0; i < np; i++ {
    fmt.Printf("%-4d | %-17s | %d\n", i+1, sortedParticipants[i].nama, sortedParticipants[i].score)
  }
  fmt.Println("------------------------------")
  fmt.Println("1. BACK")
  fmt.Scan(&n)
  if n == 1 {
    homePeserta(Ap, Ad, &np, &na, index)
  }
}


func main() {
  var npeserta, nadmin int
  var Ap arrayPeserta
  var Ad arrayAdmin
  menu(&Ap, &Ad, &npeserta, &nadmin)
}
