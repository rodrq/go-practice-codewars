/* Create a method each_cons that accepts a list and a number n,
and returns cascading subsets of the list of size n, like so:

each_cons([1,2,3,4], 2)
  #=> [[1,2], [2,3], [3,4]]

each_cons([1,2,3,4], 3)
  #=> [[1,2,3],[2,3,4]]


As you can see, the lists are cascading; ie, they overlap, but never out of order. */

package kata

func EachCons(arr []int, n int) [][]int {
	var res [][]int
	for i := 0; i <= len(arr)-n; i++ {
		res = append(res, arr[i:i+n])
	}
	return res
}

// len(arr)-n == i+n <= len(arr) y se puede definir variable res en declaracion de funcion con (res [][]int)

func EachConsBestPractice(arr []int, n int) (res [][]int) {
	for i := 0; i+n <= len(arr); i++ {
		res = append(res, arr[i:i+n])
	}
	return
}
