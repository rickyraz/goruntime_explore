package errhandle

type User struct{}

// Layer 3: Lowest level
func readDatabase(id int) (string, error) {
	// something fails here
	// return "", errors.New("database connection failed")
	return "ADA", nil
}

// Layer 2: Middle
func getUser(id int) (User, error) {
	data, err := readDatabase(id)
	if err != nil {
		return User{}, err // pass error up
	}

	_ = data       // untuk menghindari "declared and not used"
	user := User{} // buat objek user

	// process data...
	return user, nil
}

// Layer 1: Top level
func HandleRequest() error {
	user, err := getUser(123)
	if err != nil {
		return err // pass error up again
	}

	_ = user // untuk menghindari "declared and not used"
	return nil
}
