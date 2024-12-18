package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	} else {
		return d.priorityLevel
	}
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	//1 way is to write the logic for each notification type here with type switch
	//2nd way is to add a function to the notification interface, and simply call it
	//But let's assume that unknown types can be passed, as per the assignment, so we use the 1st way
	//Tried using without t, but noted that n doesn't know about its implementing types,
	//hence the type's components are unknown to it.
	switch t := n.(type) {
	case directMessage:
		return t.senderUsername, t.importance()
	case groupMessage:
		return t.groupName, t.importance()
	case systemAlert:
		return t.alertCode, t.importance()
	default:
		return "", 0
	}
}
