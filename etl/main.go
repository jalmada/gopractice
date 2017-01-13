package main
 
 import(
     "fmt"
     "encoding/csv"
     "os"
     "strconv"
     "time"     
     //"runtime"
 )

type Product struct {
    PartNumber string
    UnitCost float64
    UnitPrice float64
}

type Order struct {
    CustomerNumber int
    PartNumber string
    Quantity int

    UnitCost float64
    UnitPrice float64
}

 func main(){

     //runtime.GOMAXPROCS(4)

     start := time.Now()

     extractChannel := make(chan *Order)
     transformChannel := make(chan *Order)
     doneChannel := make(chan bool)


//     orders := extract()
     go extract(extractChannel)
//     orders = transform(orders)
     go transform(extractChannel, transformChannel)
//     load(orders)
     go load(transformChannel, doneChannel)

     <- doneChannel
     fmt.Println(time.Since(start))

 }

 func extract(ch chan *Order){
//func extract() []*Order
     //result := []*Order{}

     f, _ := os.Open("./orders.txt")
     defer f.Close()
     r := csv.NewReader(f)

    for record, err := r.Read(); err == nil; record, err = r.Read() {
        order := new(Order)
        order.CustomerNumber, _ = strconv.Atoi(record[0])
        order.PartNumber = record[1]
        order.Quantity, _ = strconv.Atoi(record[2])
        ch <- order
        
        //result = append(result, order)
    }
    close(ch)
    //return result 
 }

 //func transform(orders []*Order) []*Order{
 func transform(extractChannel, transformChannel chan *Order) {
     f, _ := os.Open("./productList.txt")
     defer f.Close()
     r := csv.NewReader(f)


     records, _ := r.ReadAll()
     productList := make(map[string]*Product)

     for _, record := range records {
        product := new(Product)
        product.PartNumber = record[0]
        product.UnitCost, _ = strconv.ParseFloat(record[1],64)
        product.UnitPrice, _ = strconv.ParseFloat(record[2],64)
        productList[product.PartNumber] = product 
     }

     numMessages := 0
     //for idx, _ := range orders {
     for o := range extractChannel {
         numMessages++
         go func(o *Order){
            time.Sleep(3 * time.Millisecond)
            //o := orders[idx]
            o.UnitCost = productList[o.PartNumber].UnitCost
            o.UnitPrice = productList[o.PartNumber].UnitPrice
            transformChannel <- o
            numMessages--
         }(o)
     }

     for numMessages > 0 {
         time.Sleep(1 * time.Millisecond)
     }

     close(transformChannel)
     //return orders
 }

// func load(orders []*Order){
 func load(transformChannel chan *Order, doneChannel chan bool){
     f, _ := os.Create("./dest.txt")
     defer f.Close()

     fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", 
     "Unit Cost", "Unit Price", "Total Cost", "Total Price")

     numMessages := 0
     //for _, o := range orders {
    for o := range transformChannel {
         numMessages++        
        go func(o *Order){
            time.Sleep(1 * time.Millisecond)
            fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
            o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
            o.UnitCost * float64(o.Quantity),
            o.UnitPrice * float64(o.Quantity))
            numMessages--            
        }(o)
     }

     for numMessages > 0 {
         time.Sleep(1 * time.Millisecond)
     }
     
     doneChannel <- true
     //close(doneChannel)   
 }