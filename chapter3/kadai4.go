package chapter3

// Eye is struct
type Eye struct {
	isOpen bool
}

// Watch is ...
func (e *Eye) Watch() {
	e.isOpen = true
}

// Nose is struct
type Nose struct {
	isOpen bool
}

// Breathe is ...
func (n *Nose) Breathe() {
	n.isOpen = true
}

// Mouth is struct
type Mouth struct {
	isOpen  bool
	hasFood bool
}

// Eat is ...
func (m *Mouth) Eat() {
	m.hasFood = true
}

// Breathe is ...
func (m *Mouth) Breathe() {
	m.isOpen = true
}

// Face is 課題4
// 上の3つのstructの機能を持つFaceを実行してください。
// ただし口と鼻両方で呼吸します。
type Face struct {
	Eye
	Nose
	Mouth
}

// Breathe is ...
func (f *Face) Breathe() {
	f.Mouth.Breathe()
	f.Nose.Breathe()
}
