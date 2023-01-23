#lang racket
;Makes a rational function
(define (make-rational . args)
  (cond [(= (length args) 1)  (list 'rational (first args) 1)]
        [(and ( = (length args) 2)  (= (first(rest args))  0)) (raise "invalid denominator")]
    [(length args) = 2 (list 'rational (first args) (first (rest args)))]))

;Returns the numerator
(define (r-numerator x)
  (first(rest x)))
;Returns the denominator
(define (r-denominator x)
  (first(rest(rest x))))
;Returns the numerator, denominator
(define (num-denom x)
  (rest x))
;Converts a rational a string
(define (to-string x)
  (string-append (~s(r-numerator x)) "/" (~s(r-denominator x))))
;Converts a rational a float
(define(to-float x)
  (exact->inexact (/ (r-numerator x)  (r-denominator x))))
;Checks of a rational is equal to another rational
(define (r= x y)
  [if (= (/ (r-numerator x)  (r-denominator x)) (/(r-numerator y)  (r-denominator y))) #t #f])
;Checks if a rational is less than another rational
(define (r< x y)
  [if (< (/ (r-numerator x)  (r-denominator x)) (/(r-numerator y)  (r-denominator y))) #t #f])
;Checks of a rational is a integer
(define (is-int? x)
  (integer? (/ (r-numerator x)  (r-denominator x))))
;Adds two rationals
(define (r+ x y) 
  (make-rational (+(* (r-numerator x)  (r-denominator y))  ( * (r-numerator y) (r-denominator x))) ( * (r-denominator x) (r-denominator y))))
;Multplies two rationals
(define (r* x y)
  (make-rational (* (r-numerator x) (r-numerator y) ) (* (r-denominator x) (r-denominator y) ))) 
;Divides two rationals
(define (r/ x y)
  (make-rational (* (r-numerator x) (r-denominator y) ) (* (r-denominator x) (r-numerator y) )))
;Inverts a rational
(define (invert x)
  (make-rational (r-denominator x) (r-numerator x )))
;helper for to-lowest-terms
(define (gcd x y)
  (if (= y 0) x (gcd y (remainder x  y ))))
;Converts a rational to lowest terms
(define (to-lowest-terms x)
  (make-rational (/(r-numerator x) (gcd (r-numerator x) (r-denominator x) )) (/(r-denominator x) (gcd (r-numerator x) (r-denominator x) )) ))
;Returns the harmonic sum of a number
(define (harmonic-sum x)
  (if (= x 1) (make-rational 1 1)
       (r+(make-rational 1 x) (harmonic-sum (- x 1)))))




        


