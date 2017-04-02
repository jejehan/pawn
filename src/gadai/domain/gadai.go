package gadai

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Gadai adalah kode utama dari domain model
// Gadai di identifikasi dengan unik transaksi id, juga dengan
// kode antrian dan kodagadai. Alur prosess gadai seperti ini :
// `Taksir-Simulasi-Ajukan-Verifikasi-Setujui-Cairkan`
// Setelah itu akan ada proses tambahan tergantung kasus yang terjadi, yaitu
// Pelunasan-Terlambat-Perpanjangan-Lelang
// Pelunasan adalah ketika nasabah mengembalikan dana yang di pinjamkan
// Terlambat adalah ketika nasabah belum mengembalikan dana hingga melebihi jatuh tempo
// Perpanjangan adalah ketika nasabah tidak bisa mengembalikan dana dan minta perpanjangan waktu
// Lelang adalah ketika nasabah tidak bisa mengembalikan dana sehingga barang yang dijaminkan
// akan dijual oleh pihak gadein.com.
/*type Gadai interface {
	//Taksir(request interface{}) TaksirResponse
	//Simulation() string
	//Ajukan() string
	//BiayaSewaModal() string
}*/

//TaksirID menghasilkan uuid
func TaksirID() string {
	u := uuid.NewV4()
	return u.String()
}

//TransaksiID mengasilkan uuid
func TransaksiID() string {
	u := uuid.NewV4()
	return u.String()
}

//DateDiffYear mengitung jarak tahun antara tanggal
func DateDiffYear(start, end time.Time) int {
	diff := start.Sub(end)
	days := int(diff.Hours() / 24)
	years := days / 365
	return years
}
