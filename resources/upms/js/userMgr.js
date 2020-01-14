$(function(){
    $.ajax({
        type: "POST",
        url: "/Logined/GetUserMgr",
        success: function(datas){
            var users=JSON.parse(datas);
            Makehtml(users);
        },
        error: function (data) {
            console.log(data);
        }
    });
    wsConnect();
})

//获取一条信息
function htmlGetHisMessage(time,type,message){return getmessage(time,type,message,"textLeft");}
//构建用户聊天按钮
function htmlUserChatButton(hasID,ID){
    var chat="<button onclick='Chat(this)'";
    if(hasID){
        chat+=" class='Red'>联系"+ ID +"</button>";
    }else{
        chat+=" class=''>联系0</button>";
    }
    return chat;
}
//构建一个用户
function htmlAddUser(user,chat) {
    return "<tr id='user"+ user['ID'] +"'> \
                <td>" + user['ID'] + "</td>\
                <td>" + user["name"] + "</td>\
                <td>" + user["email"] + "</td>\
                <td>" + user["phone"] + "</td>\
                <td>" + user["password"] + "</td>\
                <td>\
                    <button onclick='Updata(this)'>修改</button> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp\
                    <button onclick='Delete(this)'>删除</button> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp" +
                    chat +
            "</tr>";
}
//构建一个聊天
function htmlAddChatWindow(indexID,name,message){
    return "<div id='Chat"+indexID+"'>\
                <span>"+name+"</span>\
                <button style='float: right;' onclick='CloseChat()'>关闭聊天</button>\
                <hr>\
                <div class='ChatMessage'>"+
                    message +
                "</div>\
                <div class='EditMessage'>\
                    <input class='SendText' type='text' onkeypress='keypress(this,event)' style='width:80%; height:20px;' />\
                    <button class='Send' onclick='SendOnButton(this)'>发送</button>\
                    <button class='SendImage'>发送图片</button>\
                </div>\
            </div>";
}
//构建所有用户信息
function Makehtml(users){
    GetMessages(function(messages){
        var messagemap = new Map();
        var messageIDmap = new Map();
        var arry=[];
        var up;
        for (var message in messages){
            var sendID = messages[message]["SendID"];
            if (messagemap.has(sendID)){
                arry.push(message);
                messagemap.set(sendID,messagemap.get(sendID)+1);
            }else{
                if (arry.length!=0){
                    messageIDmap.set(up,arry);
                }
                up=sendID;
                arry=[message];
                messagemap.set(sendID,1);
            }
        }
        messageIDmap.set(up,arry);

        var html="";
        var messagehtml="";
        for (var data in users) {
            var ID=users[data]["ID"];
            var chat=htmlUserChatButton(messagemap.has(ID),messagemap.get(ID));
            html+=htmlAddUser(users[data],chat)
            var maphtml="";
            if(messageIDmap.has(ID)){
                messageIDmap.get(ID).forEach(function(item){
                    maphtml+=htmlGetHisMessage(messages[item]["SendTime"],messages[item]["Type"],messages[item]["Context"]);
                })
            }
            messagehtml+=htmlAddChatWindow(indexID,users[data]["name"],maphtml);
            indexID+=1;
        }
        $("#Users").html(html);
        $("#ChatMove").html(messagehtml);
    })
}
//设置对话框数据
function SetDialogData(name,email,phone,password){
    $("#name").val(name);
    $("#email").val(email);
    $("#phone").val(phone);
    $("#password").val(password);
    $("#message").text("");
}
//打开对话框
function OpenDialog(index){
    var id="user"+index;
    var text='添加';
    if(index!=0){
        text='修改';
        var children = document.getElementById(id).children;
        SetDialogData(children[1].innerText,children[2].innerText,children[3].innerText,children[4].innerText);
    }
    $('#AddUser').dialog({
        width:"400",
        height:"350",
        modal:true,
        closed:true,
        cache: false,
        buttons: [
            {
                text: text,
                click: function() {
                    $("#message").text("");
                    var name=$("#name").val();
                    var email=$("#email").val();
                    var phone=$("#phone").val();
                    var password=$("#password").val();
                    if(index==0){
                        function rback(data){
                            $("#message").text("正在添加,请稍等...");
                            $("#Users").append(htmlAddUser({ID:data,name:name,email:email,phone:phone,password:password},htmlUserChatButton(false)));
                            $("#ChatMove").append(htmlAddChatWindow(indexID,name,""))
                            $('#AddUser').dialog( "close" );
                            SetDialogData("","","","");
                            indexID+=1;
                        }
                        Register(rback)
                    }else{
                        if(IsName(name) && IsEmail(email) && IsPhone(phone)) {
                            $.ajax({
                                type: "POST",
                                url: "/Logined/Updata",
                                data: {ID:index, username:name, email:email, phone:phone, password:password},
                                success: function(datas){
                                    $("#message").text("正在修改,请稍等...");
                                    var children = document.getElementById(id).children;
                                    children[1].innerText=name;
                                    children[2].innerText=email;
                                    children[3].innerText=phone;
                                    children[4].innerText=password;
                                    SetChatWindowName(GetIndexByID(index),name);
                                    $('#AddUser').dialog( "close" );
                                    SetDialogData("","","","");
                                    console.log(datas);
                                },
                                error: function (data) {
                                    console.log(data);
                                }
                            });
                        }
                        else{
                            $("#message").text("修改信息都不能为空或者格式都要正确");
                        }
                    }
                }
            }, 
            {
                text: "关闭",
                click: function() {
                    $( this ).dialog( "close" );
                    SetDialogData("","","","");
                }
            }
        ]
    });
    $('#AddUser').dialog('open');
}
//添加用户
function Add(){
    console.log("添加");
    OpenDialog(0);
}; 
//修改用户信息
function Updata(obj){
    var index = obj.parentNode.parentNode.firstElementChild.innerText;
    console.log("修改ID为"+index);
    OpenDialog(index);
};
//删除用户
function Delete(obj){
    var index = obj.parentNode.parentNode.firstElementChild.innerText;
    console.log("删除ID为"+index);
    if(window.confirm('确定要删除ta吗？')){
        $.ajax({
            type: "POST",
            url: "/Logined/Delete",
            data: {ID:index},
            success: function(datas){
                $("#Chat"+GetIndexByID(index)).remove()
                $("#user"+index).remove();
                indexID-=1;
                console.log(datas);
            },
            error: function (data) {
                console.log(data);
            }
        });
        console.log("确定");
   }else{
        console.log("取消");
   }
};
//与用户聊天
function Chat(obj){
    var index = obj.parentNode.parentNode.children[0].innerText;

    console.log("联系ID为"+index);

    currentIndex = GetIndexByID(index);
    console.log(currentIndex)

    getChatWindow(currentIndex)
    SetScrollDown();
    SetChatButton(obj,0);
    document.getElementById("Chat").className="";
    SendReadBySendID(index);
}; 
//设置聊天按钮
function SetChatButton(obj, count){
    obj.innerHTML="联系"+count;
    if(count==0){
        obj.className="";
    }
    else{
        obj.className="Red";
    }
}
//获取聊天窗口
function getChatWindow(Index) {
    document.getElementById("ChatMove").style.top=Index*(-750)+"px";
    //$("#ChatMove").css("top", Index*-750 + "px");
}
