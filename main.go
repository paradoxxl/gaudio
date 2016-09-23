package main

import(
"azul3d.org/engine/native/al"
	"fmt"
	"log"
	"time"
	"unsafe"
)


func main() {

	captureDevices := al.StringList(al.AlcGetRawString(nil, al.ALC_CAPTURE_DEVICE_SPECIFIER))
	for i, device := range captureDevices {
		fmt.Printf("capture device %d. %q\n", i, device)
	}
	fmt.Println("")

	allDevices := al.StringList(al.AlcGetRawString(nil, al.ALC_DEVICE_SPECIFIER))
	for i, device := range allDevices {
		fmt.Printf("device %d. %q\n", i, device)
	}

	defaultDevice := al.AlcGetString(nil, al.ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER)
	fmt.Printf("Default Device: %v\n", defaultDevice)

	device,err := al.OpenDevice(captureDevices[1],nil)
	if err != nil{
		fmt.Println(err)
	}
	//defer device.Close()

	fmt.Println("Opened device")

	haveCapture := device.AlcIsExtensionPresent("ALC_EXT_CAPTURE")

	var maxSources int32
	device.AlcGetIntegerv(al.ALC_MONO_SOURCES, 1, &maxSources)
	fmt.Print("Maximum sources:", maxSources)




	if haveCapture && len(captureDevices) > 0 {
		fmt.Println("Have the ALC_EXT_CAPTURE extension.")
		err = device.InitCapture(44100, al.FORMAT_MONO16, 44100*2)
		if err != nil {
			log.Fatal(err)
		}
	}


	stopprogram:= make(chan interface{})
	stoprecording:= make(chan interface{})

	go capture(device,stoprecording,stopprogram)
	time.AfterFunc(time.Second*1,func(){
		stoprecording<-1
	})


	<-stopprogram

	device.Close()

}

func capture(device *al.Device,stoprecording chan interface{},stopprogram chan interface{}){
	device.StartCapture()
	fmt.Println("Started Capturing")

	var buffer = make([]uint32,20)
	bufptr:= unsafe.Pointer(&buffer)

	for{
		select{
		case <-stoprecording:
			device.StopCapture()
			fmt.Println("Stopped Capturing")
			stopprogram<-1
			return
		default:
		//record
		//device.AlcGetIntegerv(al.ALC_CAPTURE_SAMPLES,al.SIZE,sample)
			device.CaptureSamples(bufptr,10)
			fmt.Println(buffer)
		}
	}
}


