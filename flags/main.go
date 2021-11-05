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

type ALIAS struct{
	shortName   string
	longName    string
} 

type Flag struct {
	flagName     string
	flagSet      *flag.FlagSet
	flagAliases  []ALIAS
}

func NewFlag( name string ) (*Flag) {
	f := &Flag {
		flagName: name,
		flagSet: flag.NewFlagSet( name, flag.ContinueOnError),
	}
	return f
}

func (f *Flag) AliasByShort(name string) string {
	for _, val := range f.flagAliases {
		if val.shortName == name { return val.longName }
	}
	return ""
}

func (f *Flag) AliasByLong(name string) string {
	for _, val := range f.flagAliases {
		if val.longName == name { return val.shortName }
	}
	return ""
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
	args := arguments
	if len(args)>0 {
		for i:=0; i<len(args) ;i++ {
			if strings.HasPrefix(args[i],"-") {
				var v string
				if strings.HasPrefix(args[i],"--") { v=args[i][2:] } else { v=args[i][1:] }
				n := f.AliasByShort(v)
				if len(n)>0 { args[i] = "-"+n }
			}
		}
	}
	return f.flagSet.Parse(args)
}

func (f *Flag) addAlias(longName, shortName string) {
	if len(shortName)>1 {
		panic( f.flagName + " short name too long: " + shortName )
	}
	if len(shortName)==1 {
		if len(f.AliasByShort(shortName)) >0 {
			panic( f.flagName + " short flag redefined:" + shortName )
		}
		f.flagAliases = append( f.flagAliases, ALIAS{shortName: shortName, longName: longName} )
	}
}

func (f *Flag) PrintDefaults() {
	f.flagSet.PrintDefaults()
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

func (f *Flag) IntP(longName, shortName string, value int, usage string) *int {
	if len(shortName)>=1 { f.addAlias(longName, shortName) }
	return f.Int(longName,value,usage)
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
