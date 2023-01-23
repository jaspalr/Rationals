package main 

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"

)

type Floater64 interface {
    // Converts a value to an equivalent float64.
    toFloat64() float64
}

type Rationalizer interface {

    // Rationalizers implement the standard Stringer interface.
    fmt.Stringer

    //  Rationalizers implement the Floater64 interface.
    Floater64

    // Returns the numerator.
    Numerator() int

    // Returns the denominator.
    Denominator() int

    // Returns the numerator, denominator.
    Split() (int, int)

    // Returns true iff this value equals other.
    Equal(other Rationalizer) bool

    // Returns true iff this value is less than other.
    LessThan(other Rationalizer) bool

    // Returns true iff the value equal an integer.
    IsInt() bool

    //  Returns the sum of this value with other.
    Add(other Rationalizer) Rationalizer

    //  Returns the product of this value with other.
    Multiply(other Rationalizer) Rationalizer

    //  Returns the quotient of this value with other. The error is nil 
    // if its is successful, and a non-nil if it cannot be divided.
    Divide(other Rationalizer) (Rationalizer, error)

    // Returns the reciprocal. The error is nil if it is successful, 
    // and non-nil if it cannot be inverted.
    Invert() (Rationalizer, error)

    // Returns an equal value in lowest terms.
    ToLowestTerms() Rationalizer
} // Rationalizer interface


type Rational struct{
	numerator int
	denominator int
}

//Makes a rational function
func makeRational(a,b int) (Rational, error){
	if(b == 0){
		return Rational{a,b} ,errors.New("denomintor is zero")
	}
	return Rational{a,b}, nil
}

//Returns the numerator
func (r Rational) Numerator() int{
	return r.numerator
}
//Returns the denominator
func(r Rational) Denominator() int{
	return r.denominator
}
//Returns the numerator, denominator
func(r Rational) Split() (int, int){
	return r.numerator,r.denominator
}
//Returns a float64 representation of the rational 
func(r Rational) toFloat64() float64{
	return float64(r.numerator)/float64(r.denominator)
}
//Returns a string representation of the rational 
func(r Rational) String() string{
	if(r.denominator < 0 && r.numerator > 0){
		r.denominator = r.denominator * -1
		return fmt.Sprintf("-%v/%v", r.numerator,r.denominator)
	} else if r.denominator < 0 && r.numerator < 0{
		r.numerator = r.numerator *-1
		r.denominator = r.denominator * -1
		return fmt.Sprintf("%v/%v", r.numerator,r.denominator)
	} else{
		return fmt.Sprintf("%v/%v", r.numerator,r.denominator)
	}
}
//Checks of a rational is equal to another rational
func(r Rational) Equal(other Rational) bool{
	return (r.toFloat64() == other.toFloat64())
}
//Checks of a rational is less than to another rational
func(r Rational) LessThan(other Rational) bool{
	return (r.toFloat64() < other.toFloat64())
}
//Checks of a rational is a integer
func(r Rational) IsInt() bool{
	dec := r.toFloat64()
	return(dec == float64(int(dec)))
}
//Adds two rationals
func(r Rational) Add(other Rational) Rational{
	num := r.numerator*other.denominator + other.numerator*r.denominator
	return(Rational{num,r.denominator*other.denominator})
} 
//Multplies two rationals
func(r Rational) Multiply(other Rational) Rational{
	return(Rational{r.numerator*other.numerator, r.denominator*other.denominator})
}
//Divides two rationals
func(r Rational) Divide(other Rational) (Rational, error){
	if other.numerator == 0{
		return  Rational{0,0}, errors.New("Dividing by zero") 
	} else{
		return Rational{r.numerator*other.denominator,r.denominator*other.numerator},nil
	}
}
//Invents a rational
func(r Rational) Invert() (Rational, error){
	if r.numerator == 0{
		return Rational{0,0}, errors.New("denomintor is zero")
	}else{
		return Rational{r.denominator,r.numerator}, nil
	}
}
//helper for 14
func GCD( r Rational) int{
	a:= r.numerator
	b:= r.denominator
	if r.denominator > r.numerator{
		a = r.denominator
		b = r.numerator
	}
	mod := a%b
	for mod != 0{
		a = b
		b = mod
		mod = a%b
	}
	return b
}
//Converts a rational to lowest terms
func(r Rational) ToLowestTerms() Rational{
	
	neg := false
	if r.numerator  < 0 && r.denominator > 0{
		neg = true
		r.numerator = r.numerator * -1
	} else if r.numerator > 0 && r.denominator < 0{
		neg = true
		r.denominator= r.denominator * -1
	} else if r.numerator < 0 && r.denominator < 0 {
		r.numerator = r.numerator * -1
		r.denominator = r.denominator * -1
	}
	gdc := GCD(r)
	for gdc > 1{
		r.numerator = r.numerator / gdc
		r.denominator = r.denominator /gdc
		gdc = GCD(r) 
	}
	if neg {
		r.numerator = r.numerator * -1
	}
	return r
}
//Returns the harmonic sum of a number
func HarmonicSum(n int) Rational{
	sum:= Rational{1,1}
	for i:= 2; i <=n; i++{
		sum = sum.Add(Rational{1,i})
	}
	return sum
}

