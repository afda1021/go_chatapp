let socket = null;
let msg = $("#chatbox textarea");  // 入力されたメッセージ
let name = $("#name");             // 送信者の名前(自分の名前)
let messages = $("#msgs");         // message表示スペース
let textarea = $('#chatbox').children("textarea");
let userList = $("#user_list");  // 入室中のユーザー一覧

/*クエリを取得*/
let query = (new URL(document.location)).searchParams;
let roomId = query.get('id');

/* socketの開設 */
console.log(name.val());
socket = new WebSocket("ws://" + location.host + "/ws?id="+roomId+"&name="+name.val());

/* メッセージ受信時 */
socket.onmessage = function(e){
    let msg = eval("("+e.data+")");
    if (msg.Text == null){ //入室中のユーザー一覧を取得
        userList.children("p").remove();
        userList.append(`<p>` + name.val() + " (あなた)" + `</p>`);
        for (let i=0; i<msg.length; i++){
            if (msg[i].Name != name.val() && msg[i].RoomId == roomId){
                console.log(msg[i],"ok");
                userList.append(`<p>` + msg[i].Name + `</p>`);
            }
        }
    }
    // console.log(msg);
    let id = msg.Id;
    if (msg.Type == 'publish'){
        if (name.val() == msg.Name) { //自分の名前と投稿者名が一致する場合
            messages.append(`<div id=${"msg-outbox"+id} class="msg-outbox"></div>`);
            $(`#${"msg-outbox"+id}`).append(`<div id=${"msgbox"+id} class="my-msgbox"></div>`);
            $(`#${"msgbox"+id}`).append(`<p>` + msg.Name + `<small> -` + msg.Time + `</small></p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"msg"+id} class="msg" onclick="menu('${id}')"></div>`);
            $(`#${"msg"+id}`).append(`<p>` + msg.Text + `</p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"menu"+id} name="menu" class="my-menu menu-hidden"></div>`);
            $(`#${"menu"+id}`).append(`<input type="button" class="menu-list" value="送信取消" onclick="remove_msg('${id}');">`);
            $(`#${"menu"+id}`).append(`<input type="button" class="menu-list" value="リプライ" onclick="reply('${id}');">`);
            messages.append(`<div id=${"reply"+id} class="reply menu-hidden">`);
            //一番下まで自動スクロール
            $('#msgs').animate({scrollTop: $('#msgs')[0].scrollHeight}, 'fast');
        }else{
            messages.append(`<div id=${"msg-outbox"+id} class="msg-outbox"></div>`);
            $(`#${"msg-outbox"+id}`).append(`<div id=${"msgbox"+id} class="msgbox"></div>`);
            $(`#${"msgbox"+id}`).append(`<p>` + msg.Name + `<small> -` + msg.Time + `</small></p>`);     
            $(`#${"msgbox"+id}`).append(`<div id=${"msg"+id} class="msg" onclick="menu('${id}')"></div>`);
            $(`#${"msg"+id}`).append(`<p>` + msg.Text + `</p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"menu"+id} name="menu" class="menu menu-hidden"></div>`);
            $(`#${"menu"+id}`).append(`<input type="button" class="menu-list" value="リプライ" onclick="reply('${id}');">`);
            messages.append(`<div id=${"reply"+id} class="reply menu-hidden">`);
        }
    }else if (msg.Type == 'remove'){
        if (document.getElementById("msgbox"+id) == null){ //DB表示の場合
            location.reload() //ページを更新
        }else{ //リアルタイム表示の場合
            $(`#${"msgbox"+id}`).remove();
        }
    }else if (msg.Type == 'reply'){
        let reply = $(`#${"reply"+msg.ReplyId}`); // reply表示スペース
        if (name.val() == msg.Name) { //自分の名前と投稿者名が一致する場合
            reply.append(`<div id=${"msg-outbox"+id} class="msg-outbox"></div>`);
            $(`#${"msg-outbox"+id}`).append(`<div id=${"msgbox"+id} class="my-msgbox"></div>`);
            $(`#${"msgbox"+id}`).append(`<p>` + msg.Name + `<small> -` + msg.Time + `</small></p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"msg"+id} class="msg" onclick="menu('${id}')"></div>`);
            $(`#${"msg"+id}`).append(`<p>` + msg.Text + `</p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"menu"+id} name="menu" class="my-menu menu-hidden"></div>`);
            $(`#${"menu"+id}`).append(`<input type="button" class="menu-list" value="送信取消" onclick="remove_msg('${id}');">`);
        }else{
            reply.append(`<div id=${"msg-outbox"+id} class="msg-outbox"></div>`);
            $(`#${"msg-outbox"+id}`).append(`<div id=${"msgbox"+id} class="msgbox"></div>`);
            $(`#${"msgbox"+id}`).append(`<p>` + msg.Name + `<small> -` + msg.Time + `</small></p>`);     
            $(`#${"msgbox"+id}`).append(`<div id=${"msg"+id} class="msg" onclick="menu('${id}')"></div>`);
            $(`#${"msg"+id}`).append(`<p>` + msg.Text + `</p>`);
            $(`#${"msgbox"+id}`).append(`<div id=${"menu"+id} name="menu" class="menu menu-hidden"></div>`);
        }
    }
}

/* 自分と相手でメッセージ表示を変える */
let msgsData = document.querySelectorAll(".msg-name");
for (let i=0; i<msgsData.length; i++){
    if (name.val() == msgsData[i].value) {
        let msgId = msgsData[i].id; //名前と投稿者名が一致するid
        let divId = document.getElementById("msgbox"+msgId);
        let menuId = document.getElementById("menu"+msgId);
        divId.classList.remove("msgbox");
        divId.classList.add("my-msgbox");
        menuId.classList.remove("menu");
        menuId.classList.add("my-menu");
        let replyBtn = document.getElementById("reply-btn"+msgId);
        if (replyBtn !=null) { //返信ボタンが存在すれば自分と相手で表示を変える
            replyBtn.classList.remove("reply-btn");
            replyBtn.classList.add("my-reply-btn");
        }
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