package main
import "fmt"
import "strings"
import "jvmgo/ch06/classpath"
import "jvmgo/ch06/classfile"

func main() {
    cmd := parseCmd()
    if cmd.versionFlag {
        fmt.Println("version 0.0.1")
    } else if cmd.helpFlag || cmd.class == "" {
        printUsage()
    } else {
        startJVM(cmd)
    }
}

func startJVM(cmd *Cmd) {
    cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
func loadClass(classsName string, cp *classpath.ClassPath) *classfile.ClassFile  {
    classData, _, err := cp.ReadClass(classsName)
    if err != nil {
        panic(err)
    }
    cf, err := classfile.Parse(classData)
    if err != nil {
        panic(err)
    }
    return cf
}
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo  {
	for _,m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
// to test git routine
// to test git routine2
