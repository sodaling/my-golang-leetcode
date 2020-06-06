package problem0075

func sortColors2(a []int) {
	zeroNum :=0
	oneNum :=0
	twoNum :=0
	for _,i := range a{
		if i ==0{
			zeroNum++
		}else if i ==1{
			oneNum++
		}else if i==2{
			twoNum++
		}
	}
	for i:=0;i<zeroNum;i++{
		a[i] =0
	}
	for i :=zeroNum;i<zeroNum+oneNum;i++{
		a[i]=1
	}
	for i:=zeroNum+oneNum;i<zeroNum+oneNum+twoNum;i++{
		a[i]=2
	}
}

func sortColors(a []int) {
	//三路快排
	//三个都是包前不包后
	//原版用的是case
	//至于哪个快For just a few items, the difference is small. If you have many items you should definitely use a switch.
	//If a switch contains more than five items, it's implemented using a lookup table or a hash list. This means that all items get the same access time, compared to a list of if:s where the last item takes much more time to reach as it has to evaluate every previous condition first.
	var left = -1
	var right = len(a)
	var i = 0
	for i<right{
		if a[i]==1{
			i++
		}else if a[i]==0{
			a[left+1],a[i] = a[i],a[left+1]
			left++
			i++
		}else if a[i]==2 {
			a[right-1],a[i] = a[i],a[right-1]
			right--
		}else {panic("wrong input")}
	}
}
