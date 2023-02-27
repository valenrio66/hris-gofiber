package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfilPribadi struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama           string             `json:"nama,omitempty" bson:"nama,omitempty"`
	JenisKelamin   string             `json:"jenis_kelamin,omitempty" bson:"jenis_kelamin,omitempty"`
	TempatLahir    string             `json:"tempat_lahir,omitempty" bson:"tempat_lahir,omitempty"`
	TanggalLahir   time.Time          `json:"tanggal_lahir,omitempty" bson:"tanggal_lahir,omitempty"`
	NamaIbuKandung string             `json:"nama_ibu_kandung,omitempty" bson:"nama_ibu_kandung,omitempty"`
}
