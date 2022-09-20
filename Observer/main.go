package main

import "fmt"

type Observer interface {
	handleEvent([]string)
}

type Person struct {
	name string
}

func (p *Person) Person(name string) {
	p.name = name
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear ", p.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}

type Observable interface {
	subscribe(Observer)
	unsubscribe(Observer)
	sendAll()
}

type JobSite struct {
	subscribers []Person
	vacancies   []string
}

func (j *JobSite) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) removeVacancy(vacancy string) {
	j.vacancies = removeFromSlice(j.vacancies, vacancy)
	fmt.Println(vacancy, " removed")
	j.sendAll()
}

func (j *JobSite) subscribe(p Person) {
	j.subscribers = append(j.subscribers, p)
	fmt.Println(p, " subscribed")
}

func (j *JobSite) unsubscribe(p Person) {
	fmt.Println(p, " unsubscribed")

}

func (j *JobSite) sendAll() {
	for _, l := range j.subscribers {
		l.handleEvent(j.vacancies)
	}
}

func removeFromSlice(slice []string, item string) []string {

	var sliceRemoved []string
	for _, k := range slice {
		if k != item {
			sliceRemoved = append(sliceRemoved, k)
		}

	}

	return sliceRemoved

	// imagine that it removes item ...
}

func main() {
	hh := JobSite{}
	hh.addVacancy("Node js backend junior")

	arman := Person{name: "arman"}
	hh.subscribe(arman)

	hh.addVacancy("Something else")
	hh.removeVacancy("Node js backend junior")
}
