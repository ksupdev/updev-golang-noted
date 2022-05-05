# updev-golang-noted
noted 

GO => NON Blocking การทำงานแบบไม่ต้องรองานก่อนหน้าให้เสร็จ
go ไม่ได้ใช้การทำ callback เหมือน java หรือ node แต่จะใช้ goroutine + channel


## Compare process
Threads => แบ่งงานระดับ lowlevel
[Java,Node.js,Go]
- java fix strack memory , Cannot plus

java nodejs จะทำงานแตก threads จริง
แต่ golang จะทำงานอยู่บน goroutine ซึ่งจะทำงานอยู่บน threads เดียว ซึ่งจะทำงานได้เร็วกว่า การ switch threads

Golang จะ build machine code และสามารถ run code ที่ cpu ได้เลย เหมือนพวก ภาษา C
Java จะ run บน JVM แล้ว JVM จะจัดการ run บน CPU อีกที

Process => จะใหญ่กว่า Threads ซึ่งจะมีการใช้ Stack memory
[php]

## Interface
> ใช้สำหรับการแก้ปัญหา Polymorphism 

test