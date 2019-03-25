package main

import (
	"encoding/json"
	messages "github.com/citrullin/udp_client/generated"
	"github.com/golang/protobuf/proto"
	"github.com/op/go-logging"
	"net"
	"os"
	"strings"
	"time"
)

/**
	This application caches the last [sensorDataBufferSize] data points
 */
const sensorDataBufferSize = 10

/**
	The interfaceName is needed to route the UDP packets to the correct interface
	You can also configure radvd to have a routable address
 */
const interfaceName = "bt0"

const localPort = 8085

const debug = true

/**
	Seed configuration. A seed is used to discover other nodes.
	Every other node should register to a seed to be discoverable.
	Can also implemented as array, if a larger mesh network is used.
	Todo: Implement discover functionality
 */
var seedSensorConfig = SensorNode{
	Config: SensorConfig{
		Address: net.UDPAddr{
			IP: net.ParseIP("fe80::2ca:46ff:fed3:1967"), Port: 51037, Zone: interfaceName,
		},
	},
}

var log = logging.MustGetLogger("sensorClient")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

type EnvSensorFeatures struct {
	HasTemperature         bool
	HasHumanity            bool
	HasAtmosphericPressure bool
	HasPm2_5               bool
}

/**
	Float is used for simplicity. Consider implementing fixed-point numbers instead.
	To improve accuracy. Some MCUs also don't have an FPU.
 */
type EnvSensorData struct {
	Temperature         float32
	Humanity            float32
	AtmosphericPressure float32
	Pm2_5               float32
}

type EnvSensorDataRing struct {
	Data     [sensorDataBufferSize] EnvSensorData
	Position int
}

type SensorConfig struct {
	Address net.UDPAddr
}

type SensorNode struct {
	Config   SensorConfig
	Features EnvSensorFeatures
	DataRing EnvSensorDataRing
}

//Todo: Add neighbor discovery protocol
var sensorNodes = []SensorNode{
	seedSensorConfig,
/**
	If needed you can add other nodes here as well. Or implement a neighbor discovery protocol
 */
/*	{
		Config: SensorConfig{
			Address: net.UDPAddr{
				IP: net.ParseIP("fe80::213:afff:fe94:d75"), Port: 51037, Zone: interfaceName,
			},
		},
	},
	{
		Config: SensorConfig{
			Address: net.UDPAddr{
				IP: net.ParseIP("fe80::289:b5ff:fef4:6fa"), Port: 51037, Zone: interfaceName,
			},
		},
	},
*/
}

type SensorCommand int

const (
	requestFeatureCommand SensorCommand = iota
	requestDataCommand
	dataResponseCommand
	featureResponseCommand
	setupTestCommand
	noCommand
)

func CheckError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func getCommandName(commandByte byte) SensorCommand {
	switch commandByte {
	case 33:
		return requestFeatureCommand
	case 34:
		return requestDataCommand
	case 35:
		return dataResponseCommand
	case 36:
		return featureResponseCommand
	case 88:
		return setupTestCommand
	default:
		return noCommand
	}
}

func getCommandByte(command SensorCommand) byte {
	switch command {
	case requestFeatureCommand:
		return 33
	case requestDataCommand:
		return 34
	case dataResponseCommand:
		return 35
	case featureResponseCommand:
		return 36
	case setupTestCommand:
		return 88
	default:
		return 0
	}
}

func getEnvSensorFeature(connection *net.UDPConn, sensorNode SensorNode) {
	command := getCommandByte(requestFeatureCommand)

	if debug {
		log.Debug(
			"Send command,", command, "to address", sensorNode.Config.Address.String(),
			"from address", connection.LocalAddr().String(), "\n")
	}

	_, err := connection.WriteToUDP([]byte{command}, &sensorNode.Config.Address)
	if err != nil {
		log.Error(err)
	}
}

func getEnvSensorData(connection *net.UDPConn, sensorNode SensorNode) {
	command := getCommandByte(requestDataCommand)

	if debug {
		log.Debug(
			"Send command,", command, "to address", sensorNode.Config.Address.String(),
			"from address", connection.LocalAddr().String(), "\n")
	}

	request := messages.DataRequest{}
	request.Temperature = &sensorNode.Features.HasTemperature
	request.Pm2_5 = &sensorNode.Features.HasPm2_5
	request.Humanity = &sensorNode.Features.HasHumanity
	request.AtmosphericPressure = &sensorNode.Features.HasAtmosphericPressure

	encodedData, err := proto.Marshal(&request)
	CheckError(err)

	data := append([]byte{command}, encodedData...)

	_, err = connection.WriteToUDP(data, &sensorNode.Config.Address)
	if err != nil {
		log.Error(err)
	}
}

func addSensorFeature(response *messages.FeatureResponse, addr *net.UDPAddr) {
	for i, elem := range sensorNodes {
		if elem.Config.Address.IP.String() == addr.IP.String() {
			sensorNodes[i].Features = EnvSensorFeatures{
				HasTemperature:         *response.HasTemperature,
				HasHumanity:            *response.HasHumanity,
				HasAtmosphericPressure: *response.HasAtmosphericPressure,
				HasPm2_5:               *response.HasPm2_5,
			}
		}
	}
}

func addDataPoint(response *messages.DataResponse, addr *net.UDPAddr, ) {
	for i, elem := range sensorNodes {
		if elem.Config.Address.IP.String() == addr.IP.String() {
			position := sensorNodes[i].DataRing.Position
			if position == sensorDataBufferSize {
				position = 0
			}
			sensorNodes[i].DataRing.Data[position].AtmosphericPressure = response.GetAtmosphericPressure()
			sensorNodes[i].DataRing.Data[position].Humanity = response.GetHumanity()
			sensorNodes[i].DataRing.Data[position].Pm2_5 = response.GetPm2_5()
			sensorNodes[i].DataRing.Data[position].Temperature = response.GetTemperature()
			sensorNodes[i].DataRing.Position += 1
		}
	}
}

func handFeatureResponsePacket(inputBuffer []byte, addr *net.UDPAddr) {
	if debug {
		log.Debug("Handle feature response packet from", addr.String(), "with input", inputBuffer)
	}

	sensorFeatureResponse := &messages.FeatureResponse{}

	err := proto.Unmarshal(inputBuffer, sensorFeatureResponse)
	CheckError(err)

	addSensorFeature(sensorFeatureResponse, addr)
}

func handDataResponsePacket(inputBuffer []byte, addr *net.UDPAddr) {
	if debug {
		log.Debug("Handle data response packet from", addr.String(), "with input", inputBuffer)
	}

	dataResponse := &messages.DataResponse{}

	err := proto.Unmarshal(inputBuffer, dataResponse)
	CheckError(err)

	addDataPoint(dataResponse, addr)
}

func handleIncomingPacket(buffer []byte, length int, addr *net.UDPAddr) {
	command := getCommandName(buffer[0])
	packet := [] byte(buffer[1:length])

	log.Debug("Address:", addr.String(), "\n")
	log.Debug("Packet hex:", packet, "\n")

	switch command {
	case featureResponseCommand:
		handFeatureResponsePacket(packet, addr)
	case dataResponseCommand:
		handDataResponsePacket(packet, addr)
	}
}

func printSensorsState() {
	for _, elem := range sensorNodes {
		featureJson, err := json.Marshal(elem.Features)
		CheckError(err)
		dataJson, err := json.Marshal(elem.DataRing)
		CheckError(err)
		log.Debug(
			"\n	Address:", elem.Config.Address.String(), "|",
			"Features:", string(featureJson), "|",
			"Data:", string(dataJson), "\n",
		)
	}
}

func responseHandler(connection *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		handleIncomingPacket(buffer, n, addr)

		if err != nil {
			log.Error(err)
		}
	}
}

func initServer(localAddr *net.UDPAddr) *net.UDPConn {
	log.Info("Listen to IP address: ", localAddr.String(), "\n")
	connection, err := net.ListenUDP("udp", localAddr)
	CheckError(err)
	go responseHandler(connection)

	return connection
}

func sensorDataRequester(connection *net.UDPConn) {
	for ; ; {
		for _, elem := range sensorNodes {
			getEnvSensorData(connection, elem)
		}
		time.Sleep(time.Second * 10)
	}
}

func init() {
	localInterface, err := net.InterfaceByName(interfaceName)
	CheckError(err)

	localAddresses, err := localInterface.Addrs()
	CheckError(err)

	ipString := strings.Replace(localAddresses[0].String(), "/64", "", 1)
	localAddr := net.UDPAddr{IP: net.ParseIP(ipString), Port: localPort, Zone: interfaceName}

	connection := initServer(&localAddr)
	time.Sleep(time.Second * 2)
	go getEnvSensorFeature(connection, seedSensorConfig)
	time.Sleep(time.Second * 10)
	go sensorDataRequester(connection)
}

func statusReporter() {
	for ; ; {
		printSensorsState()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	loggingBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backend := logging.NewBackendFormatter(loggingBackend, format)
	logging.SetBackend(backend)

	go statusReporter()
	for ; ; {
	}
}
