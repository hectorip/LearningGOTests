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
