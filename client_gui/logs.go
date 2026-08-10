package main


import ()


func set_log() {
	log.RemoveAll()

	log.Add(log_text)
	log.Refresh()
}