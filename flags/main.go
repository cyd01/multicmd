package flags

import (
	"flag"
	"os"
	"strings"
	"strconv"
	"time"
)

func forgevar( cmd, name string ) string {
	return strings.ReplaceAll( strings.ToUpper(cmd+"_"+name), "-", "_" )
}

type Flag struct {
	flagName string
	flagSet *flag.FlagSet
}

func NewFlag( name string ) (*Flag) {
	f := &Flag {
		flagName: name,
		flagSet: flag.NewFlagSet( name, flag.ContinueOnError),
	}
	return f
}

func (f *Flag) Arg(i int) string {
	return f.flagSet.Arg(i)
}

func (f *Flag) Args() []string {
	return f.flagSet.Args()
}

func (f *Flag) SetUsage( fn func() ) {
	f.flagSet.Usage = fn
}

func (f *Flag) Usage() {
	f.flagSet.Usage()
}

func (f *Flag) Parse(arguments []string) error {
	return f.flagSet.Parse(arguments)
}

func (f *Flag) Bool(name string, value bool, usage string) *bool {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b {
		if strings.EqualFold(v,"1") || strings.EqualFold(v,"true") { val = true
		} else { val = false
		}
	}
	return f.flagSet.Bool(name,val,usage)
}

func (f *Flag) Duration(name string, value time.Duration, usage string) *time.Duration {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := time.ParseDuration(v); err==nil {
			val = vv
		}
	}
	return f.flagSet.Duration(name,val,usage)
}

func (f *Flag) Int(name string, value int, usage string) *int {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := strconv.Atoi(v); err==nil {
			val = vv
		}
	}
	return f.flagSet.Int(name,val,usage)
}

func (f *Flag) Float64(name string, value float64, usage string) *float64 {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := strconv.ParseFloat(v, 64); err==nil {
			val = vv
		}
	}
	return f.flagSet.Float64(name,val,usage)
}

func (f *Flag) Int64(name string, value int64, usage string) *int64 {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := strconv.ParseInt(v, 10, 64); err==nil {
			val = vv
		}
	}
	return f.flagSet.Int64(name,val,usage)
}

func (f *Flag) String(name string, value string, usage string) *string {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		val = v
	}
	return f.flagSet.String(name,val,usage)
}

func (f *Flag) NArg() int {
	return f.flagSet.NArg()
}

func (f *Flag) NFlag() int {
	return f.flagSet.NFlag()
}

func (f *Flag) Name() string {
	return f.flagSet.Name()
}

func (f *Flag) Parsed() bool {
	return f.flagSet.Parsed()
}

func (f *Flag) Set(name, value string) error {
	return f.flagSet.Set(name,value)
}

func (f *Flag) Uint(name string, value uint, usage string) *uint {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := strconv.ParseUint(v, 10, 64); err==nil {
			val = uint(vv)
		}
	}
	return f.flagSet.Uint(name,val,usage)
}

func (f *Flag) Uint64(name string, value uint64, usage string) *uint64 {
	val := value
	n := forgevar(f.flagName,name)
	if v,b :=os.LookupEnv(n); b { 
		if vv, err := strconv.ParseUint(v, 10, 64); err==nil {
			val = vv
		}
	}
	return f.flagSet.Uint64(name,val,usage)
}
