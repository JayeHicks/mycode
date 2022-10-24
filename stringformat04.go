/* string challenge */

package main

import (
    "fmt"
)

func main() {

   const url = "https://example.org:6001/v2/snacks?"

   r :="req=snickers"
   q := "quantity=1" 
   s := "size=king" 

    res := fmt.Sprintf("%s%s&%s&%s", url, r, q, s)
    fmt.Println(res) 
}

