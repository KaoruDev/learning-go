package main

import "fmt"

type vehicle struct {
    vin int
}

type car struct {
    vehicle
    model string
}

func (c car) modelName() string {
    return fmt.Sprintf("%d:%s", c.vin, c.model)
}

func main() {
    var myCar car
    z := car{vehicle{1}, "z"}
    boat := vehicle{}

    corolla := car{ vehicle: vehicle{}, model: "foobar" }

    fmt.Println(corolla)

    fmt.Println(myCar.modelName())
    fmt.Println(z.modelName())

    fmt.Println(z)
    fmt.Println(boat)
}
