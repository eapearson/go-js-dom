package dom

// type CSSStyleDeclaration interface {
// 	// attributes
// 	CSSText() string
// 	SetCSSText(string)
// 	Length() int
// 	// methods
// 	GetPropertyPriority() string
// 	GetPropertyValue() string
// 	GetItem(int) string
// 	RemoveProperty(string)
// 	SetProperty(name string, value string, priority string)

// }

type CSSImportRule interface {
}

type MediaList interface {
	Set(string)
	Length() int
	Get(int) string
	AppendMedium(string)
	DeleteMedium(string)
}

type StyleSheet interface {
	Disabled() bool
	Href() string
	Media() MediaList
	OwnerNode() Node
	ParentStyleSheet() StyleSheet
	Title() string
	Type() string
}

type CSSStyleSheetList interface {
	StyleSheets() []CSSStyleSheet
}

type CSSStyleSheet interface {
	CSSRules() []CSSRule
	OwnerRule() CSSImportRule
	DeleteRule(index int)
	InsertRule(rule string, index int) int
}

type CSSRule interface {
	// attributes
	CSSText() string
	ParentRule() CSSRule
	ParentStyleSheet() CSSStyleSheet
}

const (
	UNKNOWN_RULE = iota
	STYLE_RULE
	CHARSET_RULE
	IMPORT_RULE
	MEDIA_RULE
	FONT_FACE_RULE
	PAGE_RULE
	KEYFRAMES_RULE
	KEYFRAME_RULE
	_
	NAMESPACE_RULE
	COUNTER_STYLE_RULE
	SUPPORTS_RULE
	DOCUMENT_RULE
	FONT_FEATURE_VALUES_RULE
	VIEWPORT_RULE
)

type CSSStyleRule interface {
	SelectorText() string
	Style() *CSSStyleDeclaration
	IsType(int) bool
}
