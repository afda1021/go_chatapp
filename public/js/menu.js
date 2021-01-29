/* メニューボタン */
function menu(id){
    let menu_id = "menu" + id;
    let menus = document.getElementsByName("menu");
    for (let i=0; i<menus.length; i++){
        if (menus[i].id == menu_id){
            $(`#${menu_id}`).toggleClass('menu-hidden');
        }else{
            menus[i].classList.add("menu-hidden");
        }
    }
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
function remove_msg(msg_id){
    let room_id = $("#room_id").val();
    let result = confirm('送信を取り消しますか？');
    if (result) {
        // location.href = '/delete_msg?id='+ room_id +'&msg_id=' + msg_id;
        socket.send(JSON.stringify({
            "Id": Number(msg_id),
            "RoomId": room_id,
            "Type": 'remove'
        }));
    }
}