package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func AutoPort(Addr string, MaxPing time.Duration) int {
	Client := GetNewClient("0.0.0.0:12318")
	for x := 13478; x < 36000; x++ {
		Ping := Client.Ping(Addr+":"+fmt.Sprint(x), MaxPing)
		if Ping != -1 {
			return x
		}
	}
	return 0
}

type Client struct {
	Address string
}

func GetNewClient(Address string) Client {
	return Client{Address: Address}
}

func MilliToDuration(Milli int) time.Duration {
	x := time.Duration(0)
	for Milli > 0 {
		x += time.Millisecond * 1
		Milli--
	}
	return x
}

var ClientBuffer = make([]byte, 2048)

func (C Client) Ping(Address string, Timeout time.Duration) int {
	StartMilli := time.Now().UnixMilli()
	// Contact server and wait for answer until timeout
	HasAnswered := false
	x := func() {
		_, Error := net.ResolveUDPAddr("udp", C.Address)
		if Error != nil {
			log.Fatalln(Error.Error())
		}
		Target, Error := net.ResolveUDPAddr("udp", Address)
		if Error != nil {
			log.Fatalln(Error.Error())
		}
		conn, err := net.Dial("udp", Target.String())
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		conn.SetDeadline(time.Now().Add(Timeout))
		for {
			fmt.Fprintf(conn, "ping!")
			_, err = bufio.NewReader(conn).Read(ClientBuffer)
			if err == nil {
				if ClientBuffer[0] == 111 && ClientBuffer[1] == 107 && ClientBuffer[2] == 33 {
					HasAnswered = true
					break
				} else {
					fmt.Println("event!")
				}
			} else {
				fmt.Printf("Some error %v\n", err)
			}
			EndMilli := time.Now().UnixMilli()
			TimeElapsed := EndMilli - StartMilli
			if MilliToDuration(int(TimeElapsed)) >= Timeout {
				conn.Close()
				break
			}
		}
	}
	go x()
	EndMilli := time.Now().UnixMilli()
	TimeElapsed := EndMilli - StartMilli
	for MilliToDuration(int(TimeElapsed)) <= Timeout || HasAnswered {
		if HasAnswered {
			return int(TimeElapsed)
		}
		EndMilli = time.Now().UnixMilli()
		TimeElapsed = EndMilli - StartMilli
	}
	return -1
}

func (C Client) Listen(Address string, Duration time.Duration) {
	if Duration == 0 {
		for {
			x := func() {
				_, Error := net.ResolveUDPAddr("udp", C.Address)
				if Error != nil {
					log.Fatalln(Error.Error())
				}
				Target, Error := net.ResolveUDPAddr("udp", Address)
				if Error != nil {
					log.Fatalln(Error.Error())
				}
				conn, err := net.Dial("udp", Target.String())
				if err != nil {
					fmt.Printf("Some error %v", err)
					return
				}
				fmt.Fprintf(conn, "im hearing you to XD!")
				for {
					_, err = bufio.NewReader(conn).Read(ClientBuffer)
					if err == nil {
						if ClientBuffer[0] == 111 && ClientBuffer[1] == 107 && ClientBuffer[2] == 33 {
							break
						}
						fmt.Println("received!", string(ClientBuffer))
					} else {
						fmt.Printf("Some error %v\n", err)
					}
				}
				for {
					_, err = bufio.NewReader(conn).Read(ClientBuffer)
					if err == nil {
						fmt.Println("received!", string(ClientBuffer))
						if strings.Contains(string(ClientBuffer), "New BlockBreakEvent [e]") {
							UUID_BlockCoords := strings.Split(string(ClientBuffer), " ")[3] + " " + strings.Split(string(ClientBuffer), " ")[5]
							// UUID := strings.Replace(strings.Replace(strings.Split(UUID_BlockCoords, " ")[0], "(", "", -1), ")", "", -1)
							// fmt.Println("UUID: " + UUID + " ")
							XCoord, YCoord, ZCoord := strings.Split(strings.Split(UUID_BlockCoords, " ")[1], "/")[0], strings.Split(strings.Split(UUID_BlockCoords, " ")[1], "/")[1], strings.Split(strings.Split(UUID_BlockCoords, " ")[1], "/")[2]
							// fmt.Println("\""+XCoord+"\"", "\""+YCoord+"\"", "\""+ZCoord+"\"")

							// fmt.Println("\""+XCoord+"\"", "\""+YCoord+"\"", "\""+ZCoord+"\"")

							XCoord = strings.ReplaceAll(XCoord, "\x00", "")
							YCoord = strings.ReplaceAll(YCoord, "\x00", "")
							ZCoord = strings.ReplaceAll(ZCoord, "\x00", "")

							X, Error := strconv.Atoi(XCoord)
							if Error != nil {
								log.Fatalln(Error.Error())
							}
							Y, Error := strconv.Atoi(YCoord)
							if Error != nil {
								log.Fatalln(Error.Error())
							}
							Z, Error := strconv.Atoi(ZCoord)
							if Error != nil {
								log.Fatalln(Error.Error())
							}

							BL := BlockLocation{
								X: X,
								Y: Y,
								Z: Z,
							}
							E := Event{
								EventType: EventBlockBreak,
								EventData: []Data{Data{
									Block{Location: BL, LastUpdate: &Block{Location: BL, LastUpdate: nil}},
								}},
							}
							CallEvent(EventBlockBreak, E)
							// log.Println("Called!")
						}
					} else {
						fmt.Printf("Some error %v\n", err)
					}
				}
			}
			go x()
			for {
			}
		}
	} else {
		// conn.SetDeadline(time.Now().Add(Duration))
		// limited time!
	}
}
