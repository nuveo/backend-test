package workflow

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type JwtToken struct {
// 	Token string `json:"token"`
// }

type Exception struct {
	message string `json:"message"`
}

// Product represents an workflow item
type Workflow struct {
	Id int `json:"id"`
	// status []string `json:"status"`
	// data   []string `json:"data"`
	// steps  []string `json:"steps"`
}

// Workflows is an array of Workflow object
type Workflows []Workflow
