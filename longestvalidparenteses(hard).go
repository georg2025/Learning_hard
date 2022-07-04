// this function find our lenght of the longest valid parenteses in string.
func longestValidParentheses(s string) int {
    if s=="" || s==")"{return 0}
    n:=strings.Split(s, "")
    f:=0
    u:=0
    for {
        f= numberCount(n)
        n= sliceChange (f, n)
        if f>u {u=f}
        if len (n) == 0 {return u}
    }
return u
}


func numberCount (s []string) int {
    r:=0
    o:=0
    f:=0
    for x,y:=range s {
        if y=="("{
            r+=1
            o+=1
        }else if y==")"{
            r-=1
            o+=1
        }
        if r<0{return x}
        if r==0{f=o}

    }
return f
}


func sliceChange(s int, y []string) []string {
	if s >= len(y) {
		y = y[0:0]
		return y
	}
	y = y[s+1:]
	return y

}
