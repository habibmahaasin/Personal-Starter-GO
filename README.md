# SmartAerators - GuppyTech
Repository Tugas Akhir

<br>

## Cara Setup dan Menjalankan

- download dan install GO

```
https://go.dev/doc/install
```

- import file database yang telah disediakan / menggunakan configurasi database yang telah disediakan
- jalankan program

```
go run main.go
```

## Notes Internal Developer Testing

### Mocking step :
1. install mockgen
2. ``` export PATH=$PATH:/Users/{user}/go/bin (mac os) ```
3. ``` mockgen . {interface name}  >> mock/{directory}.go ```

### Cek coverage testing
1. go to directory
2. ``` go test -v -coverprofile cover.out ```
3. ``` go tool cover -html cover.out ```
