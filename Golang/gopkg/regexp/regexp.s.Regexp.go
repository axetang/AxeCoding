/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: regexp
 **Element: regexp.Regexp
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Regexp struct {
     // contains filtered or unexported fields
 }
 func (re *Regexp) Copy() *Regexp
 func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
 func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
 func (re *Regexp) Find(b []byte) []byte
 func (re *Regexp) FindAll(b []byte, n int) [][]byte
 func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
 func (re *Regexp) FindAllString(s string, n int) []string
 func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
 func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
 func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
 func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
 func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
 func (re *Regexp) FindIndex(b []byte) (loc []int)
 func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
 func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
 func (re *Regexp) FindString(s string) string
 func (re *Regexp) FindStringIndex(s string) (loc []int)
 func (re *Regexp) FindStringSubmatch(s string) []string
 func (re *Regexp) FindStringSubmatchIndex(s string) []int
 func (re *Regexp) FindSubmatch(b []byte) [][]byte
 func (re *Regexp) FindSubmatchIndex(b []byte) []int
 func (re *Regexp) LiteralPrefix() (prefix string, complete bool)
 func (re *Regexp) Longest()
 func (re *Regexp) Match(b []byte) bool
 func (re *Regexp) MatchReader(r io.RuneReader) bool
 func (re *Regexp) MatchString(s string) bool
 func (re *Regexp) NumSubexp() int
 func (re *Regexp) ReplaceAll(src, repl []byte) []byte
 func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
 func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
 func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
 func (re *Regexp) ReplaceAllString(src, repl string) string
 func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
 func (re *Regexp) Split(s string, n int) []string
 func (re *Regexp) String() string
 func (re *Regexp) SubexpNames() []string
 ------------------------------------------------------------------------------------
 **Description:
 Regexp is the representation of a compiled regular expression. A Regexp is safe for
 concurrent use by multiple goroutines, except for configuration methods, such as
 Longest.
 Copy returns a new Regexp object copied from re. Calling Longest on one copy does
 not affect another.
 Deprecated: In earlier releases, when using a Regexp in multiple goroutines, giving
 each goroutine its own copy helped to avoid lock contention. As of Go 1.12, using
 Copy is no longer necessary to avoid lock contention. Copy may still be appropriate
 if the reason for its use is to make two copies with different Longest settings.
 Expand appends template to dst and returns the result; during the append, Expand
 replaces variables in the template with corresponding matches drawn from src. The
 match slice should have been returned by FindSubmatchIndex.
 In the template, a variable is denoted by a substring of the form $name or ${name},
 where name is a non-empty sequence of letters, digits, and underscores. A purely
 numeric name like $1 refers to the submatch with the corresponding index; other
 names refer to capturing parentheses named with the (?P<name>...) syntax. A
 reference to an out of range or unmatched index or a name that is not present in the
 regular expression is replaced with an empty slice.
 In the $name form, name is taken to be as long as possible: $1x is equivalent to
 ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
 To insert a literal $ in the output, use $$ in the template.
 ExpandString is like Expand but the template and source are strings. It appends to
 and returns a byte slice in order to give the calling code control over allocation.
 Find returns a slice holding the text of the leftmost match in b of the regular
 expression. A return value of nil indicates no match.
 FindAll is the 'All' version of Find; it returns a slice of all successive matches
 of the expression, as defined by the 'All' description in the package comment. A
 return value of nil indicates no match.
 FindAllIndex is the 'All' version of FindIndex; it returns a slice of all successive
 matches of the expression, as defined by the 'All' description in the package comment.
 A return value of nil indicates no match.
 FindAllString is the 'All' version of FindString; it returns a slice of all
 successive matches of the expression, as defined by the 'All' description in the
 package comment. A return value of nil indicates no match.
 FindAllStringIndex is the 'All' version of FindStringIndex; it returns a slice of
 all successive matches of the expression, as defined by the 'All' description in
 the package comment. A return value of nil indicates no match.
 FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it returns a
 slice of all successive matches of the expression, as defined by the 'All'
 description in the package comment. A return value of nil indicates no match.
 FindAllStringSubmatchIndex is the 'All' version of FindStringSubmatchIndex; it
 returns a slice of all successive matches of the expression, as defined by the
 'All' description in the package comment. A return value of nil indicates no match.
 FindAllSubmatch is the 'All' version of FindSubmatch; it returns a slice of all
 successive matches of the expression, as defined by the 'All' description in the
 package comment. A return value of nil indicates no match.
 FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex; it returns a slice
 of all successive matches of the expression, as defined by the 'All' description in
 the package comment. A return value of nil indicates no match.
 FindIndex returns a two-element slice of integers defining the location of the
 leftmost match in b of the regular expression. The match itself is at
 b[loc[0]:loc[1]]. A return value of nil indicates no match.
 FindReaderIndex returns a two-element slice of integers defining the location of the
 leftmost match of the regular expression in text read from the RuneReader. The match
 text was found in the input stream at byte offset loc[0] through loc[1]-1. A return
 value of nil indicates no match.
 FindReaderSubmatchIndex returns a slice holding the index pairs identifying the
 leftmost match of the regular expression of text read by the RuneReader, and the
 matches, if any, of its subexpressions, as defined by the 'Submatch' and 'Index'
 descriptions in the package comment. A return value of nil indicates no match.
 FindString returns a string holding the text of the leftmost match in s of the
 regular expression. If there is no match, the return value is an empty string, but
 it will also be empty if the regular expression successfully matches an empty string.
 Use FindStringIndex or FindStringSubmatch if it is necessary to distinguish these
 cases.
 FindStringIndex returns a two-element slice of integers defining the location of the
 leftmost match in s of the regular expression. The match itself is at s[loc[0]:loc[1]].
 A return value of nil indicates no match.
 FindStringSubmatch returns a slice of strings holding the text of the leftmost match
 of the regular expression in s and the matches, if any, of its subexpressions, as
 defined by the 'Submatch' description in the package comment. A return value of nil
 indicates no match.
 FindStringSubmatchIndex returns a slice holding the index pairs identifying the
 leftmost match of the regular expression in s and the matches, if any, of its
 subexpressions, as defined by the 'Submatch' and 'Index' descriptions in the package
 comment. A return value of nil indicates no match.
 FindSubmatch returns a slice of slices holding the text of the leftmost match of
 the regular expression in b and the matches, if any, of its subexpressions, as
 defined by the 'Submatch' descriptions in the package comment. A return value of
 nil indicates no match.
 FindSubmatchIndex returns a slice holding the index pairs identifying the leftmost
 match of the regular expression in b and the matches, if any, of its subexpressions,
 as defined by the 'Submatch' and 'Index' descriptions in the package comment. A
 return value of nil indicates no match.
 LiteralPrefix returns a literal string that must begin any match of the regular
 expression re. It returns the boolean true if the literal string comprises the
 entire regular expression.
 Longest makes future searches prefer the leftmost-longest match. That is, when
 matching against text, the regexp returns a match that begins as early as possible
 in the input (leftmost), and among those it chooses a match that is as long as
 possible. This method modifies the Regexp and may not be called concurrently with
 any other methods.
 Match reports whether the byte slice b contains any match of the regular expression
 re.
 MatchReader reports whether the text returned by the RuneReader contains any match
 of the regular expression re.
 MatchString reports whether the string s contains any match of the regular
 expression re.
 NumSubexp returns the number of parenthesized subexpressions in this Regexp.
 ReplaceAll returns a copy of src, replacing matches of the Regexp with the
 replacement text repl. Inside repl, $ signs are interpreted as in Expand, so for
 instance $1 represents the text of the first submatch.
 ReplaceAllFunc returns a copy of src in which all matches of the Regexp have been
 replaced by the return value of function repl applied to the matched byte slice.
 The replacement returned by repl is substituted directly, without using Expand.
 ReplaceAllLiteral returns a copy of src, replacing matches of the Regexp with the
 replacement bytes repl. The replacement repl is substituted directly, without using
 Expand.
 ReplaceAllLiteralString returns a copy of src, replacing matches of the Regexp with
 the replacement string repl. The replacement repl is substituted directly, without
 using Expand.
 ReplaceAllString returns a copy of src, replacing matches of the Regexp with the
 replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for
 instance $1 represents the text of the first submatch.
 ReplaceAllStringFunc returns a copy of src in which all matches of the Regexp have
 been replaced by the return value of function repl applied to the matched substring.
 The replacement returned by repl is substituted directly, without using Expand.
 Split slices s into substrings separated by the expression and returns a slice of
 the substrings between those expression matches.
 The slice returned by this method consists of all the substrings of s not contained
 in the slice returned by FindAllString. When called on an expression that contains
 no metacharacters, it is equivalent to strings.SplitN.
 Example:
 s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
 // s: ["", "b", "b", "c", "cadaaae"]
 The count determines the number of substrings to return:
 n > 0: at most n substrings; the last substring will be the unsplit remainder.
 n == 0: the result is nil (zero substrings)
 n < 0: all substrings
 String returns the source text used to compile the regular expression.
 SubexpNames returns the names of the parenthesized subexpressions in this Regexp.
 The name for the first sub-expression is names[1], so that if m is a match slice,
 the name for m[i] is SubexpNames()[i]. Since the Regexp as a whole cannot be named,
 names[0] is always the empty string. The slice should not be modified.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
