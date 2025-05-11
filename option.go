package tbox

// Option option for engine
type Option func(t *engine)

// WithPkgName pkg name
func WithPkgName(name string) Option {
	return func(t *engine) {
		t.packageName = name
	}
}

// WithPkgPath gen code to file dir path
func WithPkgPath(pkgPath string) Option {
	return func(t *engine) {
		t.pkgPath = pkgPath
	}
}

// WithOutputCmd output struct to cmd
func WithOutputCmd() Option {
	return func(t *engine) {
		t.isOutputCmd = true
	}
}

// WithTagKey set field tag key
func WithTagKey(tag string) Option {
	return func(t *engine) {
		t.tagKey = tag
	}
}

// WithUcFirstOnly first word upper.
func WithUcFirstOnly() Option {
	return func(t *engine) {
		t.ucFirstOnly = true
	}
}

// WithEnableTableNameFunc generate TableName method.
func WithEnableTableNameFunc() Option {
	return func(t *engine) {
		t.enableTableNameFunc = true
	}
}

// WithEnableJsonTag add json tab.
func WithEnableJsonTag() Option {
	return func(t *engine) {
		t.enableJsonTag = true
	}
}

// WithNoNullField no null field when code gen.
func WithNoNullField() Option {
	return func(t *engine) {
		t.noNullField = true
	}
}

// WithTableFileSuffix set table suffix for gen file,eg:user_gen.go
func WithTableFileSuffix(suffix string) Option {
	return func(t *engine) {
		t.tableFileSuffix = suffix
	}
}
