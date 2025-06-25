package main

import "fmt"


// type UserService struct{
//     logger ConsoleLogger     // Now, UserService is hardcoded to use ConsoleLogger. If you want to test UserService, or use a different logger (like FileLogger), you’ll have to rewrite the internals.
// }

// type ConsoleLogger struct{
// }

// func(c ConsoleLogger) Log(msg string){
// 	fmt.Println(msg)
// }

// func(u UserService) CreateUser(name string){
// 	u.logger.Log("Hello," + name)
// }

// /// You are calling CreateUser directly in main(), but it's a method of UserService, not a standalone function.

// // func main(){
// // 	CreateUser("Alice")
// // }

// func main()  {
//    logger := ConsoleLogger{}
//    UserService := UserService{logger: logger}	
//    UserService.CreateUser("alice")
// }


//// instead not hard code dependency

// type Logger interface{
//     Log(msg string)
// }

// type UserService struct{
//     logger Logger     // Now, UserService is hardcoded to use ConsoleLogger. If you want to test UserService, or use a different logger (like FileLogger), you’ll have to rewrite the internals.
// }

// type ConsoleLogger struct{
// }

// func(c ConsoleLogger) Log(msg string){
// 	fmt.Println(msg)
// }

// type MockLogger struct{
// }

// func(m MockLogger) Log(msg string){
//     fmt.Println("Mock test:" + msg)
// }

// func(u UserService) CreateUser(name string){
// 	u.logger.Log("Hello," + name)
// }


// func main()  {
//    consoleLogger := ConsoleLogger{}
//    mockLogger := MockLogger{}
//    UserService1 := UserService{logger: consoleLogger}	
//    UserService2 := UserService{logger: mockLogger}	
//    UserService1.CreateUser("alice")
//    UserService2.CreateUser("alice")

// }


//// Best example

// w/o dependency

// type SqlDatabase struct{

// }

// func(d SqlDatabase) Save(){
//     fmt.Println("Saving SQL")
// }

// type UserService struct{
//   db SqlDatabase
// }

// func(u UserService) SaveDB(){
//    u.db.Save()
// }


// func main(){
//     sql := SqlDatabase{}
//     sqlService := UserService{ db: sql}
//     sqlService.SaveDB()
// }


//// now what if I want mongo also then we will be needed to create whole thing again

// 1. Create a interface

type Database interface{
    Save(data string)
}

// 2. Implement the Interface
// Implement that interface using one or more concrete types.
type SqlDatabase struct{
}

func( s SqlDatabase) Save(data string){
    fmt.Printf("Save this %s \n", data)
}

type MongoDatabase struct{
}

func(m MongoDatabase) Save(data string){
    fmt.Printf("Save this %s  \n", data)
}


//// 3. Step 4: Inject the Dependency
//// Instead of creating the dependency inside the struct, pass it in via constructor or initialization.

type UserService struct{
    db Database 
}

func (u UserService) SaveDBWithDI(db string){
    u.db.Save("{name:Rahul} to " + db)
}

//// 4. Use Dependency from Outside
//// When initializing your service (e.g., in main), inject the actual implementation.

func main(){
    sqlServiuce := UserService{db:SqlDatabase{}}
    sqlServiuce.SaveDBWithDI("sql")

    mongoServiuce := UserService{db:MongoDatabase{}}
    mongoServiuce.SaveDBWithDI("mongo")

}