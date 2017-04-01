package gadai

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
type Gadai interface {
	TaksiranID(uuid string) string
	Taksir() string
	Simulation() string
	TransaksiID(uuid string) string
	//Ajukan(db gadai.Repository) string
	BiayaSewaModal() string
}
