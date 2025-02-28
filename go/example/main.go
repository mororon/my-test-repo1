package main
import "fmt"

//ビルド時にldflagsでバージョンを埋め込む
var version string

func main() {
	fmt.Printf("Example %s\n", version)
}