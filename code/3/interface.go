// Good for private methods, when you don't want
// to delcare a new interfaces for helper methods
// or internal usage (you use it only once in the code).
package boring

// collect gets an adder node and add to it any fields
// or indexes ahead.
func (p *parser) collect(adder interface {
	add(ast.Node)
}) {
	for {
		switch tok := p.next(); tok.typ {
		case token.Field:
			adder.add(p.newField(tok))
		case token.Index:
			adder.add(p.newIndex(tok))
		default:
			p.backup()
			return
		}
	}
}
