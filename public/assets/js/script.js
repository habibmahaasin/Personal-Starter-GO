window.onscroll = () => {
    if ($(window).width() > 768) {
        if (window.scrollY > 40) {
            $(".user-information-wrapper").attr(
            "style",
            "display: none !important; transition: 0.2s;"
            );
            $('#header').addClass('hidden-header');
        } else {
            $(".user-information-wrapper").attr(
            "style",
            "display: block; transition: 0.2s;"
            );
            $('#header').removeClass('hidden-header');
        }
    }
};

function Controlling(condition){
    var conditionSplitter = condition.split('"').join("")
    Swal.fire({
        title: '<span style="font-size: 16px;font-weight: 400;">Apakah Kamu Yakin Ingin Mengubah Kondisi Perangkat?</span>',
        showDenyButton: true,
        confirmButtonText: 'Ya',
        confirmButtonColor: '#282689',
        denyButtonText: `Tidak`,
        }).then((result) => {
        if (result.isConfirmed) {
            window.location = "/control/" + conditionSplitter;
        } else if (result.isDenied) {
            window.location = "/daftar-perangkat";
        }
    })
}

$(document).ready(function () {
    $('#guppytech-table').DataTable({
      order: [[0, 'desc']],
      scrollX: true,
      "pageLength": 5,
      "lengthMenu": [
          [5,10, 25, 50, -1],
          [5,10, 25, 50, 'All'],
      ],
      "pagingType": 'full',
      "language": {
          "emptyTable": "Data Pengguna kosong",
          "zeroRecords": "Data tidak ditemukan",
          "lengthMenu": "Baris per halaman _MENU_ ",
          "info": " _START_ - _END_ dari _TOTAL_ ",
          "loadingRecords": "Loading ..",
          "paginate": {
              'previous': "<i class='bx bx-chevron-left bx-sm'></i>",
              'next': "<i class='bx bx-chevron-right bx-sm'></i>",
              "first": "<i class='bx bx-chevrons-left bx-sm'></i>",
              "last": "<i class='bx bx-chevrons-right bx-sm'></i>"
          }
      }
    });
});