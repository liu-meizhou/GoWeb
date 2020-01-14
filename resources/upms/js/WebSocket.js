var ws;

function wsConnect(){
    ws = new WebSocket("ws://193.112.207.30:8888/ws");
    ws.onopen = function (evt) {
        console.log(getCookie("MyCookie"))
        ws.send(getCookie("MyCookie"));
    }
    ws.onclose = function (evt) {
        console.log("onclose")
        ws = null;
    }
    ws.onmessage = function (evt) {
        console.log("服务器发来: " + evt.data);
        var messageData = GetMessageData()
        var data=JSON.parse(evt.data,messageData)
        var messageData = htmlGetHisMessage(data["SendTime"],data["Type"],data["Context"])
        console.log(messageData)
        IndexID=GetIndexByID(data["SendID"])
        document.getElementById("Chat"+IndexID).getElementsByClassName("ChatMessage")[0].innerHTML+=messageData;
        if(IndexID!=currentIndex && document.getElementById("Chat").className!=""){
            document.getElementsByClassName("NewMessage")[0].firstElementChild.innerText+=1;
            var chat = document.getElementById("user"+data["SendID"]).lastElementChild.lastElementChild;
            chat.className="Red";
            chat.innerText="联系1";
        }
    }
    ws.onerror = function (evt) {
        print("ERROR: " + evt.data);
    }
};

function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
 
    if(arr=document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
} 