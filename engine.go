package tbox

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// engine table to struct data define.
type engine struct {
	dsn string  // mysql dsn
	db  *sql.DB // sql DB

	packageName string // convert to struct package name
	pkgPath     string // gen code to dir path
	isOutputCmd bool   // whether to output to cmd,otherwise output to file

	addTag bool
	tagKey string // the key value of the tag field, the default is `db:xxx`

	// While the first letter of the field is capitalized
	// whether to convert other letters to lowercase, the default false is not converted
	ucFirstOnly bool

	// Table name returned by TaleName func
	// such as xorm, gorm library uses TableName method to return table name
	enableTableNameFunc bool

	// Whether to add json tag, not added by default
	enableJsonTag bool

	noNullField bool
}

// New create an entry for engine
func New(dsn string, opts ...Option) *engine {
	if dsn == "" {
		log.Fatalln("dsn is empty")
	}

	t := &engine{
		dsn:         dsn,
		packageName: "model",
		pkgPath:     "./model",
		tagKey:      "db",
		addTag:      true,
	}

	for _, o := range opts {
		o(t)
	}

	var err error
	t.db, err = t.connect()
	if err != nil {
		log.Fatalln("dsn connect mysql error: ", err)
	}

	if t.pkgPath == "" {
		log.Fatalln("pkg path is empty")
	}

	if !checkPathExist(t.pkgPath) {
		err = os.MkdirAll(t.pkgPath, 0755)
		if err != nil {
			log.Fatalln("create pkg dir error: ", err)
		}
	}

	return t
}

var noNullList = []string{
	"int", "int64", "float", "float32", "json", "string",
}

func (t *engine) getTableCode(tab string, record []columnEntry) []string {
	var header strings.Builder
	header.WriteString(fmt.Sprintf("// Package %s of db entity\n", t.packageName))
	header.WriteString(fmt.Sprintf("// Code generated by go-god/tbox. DO NOT EDIT!!!\n\n"))
	header.WriteString(fmt.Sprintf("package %s\n\n", t.packageName))

	var importBuf strings.Builder
	var tabInfoBuf strings.Builder
	tabPrefix := t.camelCase(tab)
	tabName := tabPrefix + "Table"
	structName := tabPrefix + "Entity"
	tabInfoBuf.WriteString(fmt.Sprintf("// %s for %s\n", tabName, tab))
	tabInfoBuf.WriteString(fmt.Sprintf("const %s = \"%s\"\n\n", tabName, tab))
	tabInfoBuf.WriteString(fmt.Sprintf("// %s for %s table entity struct.\n", structName, tab))
	tabInfoBuf.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, val := range record {
		dataType := getType(val.DataType) // column type
		if dataType == "time.Time" && importBuf.Len() == 0 {
			importBuf.WriteString("import (\n\t\"time\"\n)\n\n")
		}

		if t.noNullField && strInSet(dataType, noNullList) {
			val.IsNullable = "NO"
		}

		if val.IsNullable == "YES" {
			tabInfoBuf.WriteString(fmt.Sprintf("\t%s\t*%s", t.camelCase(val.Field), dataType))
		} else {
			tabInfoBuf.WriteString(fmt.Sprintf("\t%s\t%s", t.camelCase(val.Field), dataType))
		}

		if t.addTag {
			tabInfoBuf.WriteString("\t")
			tabInfoBuf.WriteString("`")
			if t.enableJsonTag {
				tabInfoBuf.WriteString(fmt.Sprintf("json:\"%s\"", strings.ToLower(val.Field)))
				tabInfoBuf.WriteString(" ")
			}

			if t.tagKey != "" {
				if t.tagKey == "xorm" {
					if val.FieldKey == "PRI" && val.Extra == "auto_increment" {
						tabInfoBuf.WriteString(fmt.Sprintf("%s:\"%s pk autoincr\"", t.tagKey,
							strings.ToLower(val.Field)))
					} else {
						tabInfoBuf.WriteString(fmt.Sprintf("%s:\"%s\"", t.tagKey, strings.ToLower(val.Field)))
					}
				} else if t.tagKey == "gorm" {
					if val.FieldKey == "PRI" && val.Extra == "auto_increment" {
						tabInfoBuf.WriteString(fmt.Sprintf("%s:\"%s primaryKey\"", t.tagKey,
							strings.ToLower(val.Field)))
					} else {
						tabInfoBuf.WriteString(fmt.Sprintf("%s:\"%s\"", t.tagKey, strings.ToLower(val.Field)))
					}
				} else {
					tabInfoBuf.WriteString(fmt.Sprintf("%s:\"%s\"", t.tagKey, strings.ToLower(val.Field)))
				}
			}

			tabInfoBuf.WriteString("`")
		}

		tabInfoBuf.WriteString("\n")
	}

	tabInfoBuf.WriteString("}\n")

	if t.enableTableNameFunc {
		tabInfoBuf.WriteString(fmt.Sprintf("\n// %s for %s", "TableName", tab))
		tabInfoBuf.WriteString(fmt.Sprintf("\nfunc (%s) TableName() string {\n", structName))
		tabInfoBuf.WriteString(fmt.Sprintf("\treturn %s\n", tabName))
		tabInfoBuf.WriteString("}\n")
	}

	return []string{
		header.String(), importBuf.String(), tabInfoBuf.String(),
	}
}

// Run exec table to struct
func (t *engine) Run(table ...string) error {
	tabColumns, err := t.GetColumns(table...)
	if err != nil {
		return err
	}

	for tab, record := range tabColumns {
		str := strings.Join(t.getTableCode(tab, record), "")
		if t.isOutputCmd {
			fmt.Println(str)
		}

		// write to file
		var fileName string
		fileName, err = filepath.Abs(filepath.Join(t.pkgPath, strings.ToLower(tab)+"_gen.go"))
		if err != nil {
			log.Println("file path error: ", err)
			continue
		}

		err = os.WriteFile(fileName, []byte(str), 0666)
		if err != nil {
			log.Println("gen code error: ", err)
			return err
		}
	}

	return nil
}

type columnEntry struct {
	TableName    string
	Field        string
	DataType     string
	FieldDesc    string
	FieldKey     string
	OrderBy      int
	IsNullable   string
	MaxLength    sql.NullInt64
	NumericPrec  sql.NullInt64
	NumericScale sql.NullInt64
	Extra        string
	FieldComment string
}

var fields = []string{
	"TABLE_NAME as table_name",
	"COLUMN_NAME as field",
	"DATA_TYPE as data_type",
	"COLUMN_TYPE as field_desc",
	"COLUMN_KEY as field_key",
	"ORDINAL_POSITION as order_by",
	"IS_NULLABLE as is_nullable",
	"CHARACTER_MAXIMUM_LENGTH as max_length",
	"NUMERIC_PRECISION as numeric_prec",
	"NUMERIC_SCALE as numeric_scale",
	"EXTRA as extra",
	"COLUMN_COMMENT as field_comment",
}

// GetColumns Get the information fields related to the table according to the table name
func (t *engine) GetColumns(table ...string) (map[string][]columnEntry, error) {
	var sqlStr = "select " + strings.Join(fields, ",") + " from information_schema.COLUMNS " +
		"WHERE table_schema = DATABASE()"
	if len(table) > 0 {
		sqlStr += " AND TABLE_NAME in (" + t.implodeTable(table...) + ")"
	}

	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"
	rows, err := t.db.Query(sqlStr)
	if err != nil {
		log.Println("read table information error: ", err.Error())
		return nil, err
	}

	defer rows.Close()

	records := make(map[string][]columnEntry, 20)
	for rows.Next() {
		col := columnEntry{}
		err = rows.Scan(&col.TableName, &col.Field, &col.DataType, &col.FieldDesc,
			&col.FieldKey, &col.OrderBy, &col.IsNullable, &col.MaxLength, &col.NumericPrec, &col.NumericScale,
			&col.Extra, &col.FieldComment,
		)

		if err != nil {
			log.Println("scan column error: ", err)
			continue
		}

		if _, ok := records[col.TableName]; !ok {
			records[col.TableName] = make([]columnEntry, 0, 20)
		}

		records[col.TableName] = append(records[col.TableName], col)
	}

	return records, nil
}

func (t *engine) connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", t.dsn)
	return db, err
}

func (t *engine) implodeTable(table ...string) string {
	var arr []string
	for _, tab := range table {
		arr = append(arr, fmt.Sprintf("'%s'", tab))
	}

	return strings.Join(arr, ",")
}

func (t *engine) camelCase(str string) string {
	var buf strings.Builder
	arr := strings.Split(str, "_")
	if len(arr) > 0 {
		for _, s := range arr {
			if t.ucFirstOnly {
				buf.WriteString(strings.ToUpper(s[0:1]))
				buf.WriteString(strings.ToLower(s[1:]))
			} else {
				lintStr := lintName(strings.Join([]string{strings.ToUpper(s[0:1]), s[1:]}, ""))
				buf.WriteString(lintStr)
			}
		}
	}

	return buf.String()
}

func checkPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func strInSet(s string, sets []string) bool {
	for _, member := range sets {
		if s == member {
			return true
		}
	}

	return false
}
