package main


func heart(usr *Player) {
	usr.MAX_HP += 20
	usr.CURRENT_HP += 20
}

func halo(usr *Player) {
	usr.ATK += 1.0
	usr.MAX_HP += 20
	usr.CURRENT_HP += 20
}

