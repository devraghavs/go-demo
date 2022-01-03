// /******************************************************************************
// Welcome to GDB Online.
// GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
// C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
// Code, Compile, Run and Debug online from anywhere in world.

// *******************************************************************************/
package main
import ("fmt"
"strings")

func main() {
    arr:= [10][3]interface{}{
        {"A",10.28,9},
        {"B",24.07,7},
        {"C",13.30,0},
        {"D",28.94,1},
        {"E",12.39,3},
        {"F",30.77,2},
        {"G",55.13,15},
        {"H",50.00,7},
        {"I",90.12,92},
        {"J",82.31,15},
    }
    total := 0.0
    res1 :=0
    // var total interface{};
    // var l [4]string;
    // for i:=0;i<4;i++{
    //     fmt.Scanf("%s", &l[i])
    // }
    // fmt.Println(l)
    // fmt.Println(len(l))
    var flag int;
    fmt.Println(arr[0][1])
    var l string;
    fmt.Scanln(&l)
    fmt.Println(len(l))
    if len(l)!=4{
        fmt.Println("Please enter string length 4 A-J")
    }else{
        for _, r := range l {
        if (r < 'a' || r > 'j') && (r < 'A' || r > 'j') {
            fmt.Println("Not Accepted")
            flag:=1
            fmt.Println(flag)
            break
        }
    }
    if flag<1{
         for i:=range l{
             fmt.Println("#######",)
             for j:=range arr{
             res:=strings.Index(arr[j][0].(string),string(l[i]))
             if res==0{
                 res1:=j
                  fmt.Println("res1",res1,arr[j][1]) 
                  total+=arr[res1][1].(float64)
                  break
             }
            //  fmt.Println("Heyyyy",i,j,res1,res,arr[j][1]) 
            // total=arr[i][1]
            }
            fmt.Println("TEMP TOTAL",res1,arr[res1][1],total) 
            // total+=arr[res1][1].(float64)
            
    }}
    fmt.Println("TOTAL",total)
           
        }
        // for i:=range l{
        //     if strings.ContainsAny(l[i],){
        //         fmt.Println(l[i])
        //     }
            
        // }
    }

// package main
// import "fmt"

// func IsLetter(s string) bool {
//     for _, r := range s {
//         if (r < 'a' || r > 'j') && (r < 'A' || r > 'j') {
//             return false
//         }
//     }
//     return true
// }
// func main() {
//     fmt.Println(IsLetter("abck"))  // true
//     fmt.Println(IsLetter("123 a")) // false

// }
