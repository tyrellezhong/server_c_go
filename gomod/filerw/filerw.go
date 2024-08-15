package filerw

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type Website struct {
	Name      string `xml:"name,attr"`
	Url       string
	Course    []string
	TimeStamp int64
}

func FileRW() {

	// json 文件
	WriteJson()
	ReadJson()

	// xml文件
	WriteXml()
	ReadXml()

	// gob 文件
	WriteGob()
	ReadGob()

	// 文本文件
	WriteText()
	ReadText()

	// 数据结构-二进制文件
	WriteBinary()
	ReadBinary()
	RWBinary()

	// zip files
	WriteZip()
	ReadZip()
}

func WriteJson() {
	info := []Website{
		{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}, 11},
		{"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}, 22},
	}
	// 创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}

func ReadJson() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()
	out, _ := io.ReadAll(filePtr)
	fmt.Println("out : ", string(out))
	var info []Website
	// 创建json解码器
	filePtr.Seek(0, 0)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

func WriteXml() {

	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}, 11}
	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer f.Close()
	//序列化到文件中
	encoder := xml.NewEncoder(f)
	encoder.Indent("", "  ")
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}

func ReadXml() {
	//打开xml文件
	file, err := os.Open("./info.xml")
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return
	}
	defer file.Close()
	info := Website{}
	readInfo, _ := io.ReadAll(file)
	fmt.Println("readinfo:", string(readInfo))
	file.Seek(0, 0)
	//创建 xml 解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

func WriteGob() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}
	name := "demo.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0766)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}

func ReadGob() {
	var M map[string]string
	File, _ := os.Open("demo.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}

func WriteText() {
	//创建一个新文件，写入内容
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}

func ReadText() {
	//打开文件
	// file, err := os.Open("./output.txt")
	file, err := os.Open("./window_text.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

type binaryStruct struct {
	Value int8
}

func WriteBinary() {
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println("文件创建失败 ", err.Error())
		return
	}
	defer file.Close()

	for i := 1; i <= 10; i++ {
		info := binaryStruct{Value: int8(i)}

		// binary.Write(file, binary.LittleEndian, info)
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}

func ReadBinary() {
	readNextBytes := func(file *os.File, number int) []byte {
		bytes := make([]byte, number)
		_, err := file.Read(bytes)
		if err != nil {
			fmt.Println("解码失败", err)
		}
		return bytes
	}

	file, err := os.Open("output.bin")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()
	for i := 1; i <= 10; i++ {
		m := binaryStruct{}
		data := readNextBytes(file, 1)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

// 定义一个结构体
type Person struct {
	Name   [20]byte
	Age    int32
	Height float32
}

func RWBinary() {
	// 创建一个 Person 实例
	p1 := Person{
		Age:    30,
		Height: 5.9,
	}
	copy(p1.Name[:], "John Doe")

	// 打开一个文件用于写入
	file, err := os.Create("person.dat")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 将结构体写入文件
	err = binary.Write(file, binary.LittleEndian, &p1)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// 打开文件用于读取
	file, err = os.Open("person.dat")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个新的 Person 实例用于存储读取的数据
	var p2 Person

	// 从文件中读取数据到结构体
	err = binary.Read(file, binary.LittleEndian, &p2)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// 打印读取的数据
	fmt.Printf("Name: %s\n", bytes.Trim(p2.Name[:], "\x00"))
	fmt.Printf("Age: %d\n", p2.Age)
	fmt.Printf("Height: %.2f\n", p2.Height)
}

func WriteZip() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
}

func ReadZip() {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Close()
	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("文件名: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			fmt.Println(err.Error())
		}
		rc.Close()
	}
}

// 定义一个结构体来包装 map 数据
type TimeMap struct {
	Entries []Entry
}

// 定义一个结构体来表示 map 的每个条目
type Entry struct {
	Key   int32
	Value time.Time
}

func MarshalXmlTest() {
	timex, _ := time.Parse("2006-01-02 15:04:05", "2024-06-01 00:00:00")
	timeMap := map[int32]time.Time{
		1: timex,
		2: timex.Add(time.Hour * 24),
	}

	// 将 map 数据转换为 TimeMap 结构体
	var entries []Entry
	for k, v := range timeMap {
		entries = append(entries, Entry{Key: k, Value: v})
	}
	timeMapStruct := &TimeMap{Entries: entries}

	// 序列化为 XML
	// 传入对象可以是结构体指针或者结构体
	xmlData, err := xml.MarshalIndent(timeMapStruct, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}

	// 打印 XML 数据
	fmt.Println(string(xmlData))
}
