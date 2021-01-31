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
        "Date": dateTime,
        "Time": time,
        "Type": 'publish'
    }));
    msg.val("");
    /* textareaの高さを元に戻す */
    let lineHeight = parseInt(textarea.css('lineHeight'));
    textarea.height(lineHeight);
    return false;
};