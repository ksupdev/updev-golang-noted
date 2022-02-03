# Wait group Example

[Full artical](https://www.youtube.com/watch?v=LvgVSSpwND8&ab_channel=JakeWright)

## Setup project
- go mod init updev.ha/labs/concurrency/basic

``` powershell
D:\> go run .
```


## Channel

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

## Select Channel

```golang

func main() {
	c := make(chan string)

	c <- "Hello"

	msg := <-c

	fmt.Println(msg)
}

===== output =====
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan send]:
main.main()
```

ที่ error นั้นเกิดจาก code นั้นทำงานอยู่ใน goroutine เดียวกันดังนั้นจะทำงานแบบเป็น sequen พอเจอคำสั่ง ``c <- "Hello"`` ก็จะมีการพยายามที่จะส่งค่า แต่ปัญหาก็คือ code ในส่วนที่จะรับค่าจาก channel ยังไม่ได้มีการทำงาน `หรืออีกนัยหนึ่งเราอาจจะบอกได้ว่าการที่จะใช้งาน channel นั้นคือท่อที่จะต้องมีทางเข้าและทางออกเสมอ ซึ่งการทำงานที่ main goroutine นั้นมันจะเป็นการทำงานที่เป็น sequence ดังนั้นมันจะเจอทางเข้าของข้อมูลและพอข้อมูลเข้าไปก็ไม่เจอทางออกเพราะในท่อนั้นมันไม่มีที่เก็บอะไร จึงทำให้ error ` ดังนั้นเราสามารถทำการแก้ code ได้ 2 วิธี

วิธีที่ 1 ทำการสร้าง goroutine เพื่อให้มีการทำงานที่แยกออกจาก main goroutine
```golang
func main() {
	c := make(chan string)

	go func() {
		c <- "Hello"
	}()

	msg := <-c

	fmt.Println(msg)
}

===== output =====
Hello
```


วิธีที่ 2 คือการกำหนด buffer ให้กับ channel นั้นก็คือการกำหนด channel ให้ทำการ buffer ค่าเอาไว้ก่อนจนกว่าจะมีการ receive ไปนันเอง
```golang
func main() {
	c := make(chan string, 1)

	c <- "Hello"
	msg := <-c

	fmt.Println(msg)
}

===== output =====
Hello
```

ข้อควรระวังในการใช้ buffer ก็คือ ถ้าเรากำหนดเป็น 1 หมายความว่ามันจะเก็บไว้ให้ได้แค่ 1 จนกว่าจะมีการดึงข้อมูลออกไปแล้วมีพื้นที่ว่างก็จะสามารถรับค่าใหม่ได้

```golang
func main() {
	c := make(chan string, 1)

	c <- "Hello"
	msg := <-c

	fmt.Println(msg)

	c <- "Hello 1"

	msg = <-c

	fmt.Println(msg)

	c <- "Hello 2"
	msg = <-c

	fmt.Println(msg)

}

===== output =====
Hello
Hello 1
Hello 2
```

จาก code มันจะเป็นการ assign แล้วก็ recieve ออกซึ่งจะสามารถทำงานได้อย่างปกติ

```golang
func main() {
	c := make(chan string, 1)

	c <- "Hello"
	msg := <-c

	fmt.Println(msg)

	c <- "Hello 1"
	c <- "Hello 2"
	msg = <-c
	fmt.Println(msg)


	msg = <-c
	fmt.Println(msg)

}

===== output =====
Hello
fatal error: all goroutines are asleep - deadlock!
```

แต่แค่เราทำการ assign Hello1 และ Hello2 โดยที่เรายังไม่ได้ทำการดึงค่าออกจาก channel หรือก็คือ Buffer ยังเต็มอยู่นั้นเองก็จะทำให้ Error ทันที

ซึ่งเราอาจจะทำการแก้ได้ โดยแค่ทำการเพิ่ม Buffer ให้กับ Channel ก็จะสามารถแก้ปัญหานั้นได้ทันที

```golang
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two Seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// Sequence running
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
===== output =====
Every 500ms
Every two Seconds
Every 500ms
Every two Seconds
Every 500ms
Every two Seconds
```

จาก code สิ่งที่ควรจะเป็นมันน่าจะเป็นการแสดงผลตามเวลา นั้นก็คือมีการ print every 500 MS ไปเลื่อยจนถึง 2 Second ถึงจะแสดง แต่การแสดงผลนั้นเป็นการแสดงผลแบบสลับกันแสดง

ที่เป็นเช่นนี้เนื่องมาจาก infinity loop มีการแสดงผลหลังจาก ``<-c1`` จะมีการดึงค่าปกติ แต่มันจะถูก Blocking ``<-c2`` คือต้องรอ 2 Second ถึงจะมีค่าส่งมาเพราะเหตุนี้จึงได้ผลลัพธ์การแสดงที่ดูสลับกัน

โดยการแก้นั้นเราจะมีการ implement `select` เพื่อเป็นตัวค่อยดูว่าค่าที่ส่งมานั้นมาจาก channel ไหนและจะไปทำงานของส่วนไหนก่อน หรือก็คือการจัดลำดับการทำงานตามที่มีการส่งข้อมูลมาจากแต่ละ channel นั้นเอง

```golang
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two Seconds"
			time.Sleep(time.Second * 4)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

===== output =====
Every two Seconds
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every 500ms
Every two Seconds
```





