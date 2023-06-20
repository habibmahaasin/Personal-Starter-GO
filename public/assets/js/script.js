window.onscroll = () => {
  if ($(window).width() > 768) {
    if (window.scrollY > 40) {
      $(".user-information-wrapper").attr(
        "style",
        "display: none !important; transition: 0.2s;"
      );
      $("#header").addClass("hidden-header");
    } else {
      $(".user-information-wrapper").attr(
        "style",
        "display: block; transition: 0.2s;"
      );
      $("#header").removeClass("hidden-header");
    }
  }
};

function ControllingFromList(condition) {
  var conditionSplitter = condition.split('"').join("");
  Swal.fire({
    title:
      '<span style="font-size: 16px;font-weight: 400;">Apakah Kamu Yakin Ingin Mengubah Kondisi Perangkat?</span>',
    showDenyButton: true,
    confirmButtonText: "Ya",
    confirmButtonColor: "#282689",
    denyButtonText: `Tidak`,
  }).then((result) => {
    if (result.isConfirmed) {
      window.location = "/control/" + conditionSplitter;
    } else if (result.isDenied) {
      window.location = "/daftar-perangkat";
    }
  });
}

function ControllingFromDetail(condition) {
  var conditionSplitter = condition.split('"').join("");
  var path_id = conditionSplitter.split('/')[1];
  Swal.fire({
    title:
      '<span style="font-size: 16px;font-weight: 400;">Apakah Kamu Yakin Ingin Mengubah Kondisi Perangkat?</span>',
    showDenyButton: true,
    confirmButtonText: "Ya",
    confirmButtonColor: "#282689",
    denyButtonText: `Tidak`,
  }).then((result) => {
    if (result.isConfirmed) {
      window.location = "/control/" + conditionSplitter;
    } else if (result.isDenied) {
      window.location = "/detail-perangkat/"+path_id;
    }
  });
}

function DeleteDevice(id_perangkat) {
  Swal.fire({
    title:
      '<span style="font-size: 16px;font-weight: 400;">Apakah Kamu Yakin Ingin Menghapus Perangkat?</span>',
    showDenyButton: true,
    confirmButtonText: "Ya",
    confirmButtonColor: "#282689",
    denyButtonText: `Tidak`,
  }).then((result) => {
    if (result.isConfirmed) {
      window.location = "/hapus-perangkat/" + id_perangkat;
    } else if (result.isDenied) {
      window.location = "/detail-perangkat/" + id_perangkat;
    }
  });
}

$(document).ready(function () {
  $("#guppytech-table.report-page").DataTable({
    order: [[0, "desc"]],
    scrollX: true,
    pageLength: 5,
    lengthMenu: [
      [5, 10, 25, 50, -1],
      [5, 10, 25, 50, "All"],
    ],
    pagingType: "full",
    language: {
      emptyTable: "Tidak ada data yang tersedia",
      zeroRecords: "Data tidak ditemukan",
      lengthMenu: "Baris per halaman _MENU_ ",
      info: " _START_ - _END_ dari _TOTAL_ ",
      loadingRecords: "Loading ..",
      paginate: {
        previous: "<i class='bx bx-chevron-left bx-sm'></i>",
        next: "<i class='bx bx-chevron-right bx-sm'></i>",
        first: "<i class='bx bx-chevrons-left bx-sm'></i>",
        last: "<i class='bx bx-chevrons-right bx-sm'></i>",
      },
    },
  });
});

$(document).ready(function () {
  $("#guppytech-table.list-device").DataTable({
    order: [[2, "desc"]],
    scrollX: true,
    pageLength: 5,
    lengthMenu: [
      [5, 10, 25, 50, -1],
      [5, 10, 25, 50, "All"],
    ],
    pagingType: "full",
    language: {
      emptyTable: "Tidak ada data yang tersedia",
      zeroRecords: "Data tidak ditemukan",
      lengthMenu: "Baris per halaman _MENU_ ",
      info: " _START_ - _END_ dari _TOTAL_ ",
      loadingRecords: "Loading ..",
      paginate: {
        previous: "<i class='bx bx-chevron-left bx-sm'></i>",
        next: "<i class='bx bx-chevron-right bx-sm'></i>",
        first: "<i class='bx bx-chevrons-left bx-sm'></i>",
        last: "<i class='bx bx-chevrons-right bx-sm'></i>",
      },
    },
  });
});


// clock
var hands = [];
hands.push(document.querySelector("#secondhand > *"));
hands.push(document.querySelector("#minutehand > *"));
hands.push(document.querySelector("#hourhand > *"));

var cx = 100;
var cy = 100;

function shifter(val) {
  return [val, cx, cy].join(" ");
}

var date = new Date();
var hoursAngle = (360 * date.getHours()) / 12 + date.getMinutes() / 2;
var minuteAngle = (360 * date.getMinutes()) / 60;
var secAngle = (360 * date.getSeconds()) / 60;

hands[0].setAttribute("from", shifter(secAngle));
hands[0].setAttribute("to", shifter(secAngle + 360));
hands[1].setAttribute("from", shifter(minuteAngle));
hands[1].setAttribute("to", shifter(minuteAngle + 360));
hands[2].setAttribute("from", shifter(hoursAngle));
hands[2].setAttribute("to", shifter(hoursAngle + 360));

for (var i = 1; i <= 12; i++) {
  var el = document.createElementNS("http://www.w3.org/2000/svg", "line");
  el.setAttribute("x1", "100");
  el.setAttribute("y1", "30");
  el.setAttribute("x2", "100");
  el.setAttribute("y2", "40");
  el.setAttribute("transform", "rotate(" + (i * 360) / 12 + " 100 100)");
  el.setAttribute("style", "stroke: #ffffff;");
  document.querySelector("svg").appendChild(el);
}

var dayElement = document.getElementById("day");
var d = new Date();
var day = new Array(7);
var month = new Array(12);

day[0] = "Sunday";
day[1] = "Monday";
day[2] = "Tuesday";
day[3] = "Wednesday";
day[4] = "Thursday";
day[5] = "Friday";
day[6] = "Saturday";

month[0] = "Januari";
month[1] = "Februari";
month[2] = "Maret";
month[3] = "April";
month[4] = "Mei";
month[5] = "Juni";
month[6] = "Juli";
month[7] = "Agustus";
month[8] = "September";
month[9] = "Oktober";
month[10] = "November";
month[11] = "Desember";

var dayNow = day[d.getDay()];
var monthNow = month[d.getMonth()];

var now = dayNow + " , " + d.getDate() + " " + monthNow + " " + d.getFullYear();

dayElement.innerHTML = now;

// end clock