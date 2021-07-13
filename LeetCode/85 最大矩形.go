package LeetCode



//给出0和1的数字1矩阵求出其中最大矩形
func maximalRectangle(matrix [][]byte) int {
	//上题单调栈扩展，多层每层进行计算
	if len(matrix)==0{
		return 0
	}
	res:=0
	arr:=make([]int,len(matrix[0])) //每行依据上一行结果进行计算
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if matrix[i][j]=='0'{
				arr[j] = 0
			}else{
				arr[j]+=1
			}
		}
		res = max(res,MonotoneStack(arr))
	}
	return res
}


//计算柱状图最大面积
func MonotoneStack(heights []int)int{
	res:=0
	if len(heights)==0{
		return 0
	}
	if len(heights)==1{
		return heights[0]
	}
	stack:=make([]int,0)
	ret:=append(heights,0)
	for i:=0;i<len(ret);i++{
		for len(stack)!=0&&ret[i]<ret[stack[len(stack)-1]]{
			high:=ret[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			area:=0
			if len(stack)==0{
				area = i*high
			}else{
				area = (i-stack[len(stack)-1]-1)*high
			}
			res = max(res,area)
		}
		stack = append(stack,i)
	}
	return  res
}

func max(x,y int)int  {
	if x>y{
		return x
	}else{
		return y
	}
}
