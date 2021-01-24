$(function(){
    let socket = null;
    let msg = $("#chatbox textarea");  // 入力されたメッセージ
    let name = $("#name");             // 送信者の名前(自分の名前)
    // let msgName = $("#msg-name");   // 送信者の名前
    let messages = $("#msgs");         // message表示スペース
    let textarea = $('#chatbox').children("textarea");

    /*クエリを取得*/
    let query = (new URL(document.location)).searchParams;
    let roomId = query.get('id');

    /* socketの開設 */
    socket = new WebSocket("ws://localhost:8080/ws?id="+roomId);

    /* 送信ボタン押下時 */
    document.getElementById("send").onclick = function(){
        // 投稿文が空、空行、スペースのときは投稿不可
        if (msg.val() == "" || !msg.val().match(/\S/g)) {
            alert('空欄では投稿できません');
            return false;
        }
        /* socketにデータを送る */
        const newDate = new Date();
        let year = newDate.getFullYear();
        let month = newDate.getMonth() + 1;
        let date = newDate.getDate();
        if (month < 10) {
            month = "0" + month;
        }
        if (date < 10) {
            date = "0" + date;
        }
        let dateTime = year + "/" + month + "/" + date;

        let hour = newDate.getHours();
        let minutes = newDate.getMinutes();
        if (hour < 10) {
            hour = "0" + hour;
        }
        if (minutes < 10) {
            minutes = "0" + minutes;
        }
        let time = hour + ":" + minutes;
        socket.send(JSON.stringify({
            "Name": name.val(),      // 送信者の名前
            "RoomId": roomId,      // メッセージ送信者のルームid
            "Text": msg.val(), // 入力されたメッセージ
            "date": dateTime,
            "time": time
        }));
        msg.val("");
        /* textareaの高さを元に戻す */
        let lineHeight = parseInt(textarea.css('lineHeight'));
        textarea.height(lineHeight);
        return false;
    };

    /* メッセージ受信時 */
    let id = 0;
    socket.onmessage = function(e){
        ++id
        let msg = eval("("+e.data+")");
        if (name.val() == msg.Name) { //自分の名前と投稿者名が一致する場合
            messages.append(`<div id=${id+"msgbox"} class="my-msgbox"></div>`);
            $(`#${id+"msgbox"}`).append(`<p>` + msg.Name + "さん -" + msg.Time + `</p>`);         
            $(`#${id+"msgbox"}`).append(`<div id=${id+"msg"} class="msg"></div>`);
            $(`#${id+"msg"}`).append(`<p>` + msg.Text + `</p>`);
        }else{
            messages.append(`<div id=${id+"msgbox"} class="msgbox"></div>`);
            $(`#${id+"msgbox"}`).append(`<p>` + msg.Name + "さん -" + msg.Time + `</p>`);         
            $(`#${id+"msgbox"}`).append(`<div id=${id+"msg"} class="msg"></div>`);
            $(`#${id+"msg"}`).append(`<p>` + msg.Text + `</p>`);
        }
    }
    /* 自分と相手でメッセージ表示を変える */
    let msgsData = document.querySelectorAll(".msg-name");

    for (let i=0; i<msgsData.length; i++){
        if (name.val() == msgsData[i].value) {
            let msgId = msgsData[i].id; //名前と投稿者名が一致するid
            let divId = document.getElementById("div-"+msgId);
            divId.classList.remove("msgbox");
            divId.classList.add("my-msgbox");
        }
    }
    /* textareaの高さ自動調整 */
    textarea.attr("rows", 1).on("input", function() {
        $(this).height(0).innerHeight(this.scrollHeight);
    });
    /* 日付の表示 */
    let dates = document.querySelectorAll(".date");
    let dt = "2006/01/23";
    for (let i=0; i<dates.length; i++){
        if (dt != dates[i].value) {
            let dtId = dates[i].id;
            document.getElementById("dt-" + dtId).classList.remove("dt-inv");
            document.getElementById("dt-" + dtId).classList.add("dt-vis");
        }
        dt = dates[i].value;
    }
});