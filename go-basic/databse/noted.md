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
query := "select id ,name from result where id=?"
row := db.QueryRow(query ,id)
```
> ตัวอย่างนี้ใช้ได้สำหรับ MYSQL เท่านั้น


```golang 
query := "select id ,name from result where id=@id"
row := db.QueryRow(query ,sql.Name("id",id))
result := Result{}
err := row.Scan(&result.Id ,$result.Name)
```

ในการใช้ row.Scan นั้นจะไม่สนใจในส่วนของชื่อ แต่จะสนใจแค่ Index หรือลำดับของ pointer ที่ส่งเข้ามาเท่านั้น หรือก็คือ ส่งมาที่ ลำดับที่ 1 ก็จะได้ถูก mapping value จาก column ที่ 1 ที่มีการ select ออกมานั้นเอง


## Create

```golang
func Create(data Result) error{
    ....
    ....
    query := "insert into result (id ,name) values(?,?)"
    result err := db.Exec(query ,data.Id data.Name)
    if err != nil{
        return err
    }

    // Use for get Id
    // result.LastInsertId 
    
    affected, err := result.RowsAffcted
    if err != nil{
        return nil
    }

    if affected <= 0{
        return errors.New("Cannot insert)
    }

    return nil
}
```



```golang
func Update(data Result) error{
    ....
    ....
    query := "update result set name=? where id=?"
    result err := db.Exec(query ,data.Name ,data.Id)
    if err != nil{
        return err
    }

    // Use for get Id
    // result.LastInsertId 
    
    affected, err := result.RowsAffcted
    if err != nil{
        return nil
    }

    if affected <= 0{
        return errors.New("Cannot Update)
    }

    return nil
}
```


```golang
func Delete(id int) error{
    ....
    ....
    query := "update from result where id=?"
    result err := db.Exec(query ,id)
    if err != nil{
        return err
    }

    // Use for get Id
    // result.LastInsertId 
    
    affected, err := result.RowsAffcted
    if err != nil{
        return nil
    }

    if affected <= 0{
        return errors.New("Cannot Delete)
    }

    return nil
}
```

Test use sqlx

```golang
func GetResultWithSqlX ([]Result ,error){
    query := "select id ,name from result"
    result := Result{}

    err := db.Select(&result ,query)
    if err != nill{
        return nil ,err
    }
    return result ,nil

}
```


## Transaction

```golang
func Create(data Result) error{

    tx, err := db.Begin()
    if err != nil{
        return err
    }
    ....
    ....
    query := "insert into result (id ,name) values(?,?)"
    result err := tx.Exec(query ,data.Id data.Name)
    if err != nil{
        return err
    }

    // Use for get Id
    // result.LastInsertId 
    
    affected, err := result.RowsAffcted
    if err != nil{
        tx.Rollback()
        return nil
    }

    if affected <= 0{
        return errors.New("Cannot insert")
    }

    err := tx.Commit()
    if err != nil{
        return err
    }
    

    return nil
}
```

