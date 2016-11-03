package lexer

// Mumber 解析数字类型到字符串
func (r *Lexer) Mumber() (string, bool) {
	if r.token.kind == tokenUndef && r.Ok() {
		r.fetchToken()
	}
	if !r.Ok() || r.token.kind != tokenNumber {
		r.errInvalidToken("number")
		return "", false

	}
	ret := bytesToStr(r.token.byteValue)
	r.consume()
	return ret, true
}

// Boolean 解析 bool 类型到字符串
func (r *Lexer) Boolean() (string, bool) {
	if r.token.kind == tokenUndef && r.Ok() {
		r.fetchToken()
	}
	if !r.Ok() || r.token.kind != tokenBool {
		r.errInvalidToken("bool")
		return "", false

	}
	ret := r.token.boolValue
	r.consume()
	if ret {
		return "true", true
	}
	return "false", true
}

// Str 解析字符串格式
func (r *Lexer) Str() (string, bool) {
	if r.token.kind == tokenUndef && r.Ok() {
		r.fetchToken()
	}
	if !r.Ok() || r.token.kind != tokenString {
		r.errInvalidToken("string")
		return "", false

	}
	ret := string(r.token.byteValue)
	r.consume()
	return ret, true
}
