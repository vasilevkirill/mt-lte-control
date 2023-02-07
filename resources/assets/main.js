setInterval(function getInfo(){
    $.ajax({
        url: "/getinfo",
        type: "GET",
        success: function (data){
            ajaxSuccess(data)
        },
        dataType: "json"

    });

}, 5000);

function ajaxSuccess(data){
    console.log(data)
}
