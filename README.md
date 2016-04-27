# People of Go
Event-based people simulation, written in go


People Of Go allow you to simulate randomized people behaviour. Each simulated person triggers events depending on its habits.
## What is it for

People Of Go is primarily meant for company day-to-day life simulation. Events are triggered by either simulated users, cutomers or employees.
The generated data then can be used to test information systems (anomaly detector for example).

## Event based simulation
An event is simply an abstract set of data, plus a trigger time. It can represent absolutely any kind of event : sending a mail, buying a car and so one.

Usually, an event is triggered by a Person. It can also be labeled as belonging to a category : 
#### Example
Sending a mail to my colleague Pog corresponds to creating the following event :
```
Event{
  Family: "send email",
  Date: 27/04/2016 3pm
  Author: EricBurel,
  Data : {
    "message":"Hi Pog, how are you ?"
    "length":42
    "to":"pog@pog.pog"
    "files":nil
}
```
People of Go aims at helping you to generate this kind of event randomly, in order to benchmark applications or to generate load tests.

## Habit

An habit is a way of generating a set of events.

## Person

A person is an entity that possesses personnal data and habits.

## Contact

Any question ? Contact @eric-burel !

