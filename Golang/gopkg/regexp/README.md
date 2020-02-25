# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package regexp
## import "regexp"

Package regexp implements regular expression search.

The syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages. More precisely, it is the syntax accepted by RE2 and described at https://golang.org/s/re2syntax, except for \C. For an overview of the syntax, run

go doc regexp/syntax
The regexp implementation provided by this package is guaranteed to run in time linear in the size of the input. (This is a property not guaranteed by most open source implementations of regular expressions.) For more information about this property, see

https://swtch.com/~rsc/regexp/regexp1.html
or any book about automata theory.

All characters are UTF-8-encoded code points.

There are 16 methods of Regexp that match a regular expression and identify the matched text. Their names are matched by this regular expression:

Find(All)?(String)?(Submatch)?(Index)?
If 'All' is present, the routine matches successive non-overlapping matches of the entire expression. Empty matches abutting a preceding match are ignored. The return value is a slice containing the successive return values of the corresponding non-'All' routine. These routines take an extra integer argument, n. If n >= 0, the function returns at most n matches/submatches; otherwise, it returns all of them.

If 'String' is present, the argument is a string; otherwise it is a slice of bytes; return values are adjusted as appropriate.

If 'Submatch' is present, the return value is a slice identifying the successive submatches of the expression. Submatches are matches of parenthesized subexpressions (also known as capturing groups) within the regular expression, numbered from left to right in order of opening parenthesis. Submatch 0 is the match of the entire expression, submatch 1 the match of the first parenthesized subexpression, and so on.

If 'Index' is present, matches and submatches are identified by byte index pairs within the input string: result[2*n:2*n+1] identifies the indexes of the nth submatch. The pair for n==0 identifies the match of the entire expression. If 'Index' is not present, the match is identified by the text of the match/submatch. If an index is negative, it means that subexpression did not match any string in the input.

There is also a subset of the methods that can be applied to text read from a RuneReader:

MatchReader, FindReaderIndex, FindReaderSubmatchIndex
This set may grow. Note that regular expression matches may need to examine text beyond the text returned by a match, so the methods that match text from a RuneReader may read arbitrarily far into the input before returning.

(There are a few other methods that do not match this pattern.)


## Index
```
func Match(pattern string, b []byte) (matched bool, err error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
func MatchString(pattern string, s string) (matched bool, err error)
func QuoteMeta(s string) string
type Regexp
func Compile(expr string) (*Regexp, error)
func CompilePOSIX(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
func MustCompilePOSIX(str string) *Regexp
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
```
## Package Files

backtrack.go exec.go onepass.go regexp.go
