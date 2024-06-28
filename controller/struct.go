package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga     float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Kategori  Kategori           `bson:"kategori,omitempty" json:"kategori,omitempty"`
	BahanBaku BahanBaku          `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
}

type Kategori struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_kategori string             `bson:"nama_kategori,omitempty" json:"nama_kategori,omitempty"`
}

type BahanBaku struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_bahan_baku string             `bson:"nama_bahan_baku,omitempty" json:"nama_bahan_baku,omitempty"`
	Jumlah          string             `bson:"jumlah,omitempty" json:"jumlah,omitempty"`
}
