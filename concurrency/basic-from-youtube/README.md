# Wait group Example

[Full artical](https://www.youtube.com/watch?v=LvgVSSpwND8&ab_channel=JakeWright)

## Setup project
- go mod init updev.ha/labs/concurrency/basic

``` powershell
D:\> go run .
```


### Channel

```golang
    func main() {
		c := make(chan string)
		go count("Sheep", c)
		msg := <-c
		fmt.Println(msg)
    }

    func count(thing string, c chan string) {
        for i := 1; true; i++ {
            c <- thing
            time.Sleep(time.Millisecond * 500)
        }
    }
```

จาก code เราได้ทำการ implement count function โดยให้มีการรับ argument ที่เป็น Channel มาด้วย เนื่องจากเราต้องการจะใช้ channel เป็นตัวกลางในการสื่อสารกันระหว่าง goroutine main และ goroutine ที่ count function ทำงาน

เมื่อ count func ทำงานจะมีการ ``c <- thing`` เพื่อทำการ assign value ให้กับ channel โดยใน func main จะทำการ blocking เพื่อรอรับค่าจาก channel ``msg := <-c`` เพื่อทำงานต่อ หรือก็คือการ print ข้อความที่ได้จาก channel  ``fmt.Println(msg)`` 

``` powershell
PS D:\basic-from-youtube> go run channel.go
Sheep
```

จากผลลัพธ์เราจะเห็นว่ามีการแสดงออกมาเพียงแค่ค่าเดียว เนื่องมาจาก

ทันทีที่มีการรับค่ามาจาก channel เรียบร้อยในส่วนของ main goroutine ก็จะทำการ print ค่าที่ได้ก็ถือว่าเป็นการจบการทำงาน นั้นก็หมายความว่า program ทำงานเรียบร้อยนั้นเอง ดังนั้น other goroutine ที่เป็ยการทำงานภายใต้ main goroutine ก็จะหยุดการทำงานไปด้วยนั้นเอง

```golang
    func main() {
		c := make(chan string)
		go count("Sheep", c)
        for {
            msg := <-c
            fmt.Println(msg)
        }
    }

    func count(thing string, c chan string) {
        for i := 1; i <= 5; i++ {
            c <- thing
            time.Sleep(time.Millisecond * 500)
        }
    }
```

แต่ถ้าเราต้องการให้มันแสดงค่าทั้งหมดละเราจะทำไง ?

เราจึงเริ่มทำการกำหนดจำนวนข้อมูลที่จะทำการ count ให้เป็น 5 ครั้ง และในส่วน func main เราได้ทำการ เพิ่ม infinity for loop เข้าไปเพื่อให้ทำการ loop มาที่ ``msg := <-c`` เพื่อทำการรอรับค่าจาก channel ที่จะมีการส่งมาใหม่ และผลที่ได้ก็คือ

```powershell
    PS D:\git-myself\updev-golang-noted\concurrency\basic-from-youtube> go run channel.go
    Sheep
    Sheep
    Sheep
    Sheep
    Sheep
    fatal error: all goroutines are asleep - deadlock!

    goroutine 1 [chan receive]:
```

มันมีการแสดงค่าเป็นจำนวน 5 ครั้งจริงๆ แต่จะเจอ error ``fatal error: all goroutines are asleep - deadlock!`` ซึ่งเหมือนจะบอกว่ามีการรอรับค่าแต่ไม่มีการส่งค่ามาแล้วนั้นเอง หรือ ตามความเข้าใจของผมก็คือ Channel ที่มีการใช้งานนั้นมันไม่ได้มีการรับค่าใหม่มาอีกแล้ว

โดยวิธีแก้เราแค่เพียงเพิ่ม ``Close()`` func เอาไว้ที่ func count เพื่อเป็นการปิดตัว Chanel เมื่อไม่มีการส่งค่ออีกแล้ว

และเพิ่มในส่วนการรับ open จาก chanel เพื่อทำการตรวจสอบว่า channel นี้ยัง open อยู่หรือเปล่าหรือก็คือถ้ามีการ Close แล้วให้ทำการ break infinity for loop นั้นเอง ซึ่ง้ราจะได้ผลลัพธ์ตามที่เราต้องการ

```golang
    func main() {
		c := make(chan string)
		go count("Sheep", c)
        for {
            msg, open := <-c
            if !open {
                break
            }
            fmt.Println(msg)
        }
    }

    func count(thing string, c chan string) {
        for i := 1; i <= 5; i++ {
            c <- thing
            time.Sleep(time.Millisecond * 500)
        }

        close(c)
    }
```

นอกจากวิธีนี้เราสามารถใช้ for range channel เพื่อทำหน้าที่ในการ blocking loop เพื่อรอการรับค่าจาก Channel หรือก็คือถ้ามีการ Close channel การจะ break loop ให้อัตโนมัตินั้นเอง โดยเราจะได้ code ตามนี้

```golang
func main() {

	c := make(chan string)
	go count("Sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}


```



