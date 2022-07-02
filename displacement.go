/*Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))*/

package main

import(
	"fmt"
	"math"
)
func GenDisplaceFn(a,v,d float64) func(float64) float64{
  fn := func (t float64) float64{
      return (a*(math.Pow(t, 2)*0.5) + (v * t) + (d))
}
	return fn
}
func main(){
	var a,v,d,t float64  
	// v = initial velocity, d = initial displacement, a = acceleration, t = time, s = displacement
	fmt.Print("Please enter acceleration:")
	    fmt.Scan(&a)
	fmt.Print("Please enter initial velocity:")
		fmt.Scan(&v)
	fmt.Print("Please enter initial displacement:")
		fmt.Scan(&d)
	fmt.Print("Please enter time:")
		fmt.Scan(&t)

	ft := GenDisplaceFn(a,v,d) // return a function which computes displacement as a function of time
    fmt.Println("Displacement:",ft(t))           //  function to calculate displacement
}