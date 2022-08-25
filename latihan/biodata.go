package latihan

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func BiodataFun() {
	var biodata = []Biodata{
		{
			Nama:      "Jajan Nurjaman",
			Alamat:    "Bandung",
			Pekerjaan: "Pelatih",
			Alasan:    "Mencari peluang kerja baru",
		},
		{
			Nama:      "Budi Gunawan",
			Alamat:    "Jakarta",
			Pekerjaan: "Pelukis",
			Alasan:    "Mencari Kesibukan baru",
		},
		{
			Nama:      "Rendi Siregar",
			Alamat:    "Medan",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Meningkatkan skill programming",
		},
	}

	var args = os.Args
	var params = args[1]
	intParams, err := strconv.Atoi(params)
	_ = err

	fmt.Printf("%+v\n", biodata[intParams])

}
