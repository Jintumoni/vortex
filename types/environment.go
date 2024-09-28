package types

type StackFrame struct {
}

type TypeEnv struct {
	Env          []*StackFrame
	CurrentLevel uint
}

func (t *TypeEnv) PushFrame() {
	t.Env = append(t.Env, new(StackFrame))
	t.CurrentLevel++
}

func (t *TypeEnv) PopFrame() {
	t.Env = t.Env[:len(t.Env)-1]
	t.CurrentLevel--
}
