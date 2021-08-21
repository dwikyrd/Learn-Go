package main

import "fmt"

type Kelas interface {
    KelasSatu() string
}

type Mahasiswa struct{}

func NewSekolahService() Kelas {
    return Mahasiswa{}
}

func (m Mahasiswa) KelasSatu() string {
    return "Hello world"
}

func main() {
    sekolah := NewSekolahService()
    fmt.Println(sekolah.KelasSatu())
}