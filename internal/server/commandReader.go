package server

func ReadCommand() string {
	switch CommandReader() {
	case "add":
		return Add()
	case "update":
		return Update()
	case "delete":
		return Delete()
	case "mark-in-progress":
		return MarkInProgress()
	case "mark-done":
		return MarkDone()
	default:
		return "Invalid command"
	}
}
