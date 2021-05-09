# Goroutine call other goroutine

> go mod init updev/goroutine-ep2

## problem
ผมต้องการที่จะทดสอบการทำงานของ goroutine กรณีที่มีการเรียนก goroutine ภายใต้ goroutine โดยตอนแรกผมเข้าใจว่ามันจะเป็นการทำงานโดยมีหลักการคล้ายๆกับ การที่ function ที่เป็น parent มีการเรียก  child function นั้นก็คือ จะต้องทำงานที่ child function ก่อนแล้วค่อยทำงานที่ parent function ต่อ

แต่จะที่ได้ลองเล่นเหมือนจะไม่เป็นที่คิดเลย

## Get started
เพื่อให้ง่ายต่อการทดสอบเราจะทดสอบทุกๆอย่างใน package main ซึ่งจะประกอบไปด้วย
- main() : โดยจะมีการเรียกใช้ func1(),func2() พร้อมกับทำการสร้าง anonymus func ซึ่งจะมีการทำงานเหมือนกับ func1,func2 แต่จะเพิ่มในส่วนเรียกใช้ func3 ภายในตัวมันเอง
- func1() func2() func3() : โดยแต่ละ func จะมีการทำงานเหมือนกันหมด นั้นก็คือ print ข้อมูล "123" โดย for loop จาก range "123" เพื่อแสดงข้อมูล 1,2,3 ตามลำดับ และมีการ set ``time.Sleep(10 * time.Millisecond)``

```powershell
karoon@Nuttakorns-MacBook-Pro goroutine-under-goroutine % go run main.go
---- Start ---- [2021-05-09 19:10:11.732069 +0700 +07 m=+0.000105711] 
---- Start program ---- [240.408µs] 
func3 output = 1 [11.619552ms]
Anonymus-1 output = 1 [11.661802ms]
func1 output = 1 [11.708334ms]
func2 output = 1 [11.632689ms]
func2 output = 2 [22.013328ms]
Anonymus-1 output = 2 [22.317375ms]
func1 output = 2 [22.712828ms]
func3 output = 2 [22.727141ms]
func2 output = 3 [33.184856ms]
Anonymus-1 output = 3 [33.385516ms]
func3 output = 3 [33.944081ms]
func1 output = 3 [33.956677ms]
---- End progra ---- [1.001740595s] 
---- End ---- [2021-05-09 19:10:11.732069 +0700 +07 m=+0.000105711]
```

จากการทดสอบจะเห็นว่า func1, func2 ,Anonymus-func ,func3 จะทำงานพร้อมกัน ถึงแม้ว่า Anonymus-1 จะเป็นคนที่ call func3 ให้ทำงาน

ดังนั้นจากผลการทดสอบผมเลยสรุปว่า การใช้งาน goroutine นั้นจะไม่สนเรื่องลำดับ หรือ level ในการเรียกใช้ ``go [ชื่อ func]`` โดยสุดท้ายทุกๆ function ก็จะทำงาน ทำงาน concurrently กัน

```golang
func main() {
	startTime := time.Now()
	fmt.Printf("---- Start ---- [%v] \n", startTime)
	fmt.Printf("---- Start program ---- [%v] \n", time.Since(startTime))
	inputData := "123"
    
	go func1(inputData, startTime)

	go func2(inputData, startTime)

    //Anonymus func
	go func(s string) {
		go func3(s, startTime)
		for _, d := range s {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Anonymus-1 output = %c [%v]\n", d, time.Since(startTime))
		}
	}(inputData)

	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("---- End progra ---- [%v] \n", time.Since(startTime))
	fmt.Printf("---- End ---- [%v] \n", startTime)

}
```

