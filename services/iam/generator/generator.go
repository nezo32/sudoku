package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load("../.env")

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if err != nil {
		fmt.Println("PORT MUST BE INTEGER")
		panic(err)
	}

	conn := postgres.DBConnection{
		Host:       os.Getenv("POSTGRES_HOST"),
		Port:       port,
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		SchemaName: "public",
		DBName:     os.Getenv("POSTGRES_DB"),
		SslMode:    "disable",
	}

	err = postgres.Generate(
		"../generated/postgres",
		conn,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									return defaultTableModelField.UseTags(
										fmt.Sprintf(`json:"%s,omitempty"`, columnMetaData.Name),
										fmt.Sprintf(`xml:"%s"`, columnMetaData.Name),
									)
								})
						}).
						UseView(func(table metadata.Table) template.ViewModel {
							return template.DefaultViewModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									return defaultTableModelField
								})
						}),
					)
			}),
	)

	if err != nil {
		panic(err)
	}
}
