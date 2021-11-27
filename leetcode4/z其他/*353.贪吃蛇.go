type snake struct {
	body [][]int
	score int
	height int 
	width  int
	food [][]int
}


func constuctor(food [][]int, height, width int) snake {
	return snake{
		body: [][]int{[0,0]}
		score: 0
		height: height
		width: width
		food: food
	}
}

func (this *snake)move(direction string) int {
	head:= this.body[0]
	tail:= this.body[len(this.body)-1]
	this.body = this.body[:len(this.body)-1]
	switch direction {
		case "U": head[0] -=1
		case "D": head[0] +=1
	    case "R": head[1] +=1
		case "L": head[1] -=1
	}
	if isContain(this.body, head) || head[0] <0 || head[0] >= this.height || 
	head[1]<0 || head[1] >=this.width {
		return -1
	}
	this.body = append([][]int{head}, this.body...)

	if len(this.food) != 0 && this.food[0] == head {
		this.food = this[1:]
		this.body = append(&this.body, tail)
		this.score++
	}
	return &this.score
}

func isContain(body [][]int, head []int) bool {
	for _, v:= range body {
		if head[0] = v[0] && head[1] == v[1] {
			return true
		}
	}
	return false
}