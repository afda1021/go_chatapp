/* メニューボタン */
function menu(menu_id){
    let menus = document.getElementsByName("menu");
    for (let i=0; i<menus.length; i++){
        if (menus[i].id == menu_id){
            $(`#${menu_id}`).toggleClass('menu-hidden');
        }else{
            menus[i].classList.add("menu-hidden");
        }
    }
}

/* メッセージ送信取消 */
function delete_msg(msg_id){
    let room_id = $("#room_id").val();
    let result = confirm('送信を取り消しますか？');
    if (result) {
        location.href = '/delete_msg?id='+ room_id +'&msg_id=' + msg_id;
    }
}