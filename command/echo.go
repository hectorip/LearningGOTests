package main
import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omitiendo nueva línea al final")
	s = flag.String("s", " ", "Separador")
)

var out io.Writer = os.Stdout

func main(){
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo %v\n", err)
		os.Exit(1)
	}
}