package tor

type Architecture uint8

const (
	X86 Architecture = iota
	AMD64
	ARM7
	ARM8
	ARM64
)

type Command struct {
	Architecture 
	FileType

	User int
	Group int


	Path string

	Basename string
	Name string
	Extension string

	MemoryFD bool
	InState State

	FD uintptr

}
