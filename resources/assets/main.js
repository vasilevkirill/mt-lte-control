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
    changeSimSlot(data["simSlot"])
    changeLteInfo(data["lteInfo"])


}

function changeLteInfo(data){
    let ob = $("#lteInfo")
    if (data ===""){
        ob.text("нет информации")
        return
    }
    let tbl = objectToTable(data)
    ob.html(tbl)
}

function changeSimSlot(data){
    let ob = $("#simSlot")
    if (data ===""){
        ob.text("нет информации")
        return
    }
    ob.text(data)
}

function objectToTable(data) {
    let table = $('<table>');
    table.addClass("table is-bordered is-striped is-narrow is-hoverable is-fullwidth")
    $.each(data, function (key, v){
        let colKey = $('<td>').text(key)
        let colV = $('<td>').text(v)
        let row = $('<tr>');
        row.append(colKey)
        row.append(colV)
        //td1 = $("<td></td>").text(key).appendTo(tr)
        //td2 = $("<td></td>").text(value).appendTo(tr)
        table.append(row)
    })

    return table;
}