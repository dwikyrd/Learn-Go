package main

import (
	"errors"
	"fmt"
)

// struct
type Mahasiswa struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}

func (m *Mahasiswa) SetNama(input string) error {
    if input == "" {
        return errors.New("nama dak boleh kosong")
    }
    m.Nama = input
    return nil
}

func main(){
    mahasiswa := Mahasiswa{
        ID: 1,
    }
    err := mahasiswa.SetNama("abda")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(mahasiswa)
}

// interface

// package integration
