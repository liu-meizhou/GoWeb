function Register(fn) {
    var name=$("#name").val();
    var email=$("#email").val();
    var phone=$("#phone").val();
    var password=$("#password").val();
    if(IsName(name) && IsEmail(email) && IsPhone(phone)) {//
        $.ajax({
            type: "POST",
            url: "/Register",
            data: {username:name, email:email, phone:phone, password:password},
            success: function(datas){
                fn(datas);
            },
            error: function (data) {
                console.log(data);
            }
        });
   }
    else{
       $("#message").text("注册信息都不能为空或者格式都要正确");
   }
}