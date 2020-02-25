/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: tempalte.Template
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Template struct {
     // The underlying template's parse tree, updated to be HTML-safe.
     Tree *parse.Tree
     // contains filtered or unexported fields
 }
 func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
 func (t *Template) Clone() (*Template, error)
 func (t *Template) DefinedTemplates() string
 func (t *Template) Delims(left, right string) *Template
 func (t *Template) Execute(wr io.Writer, data interface{}) error
 func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
 func (t *Template) Funcs(funcMap FuncMap) *Template
 func (t *Template) Lookup(name string) *Template
 func (t *Template) Name() string
 func (t *Template) New(name string) *Template
 func (t *Template) Option(opt ...string) *Template
 func (t *Template) Parse(text string) (*Template, error)
 func (t *Template) ParseFiles(filenames ...string) (*Template, error)
 func (t *Template) ParseGlob(pattern string) (*Template, error)
 func (t *Template) Templates() []*Template
 ------------------------------------------------------------------------------------
 **Description:
 Template is a specialized Template from "text/template" that produces a safe HTML document fragment.
 AddParseTree creates a new template with the name and parse tree and associates it with t.
 It returns an error if t or any associated template has already been executed.
 Clone returns a duplicate of the template, including all associated templates. The actual representation is not copied, but the name space of associated templates is, so further calls to Parse in the copy will add templates to the copy but not to the original. Clone can be used to prepare common templates and use them with variant definitions for other templates by adding the variants after the clone is made.
 It returns an error if t has already been executed.
 DefinedTemplates returns a string listing the defined templates, prefixed by the string "; defined templates are: ". If there are none, it returns the empty string. Used to generate an error message.
 Delims sets the action delimiters to the specified strings, to be used in subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template definitions will inherit the settings. An empty delimiter stands for the corresponding default: {{ or }}. The return value is the template, so calls can be chained.
 Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.
 ExecuteTemplate applies the template associated with t that has the given name to the specified data object and writes the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.
 Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed. It panics if a value in the map is not a function with appropriate return type. However, it is legal to overwrite elements of the map. The return value is the template, so calls can be chained.
 Lookup returns the template with the given name that is associated with t, or nil if there is no such template.
 Name returns the name of the template.
 New allocates a new HTML template associated with the given one and with the same delimiters. The association, which is transitive, allows one template to invoke another with a {{template}} action.
 If a template with the given name already exists, the new HTML template will replace it. The existing template will be reset and disassociated with t.
 Option sets options for the template. Options are described by strings, either a simple string or "key=value". There can be at most one equals sign in an option string. If the option string is unrecognized or otherwise invalid, Option panics.
 Known options:
 missingkey: Control the behavior during execution if a map is indexed with a key that is not present in the map.
 "missingkey=default" or "missingkey=invalid"
 	The default behavior: Do nothing and continue execution.
 	If printed, the result of the index operation is the string
 	"<no value>".
 "missingkey=zero"
 	The operation returns the zero value for the map type's element.
 "missingkey=error"
 	Execution stops immediately with an error.
 Parse parses text as a template body for t. Named template definitions ({{define ...}} or {{block ...}} statements) in text define additional templates associated with t and are removed from the definition of t itself.
 Templates can be redefined in successive calls to Parse, before the first use of Execute on t or any associated template. A template definition with a body containing only white space and comments is considered empty and will not replace an existing template's body. This allows using Parse to add new named template definitions without overwriting the main template body.
 ParseFiles parses the named files and associates the resulting templates with t. If an error occurs, parsing stops and the returned template is nil; otherwise it is t. There must be at least one file.
 When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.
 ParseFiles returns an error if t or any associated template has already been executed.
 ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with t. The pattern is processed by filepath.Glob and must match at least one file. ParseGlob is equivalent to calling t.ParseFiles with the list of files matched by the pattern.
 When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.
 ParseGlob returns an error if t or any associated template has already been executed.
 Templates returns a slice of the templates associated with t, including t itself.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
