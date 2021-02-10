/* メニューボタン */
function menu(id){
    /* 選択したメニューのみ表示 */
    let menu_id = "menu" + id;
    $(`#${menu_id}`).toggleClass('menu-hidden');
    $('.menu, .my-menu').not(`#${menu_id}`).addClass("menu-hidden");
    /* DB表示において相手の送信取消不可 */
    let msg_name = $(`#${"msgbox"+id}`).children("input");
    if (msg_name.val() != null) { //DB表示の場合
        if (name.val() != msg_name.val()){
            let remove_msg = document.getElementById("remove_msg"+id);
            remove_msg.classList.remove("menu-list");
            remove_msg.classList.add("menu-hidden");
        }
    }
}

/* メッセージ送信取消 */
function remove_msg(id){
    let room_id = $("#room_id").val();
    let result = confirm('送信を取り消しますか？');
    if (result) {
        socket.send(JSON.stringify({
            "Id": Number(id),
            "RoomId": room_id,
            "Type": 'remove'
        }));
    }
}

/* リプライ */
function reply(msg_id){
    $(`#${"reply"+msg_id}`).toggleClass("menu-hidden");
    $('.reply').not(`#${"reply"+msg_id}`).addClass("menu-hidden");
    if ($(`#${"reply"+msg_id}`).attr('class').split(' ')[1] == null){
        $("#send").addClass("menu-hidden");
        $("#reply").removeClass("menu-hidden");
    }else{
        $("#send").removeClass("menu-hidden");
        $("#reply").addClass("menu-hidden");
    }
    /* 返信ボタン押下時 */
    document.getElementById("reply").onclick = function(){
        replyButton(msg_id);
    };
}
/* 余白クリックでイベント */
$(function(){
    $('.msgs').click(function(){
        // メニュー表示終了
        $('.menu, .my-menu').addClass("menu-hidden");
        // リプライ表示終了
        $('.reply').addClass("menu-hidden");
        $("#send").removeClass("menu-hidden");
        $("#reply").addClass("menu-hidden");
    });
    //解除、動的な要素にもイベントを適用
    $('.msgs').on('click', '.msg, .reply, .menu, .my-menu, .reply-btn, .my-reply-btn', function(e){
        e.stopPropagation();
    });
});