# Gadai
API Gadai online pinjam.co.id,

##Endpoint - /gadai/v2/

1. API /taksir     - Taksir gadai online
....* POST /taksir

2. API /simulasi   - Simulasi gadai online
....* GET  /simulasi/:id
....* POST /simulasi

3. API POST /ajukan     - Pengajuan gadai online
....* GET  /ajukan/:id
....* POST /ajukan

4. API POST /verifikasi - Perifikasi pengajuan gadai online
....* POST  /verifikasi
....* POST  /verifikasi

5. API POST /setujui    - Penyetujuan setelah di verifikasi
6. API POST /cairkan    - Pencairan transaksi gadai
7. API POST /perpanjang - Perpanjang transaksi gadai
8. API POST /batalkan   - Batalkan transaksi gadai
