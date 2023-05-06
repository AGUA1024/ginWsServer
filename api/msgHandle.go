package api

// 升级get请求为webSocket协议
ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
if err != nil {
return
}
defer ws.Close() //返回前关闭

for {
// 读取ws中的数据
mt, _, err := ws.ReadMessage()
if err != nil {
break
}

remas := &act1.Student{
Name:    "tzy",
Age:     18,
Address: "127.0.0.1",
Cn:      1,
}

marsh, _ := proto.Marshal(remas)

fmt.Println("protobuf:")
fmt.Println(marsh)

ms := &act1.Student{}
proto.Unmarshal(marsh, ms)

fmt.Println("name:", ms.Name)

// 写入ws数据
err = ws.WriteMessage(mt, marsh)
if err != nil {
break
}
}