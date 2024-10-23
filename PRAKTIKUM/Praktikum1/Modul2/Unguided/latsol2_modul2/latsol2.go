package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan kelas : ")
	fmt.Scan(&kelas)

	//hasil
	fmt.Println("--- Resume Biodata Mahasiswa --- ")
	fmt.Printf("Perkenalkan saya adalah %s salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s\n", nama, kelas, nim)
}
