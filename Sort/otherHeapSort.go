package main

import "fmt"

func HeapSort(s []int) {
	N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	//下沉排序
	for N >=1 {
		swap(s, 1, N) //将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func main()  {
	s := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	HeapSort(s)
	fmt.Println(s[1:])
}