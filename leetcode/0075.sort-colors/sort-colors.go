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
	//one,two
}
