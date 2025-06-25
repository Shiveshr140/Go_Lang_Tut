package main

import (
	"fmt"
	// "log"
)

// why pointer receivers are often preferred in interfaces!, you noticed &bike{} not bike{} -> because if method is pointer reciever then
// So the method has a pointer receiver (*bike). Therefore, only a *bike (pointer) satisfies the interface, not bike (value).
// Because bike{} is a value do not have Drive method
//  Can a struct value (like bike{}) satisfy an interface? Only if the struct implements all methods of the interface using value receivers.

// type bike struct {
// }

// type car struct {
// }

// type vehicle interface {
// 	Drive()
// }

// func (c *car) Drive() {
// 	fmt.Println("I am Driving a car!")
// }

// func (b *bike) Drive() {
// 	fmt.Println("I am Driving a bike!")
// }

// func VehicleFactory(vehicleType string) vehicle {
// 	if vehicleType == "bike"{
// 		return &bike{}
// 	}

// 	if vehicleType == "car"{
// 		return &car{}
// 	}
// 	return nil

// }

// func main() {
// 	myVehicle := VehicleFactory("bike")
//     myVehicle.Drive()
// }

///// Lets make it more clear w/o & factory lets add truck

// type car struct{}
// type bike struct{}
// type truck struct{}

// func (c *car)  Drive() { fmt.Println("Driving a car") }
// func (b *bike) Drive() { fmt.Println("Riding a bike") }
// func (b *truck) Drive() { fmt.Println("Sleeping in a truck") }

// func main() {
// 	// I myself decide which concrete thing to build.
// 	myVehicle1 := &car{}
// 	myVehicle2 := &bike{}
// 	myVehicle3 := &truck{}

// 	myVehicle1.Drive()
// 	myVehicle2.Drive()
// 	myVehicle3.Drive()
// }





//// with factory & lets add truck

// Situation	“No-factory” style
// 1. One new vehicle (truck)	• Add type truck struct{} and its Drive() • Find every file that ever does &car{} or &bike{} and add new truck logic there if it might be chosen.	• Add type truck struct{} and Drive()
// Factory style -> • Add one extra case "truck": return &truck{} in the factory. Call-sites need zero edits.

// type DriveContract interface{
// 	Drive()
// }

// type car struct{}
// type bike struct{}
// type truck struct{}

// func (c *car)  Drive() { fmt.Println("Driving a car") }
// func (b *bike) Drive() { fmt.Println("Riding a bike") }
// func (t *truck) Drive() {fmt.Println("Sleeping in a truck")}

// func VehicleFactory(vType string) DriveContract{
//    if vType == "car"{
// 	 return &car{}
//    }
//    if vType == "bike"{
// 	return &bike{}
//    } 

//    if vType == "truck"{
// 	return &truck{}
//    }

//    return nil
//    }

// func main() {
// 	vehicle := VehicleFactory("truck")
// 	if vehicle == nil{
// 		log.Fatal("Invalid vehicle")
// 	}

// 	vehicle.Drive()
// }