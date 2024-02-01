function getTanggalSekarang() {
  const tanggalSekarang = new Date();
  return tanggalSekarang.toDateString();

}
function cetakTanggalSekarang() {
  console.log('Tanggal sekarang:', getTanggalSekarang());

}
cetakTanggalSekarang();
