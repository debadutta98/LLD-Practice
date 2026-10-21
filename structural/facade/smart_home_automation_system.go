package facade

import "fmt"

type lightSystem struct {
	isTurnOn  bool
	dimLevel  int
	colorTemp int
}

func (l *lightSystem) turnOn() {
	fmt.Println("Turn on light")
	l.isTurnOn = true
}

func (l *lightSystem) turnOff() {
	fmt.Println("Turn off light")
	l.isTurnOn = false
}

func (l *lightSystem) dim(level int) {
	fmt.Printf("Set dim level to %d\n", level)
	l.dimLevel = level
}

func (l *lightSystem) setColorTemperature(kelvin int) {
	fmt.Printf("Set temp %d\n", kelvin)
	l.colorTemp = kelvin
}

type audioSystem struct {
	isPowerOn bool
	volume    int
	source    string
}

func (a *audioSystem) powerOn() {
	fmt.Println("audio system power on")
	a.isPowerOn = true
}

func (a *audioSystem) powerOff() {
	fmt.Println("audio system power off")
	a.isPowerOn = false
}

func (a *audioSystem) setVolume(level int) {
	fmt.Printf("Set volumn to %d\n", level)
	a.volume = level
}

func (a *audioSystem) setInputSource(src string) {
	fmt.Printf("Set input channel %s\n", src)
	a.source = src
}

type automatedBlinds struct {
	open int
}

func (a *automatedBlinds) lower() {
	fmt.Println("lower the blinds completely")
	a.open = 0
}

func (a *automatedBlinds) raise() {
	fmt.Println("raise the blinds completely")
	a.open = 100
}

func (a *automatedBlinds) setPercentageOpen(percent int) {
	a.open = percent
	fmt.Printf("open the auto blinds %d%%\n", a.open)
}

type hvacSystem struct {
	temp  float64
	speed int
	mode  string
}

func (h *hvacSystem) setTemperature(temp float64) {
	fmt.Printf("set temperature %f\n", temp)
	h.temp = temp
}

func (h *hvacSystem) setFanSpeed(speed int) {
	fmt.Printf("set speed %d\n", speed)
	h.speed = speed
}

func (h *hvacSystem) setMode(mode string) {
	fmt.Printf("set mode %s\n", mode)
	h.mode = mode
}

type videoProjector struct {
	isTurnOn   bool
	channel    string
	aspectMode string
}

func (v *videoProjector) turnOn() {
	fmt.Println("turn on video projector")
	v.isTurnOn = true
}

func (v *videoProjector) turnOff() {
	fmt.Println("turn off video channel")
	v.isTurnOn = false
}

func (v *videoProjector) setInputChannel(channel string) {
	fmt.Printf("set video projector input channel %s\n", channel)
	v.channel = channel
}

func (v *videoProjector) setAspectMode(mode string) {
	fmt.Printf("set aspect mode input channel %s", mode)
	v.aspectMode = mode
}

type UnifiedInterface struct {
	light  *lightSystem
	audio  *audioSystem
	video  *videoProjector
	blinds *automatedBlinds
	hvac   *hvacSystem
}

func NewUnifiedInterface() *UnifiedInterface {
	return &UnifiedInterface{
		light:  new(lightSystem),
		audio:  new(audioSystem),
		video:  new(videoProjector),
		blinds: new(automatedBlinds),
		hvac:   new(hvacSystem),
	}
}

func (u *UnifiedInterface) WatchMovie() {
	u.blinds.lower()
	u.light.dim(10)
	u.light.turnOn()
	u.video.setInputChannel("HDMI 1")
	u.video.setAspectMode("16:9")
	u.audio.setVolume(30)
	u.audio.setInputSource("HDMI 1")
	u.hvac.setMode("COOL")
	u.hvac.setTemperature(21)
}

func (u *UnifiedInterface) EndMovie() {
	u.video.turnOff()
	u.audio.powerOff()
	u.blinds.raise()
	u.light.dim(100)
	u.hvac.setMode("ECO")
}

func (u *UnifiedInterface) ListenToMusic() {
	u.blinds.setPercentageOpen(50)
	u.light.setColorTemperature(2700)
	u.video.turnOff()
	u.audio.powerOn()
	u.audio.setInputSource("Bluetooth")
	u.audio.setVolume(20)
	u.hvac.setTemperature(23)
}

func (u *UnifiedInterface) LeaveHome() {
	u.light.turnOff()
	u.audio.powerOff()
	u.video.turnOff()
	u.blinds.lower()
	u.hvac.setMode("ECO")
}
