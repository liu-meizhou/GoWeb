function check() {
    var index=1;
    var username=$("#myname").val();
    if (username=="") {
        index=2;
        username=$("#myemail").val();
        if (username==""){
            index=3;
            username=$("#myphone").val();
            if (username==""){
                $("#message").text("还没输入用户名/邮箱/电话");
                return false;
            }
        }
    }
    var password=$("#password").val();
    if (password==""){
        $("#message").text("还没输入密码");
        return false;
    }
    var data = {
        index: index,
        username: username,
        password: password
    };

    $.ajax({
        type: "POST",
        url: "/Login",
        data: data,
        success: function(data){
            console.log(data)
            if (data=='OK'){
                console.log("ok");
                window.location.href='../Logined/Login';
            }
            else{
                $("#message").text(data);
            }
        },
        error: function (data) {
            alert("出现异常");
            console.log(data);
        }
    });
}