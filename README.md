# Skeleton gadein.com Gadai v.2

Domain Driven Design untuk system gadein.com, dengan domain

1. Gadai
domain gadai adalah domain yang menyelesaikan/menangangi transaksi gadai

    * API /gadai/v2/taksir

        ```
        curl -XPOST -d"barang_id=236ae212-0e23-41cf-821f-e811c49a9d21&merk=Apple&Tipe=iPhone 7 Plus 256GB&warna=Hitam&kapasita_memori=256&tahun_pembelian=2016-01-01&harga_beli=15000000" localhost:7070/gadai/v2/taksiran
        ```

2. Lelang
domain auction adalah domain yang menyelesaikan/menangangi transaksi lelang
3. Mitra
domain Parnership adalah domain yang mengurusi pendaftaran mitra
4. Anggota
domain Membership menangani service keangggotaan
5. Pengguna
domain Membership menangani service pengguna