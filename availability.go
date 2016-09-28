package main

 import (
         "fmt"
         "time"
 )

 func main() {

         now := time.Now()

         fmt.Println("Today : ", now.Format(time.ANSIC))

                                                        
         nano := now.Nanosecond()                                                                        
         millisec := nano / 1000000                                                                      
         fmt.Println("(wrong)Millisecond : ", millisec)                                                  

                                                          
         unixNano := now.UnixNano()                                                                      
         umillisec := unixNano / 1000000                                                                 
         fmt.Println("(correct)Millisecond : ", umillisec)                                               

 }  
