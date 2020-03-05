package chapter3

// Kadai1 is ...
// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

// ID is Getter
func (k Kadai1) ID() int {
	return k.id
}

// Name is Getter
func (k Kadai1) Name() string {
	return k.name
}

// SetID is Setter
func (k *Kadai1) SetID(id int) {
	k.id = id
}

// SetName is Setter
func (k *Kadai1) SetName(name string) {
	k.name = name
}
