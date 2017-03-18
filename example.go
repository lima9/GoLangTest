// example project example.go
package main

import "fmt"
import "time"
import "strconv"
import "os"

//var carTypeDist1 string = "Toyota"
//var carTypeDist2 = "Ford"

//var carTypeDist1, carTypeDist2 = "Toyota", "Ford"

var (
	carTypeDist1 string = "Toyota"
	carTypeDist2 string = "Ford"
)

const toyotaHeadQuarters, fordHeadQuarters = "Toyota Motor Corporation", "Ford Motor Company"

var newHQ string = "newHQ"

var shippingDays = 30
var shippingDaysPtr = new(int)

func main() {
	newHQ = "Ford HQ" //It has to be used
	newHQ := "Assing Variable funny way"

	var mystring = "10123"
	var x = 100
	var f float32 = 3.141516

	number, _ := strconv.Atoi(mystring)
	x = number

	fmt.Printf("%T", mystring, x, f, number)

	fmt.Println(carTypeDist1, carTypeDist2)
	fmt.Println(carTypeDist2)
	fmt.Println(toyotaHeadQuarters, fordHeadQuarters, newHQ)
	fmt.Println(time.Now())

	d := dealerShip{name: "A1 Auto", city: "Los Santos"}
	fmt.Println("city = " + d.city)

	var d2 dealerShip
	d2 = dealerShip{name: "Canada", city: "las vegas"}
	fmt.Println(d2)

	d3 := dealerShip{}
	fmt.Println("d3.name = " + d3.name)

	d1 := dealerShip{"A1 Auto", "Los Angeles"}
	dealerName := createDealerFullName(d1.name, d1.city)
	fmt.Println(dealerName)

	sold, name := calculateInvUti(10123123131421231312312414220123123, 175123131312412424121431322123132, d1)
	fmt.Println(name, sold)

	d4 := "Somethung"
	d5 := "Limousines"
	d6 := "Plastidecors"
	printDealers(d4, d5, d6)
	fmt.Println("\n")
	dealers := []string{d4, d5, d6}
	printDealers(dealers...)

	file := createFile("dealers.txt")
	defer closeFile(file)
	writeToFile(file, "A1 Auto motion")

	shippingDaysAdjustments(shippingDays)
	fmt.Println("After shippingDaysAdjustments Call", shippingDays)

	shippingDaysAdjustmentsPtr(&shippingDays)
	fmt.Println("After shippingDaysAdjustmentsPtr Call", shippingDays)
	fmt.Println("After shippingDaysAdjustmentsPtr Call", shippingDaysAdjustments)
	fmt.Println("After shippingDaysAdjustmentsPtr Call", shippingDaysAdjustmentsPtr)
	fmt.Println("After shippingDaysAdjustmentsPtr Call", &shippingDays)

	shipper := shipper{}
	shipper.id = 4000
	shipper.name = "Name faboulous"
	shipperUpdates(&shipper)
	fmt.Println("After shipperUpdates call", shipper.id)
	fmt.Println("After shipperUpdates call", shipper.name)

	*shippingDaysPtr = 545
	shippingDaysAdjustmentsPtr(shippingDaysPtr)
	fmt.Println("shippingDaysPtr After shippingDaysAdjustmentsPtr call ", *shippingDaysPtr)

}

func shippingDaysAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAdjustmentsPtr(shippingDays *int) {
	*shippingDays += 10
}

func shipperUpdates(s *shipper) {
	s.id += 1000
	s.name += "Corp"
}

type dealerShip struct {
	name string
	city string
}

type shipper struct {
	name string
	id   int
}

func createDealerFullName(s1 string, s2 string) string {
	return s1 + " of " + s2
}

func calculateInvUti(remaining float32, start float32, dealer dealerShip) (percentSold float32, dealerName string) {
	dealerName = dealerName + "sold"
	percentSold = 1 - remaining/start
	return

}

func printDealers(dealers ...string) {
	for _, dealerName := range dealers {
		fmt.Println(dealerName)
	}
}

func createFile(path string) *os.File {
	fmt.Println("Creating File")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, dealerName string) {
	fmt.Println("Writing To File")
	fmt.Fprintln(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing File")
	file.Close()
}
