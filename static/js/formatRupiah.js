// Fungsi untuk format angka ke Rupiah
function formatRupiah(angka) {
    // Menggunakan toLocaleString untuk memformat angka sebagai mata uang Indonesia
    return "Rp " + angka.toLocaleString('id-ID');
  }
  
  document.addEventListener("DOMContentLoaded", function () {
    // Seleksi elemen harga dan format angkanya
    const prices = document.querySelectorAll(".price");
    prices.forEach(price => {
      // Ambil nilai 'data-value' dan pastikan itu adalah angka
      const number = parseFloat(price.getAttribute("data-value"));
      if (!isNaN(number)) { // Pastikan number adalah angka yang valid
        price.textContent = formatRupiah(number);
      }
    });
  });
  