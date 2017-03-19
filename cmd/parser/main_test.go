package main

import (
    "time"
    "testing"
)

func TestIsNewEvent(t *testing.T) {
    emptyTitle := ""
    incorrectTitle := "abc"
    correctTitle := "прием в Ванкувере"
    oldDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    futureDate := time.Date(2017, time.May, 10, 23, 0, 0, 0, time.UTC)

    result1 := isNewEvent(emptyTitle, oldDate)
    if result1 == true {
        t.Fail()
    }

    result2 := isNewEvent(incorrectTitle, oldDate)
    if result2 == true {
        t.Fail()
    }

    result3 := isNewEvent(incorrectTitle, futureDate)
    if result3 == true {
        t.Fail()
    }

    result4 := isNewEvent(correctTitle, oldDate)
    if result4 == true {
        t.Fail()
    }

    result5 := isNewEvent(correctTitle, futureDate)
    if result5 == false {
        t.Fail()
    }
}