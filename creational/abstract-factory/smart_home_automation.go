package abstractfactory

import "errors"

type SmartLight interface {
	TurnOn()
	TurnOff()
	SetBrightnessLevel(int32) error
	GetBrightnessLevel() int32
}

type SmartThermostat interface {
	SetTemperature(float32)
	GetTemperature() float32
}

type SmartDoorLock interface {
	Lock()
	Unlock()
	GetLockState() bool
}

type HardwareFamily interface {
	GetSmartLight() SmartLight
	GetSmartThermostat() SmartThermostat
	GetSmartDoorLock() SmartDoorLock
}

func GetHardWareFamily(family string) (HardwareFamily, error) {
	switch family {
	case "zigbee":
		return &zigBeeFamily{}, nil
	case "zwave":
		return &zWaveFamily{}, nil
	default:
		return nil, errors.New("unsupported family")
	}
}

type zigBeeFamily struct{}

func (z *zigBeeFamily) GetSmartLight() SmartLight {
	return &zigBeeSmartLight{}
}

func (z *zigBeeFamily) GetSmartThermostat() SmartThermostat {
	return &zigBeeSmartThermostat{}
}

func (z *zigBeeFamily) GetSmartDoorLock() SmartDoorLock {
	return &zigBeeSmartDoorLock{}
}

type zigBeeSmartLight struct {
	isTurnedOn      bool
	brightnessLevel int32
}

func (l *zigBeeSmartLight) TurnOff() {
	l.isTurnedOn = false
}

func (l *zigBeeSmartLight) TurnOn() {
	l.isTurnedOn = true
}

func (l *zigBeeSmartLight) SetBrightnessLevel(val int32) error {
	if val < 0 || val > 100 {
		return errors.New("value out of range")
	}
	l.brightnessLevel = val
	return nil
}

func (l *zigBeeSmartLight) GetBrightnessLevel() int32 {
	return l.brightnessLevel
}

type zigBeeSmartThermostat struct {
	temprature float32
}

func (s *zigBeeSmartThermostat) SetTemperature(val float32) {
	s.temprature = val
}

func (s *zigBeeSmartThermostat) GetTemperature() float32 {
	return s.temprature
}

type zigBeeSmartDoorLock struct {
	isLocked bool
}

func (s *zigBeeSmartDoorLock) Lock() {
	s.isLocked = true
}

func (s *zigBeeSmartDoorLock) Unlock() {
	s.isLocked = false
}

func (s *zigBeeSmartDoorLock) GetLockState() bool {
	return s.isLocked
}

type zWaveFamily struct{}

func (z *zWaveFamily) GetSmartLight() SmartLight {
	return &zwaveSmartLight{}
}

func (z *zWaveFamily) GetSmartThermostat() SmartThermostat {
	return &zwaveSmartThermostat{}
}

func (z *zWaveFamily) GetSmartDoorLock() SmartDoorLock {
	return &zwaveSmartDoorLock{}
}

type zwaveSmartLight struct {
	isTurnedOn      bool
	brightnessLevel int32
}

func (l *zwaveSmartLight) TurnOff() {
	l.isTurnedOn = false
}

func (l *zwaveSmartLight) TurnOn() {
	l.isTurnedOn = true
}

func (l *zwaveSmartLight) SetBrightnessLevel(val int32) error {
	if val < 0 || val > 100 {
		return errors.New("value out of range")
	}
	l.brightnessLevel = val
	return nil
}

func (l *zwaveSmartLight) GetBrightnessLevel() int32 {
	return l.brightnessLevel
}

type zwaveSmartThermostat struct {
	temprature float32
}

func (s *zwaveSmartThermostat) SetTemperature(val float32) {
	s.temprature = val
}

func (s *zwaveSmartThermostat) GetTemperature() float32 {
	return s.temprature
}

type zwaveSmartDoorLock struct {
	isLocked bool
}

func (s *zwaveSmartDoorLock) Lock() {
	s.isLocked = true
}

func (s *zwaveSmartDoorLock) Unlock() {
	s.isLocked = false
}

func (s *zwaveSmartDoorLock) GetLockState() bool {
	return s.isLocked
}
