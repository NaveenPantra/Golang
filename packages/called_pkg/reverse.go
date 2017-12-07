package called_pkg

//Reverse() ('R' is capital case so visible to outside the package)
func Reverse() string {
	//calling reverse_two that is inside the package and not visible to outside the package
	//For privacy we are calling the reverse_two
	return reverse_two()
}