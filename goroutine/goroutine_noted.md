# Golang routine noted

## problem
1. ต้องการอยากรู้การทำงานของ goroutine 
2. อยากเข้าใจเกี่ยวกับการ Chanel Block

## Get start
เราได้ทำการสร้าง MainService struct เพื่อใช้สำหรับ สร้าง Chanel ExitChannel ขึ้นมา โดย MainService จะหยุดทำงานเมื่อมีการส่งค่า true มาที่ ExitChannel เท่านั้นซึ่งสามารถเรียกใช้ ``func (ms *MainService) Stop(isStop bool)`` เพื่อทำการ set ค่าให้กับ Chanel

### โดย MainService struct จะมี Method ดังนี้
1. ``Start()`` func โดยจะทำหน้าที่หลักๆดังนี้
- create chanel ``ms.ExitChannel = make(chan bool, 1)`` โดยกำหนดให้มี buffer เท่ากับ 1 โดยค่า default ของ channel จะไม่สามารถ buffer ข้อมูลได้ หรือก็คือถ้าได้รับมาแล้ว ถ้ายังไม่ได้เอาข้อมูลออกไปจะมาสามารถรับข้อมูลใหม่ได้นั้นเอง แต่ถ้าเราต้องการให้รับข้อมูลได้มากกว่า 1 เราจะต้องทำการกำหนด buffer และใช้ ``for range`` ในการเอาข้อมูลออกมาอีกที
- ทำการสร้าง for { } infinity loop หรือก็คือ for นี้จะทำงานไปเลื่อยๆจนกว่า และเราต้องทำการสร้าง condition โดยใช้ if เพื่อ break loop
- ภายใน for เราได้ทำการสร้าง if else เพื่อกำหนด condition เพื่อออกจาก loop โดย condition ก็คือ exit == true โดยค่า Exit นั้นจะได้มาจาก Channel นั้นเอง
- และสิ่งที่สำคัญที่สุดก็คือ เราได้ทำการสร้าง select case เพื่อให้ใช้ในการ block การทำงานจนกว่าจะได้ค่ามาจาก channel ``case <-ms.ExitChannel`` แต่ใน code เราได้มีการเขียน ``case data1 := <-ms.ExitChannel:`` เพื่อที่เราอยากจะเห็นค่าที่ได้จาก channel และทำการส่งไปให้ if check เงื่อนไขต่อ แต่มันยังไม่สะใจเรา ได้สร้าง ``data2 := <-ms.ExitChannel`` เพื่อให้ทำการ block เพื่อรอการส่งค่าอีกครั้ง พูดง่ายๆ เราทำการ พอเราเริ่ม Start มีการสร้าง channel แล้ว loop จะทำงานวนผ่าน if พอไม่ตรงเงื่อนไข ก็จะไปเจอ select case และรอจนกว่าจะมีใครส่งค่ามาที่ ExitChannel และเมื่อได้รับค่า ก็จะมีการ หยุดอีกครั้งเพื่อรอรับค่าอีกที และเมื่อทุกอย่างเรียบร้อย ก็จะ วน loop รอบต่อไป เพื่อไป check if และถ้า exit == true ก็จะทำการ break loop หรือก็คือจบการทำงานนั้นเอง

### การทำงานใน main.go
เราได้ทำการ implement Anonymus func (1) ,Anonymus func (2) เพื่อทำการจำลอง การทำงานของ goroutine โดยให้

Anonymus func (1) ทำงานเป็นเวลา time.Sleep(100 * time.Millisecond)

และ Anonymus func (2) ทำงานเป็นเวลา time.Sleep(50 * time.Millisecond)

หลังจากนั้นเราจึงได้มีการเรียก Start() ใน MainService struct เพื่อเริ่มการทำงาน
> โดยตอนแรกพยายามเอาไปใว้ก่อนจะสร้าง Anonymus fuc แต่มันไปเจอ error เนื่องจากใน method start() มันมีคำสั้งในการ block chanel อยู่ซึ่งเหมือนมันจะไม่ยอม ถ้ายังไม่ได้มี goroutine กำลังทำงาน หรือมีกระบวนการที่จะ send ค่าเข้าไปที่ channel อันนี้ก็ยัง งงๆ นิดๆ

และพอเรา ``go run main.go`` เพื่อทดสอบจะได้ผลลัพท์ดังนี้

```powershell
% go run main.go
 [line-0] Start 2021-05-09 15:59:12.370768 +0700 +07 m=+0.000091446 
 [line-1] - In for loop at time 188.075µs 
 [line-2] Waiting .... value from ExitChannel  at time 194.113µs 
 [line-3] begin Anonymus func (2) at time 197.18µs 
 [line-4] begin Anonymus func (1) at time 211.652µs 
 [line-5] end Anonymus func (2) at time 50.239342ms 
 [line-6] (1) After recieve from case <-ms.ExitChannel [false] at time 50.288305ms 
 [line-7] end Anonymus func (1) at time 100.357774ms 
 [line-8] (2) After recieve from data := <-ms.ExitChannel [true] at time 100.387284ms 
 [line-9] In for loop at time 100.393229ms 
 [line-10] exit program at time 100.396764ms 
 [line-11] End  at time 100.400017ms 

```

- line-0 : การทำงานอันนี้จะเป็นตอนที่เราเริ่มทำงาน func TestAdvGoRoutine โดยเราแค่ต้องการ logtime ออกมาดู
- line-1 : อันนี้เป็นสิ่งที่ยัง งง ไม่หาย ``In for loop at time`` มันเป็นการทำงานภายใน Start() func แต่เท่าที่เราจำได้มันจะเป็น การทำงานหลังสุดหรือก็คือมันถูกเรียกหลังจาก เราทำการ run Anonymus func 1,2 ไปแล้วแต่จาก log เราจะเห็นว่ามีการทำงานก่อน
- line-2 : `` Waiting .... value`` อันนี้จะเริ่มเข้าไปใน loop เพื่อ block channel เพื่อรอรับค่าจาก Channel เพื่อทำงานต่อ
- line-3,line-4 : Anonymus func (2) และ Anonymus func (1) เริ่มทำงาน
- line-5,line-6 : Anonymus func (2) ทำงานเสร็จและมีการ เรียก Stop() func [50.239342ms] เพื่อทำการกำหนดค่าให้ Channel เป็น false และทันทีที่ได้รับค่า ``case <-ms.ExitChannel`` ที่อยู่ใน loop จะได้รับค่าทันที โดยจะรับค่าเป็น false [50.288305ms] จาก Anonymus func (2) และก็จะเริ่มมีการ block channel อีกทีโดยทันทีเพื่อรอรับค่าอีกครั้ง ต้องย้ำว่าการรอรับค่าครั้งที่ 2 ปกติมันจะไม่มีแต่อันนี้เพื่อการศึกษาเท่านั้น
- line-7,line-8 : Anonymus func (1) ทำงานเสร็จและส่งค่า true [100.357774ms] ผ่านทางการ call Stop() func และ channel ก็จะได้รับค่า [100.387284ms] ซึ่งจะได้ค่าเป็น true และทำการกำหนดค่า ให้ exit = true 
- line-9,line-10 : เหมือนหลุดจากการ block 2 ครั้งก็จะไป loop รอบที่ 2 เพื่อทำการ check if และก็เป็นไปตามเงื่อนไข if exit == true ดังนั้นก็จะมีการ break เพื่อจบ loop และเป็นอันจบ program


