//获取验证码和MyCookie
$(function(){
    console.log("进来了")
    $.ajax({
        type: "GET",
        url: "/visitor",
        success: function(data){
            //显示图片验证码
            console.log(data);
        },
        error: function (data) {
            alert("出现异常");
            console.log(data);
        }
    });
})

//函数：验证C语言命名规则格式
function isName(strName){
    if(IsName(strName))
    {
        $('#myname').text("OK");
    }
    else{
        $('#myname').text("输入的名字格式有误!!!");
    }
}



//onblur失去焦点事件，用户离开输入框时执行 JavaScript 代码：
//函数：验证邮箱格式
function isEmail(strEmail){
    if(IsEmail(strEmail))
    {
        $('#myemail').text("OK");
    }
    else{
        $('#myemail').text("输入的邮箱格式有误!!!");
    }
}

//函数：验证电话格式
function isPhone(strPhone){
    if(IsPhone(strPhone))
    {
        $('#myphone').text("OK");
    }
    else{
        $('#myphone').text("输入的电话格式有误!!!");
    }
}
