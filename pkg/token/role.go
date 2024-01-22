package token

func IsPlayer(role int) bool {
	if role==0 {
		return true
	}else {
		return false
	}
}


func IsJudge(role int) bool {
	if role==1 {
		return true
	}else {
		return false
	}
}


func IsManager(role int) bool {
	if role==2 {
		return true
	}else {
		return false
	}
}