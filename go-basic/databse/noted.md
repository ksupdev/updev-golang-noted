# You shourd know

### การ Return data

ถ้าเป็นการ Select ข้อมูลแล้วคืนเป็น Slide เราสามารถ return เป็น Slide ได้เลย

แต่ถ้าเป็นการ select ข้อมูลแล้วคือเป็น single recoard เราควร retrun เป็น Pointer เนื่องจากเวลาที่เรา ดักพวก error แล้วให้ return error การที่เรา return เป็น struc brank มันมีการใช้ memory เยอะกว่าการ return แบบ pointer

```golang
    type Result struct {
        Id   int
        Name string
    }

    func getDatas() ([]Result , error){
        if found som error {
            return nil , err
        }

        return results, nil
    }

    func getData() (*Result ,error){
        if found some error {
            return nil , err
        }

        return result, nil
    }

```

### การใช้ Row.Scan(....)

```golang
type Result struct {
    Id   int
    Name string
}

func getData() (*Result ,error){

query := "select id ,name from result where id=@id"
row := db.QueryRow(query ,sql.Name("id",id))

result := Result{}
err := row.Scan(&result.Id ,$result.Name)

...

return &result , nil

}

```

```golang
query := "select id ,name from result where id=@id"
row := db.QueryRow(query ,sql.Name("id",id))
```

ในส่วนนี้จะเป็นการช่วยในการจัดการพวก Sql Injection เพราะถ้าเราใช้การต่อ String ปกติ อาจจะเจอปัญหานี้ได้ 
> ตัวอย่างนี้ใช้ได้สำหรับ SQL SERVER เท่านั้น


```golang 
query := "select id ,name from result where id=@id"
row := db.QueryRow(query ,sql.Name("id",id))
result := Result{}
err := row.Scan(&result.Id ,$result.Name)
```

ในการใช้ row.Scan นั้นจะไม่สนใจในส่วนของชื่อ แต่จะสนใจแค่ Index หรือลำดับของ pointer ที่ส่งเข้ามาเท่านั้น หรือก็คือ ส่งมาที่ ลำดับที่ 1 ก็จะได้ถูก mapping value จาก column ที่ 1 ที่มีการ select ออกมานั้นเอง

