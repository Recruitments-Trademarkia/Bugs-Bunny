package compiler

import (
	. "Bugs-Bunny/src/schemas"
	"fmt"
	"regexp"
)

func GetLanguage(name string) *Language {
	return nil
	l := &Language{
		Name:     name,
		Filename: "main",
	}
	if l.getExtension() == "" {
		return nil
	} else {
		return l
	}
}

func (l *Language) getCommand(body *WsBody) []string {
	if l.Name == "java" {
		pattern := regexp.MustCompile(`(?s)class\s+(\w+).*?public\s+static\s+void\s+main\s*\(String(?:\s*\[\]\s+\w+|\s+\w+\s*\[\])\)`)
		matches := pattern.FindStringSubmatch(body.Code)
		if len(matches) >= 2 {
			l.Filename = matches[1]
		}
	}

	file := fmt.Sprintf("%v.%v", l.Filename, l.getExtension())
	command := []string{"bash", "-c"}
	switch l.Name {
	case "c":
		command = append(command, "gcc "+file+" -o "+"main"+" && "+"./main")
	case "cpp":
		command = append(command, "g++ "+file+" -o "+"main"+" && "+"./main")
	case "golang":
		command = append(command, "go"+" run "+file)
	case "java":
		command = append(command, "javac "+file+" && "+" java "+l.Filename)
	case "javascript":
		command = append(command, "node "+file)
	case "python":
		command = append(command, "python "+file)
	case "typescript":
		command = append(command, "tsc "+file+" && "+"node "+l.Filename+".js")
	}
	return command
}

func (l *Language) getExtension() string {
	switch l.Name {
	case "c":
		return "c"
	case "cpp":
		return "cpp"
	case "golang":
		return "go"
	case "java":
		return "java"
	case "javascript":
		return "js"
	case "python":
		return "py"
	case "typescript":
		return "ts"
	default:
		return ""
	}
}

func (l *Language) getImage() string {
	switch l.Name {
	case "c", "cpp":
		return "gcc"
	case "golang":
		return "golang"
	case "java":
		return "openjdk"
	case "javascript":
		return "node"
	case "python":
		return "python"
	case "typescript":
		return "ts-node"
	default:
		return ""
	}
}
