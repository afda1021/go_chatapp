$(function(){
    let socket = null;
    let msg = $("#chatbox textarea");  // 入力されたメッセージ
    let name = $("#name")              // 送信者の名前
    let messages = $("#messages");     // message表示スペース
    /*クエリを取得*/
    let query = (new URL(document.location)).searchParams;
    let id = query.get('id');

    /* socketの開設 */
    socket = new WebSocket("ws://localhost:8080/ws?id="+id);

    /* 送信ボタン押下時 */
    document.getElementById("send").onclick = function(){
        /* socketにデータを送る */
        socket.send(JSON.stringify({
            "Name": name.val(),      // 送信者の名前
            "RoomId": id,      // メッセージ送信者のルームid
            "Text": msg.val()  // 入力されたメッセージ
        }));
        msg.val("");
        return false;
    };

    /* メッセージ受信時 */
    socket.onmessage = function(e){
        let msg = eval("("+e.data+")");                    
        messages.append(
            $("<p>").append(msg.Name + "さん：" + msg.Text)
        );
    }
});