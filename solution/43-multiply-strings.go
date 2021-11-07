package solution

func Multiply(num1 string, num2 string) string {
	sum:=make([]byte,len(num1)+len(num2))
	for i:=len(num1)-1;i>=0;i--{
		digit:=num1[i]-'0'
		for j:=len(num2)-1;j>=0;j--{
			tmp:=(num2[j]-'0')*digit+sum[i+j+1]
			sum[i+j+1]=tmp%10
			sum[i+j]+=tmp/10

		}
	}
	i:=-1
	for j:=range sum {
		if sum[j]!=0 && i==-1{
			i=j
		}
		sum[j]+='0'
	}
	if i==-1{
		return "0"
	}
	return string(sum[i:])
}
