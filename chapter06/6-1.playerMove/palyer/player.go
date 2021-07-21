package player

type Player struct {
	currentPos Vec2
	targetPos  Vec2
	speed      float32
}

func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

func (p *Player) Pos() Vec2 {
	return p.currentPos
}

func (p *Player) IsArrived() bool {
	return p.currentPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player) Update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currentPos).Normalize()
		newPos := p.currentPos.Add(dir.Scale(p.speed))
		p.currentPos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
