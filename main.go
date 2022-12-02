package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func main(){
    // read file
  file, _ := ioutil.ReadFile("./inputs.txt")
  calorieList := string(file)
  var list = strings.SplitAfter(strings.TrimSpace(calorieList), "\n")
  var finArray []int
  var sumItems = 0
    for k:=0;  k < len(list); k++ {
        intVal, err := strconv.Atoi(strings.TrimSpace(list[k])); 
        if err != nil{
          log.Printf("error converting %s to int",  reflect.TypeOf(list[k]))
        }
        if intVal == 0{ 
          sumItems = 0 
        }
        sumItems = sumItems + intVal
        finArray = append(finArray, sumItems)
    } 
     // get top 3 
     var maxElems = findMaxElements(finArray)
     top3Sum := topThreeSum(maxElems)

     log.Printf("max calories: %v", max(finArray))
     log.Printf("top 3 total: %v", top3Sum)
  }

func max(array []int) (int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return max
}

func findMaxElements(arr []int) []int {
   max_num := arr[0]
   var list []int
   for i:=0; i<len(arr); i++{
      if arr[i] > max_num {
         max_num = arr[i]
      }
      list = append(list, arr[i])
   }
   return list
}

func topThreeSum(arr []int) int {
   var first = -0
   var second = -0
   var third = -0
   for i:=0; i<len(arr); i++{
      if arr[i] > first {
         third = second
         second = first
         first = arr[i]
      } else if arr[i] > second && arr[i] != first {
          third = second
          second = arr[i]
      } else if  arr[i] > third && arr[i] != second { 
          third = arr[i]
      }
   }
   total := first + second + third
   return total
}
