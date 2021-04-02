package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [4096]byte
}

func (this *Transfer)ReadPkg() (mes common.Message, err error) {
	//buf := make([]byte, 8192)
	//fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("read pkg head err=",err)
		return
	}

	//将读取到的信息长度，转换成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen读取相应的数据长度
	//fmt.Println("pkgLen=",pkgLen)
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	//将读取到的信息反序列化-->Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer)WritePkg(data []byte) (err error) {

	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)

	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=",err)
		return err
	}

	//发送消息本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return err
	}
	return err
}
