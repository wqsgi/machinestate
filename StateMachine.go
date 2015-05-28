package wei2
import "errors"


type StateMachine struct {

	transitionMap map[string]map[int]*Transition;
}

type Transition struct {
	preType     string;
	currentType string;

	handle      func(event Event);

}

type Event struct {
	EventStatus int;

}


func NewStateMachine() (*StateMachine) {
	m := new(StateMachine);
	m.transitionMap=make(map[string]map[int]*Transition);
	return m;
}


func (m *StateMachine) Add(preState string, curState string, handle func(event Event), eventStatus int)  (err error){
	if &handle == nil{
		err = errors.New("handle is null")
		return err;
	}
	transitionMap := m.transitionMap[preState]
	if transitionMap == nil {
		transitionMap = make(map[int]*Transition)
		m.transitionMap[preState]=transitionMap
	}
	transition := new(Transition)
	transition.currentType=curState
	transition.preType=preState
	transition.handle=handle
	transitionMap[eventStatus]=transition
	return
}


func (m*StateMachine)Transition(preType string, event Event)(err error) {
	transitionMap := m.transitionMap[preType]
	if transitionMap==nil {
		err=errors.New("preType is not exist.")
		return
	}

	transition := transitionMap[event.EventStatus]
	if transition == nil{
		err=errors.New("preType is not exist.")
		return
	}

	transition.handle(event)
	return
}
