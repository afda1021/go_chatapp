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
/* 余白クリックでメニュー表示終了 */
$(function(){
    $('.msgs, .reply').click(function(){
        $('.menu, .my-menu').addClass("menu-hidden");
    });
    $('.msg').on('click', function(e){
        e.stopPropagation();
    });
});

/* メッセージ送信取消 */
function remove_msg(id){
    let room_id = $("#room_id").val();
    let result = confirm('送信を取り消しますか？');
    if (result) {
        // location.href = '/delete_msg?id='+ room_id +'&msg_id=' + msg_id;
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
        $("#send").removeClass("menu-hidden");
        $("#reply").addClass("menu-hidden");
    }
    $("#send").toggleClass("menu-hidden");
    $("#reply").toggleClass("menu-hidden");
    /* 返信ボタン押下時 */
    document.getElementById("reply").onclick = function(){
        replyButton(msg_id);
    };
}
/* 余白クリックでリプライ表示終了 */
$(function(){
    $('.msgs').click(function(){
        $('.reply').addClass("menu-hidden");
        $("#send").removeClass("menu-hidden");
        $("#reply").addClass("menu-hidden");
    });
    $('.reply-btn, .my-reply-btn, .reply, .menu, .my-menu').on('click', function(e){
        console.log("ok");
        e.stopPropagation();
    });
});