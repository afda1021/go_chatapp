function replyButton(msg_id){
    if (msg.val() == "" || !msg.val().match(/\S/g)) {
        alert('空欄では投稿できません');
        return false;
    }
    /* socketにデータを送る */
    socket.send(JSON.stringify({
        "Name": name.val(),      // 送信者の名前
        "RoomId": roomId,      // メッセージ送信者のルームid
        "Text": msg.val(), // 入力されたメッセージ
        "ReplyId": Number(msg_id),
        "Type": 'reply'
    }));
    msg.val("");
    /* textareaの高さを元に戻す */
    let lineHeight = parseInt(textarea.css('lineHeight'));
    textarea.height(lineHeight);
    return false;
};