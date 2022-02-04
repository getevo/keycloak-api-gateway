package settings

// @doc type 		app
// @doc name		settings
// @doc description database settings reader
// @doc author		reza

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/gpath"
	"github.com/getevo/evo/menu"
	"gopkg.in/yaml.v2"
)

func Register() {
	evo.Register(App{})
}

type App struct{}

var Parameters []Params

//var db *gorm.DB

func (App) Register() {
	fmt.Println("Settings Registered")
	//db = evo.GetDBO()
	Settings = settings{}
	//ReadSettingsFromDB()
	ReadSettingsFromYaml()
}

func Print() {
	j, _ := json.MarshalIndent(Settings, "", "  ")
	fmt.Println(string(j))
	os.Exit(0)
}

func ExportSettingsToYaml() {
	b, err := yaml.Marshal(Settings)
	if err == nil {
		f, err := gpath.Open("./settings.yml")
		if err == nil {
			f.Write(b)
			os.Exit(0)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func ReadSettingsFromYaml() {
	evo.ParseConfig("", "", &Settings)
}

/*
func ReadSettingsFromDB(){
	db.Find(&Parameters)
	for _, row := range Parameters {
		_ = row
		ref := reflect.ValueOf(&Settings).Elem()
		typ := reflect.TypeOf(Settings)
		for i := 0; i < ref.NumField(); i++ {
			obj := ref.Field(i)
			innerTyp := reflect.TypeOf(obj.Interface())
			if typ.Field(i).Tag.Get("key") == row.Class {
				for j := 0; j < obj.NumField(); j++ {
					field := obj.Field(j)
					if strings.ToLower(innerTyp.Field(j).Tag.Get("key")) == strings.ToLower(row.Name) {
						if field.CanSet() {
							switch field.Kind() {
							case reflect.Bool:
								if len(row.Value) > 0 && (row.Value[0] == '1' || row.Value[0] == 'y' || row.Value[0] == 'Y'){
									field.SetBool(true)
								}else{
									field.SetBool(false)
								}
							case reflect.Int:
								field.SetInt(lib.ParseSafeInt64(row.Value))
							case reflect.Float64:
								field.SetFloat(lib.ParseSafeFloat(row.Value))
							case reflect.String:
								field.SetString(row.Value)
							}
						}
					}
				}
			}

		}
	}
}*/

// WhenReady called after setup all apps
func (App) WhenReady() {

}

// Router setup routers
func (App) Router() {}

// Permissions setup permissions of app
func (App) Permissions() []evo.Permission { return []evo.Permission{} }

// Menus setup menus
func (App) Menus() []menu.Menu {
	return []menu.Menu{}
}

func (App) Pack() {}
