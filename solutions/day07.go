package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day07 struct{}

func (Day07) Name() string {
	return "day07"
}

type node interface {
	name() string
	size() int
	isDir() bool
}

type directory struct {
	_name      string
	cachedSize int
	parent     *directory
	children   []node
}

func (d *directory) name() string {
	return d._name
}

func (d *directory) isDir() bool {
	return true
}

func (d *directory) size() int {
	if d.cachedSize > 0 {
		return d.cachedSize
	}
	total := 0
	for _, c := range d.children {
		total += c.size()
	}
	d.cachedSize = total
	return total
}

func (d *directory) addNode(child node) {
	d.children = append(d.children, child)
}

func (d *directory) traverse(cb func(node node)) {
	cb(d)
	for _, c := range d.children {
		if c.isDir() {
			c.(*directory).traverse(cb)
		} else {
			cb(c)
		}
	}
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{_name: name, cachedSize: -1, parent: parent, children: make([]node, 0)}
}

type file struct {
	_name string
	_size int
}

func (f *file) name() string {
	return f._name
}

func (f *file) isDir() bool {
	return false
}

func (f *file) size() int {
	return f._size
}

func newFile(name string, size int) *file {
	return &file{_name: name, _size: size}
}

type fileSystem struct {
	cwd  *directory
	root *directory
}

func (fs *fileSystem) changeDir(name string) {
	if name == ".." {
		fs.cwd = fs.cwd.parent
		return
	} else if name == "/" {
		fs.cwd = fs.root
		return
	}
	for _, c := range fs.cwd.children {
		if c.name() == name && c.isDir() {
			fs.cwd = c.(*directory)
			return
		}
	}
	fs.cwd.addNode(newDirectory(name, fs.cwd))
}

func fileSystemfromInput(input string) *fileSystem {
	root := newDirectory("/", nil)
	fs := fileSystem{cwd: root, root: root}

	readNodes := false
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				readNodes = false
				fs.changeDir(tokens[2])
			} else if tokens[1] == "ls" {
				readNodes = true
			}
		} else if readNodes {
			if tokens[0] == "dir" {
				fs.cwd.addNode(newDirectory(tokens[1], fs.cwd))
			} else {
				size, _ := strconv.Atoi(tokens[0])
				fs.cwd.addNode(newFile(tokens[1], size))
			}
		}
	}
	return &fs
}

func (Day07) Solution(input string, part2 bool) string {
	fs := fileSystemfromInput(input)
	size := 0
	if !part2 {
		fs.root.traverse(func(n node) {
			if n.isDir() && n.size() <= 100000 {
				size += n.size()
			}
		})
	} else {
		unusedSpace := 70000000 - fs.root.size()
		fs.root.traverse(func(n node) {
			if n.isDir() && n.size()+unusedSpace > 30000000 && (size == 0 || n.size() < size) {
				size = n.size()
			}
		})
	}
	return fmt.Sprint(size)
}
