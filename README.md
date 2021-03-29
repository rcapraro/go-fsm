# Finite state machine (FSM) in Golang

## Gof version
following the GoF State pattern: https://en.wikipedia.org/wiki/State_pattern

## Hand-made FSM
Based on triggers and rules (a rule is map whose key is a State and values are an array of pairs of Trigger and State)

## Switch-Based FSM
Based on Go switch statement

## Complex FSM
This version allows the execution of events with a context (ie a struct containing some data to passe to the next State)
