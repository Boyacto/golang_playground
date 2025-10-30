grammar Patito;

program
  : PROGRAM ID SEMI varsOpt funcsOpt MAIN body END EOF
  ;

varsOpt
  : varsDecl
  |
  ;

funcsOpt
  : funcDecl*
  ;

varsDecl
  : VARS varGroup+
  ;

varGroup
  : idList COLON typeSpec SEMI
  ;

idList
  : ID (COMMA ID)*
  ;

typeSpec
  : INT_KW
  | FLOAT_KW
  ;

body
  : LBRACE statement* RBRACE
  ;

statement
  : assign
  | condition
  | cycle
  | fcall SEMI
  | printStmt
  | body
  ;

assign
  : ID ASSIGN expr SEMI
  ;

printStmt
  : PRINT LPAREN printArgs RPAREN SEMI
  ;

printArgs
  : expr (COMMA expr)*
  ;

condition
  : IF LPAREN expr RPAREN body (ELSE body)?
  ;

cycle
  : WHILE LPAREN expr RPAREN DO body SEMI
  ;

expr
  : expr RELOP expr
  | expr (PLUS|MINUS) expr
  | expr (MULT|DIV) expr
  | LPAREN expr RPAREN
  | fcall
  | sign? ID
  | sign? literal
  ;

sign
  : PLUS
  | MINUS
  ;

literal
  : INT_LIT
  | FLOAT_LIT
  | STRING_LIT
  ;

funcDecl
  : VOID ID LPAREN paramListOpt RPAREN LBRACE varsOpt body RBRACE SEMI
  ;

paramListOpt
  : paramList
  |
  ;

paramList
  : param (COMMA param)*
  ;

param
  : ID COLON typeSpec
  ;

fcall
  : ID LPAREN argListOpt RPAREN
  ;

argListOpt
  : argList
  |
  ;

argList
  : expr (COMMA expr)*
  ;

// ----- Lexer ----- 
//KEYWORDS
PROGRAM   : 'program';
MAIN      : 'main';
END       : 'end';
VARS      : 'vars';
VOID      : 'void';
PRINT     : 'print';
IF        : 'if';
ELSE      : 'else';
WHILE     : 'while';
DO        : 'do';
INT_KW    : 'int';
FLOAT_KW  : 'float';

//OPERATOR AND PUNCTUATION
RELOP     : '==' | '!=' | '<=' | '>=' | '<' | '>';
ASSIGN    : '=';
PLUS      : '+';
MINUS     : '-';
MULT      : '*';
DIV       : '/';
COMMA     : ',';
SEMI      : ';';
COLON     : ':';
LPAREN    : '(';
RPAREN    : ')';
LBRACE    : '{';
RBRACE    : '}';
LBRACKET  : '[';
RBRACKET  : ']';

//TOKENS
ID        : [A-Za-z_][A-Za-z_0-9]*;
INT_LIT   : [0-9]+;
FLOAT_LIT : [0-9]+ '.' [0-9]+;
STRING_LIT: '"' (~["\r\n])* '"';

//
WS        : [ \t\r\n]+ -> skip ;