$(function(){
    $.ajax({
        type: "GET",
        url: "/Message/GetCount",
        success: function(data){
            document.getElementsByClassName("Count")[0].innerHTML=data;
        },
        error: function (data) {
            console.log(data);
        }
    });
})