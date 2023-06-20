
$('#koor_loc').on('change', function() {
    if ( this.value == 'curr-loc') {
        $('.manual-loc-form').css('display','none');
        $('#map').css('display','block');
        getLocation();
    } 
    else if (this.value == 'input-loc-manual'){
        $('#latitude').val(null);
        $('#longitude').val(null);
        $('#map').css('display','none');
        $('.manual-loc-form').css('display','block');
    }
    else if (this.value == 'pin-from-maps'){
        $('#map').css('display','block');
        $('.manual-loc-form').css('display','none');
        
        leafletMap(-6.901779067052239, 107.61896902369566, 12);
    }
  });
  
  x = document.getElementById("map");
  function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else { 
        x.innerHTML = "Geolocation is not supported by this browser.";
    }
  }
  
  function showPosition(position) {
    leafletMap(position.coords.latitude, position.coords.longitude,18);
  }
  
  function leafletMap(latitude,longitude,zoomView){
    $('#latitude').val(latitude);
    $('#longitude').val(longitude);
  
    document.getElementById('wrapper-map').innerHTML = "<div id='map' style='width: 100%; height: 400px; border-radius:10px'></div>";
    const map = L.map('map').setView([latitude, longitude], zoomView);
    const tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
  
    var popup = L.popup();
  
    var theMarker = {};
    theMarker = L.marker([latitude,longitude]).addTo(map);
  
    selected_koor = document.getElementById('selected-koor');
    function onMapClick(e) {
        popup
            .setLatLng(e.latlng)
            .setContent(e.latlng.lat + "," + e.latlng.lng)
            .openOn(map);
  
        if (theMarker != undefined) {
            map.removeLayer(theMarker);
        };
        theMarker = L.marker([e.latlng.lat, e.latlng.lng]).addTo(map);
        $('#latitude').val(e.latlng.lat);
        $('#longitude').val(e.latlng.lng);
    }
    map.on('click', onMapClick);
  }