$(document).ready(function() {
    $.ajax({
        url: '/api/info',
    }).then(function(data) {
       $('.staticP').append(data.staticPJ);
       $('.jsP').append(data.jsPJ);
       $('.tableP').append(data.tablePJ);
    });
});
