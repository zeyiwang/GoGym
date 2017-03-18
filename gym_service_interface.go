package GoGym

// GymService is a service interface
type GymService interface {
	Prepare(g *Gym)
	WhoIsYourBoss(g *Gym)
	CallYourBoss() *Gym
}
