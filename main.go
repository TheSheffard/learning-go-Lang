// Every Go source file belongs to a package. The special package main is
// used to build an executable program; it can contain many functions.
package main

// fmt is a standard-library package for formatted input and output.
// Import it to use functions such as fmt.Println. It needs no separate install.
// Imports apply to this file. An ordinary unused import is a compile error.
import (
	"fmt"
	// "os/user"
)

// main is the entry-point function of this executable, after package
// initialization. Statements in its body run in order unless control flow
// changes that order. The opening brace belongs on the declaration line.
func main() {
	// 1. VARIABLES AND EXPLICIT TYPES
	// var declares a variable. The name is followed by its type and value.
	// A string stores text; an int stores whole numbers.
	var username string = "Emmanuel"
	var age int = 20

	// Println prints its arguments with spaces between them, then a newline.
	// The capital P matters: Println is exported from fmt.
	fmt.Println(username, age) // Emmanuel 20

	// = assigns to an existing variable. Its declared type stays the same.
	// The earlier print already happened, so it still showed 20.
	age = 21
	// fmt.Println(age) // Uncomment this line to print 21.

	// 2. SHORT VARIABLE DECLARATION AND TYPE INFERENCE
	// := declares and initializes variables inside a function.
	// Go infers string from "Joker" and int from 4.
	riderName := "Joker"
	completedDeliveries := 4
	fmt.Println(riderName, completedDeliveries)

	// Use = to update a variable already declared in this scope.
	// A := declaration needs at least one new non-blank variable on its left.
	// Repeating completedDeliveries := 7 here would therefore be an error.
	completedDeliveries = 7
	fmt.Println(completedDeliveries)
	space()

	// 3. CONSTANTS
	// const names a constant value that cannot be reassigned.
	// Constants need an initial constant expression, such as a literal.
	// They cannot be declared using :=.
	const appName string = "Dispatchgo"
	fmt.Println(appName)
	space()

	const platformName string = "QuickRide"
	const maxActiveDeliveries int = 3
	activeDeliveries := 1
	fmt.Println(platformName, maxActiveDeliveries, activeDeliveries)
	activeDeliveries = 2 // A variable can change; the constants above cannot.
	fmt.Println(activeDeliveries)
	space()

	// 4. FLOATING-POINT AND BOOLEAN TYPES
	// A decimal literal like 4.9 gives a := variable the type float64.
	// float64 represents floating-point numbers; some decimals are approximate.
	// bool has exactly two values: true and false, without quotation marks.
	rating := 4.9
	isOnline := false
	fmt.Println(rating, isOnline)
	space()

	var deliveryDistance float64 = 12.5
	isRiderAvailable := true
	fmt.Println(deliveryDistance, isRiderAvailable)
	isRiderAvailable = false
	fmt.Println(isRiderAvailable)

	// 5. ZERO VALUES
	// var can declare a typed variable without an explicit initial value.
	// Go initializes it to that type's zero value: "" for string, 0 for int
	// and float64, and false for bool. It is not an undefined variable.
	var clubName string
	var goals int
	fmt.Println("Club:", clubName) // The name is empty text.
	fmt.Println("Goals:", goals)   // 0
	space()

	// 6. TYPE CONVERSION
	// Arithmetic between these int and float64 variables needs conversion.
	// float64(distance) produces a converted value for this calculation.
	// It does not change distance itself, which remains an int.
	distance := 5
	extraDistance := 2.5
	totalDistance := float64(distance) + extraDistance
	fmt.Println(totalDistance) // 7.5
	space()

	completedTrips := 6
	bonusTrips := 1.5
	totalTrips := float64(completedTrips) + bonusTrips
	fmt.Println(totalTrips) // 7.5

	// Converting this float64 variable to int truncates toward zero.
	// It discards the fractional part instead of rounding: 7.5 becomes 7.
	// The original totalTrips still holds 7.5.
	fmt.Println("The converted total", int(totalTrips))
	space()

	// Call each topic's function to run its examples. Declaring a function
	// alone does not execute its body. These calls now run all topic examples.
	arithemitcs()
	comparison()
	Logicaloperators()
	conditionalStatement()
	Loop()
	space()

	// 7. FUNCTION CALLS AND ARGUMENTS
	// Arguments are the actual values supplied when calling a function.
	// They must match the parameter types and order in its declaration.
	greetRider("Emmanuel")
	greetRider("Joker")
	showPlayer("Saka", 12)
	space()

	// 8. RECEIVING RETURN VALUES
	// The function sends a result back; this variable stores it.
	// Returning does not print. Println displays the stored result separately.
	totalGoals := addGoals(2, 4)
	fmt.Println(totalGoals) // 6
	space()

	deliveryTotal := calculateDeliveryFee(4.5, 200.0)
	fmt.Println("The total", deliveryTotal) // 900
	space()

	// 9. MULTIPLE RETURN VALUES
	// Go functions can return several values, received in positional order.
	sum, difference := calculateScores(10, 4)
	fmt.Println(sum)        // 14
	fmt.Println(difference) // 6
	space()

	playerName, playerGoals := getPlayer()
	fmt.Println(playerName, playerGoals)

	// _ is the blank identifier: it discards a value and creates no variable.
	// playerName1 is new here, so := is valid.
	playerName1, _ := getPlayer()
	fmt.Println(playerName1)
	space()

	originalPrice, afterPrice := calculateOrderAmounts(5000.0, 500.0)
	fmt.Println("Original price", originalPrice, "and after discount", afterPrice)

	// If afterPrice already exists in the same scope, receive into it with =:
	// _, afterPrice = calculateOrderAmounts(5000.0, 500.0)
	// _ never counts as a new variable for a := declaration.
	space()

	LearningArray()
	LearingSlices()
	space()
	learningMap()
	space()
	learningStruct()
	space()

	learningPointer()

	space()

	customer := Customer{Name: "Emma"}

	owner := ShopOwner{Name: "Joker"}

	owner.Introduce()

	customer.Introduce()

	space()
	err := validateRiderAge(18)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	 fmt.Println("Rider age accepted")

	space()

	team, err := CreateTeam("Arsenal")

	if err != nil {
		fmt.Println("Could not create team:", err)
		return
	}

	fmt.Println("Created team:", team.Name)

	space()


	rider, err := registerRider("", 16)

	if err != nil{
		fmt.Println("An Error occurred", err)

		return
	
	}

	fmt.Println("The riders value: ",rider)


}

// arithemitcs demonstrates division. The existing function name is retained.
func arithemitcs() {
	players := 11
	teams := 2

	// Integer division produces an integer quotient, truncating toward zero.
	fmt.Println(players / teams) // 5

	// Convert BEFORE division to keep the fractional part in this example.
	// float64(players / teams) would convert the already-truncated result.
	average := float64(players) / float64(teams)
	fmt.Println(average) // 5.5

	// Other arithmetic operators learned:
	// + adds, - subtracts, * multiplies, and % gives an integer remainder.
	// For example, 11 % 2 is 1: one player remains after equal grouping.
	space()
}

// comparison demonstrates operators whose results are boolean values.
func comparison() {
	riderAge := 20
	fmt.Println(riderAge == 20) // true: equal to
	fmt.Println(riderAge != 20) // false: not equal to
	fmt.Println(riderAge > 18)  // true: greater than
	fmt.Println(riderAge < 18)  // false: less than
	fmt.Println(riderAge >= 20) // true: greater than or equal to
	fmt.Println(riderAge <= 19) // false: less than or equal to

	// Unlike JavaScript, Go has no === operator. Use == for equality.
	// = assigns a value; == compares values. String equality is case-sensitive.
}

// Logicaloperators demonstrates combining boolean conditions.
func Logicaloperators() {
	isVerified := true
	isOnline := false

	// && means AND: both conditions must be true.
	canReceiveOrders := isVerified && isOnline
	fmt.Println(canReceiveOrders) // false

	// || means OR: at least one condition must be true.
	// ! means NOT: it reverses a boolean, so !false is true.
	// Logical operands must be booleans; Go does not use JS-style truthiness.
	// Write each comparison fully: age >= 18 && age <= 60.
}

// conditionalStatement demonstrates choosing which block to execute.
func conditionalStatement() {
	customersAge := 22
	// if runs its body only when its boolean condition is true.
	// Parentheses around the condition are not required.
	if customersAge >= 16 {
		fmt.Println("You can register")
	}

	accountBalance := 1000
	bookingFee := 1500
	if accountBalance >= bookingFee {
		fmt.Println("Booking successful")
	} else {
		// else handles the false case. Keep } else { on the same line.
		fmt.Println("Insufficient balance")
	}

	playerGoals := 8
	// Conditions are checked from top to bottom. Only the first matching
	// branch runs; else handles the case where none of the conditions match.
	if playerGoals >= 10 {
		fmt.Println("Gold reward")
	} else if playerGoals >= 5 {
		fmt.Println("Silver reward")
	} else {
		fmt.Println("Keep practising")
	}
}

// Loop demonstrates repetition, increment/decrement, break, and continue.
func Loop() {
	// Initialization runs once. The condition is checked before each pass.
	// The post statement runs after each normally completed pass.
	// number++ increments by one; the loop prints 1 through 5.
	for number := 1; number <= 5; number++ {
		fmt.Println(number)
	}
	// number was declared in the loop's scope and cannot be used here.
	space()

	// -- decrements by one. This loop prints 3, 2, 1.
	for countdown := 3; countdown >= 1; countdown-- {
		fmt.Println(countdown)
	}
	space()

	// break exits this loop immediately. Since it occurs before Println,
	// this loop prints 1, 2, 3 and does not print 4.
	for ticket := 1; ticket <= 6; ticket++ {
		if ticket == 4 {
			break
		}
		fmt.Println(ticket)
	}
	space()

	// continue skips the remaining body for this pass. seat++ still runs
	// before the next condition check. Output: 1, 2, 4, 5.
	for seat := 1; seat <= 5; seat++ {
		if seat == 3 {
			continue
		}
		fmt.Println(seat)
	}
}

// greetRider accepts a string parameter named name. Parameters are local
// variables initialized from the arguments supplied by the caller.
func greetRider(name string) {
	fmt.Println("Welcome", name)
	space()
}

// showPlayer accepts two parameters with different types, in this order.
// No return type follows the parameter list, so it returns no values.
func showPlayer(name string, goals int) {
	fmt.Println(name, "scored", goals, "goals")
}

// addGoals returns one int. The int after ) declares the result type.
// return sends the result to the caller and exits the function.
func addGoals(first int, second int) int {
	return first + second
}

// calculateDeliveryFee multiplies two float64 parameters and returns a float64.
// The caller decides whether to store, print, or use the result elsewhere.
func calculateDeliveryFee(distance float64, ratePerKm float64) float64 {
	return distance * ratePerKm
}

// calculateScores returns two ints: the sum first, then the difference.
// The values and types in return must match the declared result list.
func calculateScores(first int, second int) (int, int) {
	return first + second, first - second
}

// getPlayer takes no parameters and returns values of different types.
func getPlayer() (string, int) {
	return "Saka", 12
}

// calculateOrderAmounts returns the original price and discounted price.
// This practice example subtracts the supplied discount without validation.
func calculateOrderAmounts(price float64, discount float64) (float64, float64) {
	return price, price - discount
}

// LearningArray demonstrates a fixed-size collection of one element type.
func LearningArray() {
	// [3]string means exactly three strings. Array length is part of its type.
	clubs := [3]string{"a", "b", "c"}

	// Indexes start at zero. Valid indexes here are 0, 1, and 2.
	fmt.Println(clubs[0]) // a
	fmt.Println(clubs[2]) // c

	// Replacing an element changes its value, not the array's fixed length.
	clubs[1] = "emma"
	fmt.Println(clubs)                                 // [a emma c]
	fmt.Println("The length of the array", len(clubs)) // 3

	// Use < len(clubs), not <=: the last valid index is len(clubs)-1.
	for index := 0; index < len(clubs); index++ {
		fmt.Println(clubs[index])
	}
}

// LearingSlices demonstrates slices, append, and range.
func LearingSlices() {
	// [3]string is an array; []string is a slice.
	// A slice describes a portion of an underlying array. Its length can
	// change when we assign a new slice result, such as one from append.
	clubs := [3]string{"Arsenal", "Chelsea", "Liverpool"}
	players := []string{"Saka", "Palmer", "Salah"}
	fmt.Println(clubs)

	// Indexing and element assignment work as with arrays.
	fmt.Println(players[0]) // Saka
	players[1] = "Pedri"
	fmt.Println(players) // [Saka Pedri Salah]

	// append returns a slice containing the added elements. Store the result.
	// It may reuse the underlying array or allocate a larger one.
	players = append(players, "Mbappe")
	fmt.Println(players)
	players = append(players, "Osimhen", "Lookman")

	// range supplies the index first and the element value second.
	// With only one variable, range over an array/slice supplies the index.
	footballers := []string{"Saka", "Salah", "Osimhen"}
	for index, name := range footballers {
		fmt.Println(index, name)
	}

	deliveryEarnings := []int{1500, 2000, 1200, 1800}
	// _ discards the index so this loop keeps only each earning.
	for _, item := range deliveryEarnings {
		fmt.Println(item)
	}
	// Practice still to complete: accumulate the earnings in a variable
	// declared before the loop, then print the total after the loop.
}

// learningMap demonstrates key/value storage and checking key presence.
func learningMap() {
	// map[string]int maps string keys to integer values.
	riderDeliveries := map[string]int{
		"Emmanuel": 5,
		"Joker":    3,
	}

	// The same syntax adds a missing key or updates an existing key.
	// Keys are case-sensitive: "Joker" and "joker" are different keys.
	riderDeliveries["Chidi"] = 4
	riderDeliveries["Joker"] = 7
	fmt.Println(riderDeliveries["Emmanuel"]) // 5
	fmt.Println(riderDeliveries)
	// Maps have no guaranteed iteration order. Do not rely on display order
	// to imply insertion order or to choose which entry is first.

	// delete removes a key and its value. A missing key is harmless.
	delete(riderDeliveries, "Emmanuel")
	fmt.Println(riderDeliveries)

	product := map[string]int{
		"Laptop": 4,
		"Phone":  10,
	}
	product["Keyboard"] = 6
	product["Phone"] = 8

	// A missing key returns the value type's zero value, here 0.
	// The optional second result tells us whether the key actually exists.
	// Missing Tablet: stock=0, exists=false.
	// Tablet present with zero stock: stock=0, exists=true.
	stock, exists := product["Tablet"]
	fmt.Println("Stock", stock)
	fmt.Println("Exists", exists)
}

// learningStruct demonstrates grouping named fields of different types.
func learningStruct() {
	// type defines a new named type; struct lists its fields and their types.
	// This local type is visible only within this function's scope.
	// Put a type at package level when multiple functions need to name it.
	type Rider struct {
		Name     string
		Age      int
		IsOnline bool
	}

	// A struct literal creates a value. Named fields make assignments clear.
	// Fields omitted from a literal receive their types' zero values.
	rider := Rider{
		Name:     "Emmanuel",
		Age:      24,
		IsOnline: true,
	}

	// Dot notation accesses a field on a struct value.
	fmt.Println(rider.Name)
	fmt.Println(rider.Age)
	fmt.Println(rider.IsOnline)

	type Product struct {
		Name  string
		Price float64
		stock int
	}
	// A capitalized field such as Name is exported; lowercase stock is
	// unexported and cannot be accessed directly from another package.
	// stock works here. Keep its capitalization consistent when accessing it.
	product := Product{
		Name:  "ps 5",
		Price: 1000.0,
		stock: 5,
	}
	fmt.Println(product.Name, "and price", product.Price)

	type FootballPlayer struct {
		Name  string
		Goals int
	}

	players := []FootballPlayer{
		{Name: "Saka", Goals: 12},
		{Name: "Salah", Goals: 18},
	}

	players[0].Goals = 13

	fmt.Println(players)

	for _, player := range players {
		fmt.Println(player.Name, player.Goals)
	}

	for index := range players {
		players[index].Goals = players[index].Goals + 1
	}

	type DeliveryRider struct {
		Name                string
		CompletedDeliveries int
	}

	leanringRiderSlices := []DeliveryRider{
		{Name: "Emma", CompletedDeliveries: 5},
		{Name: "Joker", CompletedDeliveries: 2},
		{Name: "Vic", CompletedDeliveries: 4},
	}

	for _, item := range leanringRiderSlices {
		fmt.Println("the rider names", item.Name)
	}

	fmt.Println(leanringRiderSlices)

}

type Driver struct {
	Name     string
	IsOnline bool
}

func learningPointer() {
	score := 10
	scorePointer := &score // use & to get the address,This points to the address

	fmt.Println(*scorePointer) // User * to get the value from the address

	*scorePointer = 15 // you can also change the original value through the pointer

	fmt.Println(score)

	updateScore(&score)

	fmt.Println(score)

	finishedDeliveries := 5

	increaseDeliveries(&finishedDeliveries)

	fmt.Println("The updated value", finishedDeliveries)

	driver := Driver{
		Name:     "Emmanuel",
		IsOnline: false,
	}

	driver2 := Driver{
		Name:     "Joker",
		IsOnline: true,
	}

	driver2.rename("Emma")

	fmt.Println("After renaming:", driver2)

	setOnlineStat(&driver)

	driver.SetOffline()

	fmt.Println(driver)

}

func (driver *Driver) rename(newName string) {
	driver.Name = newName
}

func setOnlineStat(driver *Driver) {
	(*driver).IsOnline = true
}

func updateScore(scorePointer *int) {
	*scorePointer = 20
}

func increaseDeliveries(num *int) {
	*num += 1
}

func (driver *Driver) SetOffline() {
	driver.IsOnline = false
}

// Learning Interfaces
// An interface describes methods a type must have

type Introducer interface {
	Introduce() string
}

type Customer struct {
	Name string
}

type ShopOwner struct {
	Name string
}

func (shopOwner ShopOwner) Introduce() string {
	return "Shop Owner: " + shopOwner.Name
}

func (customer Customer) Introduce() string {
	return "Customer: " + customer.Name
}

func space() {
	fmt.Println()
	fmt.Println()
}

func validateRiderAge(age int) error {
	if age < 18 {
		return fmt.Errorf("rider must be at least 18")
	}
	return nil
}

type Team struct {
	Name string
}

func CreateTeam(name string) (Team, error) {

	if name == "" {
		return Team{}, fmt.Errorf("team name is required")
	}

	return Team{Name: name}, nil
}

type RiderAccount struct {
	Name string
	Age  int
}

func registerRider(name string, age int) (RiderAccount, error) {
	if name == "" {
		return RiderAccount{}, fmt.Errorf("Riders name is required")

	}

	err := validateRiderAge(age)

	if err != nil {
		return RiderAccount{}, err
	}

	return RiderAccount{name, age}, nil

}
