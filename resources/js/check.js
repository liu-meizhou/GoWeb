//函数：验证C语言命名规则格式
function IsName(strName){
    //定义正则表达式的变量:C语言命名规则正则
    var reg=/[a-zA-Z_][a-zA-Z0-9_]*/;
    if(strName.search(reg) != -1)
    {
        return true;
    }
    else{
        return false;
    }
}

//onblur失去焦点事件，用户离开输入框时执行 JavaScript 代码：
//函数：验证邮箱格式
function IsEmail(strEmail){
    //定义正则表达式的变量:邮箱正则
    var reg=/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
    if(strEmail.search(reg) != -1)
    {
        return true;
    }
    else{
        return false;
    }
}

//函数：验证电话格式
function IsPhone(strPhone){
    //定义正则表达式的变量:电话正则
    var reg=/^[1][3,4,5,7,8,9][0-9]{9}$/;
    if(strPhone.search(reg) != -1)
    {
        return true;
    }
    else{
        return false;
    }
}