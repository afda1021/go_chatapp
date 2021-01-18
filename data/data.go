package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
)

func Encrypt(password string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	return cryptext
}

//ランダムなuuidを作成
func CreateUUID() (uuid string) {
	u := new([16]byte)        //uは要素数16のバイト配列のポインタ型  &[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	_, err := rand.Read(u[:]) //u[:]はランダムな16個の数字 16 <nil>
	//fmt.Println(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func (session *Session) DeleteByUUID() {
	db := DbInit()
	defer db.Close()
	db.Query("DELETE FROM sessions WHERE uuid = ?", session.Uuid)
}
