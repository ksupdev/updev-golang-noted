
## Build to production
- Simple build ``go build`` ซึ่งจะได้ file ที่มีชื่อเดียวกับ module โดยการ build แบบนี้จะมีขนาดใหญ่มาก (14.7 mb)

```powershell
	go build
	./updev-go-ex-stock-api
```

- Build and remove the junk file ``go build -ldflags "-s -w"`` จะได้ file ที่มีขนาดเล็กกว่าแบบแรก (11.6 mb)
```powershell
	go build -ldflags "-s -w"
	./updev-go-ex-stock-api
```