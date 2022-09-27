package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Student struct {
	ID                 uint
	Name               string
	Address            string
	Occupation         string
	WhyTakeGolangClass string
}

var Students = map[int]Student{}
var nomorAbsenDalamString string

func init() {
	Students = map[int]Student{
		1: {ID: 1, Name: "Rizki", Address: "Malang", Occupation: "Student", WhyTakeGolangClass: "Get golang skill"},
		2: {ID: 2, Name: "Pratama", Address: "Surabaya", Occupation: "Engineer", WhyTakeGolangClass: "Upskilling"},
		3: {ID: 3, Name: "Turbina", Address: "Yogyakarta", Occupation: "Fresh Graduate", WhyTakeGolangClass: "Fullfill leisure time"},
	}
}

func main() {

	if CheckOsArgument() {
		nomorAbsenDalamString = os.Args[1]
	}
	nomorAbsen, _ := strconv.Atoi(nomorAbsenDalamString)

	if !IsNumberPositive(nomorAbsen) {
		log.Fatal("Nomor absen harus bernilai positif")
	}

	fmt.Printf("Student %v: %+v\n", nomorAbsen, GetStudent(nomorAbsen))
}

func IsNumberPositive(nomor int) bool {
	return nomor > 0
}

func CheckOsArgument() bool {
	argumenOs := os.Args

	if len(argumenOs) < 2 {
		log.Fatal(
			`Silakan masukkan nomor absen untuk mencari Student pada command line dengan format "go run biodata.go <nomor absen>"`)
	}

	return true
}

func GetStudent(nomor int) Student {
	value, exist := Students[nomor]
	if !exist {
		log.Fatalf("Student dengan nomor absen %v tidak ditemukan", nomor)
	}
	return value
}
