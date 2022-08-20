package models

import "time"

type Mahasiswa struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type Admin struct {
	Id       int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Penyelenggara struct {
	Id            int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	NamaAcara     string    `json:"nama_acara"`     // Pemilihan Ketua PUB
	ThnPemilihan  string    `json:"thn_pemilihan"`  // 2020/2021
	TanggalAcara  time.Time `json:"tanggal_acara"`  // 17 Agustus 2022
	JumlahPemilih int       `json:"jumlah_pemilih"` // 90
	Sssah         int       `json:"sssah"`          // 80
	Sstdksah      int       `json:"sstdsah"`        // 5
	Sstdkterpakai int       `json:"sstdkterpakai"`  // 5
	Ketua         string    `json:"ketua"`          // Abdul Hafizh Tanjung
}

type Pemilih struct {
	Id              int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Username        string `json:"username"` // NIM
	Password        string `json:"password"` // Tahun-Bulan-Tanggal lahir
	Nama            string `json:"nama"`
	KeyValue        string `json:"key_value"`
	Email           string `json:"email"`
	Phone1          string `json:"phone1"`
	Phone2          string `json:"phone2"`
	Active          string `json:"active"`
	PenyelenggaraId int    `json:"penyelenggara_id"`
	Penyelenggara
}

type Kandidat struct {
	Id              int    `gorm:"primarykey;autoIncrement" json:"id"`
	NoUrut          string `json:"no_urut"`
	Nama            string `json:"nama"`
	Nik             string `json:"nik"`
	Tmp_lahir       string `json:"tmp_lahir"`
	Tgl_lahir       string `json:"tgl_lahir"`
	Kelamin         string `json:"kelamin"`
	Agama           string `json:"agama"`
	Photo           string `json:"photo"`
	Keterangan      string `json:"keterangan"`
	PenyelenggaraId int    `json:"penyelenggara_id"`
	Penyelenggara
}
