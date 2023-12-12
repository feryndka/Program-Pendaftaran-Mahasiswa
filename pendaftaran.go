package main

import "fmt"

const NMAX int = 10000

type mahasiswa struct {
	nama     string
	jurusan  string
	nilaiTes int
	status   string
	npm 	 int
}

type TabMahasiswa [NMAX]mahasiswa

func header() {
	/* I.S. -
	   F.S. menampilkan header dari aplikasi */
	fmt.Println("============ Selamat Datang ==============")
	fmt.Println("     Aplikasi Pendaftaran Mahasiswa 	   ")
	fmt.Println("         Algoritma Pemrograman   	       ")
	fmt.Println("------------------------------------------")
}

func menu(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa dan banyak data pada n
	   F.S. melakukan pemilihan role yaitu admin atau calon mahasiswa */
	var pilihan int

	for pilihan != 3 {
		header()
		fmt.Println("Pengguna:")
		fmt.Println("1. Admin")
		fmt.Println("2. Calon Mahasiswa")
		fmt.Println("3. Exit")
		fmt.Println("------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			// panggil procedure menu admin
			menu_admin(M, n)
		case 2:
			// panggil procedure menu mahasiswa
			menu_mahasiswa(M, n)
		case 3:
			fmt.Println("Terima kasih :)")
		}
	}
}

func menu_admin(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa dan banyak data pada n
	   F.S. menampilkan menu admin dan beberapa menu yang dapat dipilih */
	var pilihan int
	var nama, jurusan string
	var nilaiTes, npm int
	var keputusan string

	for pilihan != 5 {
		fmt.Println("==========================================")
		fmt.Println("               MENU ADMIN                 ")
		fmt.Println("==========================================")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Tampilkan Data Mahasiswa")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Exit")
		fmt.Println("-----------------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// menambahkan data mahasiswa
			fmt.Print("Nama Mahasiswa: "); fmt.Scanln(&nama)
			fmt.Print("Jurusan: "); fmt.Scanln(&jurusan)
			fmt.Print("Nilai Tes: "); fmt.Scanln(&nilaiTes)
			fmt.Print("Nomor Pendaftaran: "); fmt.Scanln(&npm)
			tambah_data(M, n, nama, jurusan, nilaiTes, npm)
			fmt.Println()
		case 2:
			// mengubah data mahasiswa
			fmt.Print("Nama mahasiswa yang ingin diubah: "); fmt.Scanln(&nama)
			ubah_data(M, n, nama, jurusan, nilaiTes)
			fmt.Println()
		case 3:
			// menampilkan data mahasiswa pada jurusan tertentu atau tidak
			fmt.Println()
			fmt.Print("Ingin Menampilkan data pada jurusan tertentu? [Y/N]: "); fmt.Scanln(&keputusan)
			if keputusan == "Y" || keputusan == "y" {
				fmt.Print("Ketik Nama Jurusan: "); fmt.Scanln(&jurusan)				
				fmt.Println("-------------------------------------------------------------------------------")
				fmt.Println("Nama     |    Jurusan    |   Nilai Tes   | Nomor Pendaftaran |    Status       ")
				fmt.Println("-------------------------------------------------------------------------------")
				for i := 0; i < *n; i++ {
					if M[i].jurusan == jurusan {
						fmt.Printf("%s\t | %s\t | \t%d\t |     \t%d\t     |  %s\n", M[i].nama, M[i].jurusan, M[i].nilaiTes, M[i].npm, M[i].status)
					}
				}
			} else {
				jenis_urutan_admin(M, n)
			}
			fmt.Println()
		case 4:
			// mengurutkan data secara ascending
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.npm < M[i-1].npm {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			print_data(M, n)
			fmt.Println()
			fmt.Print("Masukan Nomor Pendaftaran Mahasiswa yang ingin dihapus: "); fmt.Scanln(&npm)
			fmt.Print("Yakin hapus data? [Y/N]: "); fmt.Scanln(&keputusan)
			if keputusan == "Y" || keputusan == "y" {
				// menghapus data mahasiswa
				hapus_data(M, n, npm)
				fmt.Print("Data Telah Terhapus!")
				fmt.Println()
			}
			fmt.Println()
		case 5:
			fmt.Println("Back to menu")
			fmt.Println()
		}
	}
}

func menu_mahasiswa(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa dan banyak data pada n
	   F.S. menampilkan menu calon mahasiswa dan beberapa menu yang dapat dipilih */
	var pilihan int
	var nama, jurusan string
	var nilaiTes, npm int
	var keputusan string

	for pilihan != 5 {
		fmt.Println("==========================================")
		fmt.Println("           MENU CALON MAHASISWA           ")
		fmt.Println("==========================================")
		fmt.Println("1. Tambah Data Mahasiswa")
		fmt.Println("2. Ubah Data Mahasiswa")
		fmt.Println("3. Tampilkan Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Exit")
		fmt.Println("---------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// menambahkan data mahasiswa
			fmt.Print("Nama Mahasiswa: "); fmt.Scanln(&nama)
			fmt.Print("Jurusan: "); fmt.Scanln(&jurusan)
			fmt.Print("Nilai Tes: "); fmt.Scanln(&nilaiTes)
			fmt.Print("Nomor Pendaftaran: "); fmt.Scanln(&npm)
			tambah_data(M, n, nama, jurusan, nilaiTes, npm)
			fmt.Println()
		case 2:
			// mengubah data mahasiswa
			fmt.Print("Nama mahasiswa yang ingin diubah: "); fmt.Scanln(&nama)
			ubah_data(M, n, nama, jurusan, nilaiTes)
			fmt.Println()
		case 3:
			// menampilkan data mahasiswa
			jenis_urutan_mahasiswa(M, n)
		case 4:
			// mengurutkan data secara ascending
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.npm < M[i-1].npm {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			print_data(M, n)
			fmt.Println()
			fmt.Print("Masukan Nomor Pendaftaran Mahasiswa yang ingin dihapus: "); fmt.Scanln(&npm)
			fmt.Print("Yakin hapus data? [Y/N]: "); fmt.Scanln(&keputusan)
			if keputusan == "Y" || keputusan == "y" {
				// menghapus data mahasiswa
				hapus_data(M, n, npm)
				fmt.Print("Data Telah Terhapus!")
				fmt.Println()
			}
			fmt.Println()
		case 5:
			fmt.Println("Back to menu")
			fmt.Println()
		}
	}
}

func tambah_data(M *TabMahasiswa, n *int, nama, jurusan string, nilaiTes, npm int) {
	/* I.S. terdefinisi M mahasiswa, banyak data pada n, nama, jurusan, dan nilai tes
	   F.S. menambahkan data mahasiswa yaitu nama, jurusan, nilai tes, dan status mahasiswa diterima atau tidak */
	M[*n].nama = nama
	M[*n].jurusan = jurusan
	M[*n].nilaiTes = nilaiTes
	M[*n].npm = npm
	if nilaiTes >= 70 {
		M[*n].status = "Diterima"
	} else {
		M[*n].status = "Tidak Diterima"
	}
	*n++
}

func get_nama_mahasiswa(M TabMahasiswa, n int, nama string) int {
	/* mengembalikan index mahasiswa berdasarkan nama mahasiswa yang dicari 
	pada array M yang berisi n mahasiswa (pencarian dengan sequential search),
	jika tidak ketemu akan mengembalikan nilai -1 */
	idx := -1
	i := 0

	for idx == -1 && i < n {
		if M[i].nama == nama{
			idx = i
		}
		i++
	}

	return idx
}

func get_npm(M TabMahasiswa, n int, npm int) int {
	/* mengembalikan index mahasiswa berdasarkan nomor pendaftaran yang dicari 
	pada array M yang berisi n mahasiswa (pencarian dengan binary search),
	jika tidak ketemu akan mengembalikan nilai -1 */
	idx := -1
	right := n-1
	left := 0

	for left <= right && idx == -1 {
		mid := (left + right) / 2

		if npm < M[mid].npm {
			right = mid - 1
		} else if npm > M[mid].npm {
			left = mid + 1
		} else {
			idx = mid
		}
	}

	return idx
}

func ubah_data(M *TabMahasiswa, n *int, nama, jurusan string, nilaiTes int) {
	/* I.S. terdefinisi data mahasiswa berupa nama, jurusan, dan nilai tes
	   serta sejumlah n data mahasiswa pada array M
	   F.S. data mahasiswa pada array M diperbaharui. jika nama ada pada
	   array M maka update jurusan dan nilai tes, jika tidak program
	   akan mengularkan string data tidak ditemukan 
	*/
	i := get_nama_mahasiswa(*M, *n, nama)

	if i != -1 {
		fmt.Println("NOTE: Lakukan perubahan jurusan dan nilai tes")
		fmt.Print("Jurusan: "); fmt.Scanln(&jurusan)
		fmt.Print("Nilai Tes: "); fmt.Scanln(&nilaiTes)
		M[i].jurusan = jurusan
		M[i].nilaiTes = nilaiTes
		if nilaiTes >= 70 {
			M[i].status = "Diterima"
		} else {
			M[i].status = "Tidak Diterima"
		}
	} else {
		fmt.Println("Data mahasiswa tidak ditemukan")
	}
}

func hapus_data(M *TabMahasiswa, n *int, npm int) {
	/* I.S. terdefinisi nomor pendaftaran mahasiswa pada array M sebanyak n data
	   F.S. menghapus data mahasiswa berdasarkan nomor pendaftaran */
	idx := get_npm(*M, *n, npm)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		i := idx
		for i <= *n-2 {
			M[i].nama = M[i+1].nama
			M[i].jurusan = M[i+1].jurusan
			M[i].nilaiTes = M[i+1].nilaiTes
			M[i].status = M[i+1].status
			M[i].npm = M[i+1].npm
			i++
		}
		*n--
	}
}

func print_data(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. mencetak data mahasiswa sebanyak n */
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("Nama     |    Jurusan    |   Nilai Tes   | Nomor Pendaftaran |    Status       ")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < *n; i++ {
		fmt.Printf("%s\t | %s\t | \t%d\t |     \t%d\t     |  %s\n", M[i].nama, M[i].jurusan, M[i].nilaiTes, M[i].npm, M[i].status)
	}
}

func jenis_urutan_admin(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan jurusan atau status pada menu admin */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Jenis Urutan Berdasarkan:")
		fmt.Println("1. Jurusan")
		fmt.Println("2. Status Penerimaan")
		fmt.Println("3. Exit")
		fmt.Println("-------------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure metode urutan jurusan
			metode_urutan_jurusan(M, n)
			fmt.Println()
		case 2:
			// panggil procedure metode urutan status
			metode_urutan_status(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
			fmt.Println()
		}
	}
}

func jenis_urutan_mahasiswa(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan nilai tes, jurusan, dan nama mahasiswa pada menu calon mahasiswa */
	var pilihan int

	for pilihan != 4 {
		fmt.Println()
		fmt.Println("Pilih Jenis Urutan Berdasarkan:")
		fmt.Println("1. Nilai Tes")
		fmt.Println("2. Jurusan")
		fmt.Println("3. Nama Mahasiswa")
		fmt.Println("4. Exit")
		fmt.Println("-------------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure metode urutan nilai tes
			metode_urutan_nilai_tes(M, n)
			fmt.Println()
		case 2:
			// panggil procedure metode urutan jurusan
			metode_urutan_jurusan(M, n)
			fmt.Println()
		case 3:
			// panggil procedure metode urutan nama
			metode_urutan_nama(M, n)
			fmt.Println()
		case 4:
			fmt.Println("Kembali")
			fmt.Println()
		}
	}
}

func metode_urutan_nilai_tes(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan nilai tes, pengurutan dapat dilakukan
	   secara selection atau insertion yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Metode Pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Exit")
		fmt.Println("------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure selection sort nilai
			selection_sort_nilai(M, n)
			fmt.Println()
		case 2:
			// panggil procedure insertion sort nilai
			insertion_sort_nilai(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func metode_urutan_jurusan(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan jurusan, pengurutan dapat dilakukan
	   secara selection atau insertion yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Metode Pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Batal")
		fmt.Println("------------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure selection sort jurusan
			selection_sort_jurusan(M, n)
		case 2:
			// panggil procedure insertion sort jurusan
			insertion_sort_jurusan(M, n)
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func metode_urutan_nama(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan nama mahasiswa, pengurutan dapat dilakukan
	   secara selection atau insertion yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Metode Pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Batal")
		fmt.Println("-----------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure selection sort nama
			selection_sort_nama(M, n)
		case 2:
			// panggil procedure insertion sort nama
			insertion_sort_nama(M, n)
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func metode_urutan_status(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan pilihan urutan data berdasarkan status, pengurutan dapat dilakukan
	   secara selection atau insertion yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Metode Pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("3. Batal")
		fmt.Println("-----------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// panggil procedure selection sort status
			selection_sort_status(M, n)
		case 2:
			// panggil procedur insertion sort status
			insertion_sort_status(M, n)
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func selection_sort_nilai(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan selection sort berdasarkan nilai tes, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan ascending(menaik)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].nilaiTes > M[i].nilaiTes {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan descending(menurun)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].nilaiTes < M[i].nilaiTes {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func selection_sort_jurusan(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan selection sort berdasarkan jurusan, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan ascending(menaik)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].jurusan > M[i].jurusan {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan descending(menurun)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].jurusan < M[i].jurusan {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func selection_sort_nama(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan selection sort berdasarkan nama mahasiswa, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan ascending(menaik)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].nama > M[i].nama {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan descending(menurun)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].nama < M[i].nama {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func selection_sort_status(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan selection sort berdasarkan status mahasiswa, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan ascending(menaik)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].status > M[i].status {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan descending(menurun)
			pass := 1
			for pass <= *n-1 {
				idx := pass-1
				i := pass

				for i < *n {
					if M[idx].status < M[i].status {
						idx = i
					}
					i++
				}

				temp := M[pass-1]
				M[pass-1] = M[idx]
				M[idx] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func insertion_sort_nilai(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan insertion sort berdasarkan nilai tes, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------")
		fmt.Print("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan menaik(ascending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.nilaiTes < M[i-1].nilaiTes {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan menurun(descending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.nilaiTes > M[i-1].nilaiTes {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func insertion_sort_jurusan(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan insertion sort berdasarkan jurusan, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Println("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan menaik(ascending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.jurusan < M[i-1].jurusan {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan menurun(descending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.jurusan > M[i-1].jurusan {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func insertion_sort_nama(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan insertion sort berdasarkan nama mahasiswa, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Println("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan menaik (ascending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.nama < M[i-1].nama {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan menurun (descending)
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.nama > M[i-1].nama {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func insertion_sort_status(M *TabMahasiswa, n *int) {
	/* I.S. terdefinisi M mahasiswa sebanyak n data
	   F.S. menampilkan urutan data menggunakan insertion sort berdasarkan status mahasiswa, pengurutan dapat dilakukan
	   secara ascending atau descending yang dapat dipilih pengguna */
	var pilihan int

	for pilihan != 3 {
		fmt.Println()
		fmt.Println("Pilih Tipe Urutan:")
		fmt.Println("1. Ascending  (Menaik)")
		fmt.Println("2. Descending (Menurun)")
		fmt.Println("3. Batal")
		fmt.Println("-----------------------")
		fmt.Println("Pilihan: "); fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			// pengurutan menaik
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.status < M[i-1].status {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 2:
			// pengurutan menurun
			pass := 1
			for pass <= *n-1 {
				i := pass
				temp := M[pass]
				for i > 0 && temp.status > M[i-1].status {
					M[i] = M[i-1]
					i--
				}
				M[i] = temp
				pass++
			}
			// menampilkan data terurut
			print_data(M, n)
			fmt.Println()
		case 3:
			fmt.Println("Kembali")
		}
	}
}

func main() {
	// deklarasi variabel
	var m TabMahasiswa
	var n int
	// panggil procedure menu
	menu(&m, &n)
}