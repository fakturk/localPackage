package second
import(
	"fmt"
	"github.com/fakturk/localPackage/first"
)

func SecondPrint()  {
	fmt.Println("second")
	first.FirstPrint()
}