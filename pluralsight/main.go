package main
import (
"fmt"
"github.com/xaratustra/reports"
"strings"
)




func main(){
   
    PowerPlantReport()
    //sayHello("Hello", "Go", "from", "Pluralsight")
    // addFunc := func (terms ...int) (numTerms int, sum int) {
    //     for _, term := range terms {
    //         sum += term
    //     }

    //     numTerms = len(terms)

    //     return 
    // }

    // numterms, sum := addFunc(1, 2, 3, 4)
    // println(numterms, " ", sum)
    // foo := newMyStruct()
    // foo.myMap["bar"] = "baz"
    // foo.printMessage()
}

type myStruct struct {
    myMap  map[string]string

}

func (ms *myStruct) printMessage() {
    println(ms.myMap["bar"])
}

func newMyStruct() *myStruct{
    result := myStruct{}
    result.myMap = map[string]string{}

    return &result
}


func sayHello(messages ...string) {
    for _, message := range messages {
        println(message)
    }
}

func addTerms(terms ...int) (numTerms int, sum int) {
    for _, term := range terms {
        sum += term
    }

    numTerms = len(terms)

    return 
}

type PlantType string
const (
    hydro PlantType = "Hydro"
    wind PlantType = "Wind"
    solar PlantType = "Solar"
)

type PlantStatus string
const (
    active PlantStatus = "Active"
    inactive PlantStatus = "Inactive"
    unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
    plantType PlantType
    capacity float64
    status PlantStatus
}

type PowerGrid struct {
    load float64
    plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport(){
    for idx, p := range pg.plants {
        label := fmt.Sprintf("%s%d", "Plant #", idx)
        fmt.Println(label)
        fmt.Println(strings.Repeat("-", len(label))

        fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
        fmt.Printf("%-20s%.0f\n", "Load", gridLoad)
        fmt.Printf("%-20s%.1f%%\n", "Utilization", utilization*100)
    }
}

func PowerPlantReport(){
    plantCapacities := []float64{30, 30, 30, 60, 60, 100}

    activePlants := []int{0, 1}

    gridLoad := 75.

    fmt.Println("1) Generate Power Plant Report")
    fmt.Println("2) Generate Power Grid Report")
    fmt.Print("Please choose an option: ")

    var option string

    fmt.Scanln(&option)

    switch option {
        case "1":
            reports.GeneratePlantCapacity(plantCapacities...)
        case "2":
            reports.GeneratePowerGrid(activePlants, plantCapacities, gridLoad)
        default:
            fmt.Printf("Unknow Option")
    }
}

