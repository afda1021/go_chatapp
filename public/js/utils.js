function isNotEmpty(){
    // 投稿文が空、空行、スペースのときは投稿不可
    if (msg.val() == "" || !msg.val().match(/\S/g)) {
        alert('空欄では投稿できません');
        return false;
    }
};