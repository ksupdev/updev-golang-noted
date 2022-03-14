## variable
```go
    // Basic variable declaration
    var i int
    var a string
    var ok bool

    // decla and set default value
    var i int = 14
    var a string = "hello"
    var ok bool = true

    // Type inference
    var i = 14
    var a = "hello"
    var ok = true

    // Type inference in function only
    i := 14
    a := "hello"
    ok := true

```

## Constant
Constant is a immutable, เป็นการประกาศตัวแปลที่ไม่ต้องการให้มีการแก้ไขหลักจากการประกาศใช้งาน
```go
    const defaultValue int = 1
    const defaultValue2 = "value"

    const{
        errorCode = 1
        successCode = 2
    }
```
## if Else

```golang
    if n, err := strconv.Atoi("5s"); err != nil{

    }

    n, err := strconv.Atoi("5s");
    if err != nil{

    }else{

    }
```

## Loop

```golang
//A loop with 3 components
for i := 0; i < 10; i ++{

}

// A loop with a condition
for i < 10{

}

// An infinite loop
for{

}

```