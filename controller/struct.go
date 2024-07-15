package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Nasi Goreng"`
	Harga     float64            `bson:"harga,omitempty" json:"harga,omitempty" example:"15000"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty" example:"Nasi goreng dengan campuran sayuran dan ayam"`
	Kategori  Kategori           `bson:"kategori,omitempty" json:"kategori,omitempty"`
	BahanBaku BahanBaku          `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
}

type Kategori struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty" example:"Makanan"`
}

type BahanBaku struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	BahanBaku string             `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty" example:"Beras"`
	Jumlah    string             `bson:"jumlah,omitempty" json:"jumlah,omitempty" example:"1"`
}
