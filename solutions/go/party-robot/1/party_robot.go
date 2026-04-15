package partyrobot
import ("fmt"
        "strconv")
// Welcome greets a person by name.
func Welcome(name string) string {
	s := "Welcome to my party, " + name + "!"
    return s
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    ag := strconv.Itoa(age)
	str := "Happy birthday " + name + "! "  + "You are now " + ag + " years old!" 
    return str 
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    // Sprintf allows us to format the table as 3 digits (%03d) 
    // and the distance with one decimal point (%.1f)
    return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", 
        name, table, direction, distance, neighbor)
}

