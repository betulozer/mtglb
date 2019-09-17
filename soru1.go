// Read fonksiyonu içerisinde dosyanın kapalı olup olmadığını
// kontrol edin. Dosya kapalı ise "file: file is closed"
// mesajı ile birlikte yeni bir hata döndürün.

package main
import "fmt"

type File struct {
    closed bool
}
func Read(file File) {
    if file.closed {
        fmt.Println("file: file is closed")
    }
}
func main() {
    f := File{closed: true}
    Read(f)
}
