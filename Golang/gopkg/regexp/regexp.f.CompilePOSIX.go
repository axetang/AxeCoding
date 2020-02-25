/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: regexp
 **Element: regexp.CompilePOSIX
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func CompilePOSIX(expr string) (*Regexp, error)
 ------------------------------------------------------------------------------------
 **Description:
 CompilePOSIX is like Compile but restricts the regular expression to POSIX ERE
 (egrep) syntax and changes the match semantics to leftmost-longest.
 That is, when matching against text, the regexp returns a match that begins as early
 as possible in the input (leftmost), and among those it chooses a match that is as
 long as possible. This so-called leftmost-longest matching is the same semantics
 that early regular expression implementations used and that POSIX specifies.
 However, there can be multiple leftmost-longest matches, with different submatch
 choices, and here this package diverges from POSIX. Among the possible
 leftmost-longest matches, this package chooses the one that a backtracking search
 would have found first, while POSIX specifies that the match be chosen to maximize
 the length of the first subexpression, then the second, and so on from left to right.
 The POSIX rule is computationally prohibitive and not even well-defined. See
 https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
