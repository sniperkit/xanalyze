package token

func (t Type) String() string {
	switch t {
	case Error:
		return "error"
	case EOF:
		return "eof"
	case Comment:
		return "comment"
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	case LeftBracket:
		return "["
	case RightBracket:
		return "]"
	case LeftBrace:
		return "{"
	case RightBrace:
		return "}"
	case Pipe:
		return "|"
	case As:
		return "as"
	case InfixOp:
		return "infix operator"
	case Colon:
		return ":"
	case Assign:
		return "="
	case Comma:
		return ","
	case Arrow:
		return "->"
	case Identifier:
		return "identifier"
	case Op:
		return "operator"
	case String:
		return "string"
	case Int:
		return "integer"
	case Float:
		return "float"
	case Range:
		return ".."
	case Char:
		return "character"
	case True:
		return "True"
	case False:
		return "False"
	case Dot:
		return "."
	case TypeDef:
		return "type"
	case Alias:
		return "alias"
	case If:
		return "if"
	case Then:
		return "then"
	case Else:
		return "else"
	case Of:
		return "of"
	case Case:
		return "case"
	case Infix:
		return "infix"
	case Infixl:
		return "infixl"
	case Infixr:
		return "infixr"
	case Let:
		return "let"
	case In:
		return "in"
	case Module:
		return "module"
	case Exposing:
		return "exposing"
	case Import:
		return "import"
	default:
		return "invalid token"
	}
}
