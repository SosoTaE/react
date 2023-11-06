// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"
import "time"
import "github.com/google/uuid"


func main() {
  		// Generate a UUID
	uuid := uuid.New()

	// Get the current timestamp in nanoseconds
	timestamp := time.Now().UnixNano()
 
	// Combine the UUID and timestamp to create a unique ID
	uniqueID := fmt.Sprintf("%s-%d", uuid, timestamp)

	filename := fmt.Sprintf("%s.%s", uniqueID, "png")
  	fmt.Println(filename)

}