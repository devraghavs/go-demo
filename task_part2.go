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
    flag :=0;
    var flag2 int;
    var l string;
    fmt.Scanln(&l)
    // fmt.Println(len(l))
    if len(l)!=4 || !(strings.ContainsAny(l,"a | b | c") && strings.ContainsAny(l,"d | e") && strings.ContainsAny(l,"f | g | h") && strings.ContainsAny(l,"i | j")){
        fmt.Println("Invalid Input:Please enter string length 4 A-J")
    }else{
         for i:=range l{
             for j:=range arr{
             res:=strings.Index(arr[j][0].(string),strings.ToUpper(string(l[i])))
             if res==0{
                 res1:=j
                 if arr[j][2].(int)>0{
                    fmt.Println("res1",res1,arr[j][1]) 
                    total+=arr[res1][1].(float64)
                    // arr[j][2].(int)=(arr[j][2].(int))-1
                    break
                 }else{
                    //   fmt.Println("Error : Cannot create robot")
                      flag2=1
                     break
                 }
                //  if arr[j][2]==0{
                //     
                //  }else{
                //      fmt.Println("res1",res1,arr[j][1]) 
                //     total+=arr[res1][1].(float64)
                //   break
                //  }
                  
             }
            //  fmt.Println("Heyyyy",i,j,res1,res,arr[j][1]) 
            // total=arr[i][1]
            }
            if flag2<1{
				continue
                // fmt.Println("TEMP TOTAL",res1,arr[res1][1],total) 
            }else{
                break
            }
            
            // total+=arr[res1][1].(float64)
            
    }
     if flag2<1 && flag!=1{
         fmt.Println("TOTAL",total)
    }else{
         fmt.Println("Error : Cannot create robot")
    }
        
    }
           
    }