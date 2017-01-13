package reports
import "fmt"

func GeneratePlantCapacity(plantCapacities... float64){
    for idx, cap := range plantCapacities {
        fmt.Printf("Plant %d capacity: %.0f\n",idx, cap)
    }
}

func GeneratePowerGrid(activePlants []int, plantCapacities []float64, gridLoad float64){
    capacity := 0.
    for _, plantId := range activePlants {
        capacity += plantCapacities[plantId]
    }

    utilization := gridLoad/capacity

    fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
    fmt.Printf("%-20s%.0f\n", "Load", gridLoad)
    fmt.Printf("%-20s%.1f%%\n", "Utilization", utilization*100)
}