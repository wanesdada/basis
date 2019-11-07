package jwt

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestEasyToken_GetToken(t *testing.T) {
	token, _ := EasyToken{
		Username: strconv.Itoa(10),
		Expires:  time.Now().Unix() + 3600*24*30, //Segundos

	}.GetToken()
	fmt.Println(token)

	b, s, e := EasyToken{}.ValidateToken(token)
	fmt.Println(b,s,e)
	var tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIzMSIsIm5iZiI6MTU3MzAwODM2Nn0.fGWmTyY3uwPNJkiX_1U6NAHphc-dQ9MYAImZSQr_yV-7RXjrKG7h7MDbp4zvf5fVkVDjIY5C1sb4NpEKXok0tEOLMyidkHxi1suObJOX2Jdqm0H2tOLVdxLRTxoZDqiQdL22Q4jubyOarT71ckaEaA1jgPFks9Nju_0Eq-btScUKS_rLGWk3cgycrCd4XiJRjJv6QsZrWlghnUsnogWsvRu4PPIhGONK_N3W_mgsf5xTK8cJPJivNByObaUX3BCD_v5L7Hr3hUnNH5gDQT8ID4eaVRz82h8Izd_Y-zboF7zSDFAAzLPjOwxb_0T5c5CXXbRXWdEK87CEqlCnVSXXsA`
	b1, s1, e1 := EasyToken{}.ValidateToken(tokenStr)
	fmt.Println(b1,s1,e1)

}

