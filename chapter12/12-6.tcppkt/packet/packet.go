package packet

import (
	"bytes"
	"encoding/binary"
	"io"
)

// 二进制封包格式
type Packet struct {
	Size uint16
	Body []byte
}

func WritePacket(dataWriter io.Writer, data []byte) error {
	var buffer bytes.Buffer

	// 将size写入缓冲
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}
	// 写入包体数据
	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得写入的数据
	out := buffer.Bytes()

	// 写入dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}

	return nil

}

func ReadPacket(dataReader io.Reader) (pkt Packet, err error) {
	// size 为unit16类型, 占2个字节
	var sizeBuffer = make([]byte, 2)
	// 持续读取Size直到读到为止
	_, err = io.ReadFull(dataReader, sizeBuffer)
	if err != nil {
		return
	}
	// 使用bytes.Reader读取size buffer中的数据
	sizeReader := bytes.NewReader(sizeBuffer)

	// 读取小段的uint16作为size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size)
	if err != nil {
		return
	}
	// 分配包体的大小
	pkt.Body = make([]byte, pkt.Size)
	// 读取包体的数据
	_, err = io.ReadFull(dataReader, pkt.Body)

	return
}
