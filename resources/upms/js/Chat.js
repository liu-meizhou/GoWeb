//当前聊天框总量
var indexID=0;
//当前打开的聊天框下标
var currentIndex=0;



//获取当前时间
function getNowFormatDate() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1<10? "0"+(date.getMonth() + 1):date.getMonth() + 1;
    var strDate = date.getDate()<10? "0" + date.getDate():date.getDate();
    var currentdate = date.getFullYear() + seperator1  + month  + seperator1  + strDate
            + " "  + date.getHours()  + seperator2  + date.getMinutes()
            + seperator2 + date.getSeconds();
    return currentdate;
}
//获取一条信息 type=1 span  type=2 img
var getmessage = function(time,type,message,who) {    
    var cont="";
    if(type==1){
        cont="<span>"+message+"</span>"
    }else if(type==2){
        cont="<img src='"+message+"' class='SendImage'></img>"
    }
   return "<li class='"+ who +"'>"+
            time +":"+cont+
        "</li>";
}
//获取一条我方信息
var getMyMessage = function(type,message){return getmessage(getNowFormatDate(), type,message,"textRight"); }

//获取一条对方信息
var getHisMessage = function(type,message){return getmessage(getNowFormatDate(), type,message,"textLeft"); }

//发送信息
function SendMessage(ID,message){
    var $ChatMessage;
    if (! ID instanceof Object){
        $ChatMessage=$(document.getElementById("Chat"+GetIndexByID(ID)).getElementsByClassName("ChatMessage")[0]);
    }
    else{
        $ChatMessage=$(ID.getElementsByClassName("ChatMessage")[0]);
    }
    $ChatMessage.append(message);
}

//发送消息按钮
function SendOnButton(obj){
    if (ws==null){
        console.log("未连接服务器")
        return false
    }
    console.log("已经连接服务器xxxx")
    var text=obj.parentNode.firstElementChild;
    var jsondata = GetMessageJson(1,text.value)
    console.log(jsondata)
    ws.send(jsondata)
    //ws.send(jsondata)
    SendMessage(obj.parentNode.parentElement,getMyMessage(1,text.value));
    text.value="";
    SetScrollDown();
}
//按回车发送信息
function keypress(obj,event) {
    if (event.keyCode == "13") {
        if(ws==null){
            console.log("未连接服务器")
            return false
        }
        console.log("已经连接服务器xxxx")
        ws.send(GetMessageJson(1,obj.value))
        // ws.send(data);
        SendMessage(obj.parentNode.parentElement,getMyMessage(1,obj.value))
        obj.value="";
    }
}
//关闭聊天
function CloseChat(){
    document.getElementById("Chat").className="hidden";
}
//获取聊天信息
function GetMessages(fn){
    $.ajax({
        type: "POST",
        url: "/Message/GetEveryNoRead",
        success: function(datas){
            console.log(datas)
            var messages=JSON.parse(datas);
            console.log(messages);
            fn(messages);
        },
        error: function (data) {
            console.log(data);
        }
    });
}
//设置滚动条到最下方
function SetScrollDown(){
    $(".ChatMessage").scrollTop($(".ChatMessage")[0].scrollHeight);
}
//发送已读信息
function SendReadBySendID(id){
    $.ajax({
        type: "POST",
        url: "/Message/SendReadBySendID",
        data: {SendID:id},
        success: function(datas){
            console.log(datas);
        },
        error: function (data) {
            console.log(data);
        }
    });
}
//设置聊天窗口的name
function SetChatWindowName(index,name){
    document.getElementById("Chat"+index).firstElementChild.innerHTML=name;
}
//根据ID获取对话框下标index
function GetIndexByID(ID){
    object=document.getElementById("user"+ID);
    objectfather=object.parentElement.children;

    // console.log($(object).index())

    var i=0;
    while(true){
        if(objectfather[i]==object){
            return i;
        }
        i++;
    }


    // return $(object).index();

    // var ChatID = 0;
    // while((object = object.previousSibling) != null) ChatID++;
    // return ChatID;
}


function GetMessageJson(Type,Context){
    var data={
        SendID: 0,
        ReceiveID: parseInt(document.getElementById("Users").children[currentIndex].firstElementChild.innerHTML),
        Type: Type,
        Context: Context,
        SendTime: "",
        IsRead: false
    }
    return JSON.stringify(data)
}
//构建一个message类
function GetMessageData() {
    return {
        SendID: 0,
        ReceiveID: 0,
        Type: 1,
        Context: "",
        SendTime: "",
        IsRead: false
    }    
}

