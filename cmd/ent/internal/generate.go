// Copyright (C) liasica. 2021-present.
//
// Created at 2021/12/10
// Based on aurservd by liasica, magicrolan@qq.com.

package internal

import (
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "entgo.io/ent/schema/field"
    "fmt"
    "github.com/liasica/go-helpers/tools"
    "github.com/spf13/cobra"
    "log"
    "strings"
)

func init() {
    gen.Funcs["containsField"] = func(fields []*gen.Field, name string) bool {
        for _, f := range fields {
            if name == f.Name {
                return true
            }
        }
        return false
    }
}

// IDType is a custom ID implementation for pflag.
type IDType field.Type

// Set implements the Set method of the flag.Value interface.
func (t *IDType) Set(s string) error {
    switch s {
    case field.TypeInt.String():
        *t = IDType(field.TypeInt)
    case field.TypeInt64.String():
        *t = IDType(field.TypeInt64)
    case field.TypeUint.String():
        *t = IDType(field.TypeUint)
    case field.TypeUint64.String():
        *t = IDType(field.TypeUint64)
    case field.TypeString.String():
        *t = IDType(field.TypeString)
    default:
        return fmt.Errorf("invalid type %q", s)
    }
    return nil
}

// Type returns the type representation of the id option for help command.
func (IDType) Type() string {
    return fmt.Sprintf("%v", []field.Type{
        field.TypeInt,
        field.TypeInt64,
        field.TypeUint,
        field.TypeUint64,
        field.TypeString,
    })
}

// String returns the default value for the help command.
func (IDType) String() string {
    return field.TypeInt.String()
}

// GenerateCmd returns the generate command for ent/c packages.
func GenerateCmd(postRun ...func(*gen.Config)) *cobra.Command {
    var (
        cfg       gen.Config
        storage   string
        features  []string
        templates []string
        idtype    = IDType(field.TypeUint64)
        cmd       = &cobra.Command{
            Use:   "generate [flags] path",
            Short: "generate go code for the schema directory",
            Example: examples(
                "ent generate ./ent/schema",
                "ent generate github.com/a8m/x",
            ),
            // Args: cobra.ExactArgs(1),
            Run: func(cmd *cobra.Command, path []string) {
                Clean()

                if features == nil {
                    features = []string{"sql/modifier", "sql/upsert", "privacy", "entql", "sql/execquery", "schema/snapshot"}
                }
                opts := []entc.Option{
                    entc.Storage(storage),
                    entc.FeatureNames(features...),
                }
                templates = append(templates,
                    "./cmd/ent/template/upsert.tmpl",
                    "./cmd/ent/template/pointer.tmpl",
                    "./cmd/ent/template/table.tmpl",
                )
                for _, tmpl := range templates {
                    typ := "dir"
                    if parts := strings.SplitN(tmpl, "=", 2); len(parts) > 1 {
                        typ, tmpl = parts[0], parts[1]
                    }
                    switch typ {
                    case "dir":
                        opts = append(opts, entc.TemplateDir(tmpl))
                    case "file":
                        opts = append(opts, entc.TemplateFiles(tmpl))
                    case "glob":
                        opts = append(opts, entc.TemplateGlob(tmpl))
                    default:
                        log.Fatalln("unsupported template type", typ)
                    }
                }
                // If the target directory is not inferred from
                // the schema path, resolve its package path.
                if cfg.Target != "" {
                    pkgPath, err := PkgPath(DefaultConfig, cfg.Target)
                    if err != nil {
                        log.Fatalln(err)
                    }
                    cfg.Package = pkgPath
                }
                cfg.IDType = &field.TypeInfo{Type: field.Type(idtype)}
                p := defaultSchema
                if len(path) > 0 {
                    p = path[0]
                }
                // cfg.Hooks = []gen.Hook{
                //     SingularTableName(),
                // }
                if err := entc.Generate(p, &cfg, opts...); err != nil {
                    log.Fatalln(err)
                }
                for _, fn := range postRun {
                    fn(&cfg)
                }
            },
        }
    )
    cmd.Flags().Var(&idtype, "idtype", "type of the id field")
    cmd.Flags().StringVar(&storage, "storage", "sql", "storage driver to support in codegen")
    cmd.Flags().StringVar(&cfg.Header, "header", "", "override codegen header")
    cmd.Flags().StringVar(&cfg.Target, "target", "", "target directory for codegen")
    cmd.Flags().StringSliceVarP(&features, "feature", "", nil, "extend codegen with additional features")
    cmd.Flags().StringSliceVarP(&templates, "template", "", nil, "external templates to execute")
    return cmd
}

// SingularTableName ensures all nodes have a singular table name..
func SingularTableName() gen.Hook {
    return func(next gen.Generator) gen.Generator {
        return gen.GenerateFunc(func(g *gen.Graph) error {
            for _, n := range g.Nodes {
                ant := n.EntSQL()
                n.Annotations = make(map[string]interface{})

                if ant == nil {
                    ant = &entsql.Annotation{}
                    n.Annotations[ant.Name()] = ant
                }
                ant.Table = tools.StrToSnakeCase(n.Name)
            }
            return next.Generate(g)
        })
    }
}
